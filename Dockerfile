FROM golang:1.19-alpine AS build
WORKDIR /app/
COPY . .
RUN ls
RUN go build -o executable

FROM alpine:latest
WORKDIR /app/
COPY --from=build /app/executable ./
CMD ./executable