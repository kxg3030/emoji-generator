FROM golang:latest
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ENV pro
ENV GOPROXY=https://goproxy.io
COPY . .
RUN rm -f .env && cp .env.$ENV .env && go build -o main Index.go
CMD ["./main"]