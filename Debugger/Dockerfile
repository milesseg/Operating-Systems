FROM golang:alpine3.17

RUN apk add tree
RUN go env -w GO111MODULE=off
RUN go indtall --gcflags="-N -l" github.com/milesseg/hello
RUN go install github.com/milesseg/debugger
