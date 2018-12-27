package main

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

// Task struct contains data about tasks
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Store is datastore around BoltDB
type Store struct {
	db *bolt.DB
}

// NewStore sets up BoltDB
func NewStore() (*Store, error) {
	handle, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &Store{db: handle}, nil
}

// Initialize sets up all the buckets
func (s *Store) Initialize() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}

		return nil
	})
}

// GetTasks fetches all tasks from the store
func (s *Store) GetTasks() ([]*Task, error) {
	tasks := []*Task{}
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		b.ForEach(func(k, v []byte) error {
			var t Task
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}

			tasks = append(tasks, &t)
			return nil
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
