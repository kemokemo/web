FROM golang:1.13.8 AS builder
LABEL maintainer "kemokemo <t2wonderland@gmail.com>"
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o web .

FROM alpine:3.11.3
RUN adduser -D web
USER web
WORKDIR /app
COPY --from=builder /go/src/app/web .
COPY ./public ./public
ENV GIN_MODE=release
CMD ["./web"]
