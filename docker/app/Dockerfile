FROM alpine:latest

RUN apk add openssh

RUN mkdir /app
RUN mkdir /app/download
RUN mkdir /app/log

WORKDIR /app

ADD ./run /app
ADD ./config.yml /app

CMD /app/run