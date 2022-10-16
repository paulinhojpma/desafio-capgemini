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
	Service         string `json:"service" env:"SERVICE"`
	DBBizNome       string `json:"DBBizNome" env:"DB_BIZ_NOME"`
	DBBizHost       string `json:"DBBizHost" env:"DB_BIZ_HOST"`
	DBBizPorta      int    `json:"DBBizPorta" env:"DB_BIZ_PORTA"`
	DBBizUser       string `json:"DBBizUser" env:"DB_BIZ_USER"`
	DBBizSenha      string `json:"DBBizSenha" env:"DB_BIZ_SENHA"`
	DatabaseDriver  string `json:"databaseDriver" env:"DATABASE_DRIVER"`
	MessagingURL    string `json:"messagingURL" env:"MESSAGING_URL"`
	MessagingDriver string `json:"messagingDriver" env:"MESSAGING_DRIVER"`

	ReceiveRoute string `json:"receiveRoute" env:"RECEIVE_ROUTE"`
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
