package resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func Check(request CheckRequest) ([]Version, error) {

	parts := strings.SplitN(request.Source.Repository, "/", 2)
	host, repository := parts[0], parts[1]

	tags := []string{}

	switch host {
	case "quay.io":
		resp, err := http.Get(fmt.Sprintf("https://quay.io/api/v1/repository/%s/tag/?limit=100", repository))
		if err != nil {
			Fail(err)
		}
		defer resp.Body.Close()
		var data struct {
			Tags []struct {
				Name string
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			Fatal("Failed to decode response from quay.io", err)
		}
		for _, tag := range data.Tags {
			tags = append(tags, tag.Name)
		}
	default:
		resp, err := http.Get(fmt.Sprintf("https://%s/v2/%s/tags/list", host, repository))
		if err != nil {
			Fail(err)
		}
		defer resp.Body.Close()
		var data struct {
			Tags []string
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			Fatal("Failed to decode response from quay.io", err)
		}
		tags = data.Tags

	}
	if request.Source.Regex != "" {
		VersionRegex = regexp.MustCompile(request.Source.Regex)
	}
	versionGiven := request.Version.Tag != ""
	if versionGiven {
		if err := request.Version.Parse(); err != nil {
			Fatal("Invalid version given", err)
		}
	}
	versions := Versions{}
	for _, raw := range tags {
		v, err := NewVersion(raw)
		if err != nil {
			Sayf("Skipping tag %s: %s\n", raw, err)
			continue
		}

		if !versionGiven {
			versions = append(versions, v)
		} else {
			if v.GreaterThan(&request.Version) {
				versions = append(versions, v)
			}
		}
	}
	sort.Sort(versions)

	if versionGiven || len(versions) == 0 {
		return versions, nil
	} else {
		return versions[:1], nil
	}

}
