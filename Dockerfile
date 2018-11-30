FROM golang:alpine
WORKDIR /go/src/github.com/ulm0
RUN apk add --no-cache \
    git \
    upx && \
    go get -v -u github.com/golang/dep/cmd/dep && \
    go get -v -d github.com/ulm0/dyndgo/cmd/dyndgo && \
    cd dyndgo && \
    dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo && \
    upx --best dyndgo
FROM scratch
WORKDIR /app/
ADD https://raw.githubusercontent.com/containous/traefik/master/script/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /go/src/github.com/ulm0/dyndgo/dyndgo /app/dyndgo
# You need to pass your YAML file with your credentials and mount it to /app/data.yml
COPY --from=0 /go/src/github.com/ulm0/dyndgo/data.yml /app/data.yml
ENTRYPOINT [ "/app/dyndgo" ]
CMD [ "-f", "/app/data.yml"]