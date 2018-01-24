FROM golang:alpine
ADD src /go/src
RUN go install -v main
EXPOSE 5000
ENTRYPOINT ["/go/bin/main"]
CMD ["-logtostderr"]
