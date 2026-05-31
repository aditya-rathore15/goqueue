package main

import (
	"log"
	"strconv"

	"github.com/aditya-rathore15/goqueue/internal/queue"
	"github.com/aditya-rathore15/goqueue/internal/worker"
)

func main() {
	log.Println("GoQueue broker started")

	q := queue.NewQueue()

	for i := 1; i <= 3; i++ {
		w := worker.NewWorker(i, q)
		go w.Start()
	}

	for i := 1; i <= 5; i++ {
		q.Enqueue(queue.Task{
			ID:      strconv.Itoa(i),
			Payload: "Task Payload " + strconv.Itoa(i),
			Status:  queue.StatusPending,
		})
	}

	select {}
}