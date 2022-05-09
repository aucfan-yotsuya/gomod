package semaphore

import (
	"context"
	"sync"
	"time"

	"github.com/aucfan-yotsuya/gomod/common"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type (
	Semaphore struct {
		Weighted *semaphore.Weighted
		ErrGroup *errgroup.Group
		Context  struct {
			Ctx    context.Context
			Cancel context.CancelFunc
		}
		Mutex sync.Mutex
		Err   error
	}
)

var (
	s   *Semaphore
	err error
)

func New(semaphoreWeighted int64, contextTimeout time.Duration) *Semaphore {
	s = new(Semaphore)
	s.Context.Ctx, s.Context.Cancel = common.Context(contextTimeout)
	s.NewErrGroup()
	s.NewWeighted(semaphoreWeighted)
	return s
}
func (s *Semaphore) NewWeighted(i int64) {
	s.Weighted = semaphore.NewWeighted(i)
}
func (s *Semaphore) NilWeighted() bool {
	return s.Weighted == nil
}
func (s *Semaphore) NewErrGroup() {
	s.ErrGroup = new(errgroup.Group)
}
func (s *Semaphore) NilErrGroup() bool {
	return s.ErrGroup == nil
}
func (s *Semaphore) Acquire(n int64) error {
	if s.NilWeighted() {
		s.Err = &ErrWeighted{Message: "Werighted has nil"}
		return s.Err
	}
	return s.Weighted.Acquire(s.Context.Ctx, n)
}
func (s *Semaphore) Release(n int64) {
	if !s.NilWeighted() {
		s.Weighted.Release(n)
	}
}
func (s *Semaphore) Go(f func() error) error {
	defer s.Release(1)
	if err = s.Acquire(1); err != nil {
		return err
	}
	if s.NilErrGroup() {
		s.Err = &ErrErrGroup{Message: "ErrGroup has nil"}
		return s.Err
	}
	s.ErrGroup.Go(f)
	return nil
}
func (s *Semaphore) Wait() error {
	if s.NilErrGroup() {
		s.Err = &ErrErrGroup{Message: "ErrGroup has nil"}
		return s.Err
	}
	return s.ErrGroup.Wait()
}
