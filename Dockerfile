FROM golang:latest
LABEL maintainer AOAO "xuao@gmail.com"

ADD ../../src/pincloud.purchase/  $GOPATH/src/pincloud.purchase
WORKDIR $GOPATH/src/pincloud.purchase
RUN go install
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./pincloud.purchase"]