package s3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	t_s *S3
)

func T_New(t *testing.T) {
	t_s = New()
	assert.NotNil(t, t_s)
}
func T_NewSession(t *testing.T) {
	t_s.NewSession()
}
func T_GetObjectInput(t *testing.T) {
}
