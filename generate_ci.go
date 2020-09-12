package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"text/template"
)

var (
	templatePath = "config.yml.tmpl"
	versions     = []string{"latest", "1.15", "1.14", "1.13", "1.12", "1.10", "1.11"}
	// Run: curl -s https://api.github.com/orgs/gorilla/repos | jq '.[] | select(.name | contains("github") | not) | .name'
	libraries = []string{
		"handlers", "css", "feeds", "schema", "rpc", "securecookie", "sessions", "mux", "websocket", "csrf",
	}
)

func main() {
	flag.StringVar(&templatePath, "template", "config.yml.tmpl", "the path to the Go config template")
	flag.Parse()

	configTemplate, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.New("config").Parse(string(configTemplate))
	if err != nil {
		log.Fatal(err)
	}

	for _, lib := range libraries {
		buf := bytes.Buffer{}
		if err := tmpl.Execute(&buf, map[string]interface{}{
			"RepoUser": "gorilla",
			"RepoName": lib,
			"Versions": versions,
		}); err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("./%s-config.yml", lib)
		if err := ioutil.WriteFile(
			filename,
			buf.Bytes(),
			0644,
		); err != nil {
			log.Fatal(err)
		}

		log.Printf("wrote %s", filename)
	}
}
