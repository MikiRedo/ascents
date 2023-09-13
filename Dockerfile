FROM golang:1.20.6

WORKDIR /GO/go-server

RUN go build -o go-server

EXPOSE 8080

CMD [ "./go-server" ]

