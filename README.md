# Docker MTA-STS Webserver

Simple minimalistic webserver in 30 lines of Go code which serves the needed MTA-STS file for tls mail exchange enforcing.

## Example Setup with environment variables

````
STS_MODE=testing
STS_MAX_AGE=86400
STS_SERVERS=mx1.example.tld,mx2.example.tld,mx3.example.tld
````

## Example Usage in docker-compose Stack with Traefik

TODO