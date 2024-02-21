FROM golang:1.21-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o /app/bin/server ./cmd/server/main.go

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
  ca-certificates && \
  rm -rf /var/lib/apt/lists/

COPY --from=builder /app/bin/server /app/bin/server

ENV PORT=80
EXPOSE 80

CMD ["/app/bin/server"]