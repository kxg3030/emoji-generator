FROM golang
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ACTION pro
ENV GOPROXY "https://goproxy.io"
COPY . .
RUN rm -f .env && \cp .env.${ACTION} .env
RUN go mod tidy && go mod download
ENTRYPOINT ["go","run","Index.go"]