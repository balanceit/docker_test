[![Build Status](https://travis-ci.org/balanceit/docker_test.svg?branch=master)](https://travis-ci.org/balanceit/docker_test)

# docker_test

## Requirements

* ~boot2docker installed~
* docker machine is installed
* go build is installed

## Set up
* `boot2docker down` -- shutdown if running
* `brew cask install dockertoolbox` -- installs docker machine and the lastest version of virtual box
* `docker-machine create --driver virtualbox docker-vm` -- create a new VM and machine named *docker-vm*
* `eval $(docker-machine env docker-vm)` -- add needed env variables
* `docker ps` -- just a quick test to ensure it is working

* Setup the image to forward port 8000
`VBoxManage controlvm "docker-vm" natpf1 "tcp-port8000,tcp,,8000,,8000";`

* run the default make task
`make`
 * this will build an executable `web` from the `src\web` package
 * build a docker image defined by `IMAGE_NAME` (defaults to `docker_test`)
 * run a docker container publishing the exposed port 8080 -> 8000 (see `./Dockerfile`) and a container name defined by `CONTAINER_NAME` (this defaults to `go_web_server`)

## building a small container
* see the `.travis.yml`
* from https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
