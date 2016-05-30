FROM scratch
ADD ./bin/web /
EXPOSE 8080
CMD ["/web"]
