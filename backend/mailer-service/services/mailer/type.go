package mailer

import (
	"context"
	mail "github.com/xhit/go-simple-mail/v2"
)

type (
	GRPCService interface {
		SendEmail(context.Context, *MailerRequest) (*MailerResponse, error)
	}

	SimpleService struct {
		UnimplementedMailerServiceServer
		MailServer mail.SMTPServer
	}
)
