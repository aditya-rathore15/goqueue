package main

import (
	"log"
	"github.com/aditya-rathore15/goqueue/internal/queue"
)

func main() {
	log.Println("GoQueue broker started")

	q := queue.NewQueue()

	q.Enqueue(queue.Task{
		ID: "1",
		Payload: "Task 1",
		Status: queue.StatusPending,
	})

	q.Enqueue(queue.Task{
		ID: "2",
		Payload: "Task 2",
		Status: queue.StatusPending,
	})

	task, ok := q.Dequeue()

	if ok {
		log.Println("Dequeued task: ", task)
	} else {
		log.Println("Queue is empty")
	}
}