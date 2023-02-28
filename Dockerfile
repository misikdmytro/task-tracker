FROM golang:1.20.1 AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /out/app ./cmd/app/main.go

FROM scratch

WORKDIR /out
COPY --from=base /out/app /out/app
COPY config /out/config
ENTRYPOINT ["./app"]