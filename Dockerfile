FROM golang:alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/microservices

RUN mkdir /microservices
ADD . /microservices
WORKDIR /microservices

RUN go mod vendor
RUN go build

EXPOSE 8000
CMD [ "/microservices/microservices" ]