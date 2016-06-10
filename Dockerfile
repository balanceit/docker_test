FROM scratch
#FROM ubuntu
ADD ./bin/main /bin/
ENV PGPORT=5432
ENV PGHOST=localhost
EXPOSE 8080
CMD ["/bin/main"]
