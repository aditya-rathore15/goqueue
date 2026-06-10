package queue

import "sync"

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusDone Status = "done"
)

type Task struct {
	ID         string
	Payload    string
	Status     Status
	Retries    int
	MaxRetries int
}

type Queue struct {
	tasks []Task
	mu    sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		tasks: []Task{},
	}
}

func (q *Queue) Enqueue(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(q.tasks, task)
}

func (q *Queue) Dequeue() (Task, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return Task{}, false
	}

	first := q.tasks[0]
	q.tasks = q.tasks[1:]
	return first, true
}