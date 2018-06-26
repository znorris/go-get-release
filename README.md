# go-get-release

## Usage
go-get-release fetches all assets from a Github release tag.

Usage of go-get-release:
  -owner string
    	Owner of repository
  -repository string
    	Name of repository
  -tag string
    	Release tag
  -token string
    	Auth token, overrides env var GITHUB_AUTH_TOKEN

## Docker
```shell
# Build Dockerfile
$ docker build -t znorris/go-get-release .

# Run container
$ docker run --rm -v /tmp/foo:/tmp/downloads znorris/go-get-release -owner=ethereum -repository=mist -tag=v0.10.0
# Assets will be in /tmp/foo on local machine
```
