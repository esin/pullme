package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type hubDockerReplyStruct struct {
	Affiliation     interface{} `json:"affiliation"`
	CanEdit         bool        `json:"can_edit"`
	Description     string      `json:"description"`
	FullDescription interface{} `json:"full_description"`
	HasStarred      bool        `json:"has_starred"`
	IsAutomated     bool        `json:"is_automated"`
	IsMigrated      bool        `json:"is_migrated"`
	IsPrivate       bool        `json:"is_private"`
	LastUpdated     string      `json:"last_updated"`
	Name            string      `json:"name"`
	Namespace       string      `json:"namespace"`
	Permissions     struct {
		Admin bool `json:"admin"`
		Read  bool `json:"read"`
		Write bool `json:"write"`
	} `json:"permissions"`
	PullCount      int64  `json:"pull_count"`
	RepositoryType string `json:"repository_type"`
	StarCount      int64  `json:"star_count"`
	Status         int64  `json:"status"`
	User           string `json:"user"`
}

func main() {
	fmt.Println("Hey!")
	resp, err := http.Get("https://hub.docker.com/v2/repositories/es1n/pullme/")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	var hubDockerReply hubDockerReplyStruct
	err = json.Unmarshal(body, &hubDockerReply)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("I was pulled %d times\n", hubDockerReply.PullCount)

	done := make(chan bool)
	go forever()
	<-done
}

func forever() {
	for {
		time.Sleep(time.Second)
	}
}
