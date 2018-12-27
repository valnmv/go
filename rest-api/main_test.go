package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestListTasksHandler(t *testing.T) {
	store, _ := NewStore()
	store.Initialize()
	defer store.Close()
	SetStore(store)

	router := httprouter.New()
	router.GET("/", ListTasks)
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s := Server{router}
	s.ServeHTTP(w, r)

	expectedContentTypeHeader := "application/json; charset=utf-8"

	if w.Header().Get("Content-Type") != expectedContentTypeHeader {
		t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
			w.Header().Get("Content-Type"), expectedContentTypeHeader)
	}

	expectedCode := http.StatusOK
	if w.Code != expectedCode {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			w.Code, expectedCode)
	}
}

func TestCreateTask(t *testing.T) {
	store, _ := NewStore()
	store.Initialize()
	defer store.Close()
	SetStore(store)

	router := httprouter.New()
	router.POST("/", CreateTask)
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	s := Server{router}
	s.ServeHTTP(w, r)

	expectedContentTypeHeader := "application/json; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentTypeHeader {
		t.Errorf("handler returned unexpected Content-Type header: got %v want %v",
			w.Header().Get("Content-Type"), expectedContentTypeHeader)
	}

	expectedCode := http.StatusCreated
	if w.Code != expectedCode {
		t.Errorf("handler returned unexpected status code: got %v want %v", w.Code, expectedCode)
	}
}
