FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o spycats-app ./src/cmd/api/server.go

FROM golang:1.24
WORKDIR /app
COPY --from=builder /app/spycats-app .
COPY ../../.env.example .
EXPOSE 8080
CMD ["./spycats-app"]


