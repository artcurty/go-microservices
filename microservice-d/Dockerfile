FROM golang:1.15

WORKDIR /go/src/microservice-d

COPY . .

RUN go get ./...
RUN GOOS=linux go build


EXPOSE 9093

ENTRYPOINT ["./microservice-d"]