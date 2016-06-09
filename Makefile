IMAGE_NAME ?= docker_test
CONTAINER_NAME ?= go_web_server
GBGOPATH ?= $(shell pwd):$(shell pwd)/vendor
.DEFAULT_GOAL := run

run: build
	bin/main

run-docker: build-linux docker-image docker-run

bindata:
	go-bindata -o src/main/bindata.go db/migrations/

build: clean bindata
	gb build main/...

build-linux: clean bindata
	# CGO_ENABLED=0 GOOS=linux gb build main/...
	GOPATH=$(GBGOPATH) CGO_ENABLED=0 GOOS=linux go build -o bin/main -a -installsuffix cgo -ldflags '-w -extld ld -extldflags -static' -a -x main
	ls -la bin

clean:
	if [ -a bin ]; then rm -rf bin; fi;
	if [ -a pkg ]; then rm -rf pkg; fi;

docker-image:
	docker build -t $(IMAGE_NAME) .

docker-run:
	# docker run --rm -it -p 8000:8080 --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)
	# docker run --rm -it -P --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)
	# docker run --rm -it -P --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)
	# docker run --rm -it -p 8000:8080 --add-host localhost:172.17.0.1 --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)
	docker run --rm -it -p 8000:8080 --log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)

docker-run-daemon:
	docker run -d -p 8000:8080 --log-driver json-file --name  $(CONTAINER_NAME) $(IMAGE_NAME)

test:
	bash test.sh

dbcreate:
