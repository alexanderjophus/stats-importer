package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type requestBody struct {
	Link string `json:"link"`
}

func addLink(w http.ResponseWriter, r *http.Request) {
	var t requestBody
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = linkReader(t.Link)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func linkReader(link string) error {
	r, err := netClient.Get(baseString + link)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", r.Status)
	}

	err = pC.addFileToRepo(r.Body, link)
	if err != nil {
		return err
	}

	return nil
}

func (c *pachdClient) addFileToRepo(r io.Reader, path string) (err error) {
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
