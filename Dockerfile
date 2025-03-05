FROM golang:1.24.0 AS backend-builder

WORKDIR /backendcompile

COPY . .

RUN CGO_ENABLED=0 go build -o velo-mom-api main.go

FROM alpine:latest AS prod

WORKDIR /build

COPY --from=backend-builder /backendcompile/velo-mom-api .

EXPOSE 8080
ENTRYPOINT ["./velo-mom-api"]