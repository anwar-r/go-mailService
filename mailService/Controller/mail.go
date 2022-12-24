package controller

import (
	"encoding/json"
	"io/ioutil"
	model "mailService/Model"
	"net/http"
	"net/smtp"

	"gopkg.in/yaml.v2"
)

func Mail(w http.ResponseWriter, r *http.Request) {

	jsonBody := json.NewDecoder(r.Body)

	var resp = model.JsonResponse{}
	var mail = model.MailRequest{}

	err := jsonBody.Decode(&mail)
	if err != nil {
		resp.Status = err.Error()
	}

	if mail.Source == "TEST" {
		log, err := parserYmal("TEST")
		if err != nil {
			resp.Status = err.Error()

		}
		resp = mailSend(log, mail)

	} else {
		resp.Status = "error"
		resp.Message = "Send Valid Source "
	}

	json.NewEncoder(w).Encode(resp)
}

func parserYmal(source string) (model.SmtpAuth, error) {
	var log model.SmtpAuth

	data, err := ioutil.ReadFile("./Config/auth.yml")
	if err != nil {
		err.Error()
	}
	switch source {
	case "TEST":
		{
			err = yaml.Unmarshal(data, &log)
			if err != nil {
				err.Error()
			}
		}
	}

	return log, err
}

func mailSend(log model.SmtpAuth, mail model.MailRequest) (result model.JsonResponse) {
	var resp = model.JsonResponse{}

	auth := smtp.PlainAuth("", log.UserName, log.Password, log.Smpt.Host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.

	to := []string{log.To}
	msg := []byte("Subject: WebSite Email!\r\n" +
		"\r\n" +
		mail.Message)
	err := smtp.SendMail(log.Smpt.Host+":"+log.Smpt.Port, auth, log.UserName, to, msg)
	if err != nil {
		resp.Status = err.Error()

	} else {
		resp.Status = mail.Source + " " + "Mail Send Successfully"
	}
	return resp
}
