FROM ubuntu:latest
RUN apt-get update && apt-get install -y libvlc-dev golang-go git
ENV GOPATH /go/
WORKDIR /go/src/memefy/client
COPY ./ /go/src/memefy/client
RUN go get .
RUN go build -o app .
RUN mkdir /app && cp app /app/app

# Actual Image will be here
FROM ubuntu:latest
RUN apt-get update && apt-get install -y libvlc-dev ca-certificates libxext-dev libxrender-dev libxtst-dev
WORKDIR /root/
COPY --from=0 /app/app .
CMD ["./app"]  
