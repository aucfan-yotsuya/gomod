package firehose

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/aucfan-yotsuya/gomod/common"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

var deliveryStreamName = "yotsuya-gmo"

func TestMain(m *testing.M) {
	m.Run()
}
func TestPutRecord(t *testing.T) {
	b, err := json.Marshal(map[string]interface{}{
		"key": "バイナリデータ",
	})
	assert.NoError(t, err)
	err = PutRecord(&deliveryStreamName, &b)
	assert.NoError(t, err)
}
func TestPutRecordWithConfig(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("test"))
	assert.NoError(t, err)
	cfg.Region = "ap-northeast-1"
	tm := common.NowJST()
	ulidStr := ulid.MustNew(ulid.Now(), rand.Reader).String()
	{
		b, err := json.Marshal(map[string]interface{}{
			"ymd":      tm.Format("20060102"),
			"ulid":     ulidStr,
			"api":      "ExecTran",
			"mode":     "request",
			"shopID":   "xxx",
			"memberID": "xxx",
		})
		assert.NoError(t, err)
		err = PutRecord(&deliveryStreamName, &b, cfg)
		assert.NoError(t, err)
	}
	{
		b, err := json.Marshal(map[string]interface{}{
			"ymd":     tm.Format("20060102"),
			"ulid":    ulidStr,
			"api":     "ExecTran",
			"mode":    "response",
			"errCode": "EX1",
			"ErrInfo": "EX500302",
		})
		assert.NoError(t, err)
		err = PutRecord(&deliveryStreamName, &b, cfg)
	}
	{
		b, err := json.Marshal(map[string]interface{}{
			"ymd":     tm.Format("20060102"),
			"ulid":    ulidStr,
			"api":     "ExecTran",
			"mode":    "response",
			"errCode": "EX1",
			"ErrInfo": "EX500302",
		})
		assert.NoError(t, err)
		err = PutRecord(&deliveryStreamName, &b, cfg)
	}
	t.Log(tm.Format("20060102"), "time")
}
func TestPutRecordBatch(t *testing.T) {
	var bs []*[]byte
	b, err := json.Marshal(map[string]interface{}{
		"key": "バイナリデータ",
	})
	for i := 0; i < 500; i++ {
		bs = append(bs, &b)
	}
	assert.NoError(t, err)
	err = PutRecordBatch(&deliveryStreamName, bs)
	assert.NoError(t, err)
}
func TestPutRecordBatchWithConfig(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("test"))
	assert.NoError(t, err)
	cfg.Region = "ap-northeast-1"
	tm := common.NowJST()
	ulidStr := ulid.MustNew(ulid.Now(), rand.Reader).String()
	{
		b, err := json.Marshal(map[string]interface{}{
			"ymd":      tm.Format("20060102"),
			"ulid":     ulidStr,
			"api":      "ExecTran",
			"mode":     "request",
			"shopID":   "xxx",
			"memberID": "xxx",
		})
		assert.NoError(t, err)
		var bs []*[]byte
		for i := 0; i < 500; i++ {
			bs = append(bs, &b)
		}
		err = PutRecordBatch(&deliveryStreamName, bs, cfg)
		assert.NoError(t, err)
	}
	{
		b, err := json.Marshal(map[string]interface{}{
			"ymd":     tm.Format("20060102"),
			"ulid":    ulidStr,
			"api":     "ExecTran",
			"mode":    "response",
			"errCode": "EX1",
			"ErrInfo": "EX500302",
		})
		assert.NoError(t, err)
		var bs []*[]byte
		for i := 0; i < 500; i++ {
			bs = append(bs, &b)
		}
		err = PutRecordBatch(&deliveryStreamName, bs, cfg)
	}
	{
		b, err := json.Marshal(map[string]interface{}{
			"ymd":     tm.Format("20060102"),
			"ulid":    ulidStr,
			"api":     "ExecTran",
			"mode":    "response",
			"errCode": "EX1",
			"ErrInfo": "EX500302",
		})
		assert.NoError(t, err)
		var bs []*[]byte
		for i := 0; i < 500; i++ {
			bs = append(bs, &b)
		}
		err = PutRecordBatch(&deliveryStreamName, bs, cfg)
	}
	t.Log(tm.Format("20060102"), "time")
}
