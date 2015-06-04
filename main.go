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

type Request struct {
	PULL_REQUEST_AUTHOR_NAME         string
	PULL_REQUEST_TO_REPO_NAME        string
	PULL_REQUEST_TO_REPO_PROJECT_KEY string
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "https://github.com/jhaals/stash-hook-router")
		}

		if r.Method == "POST" {
			request, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("failed to read response body")
			} else {
				log.Println(string(request))

				var r Request
				json.Unmarshal(request, &r)
				scriptName := sanitize(fmt.Sprintf("%s_%s",
					r.PULL_REQUEST_TO_REPO_PROJECT_KEY,
					r.PULL_REQUEST_TO_REPO_NAME))

				// Send JSON to script
				cmd := exec.Command(path.Join(scriptDir, scriptName), string(request))
				err := cmd.Start()

				if err != nil {
					log.Println(err)
				} else {
					fmt.Fprintln(w, "OK")
				}
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
