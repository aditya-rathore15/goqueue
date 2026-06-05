package persistence

import (
	"encoding/json"

	"go.etcd.io/bbolt"

	"github.com/aditya-rathore15/goqueue/internal/queue"
)

var taskBucket = []byte("tasks")

type Store struct {
	db *bbolt.DB
}

func NewStore(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Save(task queue.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Put([]byte(task.ID), data)
	})
}

func (s *Store) LoadAll() ([]queue.Task, error) {
	var tasks []queue.Task

	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)

		return b.ForEach(func(k, v []byte) error {
			var task queue.Task

			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}

			tasks = append(tasks, task)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Store) Delete(taskID string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete([]byte(taskID))
	})
}