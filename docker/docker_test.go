package docker

import (
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var (
	d *Docker
)

func TestMain(m *testing.M) {
	d = New()
	defer d.Close()
	m.Run()
}
func TestNew(t *testing.T) {
	assert.NotNil(t, d)
	assert.NotNil(t, d.NewTarget())
	assert.NotNil(t, d.NewTarget())
}
func TestNewPool(t *testing.T) {
	opt := &Opt{
		Endpoint: "",
		DockerRunOptions: &dockertest.RunOptions{
			Repository:   "mysql",
			Tag:          "8.0",
			Env:          []string{"MYSQL_ROOT_PASSWORD=secret"},
			ExposedPorts: []string{"3306"},
		},
	}
	t.Logf("***** NewPool: *****")
	assert.NoError(t, d.GetTarget(0).NewPool(opt))
	t.Log(d.GetTarget(0).Resource.GetHostPort("tcp/3306"))
	t.Logf("***** NewPool: *****")
	assert.NoError(t, d.GetTarget(1).NewPool(opt))
	t.Log(d.GetTarget(1).Resource.GetHostPort("tcp/3306"))
}
func TestRetry(t *testing.T) {
	var tg *Target
	t.Logf("***** Retry: *****")
	tg = d.GetTarget(0)
	tg.Pool.MaxWait = 1 * time.Minute
	assert.NoError(t, tg.Retry(tg.NewGorm))
	t.Logf("***** Retry: *****")
	tg = d.GetTarget(1)
	tg.Pool.MaxWait = 1 * time.Minute
	assert.NoError(t, tg.Retry(tg.NewGorm))
}
func TestDBConnection(t *testing.T) {
	data := []map[string]interface{}{}
	for i := 0; i < 100000; i++ {
		data = append(data, map[string]interface{}{
			"name":       "Name",
			"created_at": time.Now(),
		})
	}
	t.Logf("***** DBConnection *****")
	assert.NoError(
		t, d.GetTarget(0).Gorm.Exec("create database db").
			Exec(`create table db.tbl (
		id bigint not null auto_increment primary key,
		name text not null,
		created_at datetime
	) engine=innodb`).Error)
	assert.NoError(t, d.GetTarget(0).Gorm.Table("db.tbl").Create(data).Error)
	assert.NoError(
		t, d.GetTarget(1).Gorm.Exec("create database db").
			Exec(`create table db.tbl (
		id bigint not null auto_increment primary key,
		name text not null,
		created_at datetime
	) engine=innodb`).Error)
	assert.NoError(t, d.GetTarget(1).Gorm.
		Table("db.tbl").Create(data).Error)
	time.Sleep(1 * time.Minute)
}
