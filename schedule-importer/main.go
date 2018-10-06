package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pachyderm/pachyderm/src/client"
)

var (
	netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	baseString = "https://statsapi.web.nhl.com"

	scheduleString = "/api/v1/schedule"
	repoName       = "schedule"
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

	pC := pachdClient{
		Client: client,
	}

	pC.addFileToRepo(nil, "hello")
}

func (c *pachdClient) addFileToRepo(r io.Reader, path string) error {
	// Start a commit in our "livefeed" data repo on the "master" branch.
	commit, err := c.Client.StartCommit("repoName", "master")
	if err != nil {
		return err
	}
	defer c.Client.FinishCommit(repoName, commit.ID)

	// Put a file containing the respective project name.
	if _, err := c.Client.PutFileOverwrite(repoName, commit.ID, path, r, 0); err != nil {
		return err
	}

	return nil
}

func somethingFunc(url string) (io.Reader, error) {
	r, err := netClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", r.Status)
	}
	defer r.Body.Close()

	return r.Body, nil
}
