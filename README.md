# GO Battle Royal
[![Build Status](https://travis-ci.com/jorgechato/battle-royal.svg?branch=master)](https://travis-ci.com/jorgechato/battle-royal)
[![Go Report Card](https://goreportcard.com/badge/github.com/jorgechato/battle-royal)](https://goreportcard.com/report/github.com/jorgechato/battle-royal)
[![Godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/jorgechato/battle-royal)
[![Docker Repository](https://img.shields.io/badge/docker-image-blue.svg)](https://hub.docker.com/r/jorgechato/battle-royal)

Battle royal between different languages in the backend side [GO vs JS vs Java vs Kotlin vs Vertx]

This backend can be used as a **graphql** or **rest** seed in `GO`

In these repository you will find a backend written in GO with [graphql](https://github.com/graphql-go/graphql) and API endpoints (hybrid solution).
In the following list you will find my colleagues' repos:

- [TS with nestJS](https://github.com/RecuencoJones/nestjs-perf-test) written by @RecuencoJones

## Must have

- [Go>=9](https://golang.org/)
- [dep](https://github.com/golang/dep)

## Install

```bash
$ go get -u -v https://github.com/jorgechato/battle-royal
```

## Deploy

```bash
$ docker run -p 8080:8080 -p 8081:8081 --name battle-royal -d jorgechato/battle-royal
```
