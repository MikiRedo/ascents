FROM golang:1.20.6-alpine3.15 as builder

WORKDIR /GO/go-server

RUN go build -o go-server

EXPOSE 8080

CMD [ "./go-server" ]

