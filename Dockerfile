FROM golang:1.27.1-alpine AS build

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

FROM alpine:latest AS release

WORKDIR /app
COPY --from=build /src/app .
COPY ./static ./static
COPY ./views ./views

RUN apk -U upgrade \
	&& apk add --no-cache dumb-init ca-certificates \
	&& chmod +x /app/app

EXPOSE 3000
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./app"]
