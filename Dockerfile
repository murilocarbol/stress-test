# Etapa de build
FROM golang:1.18 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o stress-test ./cmd/main.go

# Etapa de execução
FROM alpine:latest
WORKDIR /root/

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/stress-test .
RUN chmod +x ./stress-test

ENTRYPOINT ["./stress-test"]
