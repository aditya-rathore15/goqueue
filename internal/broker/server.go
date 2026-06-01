package broker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aditya-rathore15/goqueue/internal/queue"
)

type Server struct {
	queue *queue.Queue
}

type taskRequest struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}

func NewServer(q *queue.Queue) *Server {
	return &Server{
		queue: q,
	}
}

func (s *Server) Start(port string) {
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req taskRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		task := queue.Task{
			ID:      req.ID,
			Payload: req.Payload,
			Status:  queue.StatusPending,
		}

		s.queue.Enqueue(task)

		log.Printf("Task %s enqueued\n", task.ID)

		w.WriteHeader(http.StatusCreated)
	})

	log.Printf("Broker server running on port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}