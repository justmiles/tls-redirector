# tls-redirector
[![justmiles/tls-redirector](https://img.shields.io/badge/docker-justmiles/tls--redirector-brightgreen.svg)](https://hub.docker.com/r/justmiles/tls-redirector/)

2MB docker image that redirects all requests to their TLS endpoint

## Usage

    docker run -d -p 80:80 justmiles/tls-redirector:latest

## Example

    curl -I localhost/pages/xyz
    HTTP/1.1 301 Moved Permanently
    Content-Type: text/html; charset=utf-8
    Location: https://localhost/pages/xyz

## Building

    docker build -t justmiles/tls-redirector .