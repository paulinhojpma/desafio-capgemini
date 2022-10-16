package message

import (
	"fmt"
	"testing"
)

func InitMessage() IMessageClient {
	config := &OptionsMessageCLient{
		URL:    "localhost:6379",
		Driver: "redis",
		Args: map[string]interface{}{
			"DB": 0,
		},
	}

	IMessage, err := config.ConfigureMessageQueue()
	if err != nil {
		fmt.Println("Error on initialize message queue - ", err)
	}
	return *IMessage
}

func TestPublishMessage(t *testing.T) {
	IMessage := InitMessage()
	message := &MessageParam{

		Body: []byte("enviando"),
	}

	err := IMessage.PublishMessage("rota", message)
	if err != nil {
		t.Error("Expect nil, got ", err)
	}
}

func TestReceiveMessage(t *testing.T) {
	IMessage := InitMessage()
	message := &MessageParam{

		Body: []byte("enviando"),
	}

	IMessage.PublishMessage("rota", message)
	imChan, _ := IMessage.ReceiveMessage("rota")
	result := ""
	for mes := range imChan {
		result = string(mes.Body)
		break
	}
	if result != "enviando" {
		t.Error("Expect 'enviando', got - ", result)
	}

}

func TestReceiveOneMessage(t *testing.T) {
	IMessage := InitMessage()
	message := &MessageParam{

		Body: []byte("enviando"),
	}

	IMessage.PublishMessage("rota", message)
	mes, _ := IMessage.ReceiveOneMessage("rota")

	result := string(mes.Body)

	if result != "enviando" {
		t.Error("Expect 'enviando', got - ", result)
	}

}
