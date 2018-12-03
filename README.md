# Dyndgo

[![Go Report Card](https://goreportcard.com/badge/github.com/ulm0/dyndgo)](https://goreportcard.com/report/github.com/ulm0/dyndgo) [![](https://godoc.org/github.com/ulm0/dyndgo?status.svg)](http://godoc.org/github.com/ulm0/dyndgo) [![](https://images.microbadger.com/badges/image/ulm0/dyndgo.svg)](https://microbadger.com/images/ulm0/dyndgo "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/ulm0/dyndgo.svg)](https://microbadger.com/images/ulm0/dyndgo "Get your own version badge on microbadger.com") 

> Dyndgo is a tiny tool to update your A records on DNSimple

## Install dyndgo in your system

```sh
go get github.com/ulm0/dyndgo/cmd/dyndgo
```

## Use it

All you need to do is write your token and domains to [`data.yml`](data.yml), and execute the tool (`dyndgo -f /path/to/data.yml`)

```yaml
credentials:
  token: token
zones:
  # Domain must be A records only
  domain.com:
    - subdomain1
    - subdomain2
```

## Use it with crontab

Run it every eight hours using [`updater` script](updater)

```sh
0 */8 * * * /path/to/updater
```

**Note**: [`updater` script](updater) uses docker.