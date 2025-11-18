FROM golang:latest AS builder

WORKDIR /app

COPY . .
COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o question_answer ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/question_answer .
COPY common/migrations/sql/ common/migrations/sql/

EXPOSE 8080

CMD ["./question_answer"]