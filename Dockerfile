FROM golang:1.23.2-alpine3.20 AS builder
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/web

FROM alpine:3.20
WORKDIR /root/
COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/ui/html /root/ui/html
COPY --from=builder /go/src/app/ui/static /root/ui/static

EXPOSE 4000
CMD ["./app"]