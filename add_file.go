package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pachyderm/pachyderm/src/client"
)

// addFile handles /addFile endpoint
func (s *Server) addFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := s.postFile(w, r); err != nil {
			// remove once middleware is in place
			log.Println(err)
		}
	}
}

func (s *Server) postFile(w http.ResponseWriter, r *http.Request) error {
	type requestBody struct {
		Link string `json:"link"`
	}
	var t requestBody
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	err = s.linkReader(t.Link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func (s *Server) linkReader(link string) error {
	r, err := s.c.Get(baseString + link)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", r.Status)
	}

	err = s.store.storeFile(r.Body, link)
	if err != nil {
		return err
	}

	return nil
}

type pachdRepo struct {
	Client *client.APIClient
}

func (c pachdRepo) storeFile(r io.Reader, path string) (err error) {
	// Start a commit in our "livefeed" data repo on the "master" branch.
	commit, err := c.Client.StartCommit(repoName, "master")
	if err != nil {
		return err
	}
	defer func() {
		err = c.Client.FinishCommit(repoName, commit.ID)
	}()

	path = strings.Replace(path, "?", "/", -1)
	path = strings.Replace(path, "=", "_", -1)
	// Put a file containing the respective project name.
	if _, err := c.Client.PutFileOverwrite(repoName, commit.ID, path+".raw", r, 0); err != nil {
		return err
	}

	return nil
}
