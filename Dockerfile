FROM golang:1.19-alpine AS build
WORKDIR /app/
COPY . .
RUN go build -o executable bill-payment-assistant/cmd

FROM alpine:latest
WORKDIR /app/
COPY --from=build /app/executable ./
CMD ./executable