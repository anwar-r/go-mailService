package controller

import (
	"encoding/json"
	config "mailService/Config"
	model "mailService/Model"
	"net/http"
	"net/smtp"
)

func Mail(w http.ResponseWriter, r *http.Request) {

	jsonBody := json.NewDecoder(r.Body)

	var resp = model.JsonResponse{}
	var mail = model.MailRequest{}

	err := jsonBody.Decode(&mail)
	if err != nil {
		resp.Status = err.Error()
	}
	for _, src := range config.Config {
		if mail.Source == src.Source {
			resp = mailSend(src, mail)
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func mailSend(log model.SmtpAuth, mailreq model.MailRequest) (result model.JsonResponse) {
	var resp = model.JsonResponse{}
	auth := smtp.PlainAuth("", log.UserName, log.Password, log.Smpt.Host)
	addr := log.Smpt.Host + ":" + log.Smpt.Port
	to := []string{mailreq.To}
	msg := []byte("Subject:" + mailreq.Subject + "\r\n" +
		"\r\n" +
		mailreq.Body)

	err := smtp.SendMail(addr, auth, mailreq.From, to, msg)
	if err != nil {
		resp.Status = err.Error()

	} else {
		resp.Status = mailreq.Source + " " + "Mail Send Successfully"
	}
	return resp
}
