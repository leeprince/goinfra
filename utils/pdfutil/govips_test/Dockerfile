FROM golang:1.19

# Ignore APT warnings about not having a TTY
ENV DEBIAN_FRONTEND noninteractive

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

# Install ImageMagick deps
RUN apt-get -q -y install libvips-dev libglib2.0-dev libexpat1-dev \
    libtiff5-dev libgsf-1-dev \
    libjpeg-dev libpng-dev libtiff-dev \
    libgif-dev libx11-dev ghostscript --no-install-recommends

ENV LIBVIPS_VERSION=8.15.1
ENV LIBVIPS_DIR_PREFIX=/etc/vips

RUN wget https://github.com/libvips/libvips/archive/refs/tags/v${LIBVIPS_VERSION}.tar.gz

RUN apt-get -q -y  install meson

RUN tar -xvzf v${LIBVIPS_VERSION}.tar.gz && \
    cd libvips* && \
    meson setup build-dir --prefix=${LIBVIPS_DIR_PREFIX} && \
    cd build-dir && \
    ninja && \
    ninja test && \
    ninja install

apt-get -q -y  install fonts-arphic-uming fonts-arphic-ukai fonts-noto-cjk  --no-install-recommends

WORKDIR /www

VOLUME ["www"]

EXPOSE 80

CMD tail -f /dev/null