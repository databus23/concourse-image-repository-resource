package resource

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-version"
)

var VersionRegex = regexp.MustCompile("^v?(.*)$")

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
	match := VersionRegex.FindStringSubmatch(v.Tag)

	if match == nil {
		return fmt.Errorf("Tag %s filtered by regex", v.Tag)
	}
	switch len(match) {
	case 1:
		v.Version, err = version.NewVersion(match[0])
	case 2:
		v.Version, err = version.NewVersion(match[1])
	case 3:
		v.Version, err = version.NewVersion(fmt.Sprintf("%s-%s", match[1], match[2]))
	default:
		v.Version, err = version.NewVersion(fmt.Sprintf("%s-%s+%s", match[1], match[2], match[3]))
	}
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
