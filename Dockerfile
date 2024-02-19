#
# Copyright (c) 2024 Dylan O' Connor Desmond
#

FROM golang:alpine as builder

# hadolint ignore=DL3018
RUN apk update \
    && apk add --no-cache \
        ca-certificates \
        tzdata \
    && update-ca-certificates

# Create an app user
ENV USER=budgie
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app

COPY go.mod .

RUN go mod download \
    && go mod verify

COPY . .

# Build the binary
RUN GOOS=linux GOARCH=amd64 \
    go \
    build \
    -ldflags='-w -s -extldflags "-static"' \
    -a \
    -o ./backend/budgie \
    ./backend/cmd

# Build a small image
FROM scratch

WORKDIR /app

# Copy from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/backend/config/config.env /app/backend/config/config.env
COPY --from=builder /app/backend/budgie /app/backend/budgie

USER budgie:budgie

EXPOSE 8080

ENTRYPOINT [ "/app/backend/budgie" ]
