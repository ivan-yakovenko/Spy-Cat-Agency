FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o spycats-app ./src/cmd/api/server.go

FROM golang:1.24
WORKDIR /app
COPY --from=builder /app/spycats-app .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY ../../migrations ./migrations
COPY ../../.env.example .
EXPOSE 8080
CMD /usr/local/bin/goose -dir ./migrations postgres "$POSTGRES_URL" up && ./spycats-app

