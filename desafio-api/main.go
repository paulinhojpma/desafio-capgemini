package main

import (
	"desafio-api/sequence-validator/config"
	"desafio-api/sequence-validator/message"
	"desafio-api/sequence-validator/web"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
)

const (
	TimeOutSecond = 120
)

func main() {
	configNew := config.NewConfig("conf.json")

	configMessage := &message.OptionsMessageCLient{URL: configNew.MessagingURL, Driver: configNew.MessagingDriver, Args: configNew.MessagingArgs}

	mes, err := configMessage.ConfigureMessageQueue()
	if err != nil {
		fmt.Println("Fail to connect to the message service, error: ", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the message queue service")

	handler := &web.Handler{
		RoutePath: configNew.ReceiveRoute,
		Message:   *mes,
	}
	allowedParam := make(map[string][]string)
	if err := json.Unmarshal([]byte(`{"Origins":["*"],"Headers":["*"],"Methods":["GET","POST","PUT", "DELETE","OPTIONS"]}`), &allowedParam); err != nil {
		log.Println("Error on json Unmarshal do allowedOrigins. Detail:", err)
		os.Exit(1)
	}

	router := web.Router(handler)
	c := cors.New(cors.Options{
		AllowedOrigins: allowedParam["Origins"],
		AllowedHeaders: allowedParam["Headers"],
		AllowedMethods: allowedParam["Methods"],

		Debug: false,
	})
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", configNew.Port),
		Handler:      c.Handler(router),
		ReadTimeout:  TimeOutSecond * time.Second,
		WriteTimeout: TimeOutSecond * time.Second,
	}
	log.Println("Waiting connection")
	if err := s.ListenAndServe(); err != nil {
		log.Println("Error on start the Server. Error:", err)
		os.Exit(1)
	}

}
