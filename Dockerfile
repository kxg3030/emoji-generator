FROM golang:latest
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
COPY . .
RUN go mod download && go build Index.go -o main
CMD["go", "run", "main"]