# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build

ENV GOTOOLCHAIN=local

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY cmd/ cmd/
COPY internal/ internal/

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

RUN mkdir -p /out/data

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/api /api
COPY --from=build --chown=65532:65532 /out/data /data

ENV COFFER_DB_PATH=/data/coffer.db

USER nonroot

ENTRYPOINT ["/api"]
