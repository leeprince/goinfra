FROM leeprince/golang1_19:imagick_v3_7.1.1-23_v3

ARG CGO=1
ENV CGO_ENABLED=${CGO}
ENV GOARCH=amd64
ENV GOOS=linux
ENV GO111MODULE=on

ARG env=dev

WORKDIR /work

COPY . /work

# 创建目录用于存放 SSH 认证文件
RUN mkdir -p /root/.ssh
COPY common/gitssh/* /root/.ssh
# 设置 SSH 认证文件的权限
RUN chmod 600 /root/.ssh/config && \
    chmod 600 /root/.ssh/golang-git-base && \
    chmod 644 /root/.ssh/golang-git-base.pub

RUN ls /root/.ssh -la && \
    git config --global --add url.git@gitlab.yunpiaoer.com:.insteadof "http://gitlab.yunpiaoer.com" && \
    git config --global --add url.git@gitlab.yewifi.com:.insteadof "http://gitlab.yewifi.com" && \
    git config --global --add url.git@gitlab.yewifi.com:.insteadof "https://gitlab.yewifi.com"

RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GONOPROXY=*.yewifi.com,*.yunpiaoer.com,*.yunpiaoer.cn

RUN set -ex && \
    echo "---------pwd:" && pwd && echo "---------ls:" && ls && \
    cp -f format-ctl/configs/conf_${env}.yaml format-ctl/configs/conf.yaml

RUN go mod download

RUN CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOOS=${GOOS} go build -a -o format-ctl/format-ctl format-ctl/cmd/main.go

RUN mkdir -p /work/logs
VOLUME ["/work/logs"]

EXPOSE 8080
EXPOSE 50009

CMD ["./format-ctl/format-ctl", "-conf=./format-ctl/configs/conf.yaml"]