FROM golang:1.20.6

WORKDIR /GO/go-server

COPY . .

RUN go build -o go-server

EXPOSE 8080

CMD [ "./go-server" ]
