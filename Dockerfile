# ---- Build stage ----
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o server .

# ---- Final minimal image ----
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
