FROM golang:1.15.7-alpine AS builder

RUN set -ex \
    && apk add --no-cache  git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ./ci/build

FROM alpine:3.13

EXPOSE 8080

RUN set -ex \
    && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/dist/auth /app/auth

CMD ["./auth"]
