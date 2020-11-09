FROM golang:1.15

WORKDIR /go/src/microservice-a

COPY . .

RUN go get ./...
RUN GOOS=linux go build


EXPOSE 9090

ENTRYPOINT ["./microservice-a"]