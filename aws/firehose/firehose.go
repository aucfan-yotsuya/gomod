package firehose

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

func PutRecord(deliveryStreamName *string, data *[]byte) error {
	f := firehose.New(firehose.Options{})
	_, err := f.PutRecord(context.TODO(), &firehose.PutRecordInput{
		DeliveryStreamName: deliveryStreamName,
		Record: &types.Record{
			Data: *data,
		},
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}
func PutRecordWithConfig(cfg aws.Config, deliveryStreamName *string, data *[]byte) error {
	f := firehose.NewFromConfig(cfg)
	_, err := f.PutRecord(context.TODO(), &firehose.PutRecordInput{
		DeliveryStreamName: deliveryStreamName,
		Record: &types.Record{
			Data: append(*data, []byte("\n")...),
		},
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}
