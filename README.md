# Docker Webserver MTA-STS policy file

Docker Webserver for MTA-STS policy file to enforce secured mail exchange.
Simple, minimalistic 30 lines of go code to implement the mta-sts policy file as described in [RFC8461 Secion 3.2](https://datatracker.ietf.org/doc/html/rfc8461#section-3.2).

## Example Usage with docker-compose and traefik webserver

````yaml
version: '3'

services:
  mta-sts:
    image: mta-sts:1.0
    build: https://github.com/zerotens/mta-sts-webserver
    restart: always
    environment:
      - "STS_MODE=testing"
      - "STS_MAX_AGE=3600"
      - "STS_SERVERS=mx.example.org"
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.mta-sts.entrypoints=web-secure"
      - "traefik.http.routers.mta-sts.tls=true"
      - "traefik.http.routers.mta-sts.rule=Host(`mta-sts.example.org`)"
````
