package model

import (
	"io"
)

type SmtpAuth struct {
	Source   string `json:"source"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Smpt     struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"smpt"`
	To string `json:"to"`
}

type MailRequest struct {
	Source  string  `json:"source"`
	To      string  `json:"to"`
	From    string  `json:"from"`
	Cc      string  `json:"cc"`
	Bcc     string  `json:"bcc"`
	Subject string  `json:"subject"`
	Body    string  `json:"body"`
	Message Message `json:"message"`
}

type Message struct {
	Header map[string][]string
	Body   io.Reader
}
