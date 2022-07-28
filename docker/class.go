package docker

import (
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (d *Docker) NewTarget() *Target {
	var tg = &Target{}
	d.Target = append(d.Target, tg)
	return tg
}
func (d *Docker) TargetLen() int { return len(d.Target) }
func (d *Docker) GetTarget(index int) *Target {
	if d.TargetLen() < index+1 {
		return new(Target)
	}
	return d.Target[index]
}
func (d *Docker) Close() {
	for i := 0; i < d.TargetLen(); i++ {
		d.GetTarget(i).Close()
		d.Target[i] = nil
	}
	d.Target = []*Target{}
}
func (tg *Target) Close() {
	tg.Pool.Purge(tg.Resource)
}
func (tg *Target) NewPool(opt *Opt) error {
	pool, err := dockertest.NewPool(opt.Endpoint)
	if err != nil {
		return err
	}
	tg.Resource, err = pool.RunWithOptions(opt.DockerRunOptions)
	if err != nil {
		tg.Pool.Purge(tg.Resource)
		return err
	}
	tg.Pool = pool
	return nil
}
func (tg *Target) NewGorm() error {
	var err error
	tg.Gorm, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"root:secret@tcp(%s)/mysql?%s",
		tg.Resource.GetHostPort("3306/tcp"),
		"timeout=30s&charset=utf8mb4&interpolateParams=true&parseTime=true&loc=Asia%2fTokyo",
	)), &gorm.Config{
		CreateBatchSize:        1000,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	db, err := tg.Gorm.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
func (tg *Target) Retry(op func() error) error {
RetryLoop:
	time.Sleep(3 * time.Second)
	if err := tg.Pool.Retry(op); err != nil {
		goto RetryLoop
	}
	return nil
}
