package main

import (
	"encoding/binary"
	"encoding/json"

	"github.com/boltdb/bolt"
)

// Task struct contains data about tasks
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Store is datastore interface around BoltDB
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

// CreateTask persists a task
func (s *Store) CreateTask(t *Task) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		id, _ := b.NextSequence()
		t.ID = int(id)
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(itob(t.ID), buf)
	})
}

// Close closes the database
func (s *Store) Close() error {
	return s.db.Close()
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
