FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache postgresql-client ca-certificates

COPY --from=builder /app/main ./

ENV SERVER_PORT=${SERVER_PORT}

EXPOSE ${SERVER_PORT}

CMD ["./main"]
