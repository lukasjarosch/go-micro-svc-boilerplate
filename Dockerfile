FROM golang:1.11 as builder

WORKDIR /go-modules
COPY . .
# Building using -mod=vendor, which will utilize the vendor directory
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o service

FROM alpine:3.8

COPY --from=builder /go-modules/service .
ENTRYPOINT ["/service"]