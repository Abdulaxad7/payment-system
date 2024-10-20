FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/payment-system .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/payment-system /usr/local/bin/payment-system
EXPOSE 8080
CMD ["payment-system"]