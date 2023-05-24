FROM golang:1.20-alpine as builder

ENV APP_HOME /go/src/app

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

#Set github login
# RUN go env -w GOPRIVATE=github.com/oasis-prime/oas-platform-core,github.com/oasis-prime/oas-platform-firebase-core && \
#     git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com".insteadOf "https://github.com"

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o goapp ./cmd/api/main.go

##########################################

FROM alpine as release

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add git && \
    apk add build-base && \
    apk add tzdata

#Set Env
ARG ENVIRONMENT \
    VERSION \
    MONGODB_DB_NAME \
    MONGODB_HOSTS \
    MONGODB_TIMEOUT \
    MONGODB_DIAL_TIMEOUT \
    MONGODB_POOL_SIZE \
    MONGODB_USERNAME \
    MONGODB_PASSWORD \
    MONGODB_REPLICA_SET \
    MONGODB_AUTH_SOURCE \
    MONGODB_URI \
    CUSTOM_ID_PREFIX

ENV TZ=Asia/Bangkok

ENV ENVIRONMENT=$ENVIRONMENT \
    VERSION=$VERSION \
    MONGODB_DB_NAME=$MONGODB_DB_NAME \
    MONGODB_HOSTS=$MONGODB_HOSTS \
    MONGODB_TIMEOUT=$MONGODB_TIMEOUT \
    MONGODB_DIAL_TIMEOUT=$MONGODB_DIAL_TIMEOUT \
    MONGODB_POOL_SIZE=$MONGODB_POOL_SIZE \
    MONGODB_USERNAME=$MONGODB_USERNAME \
    MONGODB_PASSWORD=$MONGODB_PASSWORD \
    MONGODB_REPLICA_SET=$MONGODB_REPLICA_SET \
    MONGODB_AUTH_SOURCE=$MONGODB_AUTH_SOURCE \
    MONGODB_URI=$MONGODB_URI \
    CUSTOM_ID_PREFIX=$CUSTOM_ID_PREFIX

COPY --from=builder ./app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]