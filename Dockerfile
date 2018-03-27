FROM golang:1.8.3-alpine as builder

ARG APP_VERSION=unkown

# copy code
ADD . /go/src/github.com/malkir/mongo-loadtest

# solution root
WORKDIR /go/src/github.com/malkir/mongo-loadtest

# pull deps
RUN apk add --no-cache --virtual git
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# output
RUN mkdir /go/dist
RUN go build -ldflags "-X main.version=$APP_VERSION" \
    -o /go/dist/loadtest github.com/malkir/mongo-loadtest

FROM alpine:latest

COPY --from=builder /go/dist/loadtest /mongo-loadtest

RUN chmod 777 /mongo-loadtest

EXPOSE 9999
WORKDIR /
ENTRYPOINT ["/mongo-loadtest"]
