package firehose

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

func PutRecord(deliveryStreamName *string, data *[]byte, cfg ...aws.Config) error {
	var f *firehose.Client
	if len(cfg) == 0 {
		f = firehose.New(firehose.Options{})
	} else {
		f = firehose.NewFromConfig(cfg[0])
	}
	_, err := f.PutRecord(context.TODO(), &firehose.PutRecordInput{
		DeliveryStreamName: deliveryStreamName,
		Record: &types.Record{
			Data: append(*data, []byte("\n")),
		},
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}
