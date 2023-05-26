package builder

import "mailer-service/services/mailer"

type Services struct {
	MailerService mailer.GRPCService
}

func BuildServices() *Services {
	mailServer := mailer.NewSMTPServer()

	mailerService := &mailer.SimpleService{
		UnimplementedMailerServiceServer: mailer.UnimplementedMailerServiceServer{},
		MailServer:                       *mailServer,
	}

	return &Services{
		MailerService: mailerService,
	}
}
