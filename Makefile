IMAGE_NAME ?= docker_test
CONTAINER_NAME ?= go_web_server
.DEFAULT_GOAL := all

all: build docker-image

build:
	gb build web/...

build-linux:
	# CGO_ENABLED=0 GOOS=linux gb build web/...
	GOPATH=$(pwd) CGO_ENABLED=0 GOOS=linux go build -o bin/web -a -installsuffix cgo -ldflags '-w -extld ld -extldflags -static' -a -x src/web/main.go
	ls -la bin

clean:
	rm -r bin
	rm -r pkg

docker-image:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run --rm -it -p 8000:8080 --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)

docker-run-daemon:
	docker run -d -p 8000:8080 --log-driver json-file --name  $(CONTAINER_NAME) $(IMAGE_NAME)

test:
	bash test.sh
