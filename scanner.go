package scanner

import (
	"errors"
	"fmt"
	"sync"

	"github.com/panjf2000/ants"
)

type Task interface {
	Action()
}

type Scanner struct {
	wg sync.WaitGroup
	p  *ants.PoolWithFunc
}

func New(size int, fn func(task interface{})) (*Scanner, error) {
	if size <= 0 {
		return nil, errors.New("Size value to small")
	}

	scanner := Scanner{
		wg: sync.WaitGroup{},
	}
	p, _ := ants.NewPoolWithFunc(size, func(task interface{}) {
		fn(task)
		scanner.wg.Done()
	})
	scanner.p = p
	return &scanner, nil
}

func (s *Scanner) PushTask(task interface{}) {
	s.wg.Add(1)
	_ = s.p.Invoke(task)
}

func (s *Scanner) Close() {
	s.p.Release()
}

func (s *Scanner) Run() {
	fmt.Printf("running goroutines: %d\n", s.p.Running())
}

func (s *Scanner) Wait() {
	s.wg.Wait()
}
