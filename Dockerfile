FROM golang
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV ACTION pro
COPY . .
# Timezone
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' > /etc/timezone
# Libs
RUN apt-get update \
    && apt-get install -y \
        curl \
        wget \
        git \
        zip \
        libz-dev \
        libssl-dev \
        libnghttp2-dev \
        libpcre3-dev \
		openssh-server \
    && apt-get clean \
    && apt-get autoremove
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.io
RUN rm -f .env \
   && \cp .env.${ACTION} .env \
   && go mod download \

CMD ["go", "run", "Index.go"]