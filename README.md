# go-get-release

## Usage
```shell
go-get-release fetches all assets from a Github release tag.

Usage of go-get-release:
  -owner string
    	Owner of repository
  -repository string
    	Name of repository
  -tag string
    	Release tag
  -token string
    	Auth token, overrides env var $GITHUB_AUTH_TOKEN
```

## Docker
```shell
# Build Dockerfile
$ docker build -t znorris/go-get-release .

# Run container
$ docker run --rm -v /tmp/foo:/tmp/downloads znorris/go-get-release -owner=znorris -repository=go-get-release -tag=v0.1.0-alpha

# Assets will be in /tmp/foo on local machine
$ ls -l /tmp/foo
total 14384
-rw-r--r--  1 user  wheel  6667852 Jun 26 09:56 go-get-release-linux-amd-64
```
