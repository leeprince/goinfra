FROM golang AS builder

ARG CGO=0
ENV CGO_ENABLED=${CGO}
ENV GOARCH=amd64
ENV GOOS=linux
ENV GO111MODULE=on

ARG env=dev

WORKDIR /build

COPY . /build
RUN set -ex && \
    cp -f conf_${env}.yaml conf.yaml && \
    go mod tidy && \
    CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} go build -a -o gowebsocket ./main.go

FROM debian

WORKDIR /work

COPY --from=builder /build/gowebsocket .
COPY --from=builder /build/conf.yaml .

VOLUME ["/work/runtime"]

EXPOSE 8080

CMD ["./gowebsocket", "-config=conf.yaml"]