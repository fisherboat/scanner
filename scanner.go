package scanner

import (
	"sync"
)

type Task interface {
	Action()
}

type Scanner struct {
	wg      sync.WaitGroup
	runTask func(task Task)
}

func New(fn func(task Task)) *Scanner {
	scanner := Scanner{
		wg:      sync.WaitGroup{},
		runTask: fn,
	}
	return &scanner
}

func (s *Scanner) PushTask(task Task) {
	s.wg.Add(1)
	go func(task Task) {
		defer s.wg.Done()
		s.runTask(task)
	}(task)
}

func (s *Scanner) Close() {
	s.wg.Wait()
}
