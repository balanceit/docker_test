IMAGE_NAME ?= docker_test
CONTAINER_NAME ?= go_web_server
GBGOPATH ?= $(shell pwd):$(shell pwd)/vendor
CONTAINER_PGPORT ?= 5432
CONTAINER_PGHOST ?= 10.0.2.2


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
	docker run --rm -it -p 8000:8080 \
							-e PGHOST=$(CONTAINER_PGHOST) \
							-e PGPORT=$(CONTAINER_PGPORT) \
							--log-driver json-file --name $(CONTAINER_NAME) $(IMAGE_NAME)

docker-run-daemon:
	docker run -d -p 8000:8080 -e PGHOST=$(CONTAINER_PGHOST) --log-driver json-file --name  $(CONTAINER_NAME) $(IMAGE_NAME)

test:
	bash test.sh

dbcreate:
