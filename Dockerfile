FROM golang:alpine as builder

WORKDIR /backend

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /backend/build/app ./app

FROM alpine:latest

WORKDIR /backend

ENV GIN_MODE="release"
ENV MODELS_HOST="models"

COPY --from=builder /backend/build/app /backend/build/app
COPY --from=builder /backend/config /backend/config

EXPOSE 8080
ENTRYPOINT [ "/backend/build/app" ]