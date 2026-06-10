package worker

import (
	"log"
	"time"

	"github.com/aditya-rathore15/goqueue/internal/persistence"
	"github.com/aditya-rathore15/goqueue/internal/queue"
)

type Worker struct {
	id    int
	queue *queue.Queue
	store *persistence.Store
}

func NewWorker(id int, q *queue.Queue, store *persistence.Store) *Worker {
	return &Worker{
		id:    id,
		queue: q,
		store: store,
	}
}

func (w *Worker) Start() {
	for {
		task, ok := w.queue.Dequeue()

		if ok {
			log.Printf("Worker %d processing task: %s\n", w.id, task.ID)

			err := w.store.Delete(task.ID)
			if err != nil {
				log.Printf("Failed to delete task %s: %v\n", task.ID, err)
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}