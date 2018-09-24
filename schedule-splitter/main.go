package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pachyderm/pachyderm/src/client"
)

const repoName = "livefeed"

func main() {
	// connect to pachyderm
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}

	// // Create a repo called "livefeed."
	// if err := c.CreateRepo(repoName); err != nil {
	// 	log.Fatal(err)
	// }

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(path + "/data")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		err := addFileToRepo(c, path+"/data/"+f.Name())
		if err != nil {
			log.Printf("Couldn't add file: %v", err)
		}
	}
}

func addFileToRepo(c *client.APIClient, filePath string) error {
	// Start a commit in our "livefeed" data repo on the "master" branch.
	commit, err := c.StartCommit(repoName, "master")
	if err != nil {
		return err
	}
	defer c.FinishCommit(repoName, commit.ID)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Put a file containing the respective project name.
	if _, err := c.PutFileOverwrite(repoName, commit.ID, filepath.Base(filePath), file, 0); err != nil {
		return err
	}

	return nil
}
