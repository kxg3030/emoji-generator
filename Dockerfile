FROM golang:latest
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ACTION pro
COPY . .
RUN export GOPROXY=https://goproxy.io
RUN rm -f .env && cp .env.$ACTION .env && go build -o main Index.go
CMD ["./main"]