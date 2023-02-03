package ses

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/stretchr/testify/assert"
)

var s *Ses

func TestMain(m *testing.M) {
	s = &Ses{
		From:    "yotsuya@aucfan.com",
		To:      []string{"to1@example.jp", "to2@example.jp"},
		Subject: "件名",
		Body:    "ここに本文が入ります",
	}
	m.Run()
}
func TestSendmail(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	cfg.Region = region
	assert.NoError(t, err)
	res, err := s.Sendmail(sesv2.Options{
		Region:      "us-east-1",
		Credentials: cfg.Credentials,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, *res.MessageId)
	t.Log(*res)
}
func TestSendmailWithConfig(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	cfg.Region = region
	assert.NoError(t, err)
	res, err := s.SendmailWithConfig(cfg)
	assert.NoError(t, err)
	assert.NotEmpty(t, *res.MessageId)
}
