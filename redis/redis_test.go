package redis

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	Master int = iota
	Reader
)

func newConn() (*Conn, error) {
return r.NewTarget().NewConn(&RedisConnOpt{
		Protocol:   "tcp",
		Address:    os.Getenv("Target"),
		RetryCount: 3,
		Timeout:    10 * time.Second,
	})
}
func newPool() *Conn {
	return r.NewTarget().NewPool(&RedisConnOpt{
		Protocol:      "tcp",
		Address:       os.Getenv("Target"),
		RetryCount:    3,
		PoolMaxIdle:   3,
		PoolMaxActive: 3,
		Timeout:       10 * time.Second,
	})
}

func T_LaodEnv(t *testing.T) {
	assert.NoError(t, godotenv.Load(".env"))
}
func T_New(t *testing.T) {
	r = New()
	assert.NotNil(t, r)
}
func T_Close(t *testing.T) {
	assert.Len(t, r.Target, 1)
	r.Close()
	assert.Len(t, r.Target, 0)
}
func T_NewTarget(t *testing.T) {
	defer r.Close()
	var target = r.NewTarget()
	assert.True(t, assert.ObjectsAreEqual(r.GetTarget(Master), target))
}
func T_NewConn(t *testing.T) {
	defer r.Close()
	newConn(),
	assert.Nil(t, err)
}
func T_NewPool(t *testing.T) {
	defer r.Close()
	newPool()
	assert.False(t, assert.Nil(t, r.GetTarget(Master).Pool))
}
func T_Ping(t *testing.T) {
	defer r.Close()
	newConn()
	assert.True(t, r.GetTarget(Master).Ping())
}
func T_HSet(t *testing.T) {
	defer r.Close()
	newConn()
	assert.NoError(t, r.GetTarget(Master).HSet(
		"___hsetkey", map[string][]byte{
			"key1": []byte("value1"),
			"key2": []byte("value2"),
			"key3": []byte("value3"),
		}))
}
func T_HSetString(t *testing.T) {
	defer r.Close()
	newConn()
	assert.NoError(t, r.GetTarget(Master).HSetString(
		"___hsetstringkey", map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}))
}
func T_Expire(t *testing.T) {
	defer r.Close()
	newConn()
	assert.NoError(t, r.GetTarget(Master).SetExpire("___hsetkey", 2))
	assert.NoError(t, r.GetTarget(Master).SetExpire("__hsetstringkey", 2))
}
func T_HGetAll(t *testing.T) {
	defer r.Close()
	newConn()
	var kv map[string][]byte
	kv, err = r.GetTarget(Master).HGetAll("___hsetkey")
	assert.NoError(t, err)
	assert.EqualValues(t, kv["key1"], []byte("value1"))
	kv, err = r.GetTarget(Master).HGetAll("___hsetstringkey")
	assert.NoError(t, err)
	assert.EqualValues(t, kv["key1"], []byte("value1"))
}
func T_HIncrBy(t *testing.T) {
	defer r.Close()
	newConn()
	var resp int
	resp, err = r.GetTarget(Master).HIncrBy("___hincby", "key1", 10)
	assert.NoError(t, err)
	assert.EqualValues(t, resp, 1)
	resp, err = r.GetTarget(Master).HIncrBy("___hincby", "key1", 20)
	assert.NoError(t, err)
	assert.EqualValues(t, resp, 3)
	assert.NoError(t, r.GetTarget(Master).SetExpire("___hincrby", 20))
}
func T_Subscribe(t *testing.T) {
	defer r.Close()
	newConn()
	r.GetTarget(Master).Subscribe("hoge")
}
func T_SetSelialize(t *testing.T) {
	defer r.Close()
	assert.NoError(t, newConn())
	type (
		Vector struct {
			Num int
			Val string
		}
	)
	var vector = Vector{
		Num: 1,
		Val: "hoge",
	}
	assert.NoError(t, r.NewEncoder().Encode(vector))
	assert.NoError(t, r.GetTarget(Master).Set("___selialize", r.ReadFromBuffer()))
}
func TestKeys(t *testing.T) {
	//var keys []string
	//keys, err = r.Keys("auth_sess_111*")
	//fmt.Println(keys)
	//assert.Nil(t, err)
	//assert.True(t, assert.ObjectsAreEqual(keys, []string{}))
}
