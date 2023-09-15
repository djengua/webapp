#build stage
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .


EXPOSE 8081
CMD [ "/app/main" ]
