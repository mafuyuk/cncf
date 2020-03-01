FROM golang:1.13-alpine AS builder

WORKDIR /opt/app-b
COPY go.mod go.sum /opt/app-b/
RUN go mod download

COPY . /opt/app-b/
RUN go build

CMD ["/opt/app-b/app-b"]

# --- #
FROM alpine:3.10

COPY --from=builder /opt/app-b/app-b /usr/local/bin

CMD ["/usr/local/bin/app-b"]
