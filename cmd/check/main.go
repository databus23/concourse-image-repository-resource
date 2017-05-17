package main

import (
	"encoding/json"
	"os"

	"github.com/databus23/concourse-image-repository-resource/pkg/resource"
)

func main() {
	var request resource.CheckRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		resource.Fatal("reading request from stdin", err)
	}

	response, err := resource.Check(request)
	if err != nil {
		resource.Fail(err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		resource.Fatal("writing response to stdout", err)
	}
}
