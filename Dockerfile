## STAGE 1 - build env.
FROM golang:1.17-alpine as builder

RUN apk add --update git

WORKDIR /app

COPY . .

RUN go build -v -o marvel .

## STAGE 2 - runtime env.
FROM gcr.io/distroless/static

ENV SERVICE_NAME=marvel

COPY --from=builder /app/$SERVICE_NAME /app/$SERVICE_NAME

RUN chmod +x /app/$SERVICE_NAME

ENTRYPOINT /bin/sh -c /app/$SERVICE_NAME