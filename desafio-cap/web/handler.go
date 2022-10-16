package web

import (
	"defafio-cap/sequence-validator/database"
	"defafio-cap/sequence-validator/message"
	"defafio-cap/sequence-validator/validate"
	"encoding/json"
	"fmt"
	"strings"
)

type Handler struct {
	DB        database.IDataBase
	Message   message.IMessageClient
	Validator validate.SequecenValidator
}

type Letters struct {
	Letters []string `json:"letters"`
}

func (h *Handler) CreateSequence(m *message.MessageParam) *message.MessageParam {

	letters := &Letters{}
	sequence := &database.Sequence{}
	fmt.Println("body - ", string(m.Body))
	err := json.Unmarshal(m.Body, letters)

	if err != nil {
		fmt.Println("It was not possible to umMarshal the data, err - ", err)
		return &message.MessageParam{ID: m.ID, Body: []byte(fmt.Sprintf(`{"is_valid": %t}`, false))}

	}

	sequence.IsValid = h.Validator.ValidateSequence(letters.Letters)
	sequence.Letters = strings.Join(letters.Letters, ",")

	err = h.DB.CreateSequence(sequence)
	if err != nil {
		fmt.Println("It was not possible to create the data, err - ", err)
		return &message.MessageParam{ID: m.ID, Body: []byte(fmt.Sprintf(`{"is_valid": %t}`, false))}
	}

	return &message.MessageParam{ID: m.ID, Body: []byte(fmt.Sprintf(`{"is_valid": %t}`, sequence.IsValid))}

}

func (h *Handler) GetInfoSequences(m *message.MessageParam) *message.MessageParam {

	info, err := h.DB.GetInfoSequences()
	if err != nil {
		return &message.MessageParam{ID: m.ID, Body: nil}
	}

	bit, _ := json.Marshal(info)
	fmt.Println("Retornado - ", string(bit))
	return &message.MessageParam{ID: m.ID, Body: bit}
}
