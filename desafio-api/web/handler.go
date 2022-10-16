package web

import (
	"desafio-api/sequence-validator/core"
	"desafio-api/sequence-validator/message"
	"desafio-api/sequence-validator/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const (
	VALIDATEMETHOD = "POST"
	INFOMETHOD     = "GET"
)

type Handler struct {
	Message   message.IMessageClient
	RoutePath string
}

func (h *Handler) ValidateSequence(w http.ResponseWriter, r *http.Request) {
	letters := &Letters{}

	err := json.NewDecoder(r.Body).Decode(letters)
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Invalid request body"), "Error on validate a sequence", http.StatusBadRequest)
		return
	}

	if !utils.CheckIfSliceIsSquare(letters.Letters) {
		h.CoreRespondErro(w, r, "", errors.New("Array is not simetric"), "Error on validate a sequence", http.StatusBadRequest)
		return
	}
	receiveRoute := uuid.New().String()
	bit, _ := json.Marshal(letters)
	mes := message.MessageParam{
		ID:     receiveRoute,
		Method: VALIDATEMETHOD,
		Body:   bit,
	}

	err = h.Message.PublishMessage(h.RoutePath, &mes)
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Request cannot be processed"), "Error on validate a sequence", http.StatusInternalServerError)
		return
	}
	mesResponse, err := h.Message.ReceiveOneMessage(receiveRoute)
	if err != nil {
		fmt.Println("Error on receive the validate response - ", err)
		h.CoreRespondErro(w, r, "", errors.New("Request cannot be processed"), "Error on validate a sequence", http.StatusInternalServerError)
		return
	}
	response := &SequenceValidatorResponse{}
	err = json.Unmarshal(mesResponse.Body, response)
	if err != nil {
		fmt.Println("Error on Ummarshal the validate response - ", err)
		h.CoreRespondErro(w, r, "", errors.New("Request cannot be processed"), "Error on validate a sequence", http.StatusInternalServerError)
		return
	}
	h.CoreRespondSucess(w, r, http.StatusOK, response)
	return

}

func (h *Handler) GetInfoSequences(w http.ResponseWriter, r *http.Request) {
	receiveRoute := uuid.New().String()
	err := h.Message.PublishMessage(h.RoutePath, &message.MessageParam{ID: receiveRoute, Method: INFOMETHOD})
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Request cannot be processed"), "Error on get information about the sequence a sequence", http.StatusInternalServerError)
		return
	}
	messageResponse, err := h.Message.ReceiveOneMessage(receiveRoute)
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Request cannot be processed"), "Error on get information about the sequence a sequence", http.StatusInternalServerError)
		return
	}
	response := &InfoSequence{}
	err = json.Unmarshal(messageResponse.Body, response)
	if err != nil {
		h.CoreRespondErro(w, r, "", errors.New("Request cannot be processed"), "Error on get information about the sequence a sequence", http.StatusInternalServerError)
		return
	}

	h.CoreRespondSucess(w, r, http.StatusOK, response)
	return
}

func (h *Handler) CoreRespondErro(w http.ResponseWriter, r *http.Request, idOperation string, erro error, message string, codeError int) {
	core.Respond(w, r, codeError, core.ErrDetail{
		Resource: message, Code: strconv.Itoa(codeError), Message: erro.Error(), IDOperation: idOperation,
	})
}

func (h *Handler) CoreRespondSucess(w http.ResponseWriter, r *http.Request, code int, object interface{}) {
	core.Respond(w, r, code, object)
}
