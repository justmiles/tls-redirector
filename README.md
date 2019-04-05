# tls-redirector
[![justmiles/tls-redirector](https://img.shields.io/badge/docker-justmiles/tls--redirector-brightgreen.svg)](https://hub.docker.com/r/justmiles/tls-redirector/)

2MB docker image that redirects all requests to the TLS endpoint of the same name.

## Usage

    docker run -d -p 80:80 justmiles/tls-redirector:latest

## Example

    curl -I example.milesmaddox.com
    HTTP/1.1 301 Moved Permanently
    Content-Type: text/html; charset=utf-8
    Location: https://example.milesmaddox.com/
    Date: Fri, 05 Apr 2019 19:32:11 GMT


## Building

    docker build -t justmiles/tls-redirector .
