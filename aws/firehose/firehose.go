package firehose

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
)

// PutRecord レコード送出
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
			Data: append(*data, []byte("\n")...),
		},
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}

// PutRecordBatch 複数レコード送出
func PutRecordBatch(deliveryStreamName *string, data *[][]byte, cfg ...aws.Config) error {
	var f *firehose.Client
	if len(cfg) == 0 {
		f = firehose.New(firehose.Options{Region: "ap-northeast-1"})
	} else {
		f = firehose.NewFromConfig(cfg[0])
	}
	var records []types.Record
	for _, v := range *data {
		records = append(records,
			types.Record{
				Data: append(v, []byte("\n")...),
			})
	}
	_, err := f.PutRecordBatch(context.TODO(), &firehose.PutRecordBatchInput{
		DeliveryStreamName: deliveryStreamName,
		Records:            records,
	})
	if err != nil {
		return errors.New("firehose putRecord: " + err.Error())
	}
	return nil
}
