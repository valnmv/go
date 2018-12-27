package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
	fmt.Fprint(w, "ListTasks\n")
}

// CreateTask serves POST requests
func CreateTask(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "CreateTasks\n")
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

func main() {
	router := httprouter.New()
	router.GET("/", ListTasks)
	router.POST("/", CreateTask)
	router.GET("/:id", ReadTask)
	router.PUT("/:id", UpdateTask)
	router.DELETE("/:id", DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", &Server{router}))

}
