FROM golang:alpine

LABEL base.name="back-atv"

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY . .

RUN go mod vendor
RUN go build -o main .

EXPOSE 5000
ENTRYPOINT ["./main"]