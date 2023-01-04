FROM golang:1.19-alpine AS builder
WORKDIR /usr/src/usdrub-bot
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o usdrub-bot ./cmd/app

FROM alpine:3.17.0
RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow
COPY --from=builder /usr/src/usdrub-bot/usdrub-bot /usr/local/bin/usdrub-bot
CMD ["usdrub-bot"]