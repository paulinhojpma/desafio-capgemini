package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	env "github.com/caarlos0/env"
)

var config *Configuracoes

// Configuracoes ...
type Configuracoes struct {
	Port            int                    `json:"port" env:"PORT"`
	AllowedParam    string                 `json:"allowedParam" env:"ALLOWED_PARAM"`
	AWSRegion       string                 `json:"AWSRegion" env:"AWS_REGION"`
	MessagingURL    string                 `json:"messagingURL" env:"MESSAGING_URL"`
	MessagingDriver string                 `json:"messagingDriver" env:"MESSAGING_DRIVER"`
	MessagingArgs   map[string]interface{} `json:"MessagingArgs" env:"MESSAGING_ARGS"`
	ReceiveRoute    string                 `json:"receiveRoute" env:"RECEIVE_ROUTE"`
}

// NewConfig ...
func NewConfig(file string) *Configuracoes {
	var erro error

	conf := &Configuracoes{}

	if file != "" {
		fmt.Println(file)

		bufConf, err := ioutil.ReadFile(file)
		if err == nil {
			erro = json.Unmarshal(bufConf, conf)
			if erro != nil {
				log.Println(erro)
			}
		}
	}

	if erro = env.Parse(conf); erro != nil {
		log.Println(erro)
	}

	return conf
}

// Config ...
func Config() *Configuracoes {
	return config
}
