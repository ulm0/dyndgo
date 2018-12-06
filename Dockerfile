FROM golang:alpine
LABEL maintainer "Pierre Ugaz <ulm0@innersea.xyz>"
COPY . /go/src/github.com/ulm0/dyndgo
WORKDIR /go/src/github.com/ulm0/dyndgo/cmd/dyndgo
RUN apk add --no-cache \
    upx && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo && \
    upx --ultra-brute dyndgo
FROM scratch
WORKDIR /app/
ADD https://raw.githubusercontent.com/containous/traefik/master/script/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /go/src/github.com/ulm0/dyndgo/cmd/dyndgo/dyndgo /app/dyndgo
# You need to pass your YAML file with your credentials and mount it to /app/data.yml
COPY --from=0 /go/src/github.com/ulm0/dyndgo/data.yml /app/data.yml
ENTRYPOINT [ "/app/dyndgo" ]
CMD [ "-f", "/app/data.yml"]
