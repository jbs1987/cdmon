FROM golang:1.8.0-alpine
COPY main.go /go/src/
RUN go build /go/src/main.go
CMD /go/main
