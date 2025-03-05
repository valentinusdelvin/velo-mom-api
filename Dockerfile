FROM golang:1.23.1@sha256:2fe82a3f3e006b4f2a316c6a21f62b66e1330ae211d039bb8d1128e12ed57bf1 AS backend-builder

WORKDIR /backendcompile

COPY . .

RUN CGO_ENABLED=0 go build -o velo-mom-api main.go

FROM alpine:latest AS prod

WORKDIR /build

COPY --from=backend-builder /backendcompile/velo-mom-api .

EXPOSE 3014
ENTRYPOINT ["./velo-mom-api"]