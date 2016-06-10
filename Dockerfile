FROM scratch
#FROM ubuntu
ADD ./bin/main /bin/
EXPOSE 8080
CMD ["/bin/main"]
