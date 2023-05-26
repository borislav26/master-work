package mailer

import (
	"context"
	"cv-maker-service/environement"
	error2 "cv-maker-service/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s SimpleService) SendEmail(ctx context.Context, req *MailerRequest) (*MailerResponse, error) {
	conn, err := grpc.Dial(environement.GRPC_MAILER_SERVICE_HOST.GetValue(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	mailerClient := NewMailerServiceClient(conn)
	res, err := mailerClient.SendEmail(ctx, req)
	if err != nil {
		return res, error2.ServiceLayerError.
			WrapWithUserMessage(err, "Failed to send email to user").
			With("email", req.GetMailEntry().ToAddress)
	}

	if !res.Result {
		return res, error2.ServiceLayerError.Wrap(err)
	}

	return res, nil
}

func (s SimpleService) SendEmailWithParameters(ctx context.Context, to, firstName, lastName string, createdCVBytes []byte) error {
	request := &MailerRequest{
		MailEntry: &MaiData{
			ToAddress:     to,
			FirstName:     firstName,
			LastName:      lastName,
			CreatedCVData: createdCVBytes,
		},
	}
	res, err := s.SendEmail(ctx, request)
	if err != nil || !res.Result {
		return error2.ServiceLayerError.
			WrapWithUserMessage(err, "Failed to send email to user").
			With("email", to)
	}

	return nil
}
