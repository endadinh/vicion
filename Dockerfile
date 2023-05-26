FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /vicion
RUN cd /vicion && make vicion

FROM alpine:latest

WORKDIR /vicion

COPY --from=builder /vicion/build/bin/vicion /usr/local/bin/vicion

RUN chmod +x /usr/local/bin/vicion

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/vicion"]

CMD ["--help"]
