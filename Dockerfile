FROM golang:1.9

WORKDIR /src

ENV GOPATH /

RUN go get -u github.com/nsf/termbox-go
RUN go get -u github.com/gorilla/mux
