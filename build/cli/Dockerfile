FROM golang:1.18

WORKDIR /go/app
RUN git clone "https://github.com/nicolebroyak/niqurldev/"
WORKDIR /go/app/niqurldev/cmd/cli
RUN go mod tidy
RUN go build -o niqurl

EXPOSE 8081
CMD [ "./niqurl" ]