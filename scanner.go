package scanner

import (
	"errors"
	"fmt"
)

type Task interface {
	Action()
}

type Scanner struct {
	tasks   chan *Task
	runTask func(task *Task)
}

func New(size int, fn func(task *Task)) (*Scanner, error) {
	if size <= 0 {
		return nil, errors.New("Size value to small")
	}
	scanner := Scanner{
		tasks:   make(chan *Task, size),
		runTask: fn,
	}
	return &scanner, nil
}

func (s *Scanner) PushTask(task Task) {
	go func(task Task) {
		s.tasks <- &task
	}(task)
}

func (s *Scanner) Running() {
	fmt.Println("Running")
	for task := range s.tasks {
		go s.runTask(task)
	}
}

func (s *Scanner) Close() {
	close(s.tasks)
}
