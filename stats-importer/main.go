package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pachyderm/pachyderm/src/client"
)

var (
	baseString = "https://statsapi.web.nhl.com"
	repoName   = "statsapi"
	netClient  = &http.Client{
		Timeout: time.Second * 10,
	}

	pC pachdClient
)

type pachdClient struct {
	Client *client.APIClient
}

func main() {
	// connect to pachyderm
	pachydermAddress := "192.168.99.100:30650"

	client, err := client.NewFromAddress(pachydermAddress)
	if err != nil {
		log.Fatalf("unable to get client: %v", err)
	}

	pC = pachdClient{
		Client: client,
	}

	r := mux.NewRouter()
	r.HandleFunc("/", addLink).Methods("POST")

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Printf("Server running at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
