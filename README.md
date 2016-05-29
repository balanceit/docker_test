# docker_test

## Requirements

* boot2docker installed
* go build is installed

## Set up

* Start boot2docker
`boot2docker up`

* Setup the image to forward port 8000
`VBoxManage controlvm "boot2docker-vm" natpf1 "tcp-port8000,tcp,,8000,,8000";`

* run the default make task
`make`
 * this will build an executable `web` from the `src\web` package
 * build a docker image defined by `IMAGE_NAME` (defaults to `docker_test`)
 * run a docker container publishing the exposed port 8080 -> 8000 (see `./Dockerfile`) and a container name defined by `CONTAINER_NAME` (this defaults to `go_web_server`)
