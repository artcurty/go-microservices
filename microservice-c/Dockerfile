FROM golang:1.15

WORKDIR /go/src/microservice-c

COPY . .

RUN go get ./...
RUN GOOS=linux go build


EXPOSE 9092

ENTRYPOINT ["./microservice-c"]