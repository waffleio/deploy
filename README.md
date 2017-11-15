# deploy

[![Go Report Card](https://goreportcard.com/badge/waffleio/deploy)](https://goreportcard.com/report/waffleio/deploy)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)


## Purpose
* To deploy waffle SaaS applications

## To use
### Install
* `go get github.com/waffleio/deploy`

### Configure
* Navigate to the repo for which you'd like to perform a deploy
* Drop a configuration file
  * Follow the [example](./doc/deploy.yaml) if you'd like
  * Set the necessary env vars - right now they are very specific to circleci

### Run
* `deploy`

## To Hack on
* I'm using [`dep`](https://github.com/golang/dep) for package management
* clone/fork this repo
* `cd` into the clone
* `dep ensure`
