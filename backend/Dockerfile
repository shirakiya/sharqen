FROM golang:latest
LABEL maintainer="shirakiya"

WORKDIR /go/src/github.com/shirakiya/sharqen/backend

ADD . .

RUN go get -u github.com/golang/dep/cmd/dep \
    && make deps

CMD ["make", "run"]
