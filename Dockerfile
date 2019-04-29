FROM golang:1.12-alpine AS builder

ENV GO111MODULE=on

LABEL maintainer="savchukyarpolk@gmail.com"

WORKDIR /go/src/github.com/ysv/pick

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git bash gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go get -u github.com/gobuffalo/packr/packr \
 && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $GOPATH/bin/packr build -o pick pick.go

CMD ["./pick"]

# TODO: Fix multistage build.
#FROM alpine
#COPY --from=builder /go/src/github.com/ysv/pick/pick /
#EXPOSE 8080
#ENTRYPOINT ["/pick"]
#CMD ["server"]
