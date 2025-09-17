# SPDX-FileCopyrightText: 2023 Sidings Media
# SPDX-License-Identifier: MIT

FROM golang:latest AS build

## Build
WORKDIR /build

COPY go.mod /build
COPY go.sum /build

# Download go modules
RUN go mod download

# Copy all files
COPY . /build

# Compile binary
RUN CGO_ENABLED=0 go build -a -o service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build/service /service

ENV GIN_MODE=release

EXPOSE 3000/tcp

USER nonroot:nonroot

ENTRYPOINT ["/service", "--config", "/etc/server/config.yaml"]
