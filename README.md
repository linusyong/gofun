# Introduction
This project just tested using Docker to build Go binary with a Makefile that will help.  It allows you to do **go** development without having to install **go** on your system.

# Requirements
*  [jq](https://stedolan.github.io/jq/)
*  [curl](https://curl.haxx.se/)
*  [musl libc](https://www.musl-libc.org/)
_On Ubuntu simply `sudo apt install jq curl musl`_

# Quick Start
*  tested on Ubuntu 18.04, Docker 18.06.0-ce 
*  first you need to have docker running
*  `make dockerimage` will create a docker image golangdep:latest from [golang:alpine](https://hub.docker.com/_/golang/) + [dep](https://golang.github.io/dep/) + [golint](https://github.com/golang/lint)
*  `make` will (perform all these in docker container):
   *  perform `depinit` if no Gopkg.toml exists to manage the dependencies
   *  perform `dep` to update the dependencies packages
   *  perform `lint` to use **golint** check the style mistake
   *  perform `build` if the **app** binary is not up-to-date with the **main.go** source file
