package main

// Written by Johan Haals <jhaals@ooyala.com>

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

type Hook struct {
	AuthorName string `json:"PULL_REQUEST_AUTHOR_NAME"`
	RepoName   string `json:"PULL_REQUEST_TO_REPO_NAME"`
	ProjectKey string `json:"PULL_REQUEST_TO_REPO_PROJECT_KEY"`
}

// replace all non [^A-Za-z0-9_] and return lowercase result
func sanitize(input string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9_]+")
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToLower(reg.ReplaceAllString(input, ""))
}

func main() {

	scriptDir := os.Getenv("STASH_SCRIPT_DIR")
	if scriptDir == "" {
		log.Fatal("need STASH_SCRIPT_DIR in order to execute stuff")
	}

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			fmt.Fprintf(response, "https://github.com/jhaals/stash-hook-router")
			return
		}

		requestBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(response, "Invalid json", 500)
			return
		}
		log.Println(string(requestBody))

		var hook Hook
		json.Unmarshal(requestBody, &hook)
		scriptName := sanitize(fmt.Sprintf("%s_%s",
			hook.ProjectKey,
			hook.RepoName))

		// Send JSON to script
		cmd := exec.Command(path.Join(scriptDir, scriptName), string(requestBody))
		err = cmd.Start()
		if err != nil {
			log.Println(err)
			http.Error(response, "Failed to execute command", 500)
			return
		}
		fmt.Fprintln(response, "OK")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
