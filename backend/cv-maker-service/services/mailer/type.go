package mailer

import "context"

type (
	GRPCService interface {
		SendEmail(context.Context, *MailerRequest) (*MailerResponse, error)
	}

	Service interface {
		SendEmailWithParameters(ctx context.Context, to, firstName, lastName string, createdCVBytes []byte) error
	}

	SimpleService struct {
	}
)
