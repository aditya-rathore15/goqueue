package main

import (
	"log"

	"github.com/aditya-rathore15/goqueue/internal/broker"
	"github.com/aditya-rathore15/goqueue/internal/persistence"
	"github.com/aditya-rathore15/goqueue/internal/queue"
	"github.com/aditya-rathore15/goqueue/internal/worker"
)

func main() {
	log.Println("GoQueue broker started")

	q := queue.NewQueue()

	store, err := persistence.NewStore("goqueue.db")
	if err != nil {
		log.Fatal(err)
	}

	tasks, err := store.LoadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, task := range tasks {
		q.Enqueue(task)
	}

	for i := 1; i <= 3; i++ {
		w := worker.NewWorker(i, q, store)
		go w.Start()
	}

	server := broker.NewServer(q, store)
	go server.Start("8080")

	select {}
}