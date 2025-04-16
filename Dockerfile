FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go build -o /usr/src/app/app .

EXPOSE 8080

CMD ["/usr/src/app/app"]
