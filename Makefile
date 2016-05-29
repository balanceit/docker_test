IMAGE_NAME ?= docker_test
CONTAINER_NAME ?= go_web_server
.DEFAULT_GOAL := all

all: build docker-image

build:
	gb build web/...

docker-image:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run --rm -it -p 8000:8080 --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)
