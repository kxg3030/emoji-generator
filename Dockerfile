FROM golang:latest
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ENV pro
COPY . .
RUN rm -f .env && cp .env.$ENV .env && go mod download && go build -o main Index.go
CMD ["./main"]