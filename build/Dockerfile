FROM golang:alpine AS builder
WORKDIR /go/src/code-challenge-hl
COPY . .
RUN go get -tags "musl dynamic" -d -v ./...
RUN go build -tags "musl dynamic" -o /go/bin/code-challenge-hl cmd/main.go

FROM alpine
COPY --from=builder /go/bin/code-challenge-hl /go/bin/code-challenge-hl
COPY --from=builder /go/src/code-challenge-hl/configs/config.yaml /go/bin/config.yaml
WORKDIR /go/bin/
EXPOSE 8080
ENTRYPOINT ["/go/bin/code-challenge-hl"]