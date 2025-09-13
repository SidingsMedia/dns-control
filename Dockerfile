# SPDX-FileCopyrightText: 2023 Sidings Media
# SPDX-License-Identifier: MIT

FROM golang:latest as build

## Build
WORKDIR /build

COPY go.mod /build
COPY go.sum /build

# Download go modules
RUN go mod download

# Copy all files
COPY . /build

# Compile binary
RUN CGO_ENABLED=0 go build -a -o server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build/server /server

ENV GIN_MODE=release

EXPOSE 3000/tcp

USER nonroot:nonroot

ENTRYPOINT ["/server"]
