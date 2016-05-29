FROM golang

ADD . /go/
RUN go install web
ENTRYPOINT /go/bin/web
EXPOSE 8080
