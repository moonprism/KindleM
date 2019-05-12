FROM alpine:latest

RUN apk add bash
RUN apk add webp
RUn apk add wget

RUN wget http://kindlegen.s3.amazonaws.com/kindlegen_linux_2.6_i386_v2_9.tar.gz
# todo install kindleGen

RUN mkdir /app
RUN mkdir /app/download
RUN mkdir /app/log

WORKDIR /app

ADD ./run /app
ADD ./config.yml /app

CMD /app/run