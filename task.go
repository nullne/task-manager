package task

import "time"

type TaskFn func() error

type Task struct{}

func (t *Task) Done() bool {
	return false
}

func (t *Task) Wait() {
}

type Manager struct{}

func New() *Manager {
	return &Manager{}
}

func (m *Manager) AddTask(fn TaskFn, timeout time.Duration, maxRetry int) *Task {
	return &Task{}
}

type temporary interface {
	Temporary() bool
}

func isTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}
