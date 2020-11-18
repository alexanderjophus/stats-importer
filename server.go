package main

import (
	"io"
	"net/http"
	"time"
)

// A Server is a server
type Server struct {
	store  FileStorer
	router *http.ServeMux
	c      *http.Client
}

// FileStorer represents the ability to store files
type FileStorer interface {
	storeFile(r io.Reader, path string) (err error)
}

// NewServer creates a new server struct, initialised with the routing set
func NewServer(pr pachdRepo) Server {
	s := Server{
		store:  pr,
		router: http.NewServeMux(),
		c: &http.Client{
			Timeout: time.Second * 10,
		},
	}
	s.router.HandleFunc("/", s.addFile)
	return s
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
