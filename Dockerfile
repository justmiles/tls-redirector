FROM golang:1.10-stretch as builder
RUN mkdir -p /go/src/
WORKDIR /go/src/
COPY main.go /go/src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tls-redirector main.go
RUN md5sum tls-redirector

FROM scratch
COPY --from=builder /go/src/tls-redirector /tls-redirector

ENTRYPOINT ["/tls-redirector"]