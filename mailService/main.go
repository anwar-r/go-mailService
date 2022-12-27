package main

import (
	cgf "mailService/Config"
	model "mailService/Model"
	rout "mailService/Router"
)

var smtpAuth []model.SmtpAuth

func main() {
	cgf.Init()
	rout.HandleFunc()
}
