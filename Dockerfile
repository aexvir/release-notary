FROM golang:1.13.3-alpine as builder
RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache make git gcc musl-dev

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum Makefile /app/
RUN make install_deps

COPY . /app/
RUN make build/docker

FROM alpine:3.10.2
RUN  apk add --no-cache --virtual=.run-deps ca-certificates git &&\
    mkdir /app

WORKDIR /app
COPY --from=builder /app/build/release-notary ./release-notary

RUN ln -s $PWD/release-notary /usr/local/bin

CMD [ "release-notary", "publish" ]