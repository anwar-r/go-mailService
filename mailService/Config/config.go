package config

import (
	"io/ioutil"
	model "mailService/Model"

	"gopkg.in/yaml.v2"
)

var Config []model.SmtpAuth

func Init() {
	data, err := ioutil.ReadFile("./Config/auth.yml")
	if err != nil {
		err.Error()
	}
	// var src model.SmtpAuth
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		err.Error()
	}
}
