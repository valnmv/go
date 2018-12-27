package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var store *Store

// Server is wrapper structure for httprouter
type Server struct {
	r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s.r.ServeHTTP(w, r)
}

// ListTasks serves GET requests
func ListTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks, err := store.GetTasks()
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(tasks)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// CreateTask serves POST requests
func CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := Task{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}

	err = store.CreateTask(&t)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// ReadTask serves GET/:id requests
func ReadTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "ReadTask\n")
}

// UpdateTask serves PUT/:id requests
func UpdateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "UpdateTask\n")
}

// DeleteTask serves DELETE/:id requests
func DeleteTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "DeleteTask\n")
}

// SetStore is used to provide a store initialized from outside.
// The provider is responsible for closing the store.Used for testing.
func SetStore(s *Store) {
	store = s
}

func main() {
	router := httprouter.New()
	router.GET("/", ListTasks)
	router.POST("/", CreateTask)
	router.GET("/:id", ReadTask)
	router.PUT("/:id", UpdateTask)
	router.DELETE("/:id", DeleteTask)

	var err error
	store, err = NewStore()
	if err != nil {
		log.Fatal(err)
	}

	store.Initialize()
	defer store.Close()

	log.Fatal(http.ListenAndServe(":8080", &Server{router}))
}
