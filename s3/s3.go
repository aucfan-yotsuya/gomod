package s3

import (
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	Session        *awss3.S3
	GetObjectInput *awss3.GetObjectInput
}

func New() *S3 {
	var s = new(S3)
	s.GetObjectInput = new(awss3.GetObjectInput)
	s.Session = s.NewSession()
	return s
}
func (s *S3) NewSession() *awss3.S3 {
	return awss3.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("ap-northeast-1"),
		},
	})))
}
func (s *S3) CopyObjectInput(arg *map[string]*string) *awss3.CopyObjectInput {
	return &awss3.CopyObjectInput{
		Bucket:            (*arg)["Bucket"],
		Key:               (*arg)["Key"],
		ContentType:       (*arg)["ContentType"],
		CopySource:        aws.String(url.QueryEscape(*(*arg)["Bucket"] + *(*arg)["Key"])),
		MetadataDirective: (*arg)["MetadataDirective"],
	}
}
func (s *S3) CopyObject(arg *map[string]*string) (*awss3.CopyObjectOutput, error) {
	return s.Session.CopyObject(s.CopyObjectInput(arg))
}
func (s *S3) GetObject(getObjectInput awss3.GetObjectInput) (*awss3.GetObjectOutput, error) {
	return s.GetObject(getObjectInput)
}
