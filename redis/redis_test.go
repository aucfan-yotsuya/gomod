package redis

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func T_New(t *testing.T) {
	r = New()
	assert.NotNil(t, r)
}
func T_NewTarget(t *testing.T) {
	var target = r.NewTarget()
	assert.True(t, assert.ObjectsAreEqual(r.GetTarget(0), target))
}
func T_Close(t *testing.T) {
	r.Close()
	assert.Len(t, r.Target, 0)
}
func T_NewConn(t *testing.T) {
	defer r.GetTarget(0).Close()
	r.GetTarget(0).NewConn(&RedisConnOpt{
		Protocol:   "tcp",
		Address:    os.Getenv("TargetIP"),
		RetryCount: 3,
		Timeout:    10 * time.Second,
	})
	assert.Nil(t, err)
}
func T_NewPool(t *testing.T) {
	defer r.Close()
	r.GetTarget(0).NewPool(&RedisConnOpt{
		Protocol:      "tcp",
		Address:       os.Getenv("TargetIP"),
		RetryCount:    3,
		PoolMaxIdle:   3,
		PoolMaxActive: 3,
		Timeout:       10 * time.Second,
	})
	assert.False(t, assert.Nil(t, r.GetTarget(0).Pool))
}
func T_Ping(t *testing.T) {
	assert.True(t, r.GetTarget(0).GetConn().Ping())
}
func T_HSet(t *testing.T) {
	assert.NoError(t, r.GetTarget(0).HSet(
		"___hsetkey", map[string][]byte{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
			"key3": []byte("value3"),
		}))
}
func T_HSetString(t *testing.T) {
	assert.NoError(t, r.GetTarget(0).HSetString(
		"___hsetstringkey", map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}))
}
func T_Expire(t *testing.T) {
	assert.NoError(t, r.GetTarget(0).Expire(2, "___hsetkey", "___hsetstringkey"))
}
func T_HGetAll(t *testing.T) {
	var kv map[string][]byte
	kv, err = r.GetTarget(0).HGetAll("___hsetkey")
	assert.NoError(t, err)
	assert.EqualValues(t, kv["key1"], []byte("value1"))
	kv, err = r.GetTarget(0).HGetAll("___hsetstringkey")
	assert.NoError(t, err)
	assert.EqualValues(t, kv["key1"], []byte("value1"))
}
func T_HIncrBy(t *testing.T) {
	var resp int
	resp, err = r.GetTarget(0).HIncrBy("___hincby", "key1", 1)
	assert.NoError(t, err)
	assert.EqualValues(t, resp, 1)
	resp, err = r.GetTarget(0).HIncrBy("___hincby", "key1", 2)
	assert.NoError(t, err)
	assert.EqualValues(t, resp, 3)
	assert.NoError(t, r.GetTarget(0).Expire(2, "___hincrby"))
}
func TestKeys(t *testing.T) {
	//var keys []string
	//keys, err = r.Keys("auth_sess_111*")
	//fmt.Println(keys)
	//assert.Nil(t, err)
	//assert.True(t, assert.ObjectsAreEqual(keys, []string{}))
}
