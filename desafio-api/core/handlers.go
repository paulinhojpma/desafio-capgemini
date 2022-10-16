package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	mux "github.com/gorilla/mux"
	// "github.com/pkg/errors"
)

// RequestTest ..
func RequestTest(router *mux.Router, metodo, url string, bodyObj, responseObj interface{}) (int, error) {
	bodyRequestJSON, erroReq := requestEncodeJSON(bodyObj)
	if erroReq != nil {
		return 0, erroReq
	}

	r, err := http.NewRequest(metodo, url, bodyRequestJSON)
	if err != nil {
		return 0, err
	}
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	erroBody := responseDecodeJSON(w.Body, responseObj)
	if erroBody != nil {
		return w.Code, erroReq
	}

	return w.Code, nil
}

// responseDecodeJSON .. usado RequestTest
func responseDecodeJSON(bodyResponse io.Reader, response interface{}) error {
	var body, errBody = ioutil.ReadAll(bodyResponse)
	if errBody != nil {
		return errBody
	}

	errJSON := json.Unmarshal(body, response)
	if errJSON != nil {
		return errJSON
	}

	return nil
}

// requestEncodeJSON .. usado RequestTest
func requestEncodeJSON(objRequest interface{}) (*bytes.Buffer, error) {
	bodyRequestJSON := new(bytes.Buffer)
	encodeJSON, erro := json.Marshal(objRequest)
	if erro != nil {
		return nil, erro
	}
	bodyRequestJSON.Write(encodeJSON)

	return bodyRequestJSON, nil
}

//DecodeBodyJSON ...
func DecodeBodyJSON(r *http.Request, v interface{}) error {
	conteudo, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		return errors.New(ErrorReadAllBuffer)
	}
	// logger.Trace(fmt.Sprintf("Request: %s", string(conteudo)))

	if erro = json.Unmarshal(conteudo, v); erro != nil {
		return errors.New(ErrorJSONUnmarshal)
	}

	return nil
}

//Respond ...
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Println("Erro encode:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("Erro copy buf:", err)
	}

	log.Println(r.URL, "status:", status)
}

//handleNotFound ...
func handleNotFound(w http.ResponseWriter, r *http.Request) {

	body := fmt.Sprintf(`{"Message": "URL não encontrada", "Code":"%d"}`, http.StatusNotFound)
	Respond(w, r, http.StatusNotFound, body)

}
