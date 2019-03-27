FROM golang
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ACTION pro
COPY . .
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.io && source /etc/profile
RUN rm -f .env && \cp .env.${ACTION} .env && chmod +x main
CMD ["./main"]