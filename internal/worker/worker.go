package worker

import (
	"log"
	"time"

	"github.com/aditya-rathore15/goqueue/internal/queue"
)

type Worker struct {
	id int
	queue *queue.Queue
}

func NewWorker(id int, q *queue.Queue) *Worker {
	return &Worker{
		id:    id,
		queue: q,
	}
}

func (w *Worker) Start() {
	for {
		task, ok := w.queue.Dequeue()

		if ok {
			log.Printf("Worker %d processing task: %s\n", w.id, task.ID)
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}