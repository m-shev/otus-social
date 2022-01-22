FROM golang:1.17.2-alpine as builder
MAINTAINER mshev
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -mod=vendor -o social ./main.go

FROM alpine
COPY --from=builder /build/social /opt/social/
COPY --from=builder /build/config/ /opt/social/config
COPY --from=builder /build/migrations/ /opt/social/migrations
WORKDIR /opt/social
RUN apk update && apk add bash
RUN apk add --no-feed bash
RUN chmod +x ./social
CMD ["./social"]