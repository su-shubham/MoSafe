FROM golang:1.21-alpine AS builder

WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/account-service

FROM alpine:latest
COPY --from=builder /bin/account-service /bin/account-service
RUN chmod +x /bin/account-service

EXPOSE 8081
CMD ["/bin/account-service"]
