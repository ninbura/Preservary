#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY ../preservary ./preservary
WORKDIR /go/src/app/preservary
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates docker-cli tar
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
LABEL Name=preservary Version=0.0.1