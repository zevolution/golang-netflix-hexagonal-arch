FROM golang:1.20-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/golang-netflix-hexagonal-arch

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/golang-netflix-hexagonal-arch ./bootstrap

FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/golang-netflix-hexagonal-arch/out/golang-netflix-hexagonal-arch /app/golang-netflix-hexagonal-arch

CMD ["/app/golang-netflix-hexagonal-arch"]