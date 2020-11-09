FROM golang:1.15

WORKDIR /go/src/microservice-b

COPY . .

RUN go get ./...
RUN GOOS=linux go build


EXPOSE 9091

ENTRYPOINT ["./microservice-b"]