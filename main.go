package main

import (
	"log"
	"net/http"
)

type server struct {
}

func (s *server) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	server := &server{}
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
