FROM golang:1.18

WORKDIR /go/app
RUN git clone "https://github.com/nicolebroyak/niqurldev/"
WORKDIR /go/app/niqurldev/cmd/server
RUN go mod tidy
RUN go build -o server

EXPOSE 8081
CMD [ "./server" ]