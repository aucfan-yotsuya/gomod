package firehose

import (
	"context"
	"errors"
	"rebill/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

func PutRecord(deliveryStreamName *string, c *config.Config, data *[]byte) error {
	f := firehose.New(firehose.Options{})
	_, err := f.PutRecord(context.TODO(), &firehose.PutRecordInput{
		DeliveryStreamName: &c.Toml.Firehose.DeliveryStreamName,
		Record: &types.Record{
			Data: *data,
		},
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}
func PutRecordWithConfig(cfg aws.Config, c *config.Config, data *[]byte) error {
	f := firehose.NewFromConfig(cfg)
	_, err := f.PutRecord(context.TODO(), &firehose.PutRecordInput{
		DeliveryStreamName: &c.Toml.Firehose.DeliveryStreamName,
		Record: &types.Record{
			Data: *data,
		},
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}
