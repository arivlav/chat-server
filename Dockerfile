FROM golang:1.22.0-alpine AS builder

COPY . /github.com/arivlav/chat-server/source/
WORKDIR /github.com/arivlav/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/arivlav/chat-server/source/bin/chat_server .

CMD ["./chat_server"]