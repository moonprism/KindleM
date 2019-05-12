FROM alpine:latest

RUN apk add bash
RUN apk add webp

RUN mkdir /app
RUN mkdir /app/download
RUN mkdir /app/log

WORKDIR /app

ADD ./run /app
ADD ./config.yml /app

CMD /app/run