FROM golang:1.15-alpine as builder

WORKDIR /

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags="-s -w" -o app

FROM gcr.io/distroless/base

WORKDIR /

COPY --from=builder /app .

CMD [ "./app" ]