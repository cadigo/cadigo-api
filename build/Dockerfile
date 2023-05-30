FROM golang:1.20-alpine as builder

WORKDIR /app

RUN set -ex; \
    apk update; \
    apk add --no-cache \
    tzdata \
    git \
    gcc \
    musl-dev \
    make

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o goapp ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o goapp ./cmd/api/main.go

RUN strip goapp

##########################################

FROM alpine as release

WORKDIR /app

RUN apk apk add tzdata && \
    add --no-cache ca-certificates

ENV TZ=Asia/Bangkok

COPY --from=builder ./app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]