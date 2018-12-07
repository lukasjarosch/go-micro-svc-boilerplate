FROM golang:1.11 as builder

ENV GO111MODULE=on
WORKDIR /build

# warm up dependency cache
COPY go.mod .
COPY go.sum .
RUN go mod download

# build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /build/service

FROM alpine:3.8
COPY --from=builder /build/service /service
COPY config.json .
ENTRYPOINT ["/service"]