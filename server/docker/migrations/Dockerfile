FROM golang:1.17-buster as builder

# Create and change to the app directory.
WORKDIR /app

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
# COPY ./docker/migrations/init.sql ./init.sql
COPY ./docker/migrations .
RUN go mod download
RUN go build -v -o main

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY . ./
COPY --from=builder /app/main /main
COPY --from=builder /app/init.sql /init.sql

CMD ["/main"]


