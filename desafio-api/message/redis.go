package message

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	TIMEOUT = 0
)

type redisMessage struct {
	client *redis.Client
	subs   *redis.PubSub
}

func (m *redisMessage) connectService(config *OptionsMessageCLient) error {
	fmt.Println("ARGS ", config.Args["DB"])
	clientRedis := redis.NewClient(&redis.Options{
		Addr:        config.URL,
		Password:    config.Password,
		DB:          0, // use default DB
		PoolSize:    1000,
		PoolTimeout: 10 * time.Second,
	})

	res, err := clientRedis.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Error on init redis . Error=", err)
		return err

	}
	fmt.Println("Redis connected - ", res)
	m.client = clientRedis
	m.subs = clientRedis.Subscribe(context.Background())
	return nil
}

func (m *redisMessage) PublishMessage(routing string, params *MessageParam) error {
	bit, err := json.Marshal(params)
	if err != nil {
		return err
	}

	return m.client.RPush(context.Background(), routing, string(bit)).Err()
}

func (m *redisMessage) ReceiveMessage(routing string) (<-chan MessageParam, error) {
	messageParamChan := make(chan MessageParam)

	go func() {

		for {

			res, _ := m.client.BLPop(context.Background(), 0, routing).Result()
			messageParam := &MessageParam{}
			json.Unmarshal([]byte(res[1]), messageParam)
			messageParamChan <- *messageParam
		}

	}()

	return messageParamChan, nil
}

func (m *redisMessage) ReceiveOneMessage(routing string) (*MessageParam, error) {
	messageParam := &MessageParam{}
	i := 0
	for {
		res, err := m.client.BLPop(context.Background(), time.Second*TIMEOUT, routing).Result()
		if err != nil {
			if i >= 2 {
				return nil, err
			}

		} else {
			json.Unmarshal([]byte(res[1]), messageParam)
			break
		}
		i++
	}

	return messageParam, nil
}
