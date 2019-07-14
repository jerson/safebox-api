FROM jerson/go:1.12 AS builder

ENV WORKDIR ${GOPATH}/src/safebox.jerson.dev/api
WORKDIR ${WORKDIR}

COPY Gopkg.toml Gopkg.lock Makefile ./
RUN make deps

USER root

COPY config.toml-dist config.toml
COPY . .

RUN make build

FROM jerson/base:1.1

LABEL maintainer="jeral17@gmail.com"

ENV BUILDER_PATH /srv/go/src/safebox.jerson.dev/api
ENV WORKDIR /app
WORKDIR ${WORKDIR}

COPY --from=builder ${BUILDER_PATH}/config.toml .
COPY --from=builder ${BUILDER_PATH}/api-server .

EXPOSE 8000
VOLUME ["${WORKDIR}/logs"]

ENTRYPOINT ["/app/api-server"]