FROM golang:1.14.1 as builder

ADD . ./
RUN go build ./server.go && chmod +x server

ENTRYPOINT [ "./server" ]