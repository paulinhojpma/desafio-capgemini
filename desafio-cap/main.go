package main

import (
	"defafio-cap/sequence-validator/config"
	"defafio-cap/sequence-validator/database"
	"defafio-cap/sequence-validator/message"
	"defafio-cap/sequence-validator/utils"
	"defafio-cap/sequence-validator/validate"
	"defafio-cap/sequence-validator/web"
	"fmt"
	"os"
)

func main() {

	configNew := config.NewConfig("conf.json")

	OptionsDBClient := &database.OptionsDBClient{Driver: configNew.DatabaseDriver, URL: configNew.DBBizHost}
	db, err := OptionsDBClient.ConfigDatabase()
	if err != nil {
		fmt.Println("Fail to connect to the database, error: ", err)
		os.Exit(1)
	}

	fmt.Println("Connected to the database")

	configMessage := &message.OptionsMessageCLient{URL: configNew.MessagingURL, Driver: configNew.MessagingDriver}

	mes, err := configMessage.ConfigureMessageQueue()
	if err != nil {
		fmt.Println("Fail to connect to the message service, error: ", err)
		os.Exit(1)
	}
	fmt.Println("Connected to the message queue service")

	validator := validate.CreateValidator(validate.ValidatorAssync)
	fmt.Println("New sequence validator created")

	handler := &web.Handler{
		DB:        *db,
		Message:   *mes,
		Validator: validator,
	}

	manager := web.NewRoutes(handler)
	msgChan, err := handler.Message.ReceiveMessage(configNew.ReceiveRoute)
	if err != nil {
		fmt.Println("Fail to receive from the message queue, error: ", err)
		os.Exit(1)
	}

	for msg := range msgChan {
		fmt.Println("Receiving messages")
		messageResponse := manager.CallService(utils.PathMethod(msg.Method), &msg)
		go func() {
			err := manager.ManagerMessage(handler.Message, messageResponse)
			if err != nil {
				fmt.Println("Error on respond the request - ", err)
			}
		}()
	}

}
