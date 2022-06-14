FROM golang:1.18

WORKDIR /go/app

COPY . .

EXPOSE 8081
RUN go install niqurl-server.go
CMD "niqurl-server"