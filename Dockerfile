# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o weather ./cmd/app/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/weather .
COPY ./config/config.yml ./config/config.yml
COPY ./city.csv ./city.csv
COPY ./migrations ./migrations
CMD ["/app/weather"]