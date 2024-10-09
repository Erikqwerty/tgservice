FROM golang:1.23-alpine AS builder

COPY . /github.com/erikqwerty/tgservice/
WORKDIR /github.com/erikqwerty/tgservice/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/erikqwerty/tgservice/bin/crud_server .

CMD ["./crud_server", "-config-path", "./"]