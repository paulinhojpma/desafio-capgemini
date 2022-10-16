package web

import (
	"bytes"
	"desafio-api/sequence-validator/message"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

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

	handler.Message = InitMessage()
	handler.RoutePath = "api.route"

	return handler

}

func TestValidateSequence(t *testing.T) {
	cases := []struct {
		should     string
		errMessage string
		scenario   *Letters
	}{
		{
			errMessage: "Should return true because there is a vertical sequence",
			should:     `{"is_valid":true}`,
			scenario: &Letters{
				Letters: []string{
					"DUHHDB",
					"DUBUHD",
					"UBUUHU",
					"BHBDHH",
					"UDBDUH",
					"UHDDDD"},
			},
		},
		{
			errMessage: "Should return true because there is a horizontal sequence",
			should:     `{"is_valid":true}`,
			scenario: &Letters{
				Letters: []string{
					"DUHHDB",
					"DUBUHD",
					"UBUUHH",
					"BHBDHH",
					"UDBDUH",
					"UHDDHH"},
			},
		},
		{
			errMessage: "Should return true because there is a left to right diagonal sequence",
			should:     `{"is_valid":true}`,
			scenario: &Letters{
				Letters: []string{
					"DUHHDB",
					"DUBUHD",
					"UBUUHU",
					"BUBDHH",
					"UDUDUH",
					"UHDUHD"},
			},
		},
		{
			errMessage: "Should return true because there is a right to left diagonal sequence",
			should:     `{"is_valid":true}`,
			scenario: &Letters{
				Letters: []string{
					"DUHHDB",
					"DUBUHD",
					"UBUUHH",
					"BHBDHH",
					"UDBHUH",
					"UHHDHD"},
			},
		},

		{
			errMessage: "Should return false because there is no valid sequence ",
			should:     `{"is_valid":false}`,
			scenario: &Letters{
				Letters: []string{
					"DUHHDB",
					"DUBUHD",
					"UBUUHU",
					"BHBDHH",
					"UDBDUH",
					"UHDDHD"},
			},
		},
	}

	for _, cas := range cases {
		bit, _ := json.Marshal(cas.scenario)
		fmt.Println(string(bit))
		r := httptest.NewRequest("POST", "/sequence", bytes.NewReader(bit))

		w := httptest.NewRecorder()
		hWeb := initHandler()
		RouterTest(hWeb).ServeHTTP(w, r)
		res, _ := ioutil.ReadAll(w.Body)
		val := strings.TrimSpace(string(res))
		if val != cas.should {
			t.Errorf("expect %s, got - %s ", cas.should, string(res))
		}
	}

}

func TestGetInfoSequence(t *testing.T) {
	r := httptest.NewRequest("GET", "/stats", nil)

	w := httptest.NewRecorder()
	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	res, errBody := ioutil.ReadAll(w.Body)
	fmt.Println(string(res))
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
}

// func TestListTransacoes(t *testing.T) {
// 	hWeb := initHandler()
// 	r := httptest.NewRequest("GET", "/transactions", nil)
// 	w := httptest.NewRecorder()
// 	RouterTest(hWeb).ServeHTTP(w, r)
// 	data, errBody := ioutil.ReadAll(w.Body)
// 	fmt.Println(string(data))
// 	if errBody == nil {
// 		t.Error("expect nil, got - ", errBody)
// 	}
// }
