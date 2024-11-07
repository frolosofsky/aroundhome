FROM golang:1.21.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /app/matcher ./cmd/matcher

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/matcher /app/matcher

CMD ["/app/matcher"]
