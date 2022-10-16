package web

import (
	"defafio-cap/sequence-validator/database"
	"defafio-cap/sequence-validator/message"
	"defafio-cap/sequence-validator/validate"
	"encoding/json"
	"fmt"
	"testing"
)

func InitDataBase() (database.IDataBase, error) {
	DBOpt := &database.OptionsDBClient{
		URL:    "postgresql://root:luke@localhost:5434/app",
		Driver: "postgres",
	}
	db, err := DBOpt.ConfigDatabase()
	return *db, err

}

func InitMessage() message.IMessageClient {
	config := &message.OptionsMessageCLient{
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

func initHandler() *Handler {
	handler := &Handler{}

	handler.Validator = validate.CreateValidator(validate.ValidatorAssync)
	handler.Message = InitMessage()
	handler.DB, _ = InitDataBase()

	return handler

}

func TestCreateSequenceValid(t *testing.T) {
	h := initHandler()

	letters := &Letters{
		Letters: []string{"DUHHDB", "DUBUHD", "UBUUHU", "BHBHHH", "UDBDUH", "UHDDDD"},
	}
	b, _ := json.Marshal(letters)
	messageParam := message.MessageParam{
		ID:   "10",
		Body: b,
	}
	messageResponse := h.CreateSequence(&messageParam)
	fmt.Println(string(messageResponse.Body))
	if string(messageResponse.Body) != `{"is_valid": true}` {
		t.Error(`Expect- {"is_valid": true}, got - `, string(messageResponse.Body))
	}

}

func TestCreateSequenceInValid(t *testing.T) {
	h := initHandler()

	letters := &Letters{
		Letters: []string{"DUHHDB", "DUBUHD", "UBUDHU", "BHBHHH", "UDBDUH", "UHDUUD"},
	}
	b, _ := json.Marshal(letters)
	messageParam := message.MessageParam{
		ID:   "10",
		Body: b,
	}
	messageResponse := h.CreateSequence(&messageParam)
	fmt.Println(string(messageResponse.Body))
	if string(messageResponse.Body) != `{"is_valid": false}` {
		t.Error(`Expect- {"is_valid": true}, got - `, string(messageResponse.Body))
	}

}

func TestGetInfoSequence(t *testing.T) {
	h := initHandler()

	messageResponse := h.GetInfoSequences(&message.MessageParam{})
	fmt.Println(string(messageResponse.Body))

	if messageResponse.Body == nil {
		t.Error("Expect response full, got - ", string(messageResponse.Body))
	}
}
