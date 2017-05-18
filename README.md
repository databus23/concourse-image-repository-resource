Concourse image-repository-resource
===================================

This repo contains a rough version of a concourse resource for monitoring docker image repositories.

It emits versions based on the tags it can discover in the given image repository.
It only considers semver-ish complaints tags


## Source Configuration

* `repository`: *Required.* The image repository to watch (e.g. quay.io/coreos/etcd)

* `regex`: Filter image tags using this regex. This is also useful for extracting the version information if the tag is not semver compliant. It the regex contains no subexpressions the match is taken as the version. If it contains 1 subexpression it is considered as the version. An optional second and third subexpression are interpreted as pre-release and patch information.


### `check`: Discover image versions 

Discovers new tags in the given repository. Sorting is done using `github.com/hashicorp/go-version`.

Tags that can't be parsed as a version by `github.com/hashicorp/go-version` are ignored (Leading `v` is allowed.)

### `in`: Get image version
Creates the following files:

 * `repository`: the repository from the source configuration. 
 * `tag`: the image tag 

### `out`: not implemented


### Resource

``` yaml
- name: etcd.repository
  type: swift
  source:
    repository: quay.io/coreos/etcd
    regex: "^v[3][.0-9]*$" # (optional) only track v3.x images
```

### Plan

``` yaml
- get: etcd.repository
  version: every # (optional) if you want to trigger the pipeline for every tag found.
```

