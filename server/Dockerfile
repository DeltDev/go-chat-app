FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache make

COPY --from=builder /app/main .
COPY --from=builder /app/Makefile .
COPY --from=builder /app/db/migrations ./db/migrations

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

EXPOSE 8080
CMD ["./main"]