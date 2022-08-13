package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func New(AccessKeyID, SecretAccessKey, RoleArn string) *session.Session {
	s := session.Must(session.NewSession(NewConfig(AccessKeyID, SecretAccessKey)))
	return session.Must(session.NewSession(&aws.Config{
		Region:      s.Config.Region,
		Credentials: stscreds.NewCredentials(s, RoleArn),
	}))
}
func NewConfig(AccessKeyID, SecretAccessKey string) *aws.Config {
	return &aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
	}
}
