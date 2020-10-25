package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pachyderm/pachyderm/src/client"
)

const (
	port = "8080"

	baseString = "https://statsapi.web.nhl.com"
	repoName   = "statsapi"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	client, err := client.NewInCluster()
	if err != nil {
		return fmt.Errorf("unable to get client: %w", err)
	}

	if err := client.UpdateRepo(repoName); err != nil {
		return fmt.Errorf("cannot create repo %s: %w", repoName, err)
	}

	s := NewServer(pachdRepo{Client: client})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("running on address: %s", addr)
	if err := http.ListenAndServe(addr, s); err != nil {
		return err
	}

	return nil
}
