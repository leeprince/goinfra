FROM golang:1.19

# Ignore APT warnings about not having a TTY
ENV DEBIAN_FRONTEND noninteractive

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

# Install ImageMagick deps
RUN apt-get -q -y install libjpeg-dev libpng-dev libtiff-dev \
    libgif-dev libx11-dev ghostscript --no-install-recommends

# pdf 生成的图片有乱码或者存在空白
# 发现加上 fonts-arphic-uming 依赖之后会影响处理小数点的问题，会出现小数点模糊和偏移的情况
RUN apt-get -q -y install fonts-arphic-uming fonts-arphic-ukai fonts-noto-cjk --no-install-recommends

ENV IMAGEMAGICK_VERSION=7.1.1-23

RUN set -ex && \
	wget https://github.com/ImageMagick/ImageMagick/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
	tar xvzf ${IMAGEMAGICK_VERSION}.tar.gz && \
	cd ImageMagick* && \
	./configure \
	    --without-magick-plus-plus \
	    --without-perl \
	    --disable-openmp \
	    --with-gvc=no \
	    --disable-docs && \
	make -j$(nproc) && make install && \
	ldconfig /usr/local/lib

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /www

VOLUME ["/www"]

EXPOSE 80

CMD tail -f /dev/null