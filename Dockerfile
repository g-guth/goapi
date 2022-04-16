#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/app /app/main
ENTRYPOINT /app/main
LABEL Name=goapi Version=0.0.1
EXPOSE 8080
