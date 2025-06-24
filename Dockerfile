FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git && \
    go install github.com/cosmtrek/air@v1.40.4

COPY . .

CMD ["air"]