FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o user-api ./cmd/server

FROM scratch
COPY ./config /config
COPY --from=builder /build/user-api /

ENTRYPOINT [ "/user-api", "config/local.yaml" ]