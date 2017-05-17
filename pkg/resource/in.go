package resource

import (
	"io/ioutil"
	"path"
)

func In(request InRequest, destinationDir string) (*InResponse, error) {

	ioutil.WriteFile(path.Join(destinationDir, "tag"), []byte(request.Version.Tag), 0644)
	ioutil.WriteFile(path.Join(destinationDir, "repository"), []byte(request.Source.Repository), 0644)
	if request.Version.Digest != "" {
		ioutil.WriteFile(path.Join(destinationDir, "digest"), []byte(request.Version.Digest), 0644)
	}

	return &InResponse{
		Version: request.Version,
	}, nil
}
