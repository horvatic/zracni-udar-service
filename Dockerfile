FROM golang:1.16-buster AS builder
ENV GOOS=linux \
    GOARCH=arm64

WORKDIR /build
COPY . .
RUN go build -o bin/zracni-udar-service cmd/zracni-udar-service/main.go

FROM golang:1.16-buster
ENV NAMESPACE=default \
		SERVICE=default \
        MONGO_CONNECTION_STRING=default \
        MONGO_DATABASE=default \
        MONGO_COLLECTION=default \
        FRONT_END_HOST=default

WORKDIR /dist
COPY --from=builder /build/bin/zracni-udar-service .
EXPOSE 8080
CMD ["/dist/zracni-udar-service"]
