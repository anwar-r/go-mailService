package model

type SmtpAuth struct {
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Smpt     struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"smpt"`
	To string `json:"to"`
}

type MailRequest struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}
