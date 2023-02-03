package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const region = "us-east-1"

type (
	Ses struct {
		From, Subject, Body string
		To                  []string
	}
)

func (s *Ses) SendmailWithConfig(cfg aws.Config) (*sesv2.SendEmailOutput, error) {
	return s.exec(sesv2.NewFromConfig(cfg))
}
func (s *Ses) Sendmail(opt sesv2.Options) (*sesv2.SendEmailOutput, error) {
	return s.exec(sesv2.New(opt))
}
func (s *Ses) exec(client *sesv2.Client) (*sesv2.SendEmailOutput, error) {
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &s.From,
		Destination: &types.Destination{
			ToAddresses: s.To,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &s.Body,
					},
				},
				Subject: &types.Content{
					Data: &s.Subject,
				},
			},
		},
	}
	res, err := client.SendEmail(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
