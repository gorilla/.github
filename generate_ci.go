package main

import (
	"log"
	"os"
	"text/template"
)

var (
	versions  = []string{"1.12", "1.11", "1.10", "1.9", "1.8", "1.7"}
  # also: curl -s https://api.github.com/orgs/gorilla/repos | jq '.[] | select(.name | contains("github") | not) | .name'
	libraries = []string{"handlers", "css", "feeds", "schema", "rpc", "securecookie", "sessions", "mux"}
)

func main() {
  // load configTemplate from a file
	tmpl, err := template.New("config").Parse(string(configTemplate))
	if err != nil {
		log.Fatal(err)
	}

	for _, lib := range libraries {
		if err := tmpl.Execute(os.Stdout, map[string]interface{}{
			"RepoUser": "gorilla",
			"RepoName": lib,
			"Versions": versions,
		}); err != nil {
			log.Fatal(err)
		}
	}
}
