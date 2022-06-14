FROM golang:1.18

WORKDIR /go/app

COPY . .

RUN go install niqurl-cli.go
CMD "niqurl-cli"