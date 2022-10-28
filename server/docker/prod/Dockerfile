FROM golang:1.18-alpine
RUN apk add build-base
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
EXPOSE 50051
ENV HOST 0.0.0.0