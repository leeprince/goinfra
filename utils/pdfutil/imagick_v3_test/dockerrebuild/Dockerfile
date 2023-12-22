# syntax=docker/dockerfile:1
FROM leeprince/golang1_19:imagick_v3_7.1.1-23

ARG CGO=1
ENV CGO_ENABLED=${CGO}
ENV GOARCH=amd64
ENV GOOS=linux
ENV GO111MODULE=on

ARG env=dev

WORKDIR /work

COPY . /work

RUN set -ex && \
    go mod download && \
    CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} go build -a -o format-ctl main.go

RUN mkdir -p /work/logs
VOLUME ["/work/logs"]

EXPOSE 8080
EXPOSE 50009

CMD ["./format-ctl"]
