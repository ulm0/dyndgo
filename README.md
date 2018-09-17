# Dyndgo

> Dynamic DNS record updater for DNSimple

Dyndgo is a tiny tool to update your A records on DNSimple

## Install dyndgo in your system

```sh
go get github.com/ulm0/dyndgo
```

## Use it

All you need to do is write your token and domains to [`data.yml`](data.yml), and execute the tool

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