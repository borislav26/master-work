package mailer

import (
	"context"
	mail "github.com/xhit/go-simple-mail/v2"
	error2 "mailer-service/error"
)

func (s SimpleService) SendEmail(ctx context.Context, request *MailerRequest) (*MailerResponse, error) {
	res := &MailerResponse{
		Result: false,
	}

	emailClientConnection, err := s.MailServer.Connect()
	if err != nil {
		return res, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to connect to mailer service")
	}

	email := mail.NewMSG().
		AddTo(request.MailEntry.GetToAddress()).
		SetFrom("borislavpetric66@gmail.com").
		SetSubject("Created CV").
		SetBody(mail.TextPlain, "Thank you for using our service.")
	email.AddAttachmentData(request.MailEntry.GetCreatedCVData(), "cv.pdf", "application/pdf")

	err = email.Send(emailClientConnection)
	if err != nil {
		return res, error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to send mail to user").With("email", request.MailEntry.ToAddress)
	}
	res.Result = true
	return res, nil
}
