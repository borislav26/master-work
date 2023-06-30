package mailer

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"time"
)

func NewSMTPServer() *mail.SMTPServer {
	server := mail.NewSMTPClient()
	server.KeepAlive = false
	server.SendTimeout = 10 * time.Second
	server.ConnectTimeout = 10 * time.Second
	server.Host = "mailhog"
	server.Port = 1025
	server.Username = ""
	server.Password = ""
	server.Encryption = mail.EncryptionNone

	return server
}
