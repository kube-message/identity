FROM golang:alpine AS build-env
RUN mkdir -p /src/github.com/im-auld/users
ENV GOPATH /
ADD . /src/github.com/im-auld/users
RUN cd /src/github.com/im-auld/users && CGO_ENABLED=0 GOOS=linux go build -o user-server-bin -a -tags netgo -ldflags '-w' main.go

FROM alpine
MAINTAINER Ian Auld<imauld@gmail.com>

COPY --from=build-env /src/github.com/im-auld/users/user-server-bin /user-server-bin
