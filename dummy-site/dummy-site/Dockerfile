FROM golang:alpine

WORKDIR /usr/src/app

COPY main.go ./

RUN go build -o main .

CMD ["./main"]
