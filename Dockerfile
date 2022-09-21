# This docker file is for local dev, the official Dockerfile is at
# https://github.com/trickstercache/trickster-docker-images/

FROM golang:1.18 as builder
COPY . /go/src/github.com/trickstercache/trickster
WORKDIR /go/src/github.com/trickstercache/trickster

RUN GOOS=linux CGO_ENABLED=0 make build

FROM alpine:3
LABEL maintainer "The Trickster Authors <trickster-developers@googlegroups.com>"

COPY --from=builder /go/src/github.com/trickstercache/trickster/OPATH/trickster /usr/local/bin/trickster
COPY examples/conf/example.full.yaml /etc/trickster/trickster.yaml
RUN chown nobody /usr/local/bin/trickster
RUN chmod +x /usr/local/bin/trickster

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

USER nobody
ENTRYPOINT ["trickster"]
