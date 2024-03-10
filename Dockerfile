FROM golang:alpine as builder

WORKDIR /backend

COPY go.* ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN cd app && swag init
RUN go build -o /backend/build/app /backend/app

FROM alpine:latest

WORKDIR /backend

ARG GIN_MODE
ARG MODELS_HOST

COPY --from=builder /backend/build/app /backend/build/app

EXPOSE 8080
ENTRYPOINT [ "/backend/build/app" ]