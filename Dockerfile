FROM golang:1.12-alpine AS builder

ENV GO111MODULE=on

LABEL maintainer="savchukyarpolk@gmail.com"

WORKDIR /go/src/github.com/ysv/pick

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git bash gcc libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/pick pick.go

FROM alpine
COPY --from=builder /go/bin/pick /
EXPOSE 8080
ENTRYPOINT ["/pick"]
CMD ["server"]
