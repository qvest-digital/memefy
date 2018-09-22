FROM alpine:latest
RUN apk update && apk upgrade && apk add go musl-dev
ENV GOPATH /go/
WORKDIR /go/src/memefy/server
COPY ./ /go/src/memefy/server
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app cmd/server/main.go
RUN mkdir /app && cp app /app/app && cp config.yml /app/config.yml

# Actual Image will be here
FROM alpine:latest
RUN apk update && apk upgrade && apk add ffmpeg ca-certificates
WORKDIR /root/
COPY --from=0 /app/app .
COPY --from=0 /app/config.yml .
CMD ["./app"]
EXPOSE 8080
