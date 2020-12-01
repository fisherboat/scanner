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
	workers chan Task
}

func New(size int, fn func(task Task)) *Scanner {
	scanner := Scanner{
		wg:      sync.WaitGroup{},
		runTask: fn,
		workers: make(chan Task, size),
	}

	for i := 0; i < size; i++ {
		go func() {
			for task := range scanner.workers {
				scanner.runTask(task)
				scanner.wg.Done()
			}
		}()
	}

	return &scanner
}

func (s *Scanner) PushTask(task Task) {
	s.wg.Add(1)
	go func(task Task) {
		s.workers <- task
	}(task)
}

func (s *Scanner) Close() {
	s.wg.Wait()
}
