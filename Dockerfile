# ---------- Frontend ----------
FROM node:22-alpine AS frontend

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build


# ---------- Go ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux \
    go build -o seabian-dashboard ./cmd/server


# ---------- Runtime ----------
FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/seabian-dashboard .
COPY --from=frontend /app/web ./web

EXPOSE 80

CMD ["./seabian-dashboard"]
