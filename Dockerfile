FROM golang:latest
LABEL maintainer AOAO "xuao@gmail.com"

ADD ./purchaseApp/  $GOPATH/src/pincloud.purchase/purchaseApp
WORKDIR $GOPATH/src/pincloud.purchase/purchaseApp/
RUN go get
RUN go build .

EXPOSE 9401

ENTRYPOINT ["./purchaseApp"]