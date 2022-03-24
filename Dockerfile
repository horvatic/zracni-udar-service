FROM golang:alpine AS builder
ENV GOOS=linux

WORKDIR /build
COPY . .
RUN go build -o bin/zracni-udar-service cmd/zracni-udar-service/main.go

FROM alpine:3
ENV MONGO_CONNECTION_STRING=default \
        MONGO_DATABASE=default \
        MONGO_COLLECTION=default \
        FRONT_END_HOST=default \
        GITHUB_PAT=default

RUN apk --no-cache add ca-certificates
WORKDIR /dist
COPY --from=builder /build/bin/zracni-udar-service .
EXPOSE 8080
CMD ["/dist/zracni-udar-service"]
