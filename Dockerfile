FROM golang:alpine AS build

WORKDIR /Activity-Relay
COPY . /Activity-Relay

RUN  mkdir -p /rootfs/usr/bin && \
     apk add -U --no-cache git && \
     go build -o /rootfs/usr/bin/server . && \
     go build -o /rootfs/usr/bin/worker ./worker && \
     go build -o /rootfs/usr/bin/ar-cli ./cli

FROM alpine

COPY --from=build /rootfs/usr/bin /usr/bin
RUN  chmod +x /usr/bin/server /usr/bin/worker /usr/bin/ar-cli && \
     apk add -U --no-cache ca-certificates
