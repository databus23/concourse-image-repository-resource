package resource

import (
	"strings"

	"github.com/hashicorp/go-version"
)

type Source struct {
	Repository string `json:"repository"`
	Regex      string `json:"regex"`
}

type Version struct {
	Tag     string           `json:"tag,omitempty"`
	Digest  string           `json:"digest,omitempty"`
	Version *version.Version `json:"-"`
}

func (v *Version) Parse() (err error) {
	v.Version, err = version.NewVersion(strings.TrimLeft(v.Tag, "v"))
	return
}

func (v *Version) GreaterThan(other *Version) bool {
	return v.Version.GreaterThan(other.Version)
}

func NewVersion(raw string) (Version, error) {
	v := Version{Tag: raw}
	err := v.Parse()
	return v, err
}

type Versions []Version

func (e Versions) Len() int {
	return len(e)
}

func (e Versions) Less(i int, j int) bool {
	return e[i].Version.LessThan(e[j].Version)
}

func (e Versions) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}
type InResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}
type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
