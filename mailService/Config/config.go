package config

import (
	"io/ioutil"
	model "mailService/Model"

	"gopkg.in/yaml.v2"
)

var config model.SmtpAuth

func Init() {
	data, err := ioutil.ReadFile("./Config/auth.yml")
	if err != nil {
		err.Error()
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		err.Error()
	}
}
