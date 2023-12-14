FROM golang:1.18.10-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/api/api.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./migrate ./cmd/migrate/migrate.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./seed ./cmd/seed/seed.go

FROM alpine:3 as migration

WORKDIR /app

RUN apk add --no-cache make

COPY --from=builder /app/migrate /app/
COPY --from=builder /app/seed /app/
COPY --from=builder /app/Makefile /app/

CMD make migration_docker

FROM alpine:3 as dev

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 8080

CMD ./main
