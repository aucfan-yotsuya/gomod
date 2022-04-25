package redis

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ep = New()
	assert.NotNil(t, ep)
}
func TestEPNewRedis(t *testing.T) {
	r = ep.NewRedis()
	assert.True(t, assert.ObjectsAreEqual(ep.GetRedis(0), r))
}
func TestEPClose(t *testing.T) {
	ep.Close()
	assert.Len(t, ep.Redis, 0)
}
func TestRedisSetContext(t *testing.T) {
	var cnt = 0
	r.SetContext(3 * time.Second)
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-time.After(1 * time.Second):
			cnt++
		}
	}
	assert.InDelta(t, 1, cnt, 3.0)
}
func TestRedisNewConn(t *testing.T) {
	defer ep.GetRedis(0).Close()
	err = ep.GetRedis(0).NewConn(
		"172.31.15.252", 6379, 3*time.Second,
	)
	assert.Nil(t, err)
}
func TestRedisNewPool(t *testing.T) {
	defer ep.Close()
	ep.GetRedis(0).NewPool(
		"172.31.15.252", 6379, 5, 5, 3*time.Second,
	)
	assert.Nil(t, err)
}
func TestPing(t *testing.T) {
	err = r.NewConn(
		"172.31.15.252", 6379, 3*time.Second,
	)
	assert.Nil(t, err)
	assert.True(t, r.Ping())
}
func TestHSet(t *testing.T) {
	r.HSet("___hsetkey", map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
		"key3": []byte("value3"),
	})
}
func TestExpire(t *testing.T) {
	assert.Nil(t, r.Expire(10, "___hsetkey"))
}
func TestHGetAll(t *testing.T) {
	var (
		kv map[string][]byte
	)
	kv, err = r.HGetAll("___hsetkey")
	assert.Nil(t, err)
	assert.EqualValues(t, kv["key1"], []byte("value1"))
}
func TestKeys(t *testing.T) {
	//var keys []string
	//keys, err = r.Keys("auth_sess_111*")
	//fmt.Println(keys)
	//assert.Nil(t, err)
	//assert.True(t, assert.ObjectsAreEqual(keys, []string{}))
}
