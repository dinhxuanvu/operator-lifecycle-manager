FROM golang:1.13-alpine AS builder

RUN apk update && apk add sqlite build-base git mercurial bash
WORKDIR /build

COPY vendor vendor
COPY cmd cmd
COPY pkg pkg
COPY Makefile Makefile
COPY go.mod go.mod
RUN make static
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-$(go env GOARCH) && \
    chmod +x /bin/grpc_health_probe

FROM scratch
COPY --from=builder /build/bin/opm /bin/opm
COPY --from=builder /build/bin/initializer /bin/initializer
COPY --from=builder /build/bin/appregistry-server /bin/appregistry-server
COPY --from=builder /build/bin/configmap-server /bin/configmap-server
COPY --from=builder /build/bin/registry-server /bin/registry-server
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
