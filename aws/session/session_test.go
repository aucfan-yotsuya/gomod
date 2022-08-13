package session

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig("accesskeyid", "secretkey")
	assert.Equal(t, *c.Region, "ap-northeast-1")
}
func TestNew(t *testing.T) {
	if os.Getenv("accessskeyid") == "" ||
		os.Getenv("secretkey") == "" ||
		os.Getenv("rolearn") == "" {
		return
	}
	s := New(os.Getenv("accesskeyid"), os.Getenv("secretkey"), os.Getenv("rolearn"))
	v, err := s.Config.Credentials.Get()
	assert.Equal(t, v.AccessKeyID, os.Getenv("accesskeyid"))
	assert.NoError(t, err)
}
