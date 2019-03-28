FROM golang
MAINTAINER "kxg3030@sina.com"
WORKDIR /home/emoji
ENV FFMPEG 4.1
# Libs
RUN apt-get update \
    && apt-get install -y \
         tar \
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

RUN wget https://ffmpeg.org/releases/ffmpeg-${FFMPEG}.tar.bz2 -O ./ffmpeg.tar.bz2 \
    && mkdir -p ./ffmpeg \
    && tar -jxvf ffmpeg.tar.bz2 -C ./ffmpeg \
    && rm ./ffmpeg.tar.bz2 \
    && cd ./ffmpeg \
    && ./configure \
    && make && make install

ENV ACTION pro
ENV GOPROXY "https://goproxy.io"
COPY . .
RUN rm -f .env && \cp .env.${ACTION} .env
RUN go mod tidy && go mod download
ENTRYPOINT ["go","run","Index.go"]