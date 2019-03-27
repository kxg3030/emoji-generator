FROM golang
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ACTION pro
COPY main .
RUN rm -f .env && \cp .env.${ACTION} .env
ENTRYPOINT ["./main"]