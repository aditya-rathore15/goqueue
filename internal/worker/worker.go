package worker

import (
	"errors"
	"log"
	"math/rand"
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

		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("Worker %d processing task: %s\n", w.id, task.ID)

		err := process(task)

		// Success
		if err == nil {
			log.Printf("Task %s completed successfully\n", task.ID)

			if err := w.store.Delete(task.ID); err != nil {
				log.Printf("Failed to delete task %s: %v\n", task.ID, err)
			}

			continue
		}

		// Failure with retries remaining
		if task.Retries < task.MaxRetries {
			task.Retries++

			log.Printf(
				"Task %s failed, retry %d/%d\n",
				task.ID,
				task.Retries,
				task.MaxRetries,
			)

			if err := w.store.Save(task); err != nil {
				log.Printf("Failed to save task %s: %v\n", task.ID, err)
				continue
			}

			w.queue.Enqueue(task)
			continue
		}

		// Failure, max retries exceeded
		log.Printf(
			"Task %s failed permanently after %d retries\n",
			task.ID,
			task.MaxRetries,
		)

		if err := w.store.Delete(task.ID); err != nil {
			log.Printf("Failed to delete task %s: %v\n", task.ID, err)
		}
	}
}

func process(task queue.Task) error {
	// 30% chance of failure
	if rand.Intn(10) < 3 {
		return errors.New("simulated task failure")
	}

	return nil
}