FROM golang
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ACTION pro
COPY . .
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.io
RUN rm -f .env \
   && \cp .env.${ACTION} .env \
   && go mod download

CMD ["go", "run", "Index.go"]