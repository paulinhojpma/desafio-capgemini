package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	redis "github.com/go-redis/redis/v8"
)

type Letters struct {
	Letters []string `json:"letters"`
}

type MessageParam struct {
	ID     string `json:"id"`
	Body   []byte `json:"body"`
	Method string `json:"method"`
}

func main() {
	clientRedis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB

	})
	forever := make(chan bool)
	fmt.Println(clientRedis.Ping(context.Background()).Val())
	for i := 0; i < 10000; i++ {
		go func(i int) {
			for j := 0; j < 5; j++ {
				MakeRequest(i + j + 1)
			}

		}(i)

	}
	//MakeRequest()
	// for j := 0; j < 5; j++ {
	// 	go func(j int) {
	// 		i := 0
	// 		for {

	// 			res, _ := clientRedis.BLPop(context.Background(), 0, fmt.Sprintf("vicente%d", j)).Result()
	// 			messageParam := &MessageParam{}
	// 			json.Unmarshal([]byte(res[1]), messageParam)
	// 			fmt.Println("Receiving - ", i*(j+1), " - ", string(messageParam.Body))
	// 			i++
	// 		}
	// 	}(j)
	// }
	// for j := 0; j < 5; j++ {
	// 	go func(j int) {
	// 		for i := 0; i < 10000; i++ {

	// 			bit, _ := json.Marshal(&Letters{Letters: generateArray()})

	// 			bit, _ = json.Marshal(&MessageParam{ID: fmt.Sprintf("vicente%d", j), Body: bit, Method: "POST"})

	// 			clientRedis.RPush(context.Background(), "api.route", string(bit)).Err()
	// 			fmt.Println("Sending - ", i*(j+1))
	// 		}
	// 	}(j)
	// }

	// bit, _ := json.Marshal(&MessageParam{ID: "vicente", Body: nil, Method: "GET"})

	// clientRedis.RPush(context.Background(), "api.route", string(bit)).Err()
	// res, _ := clientRedis.BLPop(context.Background(), 0, "vicente").Result()
	// messageParam := &MessageParam{}
	// json.Unmarshal([]byte(res[1]), messageParam)
	// fmt.Println(string(messageParam.Body))

	<-forever
}

func generateArray() []string {
	charValids := []string{"B", "U", "D", "H"}

	arr := make([]string, 6)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {

			arr[i] += charValids[rand.Intn(4)]
		}
	}

	return arr
}

func MakeRequest(i int) {
	val := generateArray()
	//fmt.Printf("Enviando - %d\n", i)
	bit, _ := json.Marshal(&Letters{Letters: val})
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8890/validateSequences", bytes.NewReader(bit))
	_, _ = http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("client: error making http request: %s\n", err)
	// 	//panic(err)
	// }

	//fmt.Println("Response - ", i, " - ", res.StatusCode)
}
