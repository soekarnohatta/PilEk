#
# 1. Build Go App
#
FROM golang:1.16-alpine AS build-go
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_CFLAGS="-g -O2 -Wno-return-local-addr"
RUN apk add build-base
WORKDIR /app/backend
COPY go.mod /app/backend
COPY go.sum /app/backend
RUN go mod download
COPY . /app/backend
RUN go build -o ./PilEk

#
# 2. Build Vue App
#
FROM node:lts-alpine as build-vue
WORKDIR /app/frontend
COPY client /app/frontend
RUN npm install
RUN npm run build

#
# 3. Runtime Container
#
FROM alpine
LABEL maintainer="Nayef Haidir <nayefhaidir@outlook.co.id>"
ENV TZ=Asia/Jakarta \
    PATH="/app:${PATH}"
RUN apk add --update --no-cache \
    sqlite \
    tzdata \
    ca-certificates \
    bash \
    && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /app
COPY --from=build-vue /app/frontend/dist /app/dist/
COPY --from=build-go /app/backend/PilEk /app/
COPY --from=build-go /app/backend/voting.db /app/
EXPOSE 8080
CMD ["./PilEk"]
