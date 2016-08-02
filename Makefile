IMAGE_NAME ?= testing_image
CONTAINER_NAME ?= testing_container
GBGOPATH ?= $(shell pwd):$(shell pwd)/vendor
CONTAINER_PGPORT ?= 5432
CONTAINER_PGHOST ?= 10.0.2.2
DB_CONNECTION_STRING ?= "dbname=docker_test_developement user=postgres"
TMPDIR ?= "/tmp"

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
	if [ -a src/main/bindata.go ]; then rm -f src/main/bindata.go; fi;
	if [ -a bin ]; then rm -rf bin; fi;
	if [ -a pkg ]; then rm -rf pkg; fi;

docker-image:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run --rm -it -p 8000:8080 -e DB_CONNECTION_STRING=$(DB_CONNECTION_STRING) -e PGHOST=${CONTAINER_PGHOST} --log-driver json-file --name  $(CONTAINER_NAME) $(IMAGE_NAME)

docker-run-daemon:
	docker run -d -p 8000:8080 -e DB_CONNECTION_STRING=$(DB_CONNECTION_STRING) -e PGHOST=${CONTAINER_PGHOST} --log-driver json-file --name  $(CONTAINER_NAME) $(IMAGE_NAME)
	#docker run -d -p 8000:8080 -e PGHOST=$(CONTAINER_PGHOST) --log-driver json-file --name  $(CONTAINER_NAME) $(IMAGE_NAME)

test:
	bash test.sh

dbcreate:

test1:
	@echo $(TMPDIR)
	@echo $(DB_CONNECTION_STRING)
