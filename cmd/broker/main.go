package main

import (
	"log"

	"github.com/aditya-rathore15/goqueue/internal/broker"
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

	server := broker.NewServer(q)
	go server.Start("8080")

	select {}
}