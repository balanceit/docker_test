FROM scratch
#FROM ubuntu
ADD ./bin/main /bin/
ENV PGPORT=5432
ENV PGHOST=10.0.2.2
EXPOSE 8080
CMD ["/bin/main"]
