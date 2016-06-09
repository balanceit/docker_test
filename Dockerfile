FROM scratch
#FROM ubuntu
ADD ./bin/main /bin/
ENV PGPORT=5454
ENV PGHOST=192.168.99.101
EXPOSE 8080
CMD ["/bin/main"]
