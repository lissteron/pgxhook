FROM golang:1.22.2-alpine

ENV WORKDIR=/src
RUN apk add git make gcc g++

RUN mkdir -p ${WORKDIR}

WORKDIR ${WORKDIR}
