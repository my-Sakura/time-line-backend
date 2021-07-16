ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS builder
ENV GOPROXY="https://goproxy.io,direct"
ENV CGO_ENABLED=0
ENV GIN_MODE=release
WORKDIR /go/release
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app ./cmd/main.go

FROM scratch as prod
COPY --from=builder /go/release/app .
EXPOSE 10002

CMD ["/app"]