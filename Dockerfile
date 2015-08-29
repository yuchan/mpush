FROM golang:1.5-onbuild
MAINTAINER Yusuke Ohashi <github@junkpiano.me>

ADD . /code
WORKDIR /code
RUN go get github.com/timehop/apns
RUN go get github.com/gin-gonic/gin
RUN go build -o mpush
CMD ./mpush
