FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./  
RUN go mod download  

COPY . ./  

RUN go build -o pinger ./cmd/main.go  

FROM alpine:latest  

WORKDIR /root/  

COPY --from=builder /app/pinger  ./

CMD ["./pinger"]
