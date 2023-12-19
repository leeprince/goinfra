# pdf操作
---

# 使用 “gopkg.in/gographics/imagick.v2/imagick” 库
> 测试成功后已经打包的镜像：``docker pull leeprince/golang1_19:imagick_v2_test`

## 下载并编译安装 leptonica，用于处理图片
leptonica是一个开源的图像处理库，它提供了一系列用于图像处理和分析的函数和工具。它可以在Linux系统中用于处理和操作图像，包括图像的读取、写入、转换、缩放、旋转、滤波、分割等操作。
具体来说，leptonica可以用于以下方面：

- 图像格式转换：leptonica支持多种常见的图像格式，可以将图像从一种格式转换为另一种格式，以满足不同应用的需求。
- 图像处理：leptonica提供了各种图像处理函数，如图像平滑、边缘检测、二值化、灰度化、直方图均衡化等，可以对图像进行各种处理和增强操作。
- 文字识别：leptonica提供了一些用于文字识别的函数和工具，可以用于文字的分割、识别和提取。
- 图像分割：leptonica提供了一些图像分割算法，可以将图像分割成不同的区域或对象，用于图像分析和处理。

总之，leptonica在Linux系统中是一个功能强大的图像处理库，可以用于各种图像处理和分析任务。
```
wget http://www.leptonica.org/source/leptonica-1.78.0.tar.gz
tar -xzvf leptonica-1.78.0.tar.gz
cd leptonica-1.78.0
./configure
make && make instal
```

## 安装依赖其它pdf、html、图片依赖
```
go get github.com/SebastiaanKlippert/go-wkhtmltopdf
go get github.com/disintegration/imaging
go get github.com/fogleman/gg
go get github.com/ninetwentyfour/go-wkhtmltoimage
go get github.com/panjf2000/ants
go get github.com/pieterclaerhout/go-html
```

## 安装图片处理工具 wkhtmltopdf
根据参考资料，我了解到：

- wkhtmltopdf是一个开源的命令行工具，它可以将HTML页面转换为PDF文档。它使用WebKit渲染引擎来呈现HTML内容，并将其转换为PDF格式。
- wkhtmltox是wkhtmltopdf的一个扩展版本，它不仅可以将HTML页面转换为PDF文档，还可以将HTML页面转换为图像格式（如PNG、JPEG等）。

因此，可以说wkhtmltopdf和wkhtmltox都是用于将HTML页面转换为其他格式（如PDF、图像）的工具。wkhtmltox是wkhtmltopdf的扩展版本，它提供了更多的功能，可以将HTML页面转换为多种图像格式。
总之，wkhtmltopdf和wkhtmltox都是非常有用的工具，可以帮助我们将HTML页面转换为其他格式，以满足不同的需求。

```
sudo yum install wkhtmltopdf
sudo yum install xorg-x11-server-Xvfb
```

解决：wkhtmltopdf：cannot connect to X server问题：
在 /usr/bin/ 目录下生成脚本 wkhtmltopdf.sh 并写入命令：
```
xvfb-run -a --server-args="-screen 0, 1024x768x24" /usr/bin/wkhtmltopdf -q $*
chmod a+x /usr/bin/wkhtmltopdf.sh
ln -s /usr/bin/wkhtmltopdf.sh /usr/local/bin/wkhtmltopdf
```
> 注：wkhtmltopdf是wkhtmltox的低版本，在ubuntu下，如果安装了wkhtmltopdf，再安装wkhtmltox，相关动态库自动忽略。故没有安装wkhtmltopdf


## 本地安装图片处理工具 wkhtmltox
安装包下载地址：https://github.com/adrg/go-wkhtmltopdf/wiki/Install-on-Linux
> 注：我是在centos下测试的，所以下载了wkhtmltox-0.12.6-1.centos7.x86_64.rpm
> 注：ubuntu可能还需要其它依赖：https://blog.csdn.net/qq_15378385/article/details/107456644
```
wget https://github.com/adrg/go-wkhtmltopdf/wiki/Install-on-Linux
sudo  yum localinstall wkhtmltox-0.12.6-1.centos7.x86_64.rpm
go get github.com/adrg/go-wkhtmltopdf
```


## 安装 ImageMagick、ghostscript、libmagickwand-dev
> ImageMagick
ImageMagick 是一个强大的图像处理库，它可以用于创建、编辑、组合或转换各种格式的位图图像。它可以读取、转换和写入多达200种以上的图像格式，包括PNG、JPEG、JPEG-2000、GIF、WebP、Postscript、PDF等。
ImageMagick的主要功能包括：

- 格式转换：可以将图像从一种格式转换为另一种格式。
- 图像编辑：可以调整图像的大小、模糊、裁剪、去除斑点、旋转、应用各种特效等。
- 图像绘制：可以在图像上添加线条、文字、多边形、椭圆、曲线等。
- 图像计算：可以进行图像的数学运算，如加、减、乘、除等。
- 动画处理：可以创建、编辑和处理各种动画图像。
- 处理大图像：可以处理大于内存大小的图像。

ImageMagick是开源的，可以在各种操作系统上使用，包括Linux、Windows、Mac OS X等。它提供了命令行工具和开发库，可以通过编程语言（如C、C++、Python、Ruby等）调用其功能。

> leptonica、ImageMagick、ghostscript作用是什么，关系是什么
根据参考资料，我了解到：
- leptonica是一个开源的图像处理库，它提供了一系列用于图像处理和分析的函数和工具。它可以用于图像的读取、写入、转换、缩放、旋转、滤波、分割等操作，以及文字识别和图像分析等任务。
- ImageMagick是一个强大的图像处理库，它可以用于创建、编辑、组合或转换各种格式的位图图像。它支持多种图像格式，并提供了丰富的功能，包括格式转换、图像编辑、图像绘制、图像计算、动画处理等。
- ghostscript是一个用于解释和渲染PostScript和PDF文件的开源引擎。它可以将PostScript和PDF文件转换为图像文件，也可以进行一些图像处理操作，如合并、分割、压缩等。

关于它们的关系，leptonica、ImageMagick和ghostscript是三个独立的图像处理库，它们有一些共同的功能，但也有一些区别。leptonica主要关注图像处理和分析，ImageMagick提供了更全面的图像处理功能，而ghostscript专注于解释和渲染PostScript和PDF文件。
在某些情况下，这三个库可以结合使用。例如，可以使用leptonica读取和处理图像，然后使用ImageMagick进行格式转换或进一步的图像处理，最后使用ghostscript将结果转换为PDF或其他格式。
总结起来，leptonica、ImageMagick和ghostscript都是用于图像处理的库，它们各自有不同的功能和特点，可以根据具体需求选择合适的库进行使用。

> ImageMagick 和 MagickWand 作用是什么，关系是什么
ImageMagick 是一个开源的软件套件，用于创建、编辑、合成和转换位图图像。它支持众多图片格式，并提供了丰富的命令行工具，可以用来处理图像，比如改变尺寸、旋转、裁剪、添加特效等等。
MagickWand 是 ImageMagick 的一个编程接口，允许开发者使用 C、C++、Python、PHP 等编程语言调用 ImageMagick 的功能，实现图像处理、编辑和生成。
MagickWand 可以被视为 ImageMagick 的一个库或接口，使得开发者能够在自己的应用程序中方便地利用 ImageMagick 的强大功能。

```
apt-get update
apt-get remove imagemagick
apt-get install imagemagick
apt-get install  ghostscript


# 解决报错：Package MagickWand was not found in the pkg-config search path.
apt-get install libmagickwand-dev
```

## 遇到的其他问题
### 关于 authorized 的问题
```
ImageMagick "not authorized" PDF errors
```
```
ERROR_POLICY: attempt to perform an operation not allowed by the security policy `PDF' @ error/constitute.c/IsCoderAuthorized/421
```

参考：https://cromwell-intl.com/open-source/pdf-not-authorized.html
参考：https://stackoverflow.com/questions/52998331/imagemagick-security-policy-pdf-blocking-conversion

解决操作：
实际上原始的`ImageMagick-6`在复制到容器或者pod的`/etc/ImageMagick-6`后，
在上面操作过程中`/etc/ImageMagick-6/policy.xml`会被重写，
所以解决仅需重新将`ImageMagick-6/policy.xml`复制到`/etc/ImageMagick-6/policy.xml`。

## 基于镜像构建
1. cd imagick_v2_test
2. 构建镜像。目前已经测试成功，并上传到 docker hub，无需再构建镜像，仅需直接拉取：`docker pull leeprince/golang1_19:imagick_v2`。感兴趣可以重新在本地构建
```
docker build -t leeprince/golang1_19:imagick_v2
```
3. 基于镜像构建 imagick_v2 的运行环境的容器
```
docker run -d -v .:/www -p 8081:80 --name imagick_v2_1 leeprince/golang1_19:imagick_v2
```
4. 进入容器
```
docker exec -it imagick_v2_1 bash
```

# pkg-config
pkg-config是一个在编译过程中帮助程序找到库文件的工具。当你在编译一个程序时，如果它依赖于某些库文件，那么你需要告诉编译器这些库文件的位置。这就是pkg-config的作用。
具体来说，pkg-config可以做以下几件事：

- 提供库文件的路径：pkg-config可以告诉编译器库文件的位置，这样编译器就可以找到这些文件并将它们链接到你的程序中。
- 提供库文件的版本信息：pkg-config可以检查库文件的版本，确保你的程序使用的是正确的版本。
- 提供编译参数：pkg-config可以提供编译参数，这些参数可以告诉编译器如何编译你的程序。

在Linux系统中，pkg-config是一个非常重要的工具，它可以帮助你更容易地编译和安装程序。

# build-essential
build-essential 是针对 Debian 和 Ubuntu 等基于 Debian 的 Linux 发行版的软件包。它实际上是一个包含了用于编译软件的基本工具集合。

主要组成部分包括：
- GNU Compiler Collection (GCC)：这是一个非常重要的编译器集合，支持多种编程语言，如C、C++、Objective-C、Fortran等。它是开发C语言和许多其他语言程序的主要工具。
- GNU Make：是一个构建自动化工具，用于管理源代码的编译过程。它读取一个名为 Makefile 的文件，其中包含了关于如何编译和链接代码的规则。
- dpkg-dev：包含了一些工具和库，用于构建和管理Debian软件包（.deb文件）。

这些工具的集合使得在Linux系统上进行软件开发更加方便，因为它们提供了编译和构建程序所需的基本组件。使用这些工具，开发者可以编译源代码、创建可执行文件，并将软件打包成易于分发和安装的软件包。

# 使用“gopkg.in/gographics/imagick.v3/imagick”
> 官方文档：https://pkg.go.dev/gopkg.in/gographics/imagick.v3@v3.5.0#section-readme
> github仓库：https://github.com/gographics/imagick/tree/v3.5.0
> 官方示例：https://github.com/gographics/imagick/tree/v3.5.0/examples

## 须知
o Imagick is a Go bind to ImageMagick's MagickWand C API.

We support two compatibility branches:
```
im-7   (tag v3.x.x): 7.0.8-41 <= ImageMagick <= 7.x
master (tag v2.x.x): 6.9.1-7  <= ImageMagick <= 6.9.9-35
legacy (tag v1.x.x): 6.7.x    <= ImageMagick <= 6.8.9-10
```


They map, respectively, through gopkg.in:
```
gopkg.in/gographics/imagick.v3/imagick
gopkg.in/gographics/imagick.v2/imagick
gopkg.in/gographics/imagick.v1/imagick
```

## 注意
1. 官方示例：`https://github.com/gographics/imagick/tree/v3.5.0/examples/docker` 的示例 `Dockerfile` 文件中的<IMAGEMAGICK_VERSION>参数一定要匹配上 <## 须知> 的版本。
- 官方原版
```
FROM golang:1.11
...

ENV IMAGEMAGICK_VERSION=7.0.8-11

...

WORKDIR /go/projects/resizer
COPY . .

RUN go install
CMD /go/bin/resizer
```

- 修改后
```
FROM golang:1.19

...

ENV IMAGEMAGICK_VERSION=7.1.1-23

...

WORKDIR /www

VOLUME ["www"]

EXPOSE 80

CMD tail -f /dev/null
```

## 基于镜像构建
1. cd imagick_v3_test
2. 构建镜像。目前已经测试成功，并上传到 docker hub，无需再构建镜像，仅需直接拉取：`docker pull leeprince/golang1_19:imagick_v3_7.1.1-23`。感兴趣可以重新在本地构建
```
docker build -t leeprince/golang1_19:imagick_v3_7.1.1-23
```
3. 基于镜像构建 imagick_v3 的运行环境的容器
```
docker run -d -v .:/www -p 8082:80 --name imagick_v3_1 leeprince/golang1_19:imagick_v3_7.1.1-23
```
4. 进入容器
```
docker exec -it imagick_v3_1 bash
```

### 完整 Dockerfile
```
FROM golang:1.19

# Ignore APT warnings about not having a TTY
ENV DEBIAN_FRONTEND noninteractive

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

# Install ImageMagick deps
RUN apt-get -q -y install libjpeg-dev libpng-dev libtiff-dev \
    libgif-dev libx11-dev ghostscript --no-install-recommends

ENV IMAGEMAGICK_VERSION=7.1.1-23

RUN cd && \
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

WORKDIR /www

VOLUME ["www"]

EXPOSE 80

CMD tail -f /dev/null
```

# 使用 "github.com/davidbyttow/govips/v2/vips"
> 官方文档：https://github.com/davidbyttow/govips
> 关于libvips：https://github.com/libvips/libvips

## 基于镜像构建
1. cd govips_test
2. 构建镜像。目前已经测试成功，并上传到 docker hub，无需再构建镜像，仅需直接拉取：`docker pull leeprince/golang1_19:govips_v2`。感兴趣可以重新在本地构建
```
docker build -t leeprince/golang1_19:govips_v2
```
3. 基于镜像构建 govips_v2 的运行环境的容器
```
docker run -d -v .:/www -p 8083:80 --name govips_v2_1 leeprince/golang1_19:govips_v2
```
4. 进入容器
```
docker exec -it govips_v2_1 bash
```

### 完整 Dockerfile
> 官方：https://github.com/davidbyttow/govips/tree/master/build
> 后来才发现官方有基于ubuntu构建的环境，发现原来 `libvips` 可以直接通过 `apt-get install -y libvips-dev`，并包含其他依赖，如：imagemagick、libmagickwand、

```
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

RUN apt-get -q -y install meson

RUN tar -xvzf v${LIBVIPS_VERSION}.tar.gz && \
    cd libvips* && \
    meson setup build-dir --prefix=${LIBVIPS_DIR_PREFIX} && \
    cd build-dir && \
    ninja && \
    ninja test && \
    ninja install

WORKDIR /www

VOLUME ["www"]

EXPOSE 80

CMD tail -f /dev/null
```


## 问题
### libjpeg-turbo8-dev
`Package libjpeg-turbo8-dev is not available, but is referred to by another package.`

解决：
```
libjpeg-dev 代替 libjpeg-turbo8-dev
```

### meson
安装 libvips 时需要使用 `meson`, 所以需要提前安装
```
apt-get install -y meson
```

### error: cannot find builtin CJK font
```text
error: cannot find builtin CJK font
warning: unrecoverable error; ignoring rest of page
```
效果：pdf 文件转图片时，除了签章，中文部分都没有显示在图片上

解决：
```text
apt-get install -y fonts-noto-cjk
```
> fonts-noto-cjk 包含 cjk 语言包，但是安装后不报错了，但是问题并没有解决，只是不报错了！

继续解决：
```text
apt-get install -y fonts-arphic-ukai
```
> fonts-arphic-ukai 与 fonts-noto-cjk 一样的情况

最终解决：
```
apt-get install -y fonts-arphic-uming
```

建议：一次性把这些语言包都装上
```
apt-get -q -y install fonts-arphic-uming fonts-arphic-ukai fonts-noto-cjk  --no-install-recommends
```

## 语言包
- fonts-noto-cjk: Noto 字体系列，包括中文、日文、韩文等各种语言的字体。
- fonts-arphic-ukai: AR PL UKai 中文字体。
- fonts-arphic-uming: AR PL UMing 中文字体。


## lsb-release
lsb-release 是一个 Linux Standard Base（LSB）的工具，用于显示有关 Linux 发行版的信息。它通常位于大多数基于 Linux 的系统中，并提供了一种简单的方式来确定当前系统所使用的 Linux 发行版信息。

通过运行 lsb_release -a 命令，你可以获取包括发行版代号、发行版描述、发行版号码以及其他有关 Linux 发行版的信息。例如：
```
lsb_release -a
```

这将显示类似以下的输出：
```text
Distributor ID: Ubuntu
Description:    Ubuntu 20.04.3 LTS
Release:        20.04
Codename:       focal
```

## devscripts
devscripts 是一个为 Debian 开发者设计的软件包，提供了一系列用于简化 Debian 软件包开发和维护的工具集。

这些工具涵盖了各种开发任务，包括但不限于：
- 包构建和管理： 提供了一些工具来构建 Debian 软件包、进行版本控制和管理包的发布。
- 质量保证： 包含了用于检查 Debian 软件包是否符合 Debian 政策的工具，例如 checkbashisms 用于检查脚本是否符合 Bashism 的规范。
- Bug 跟踪和维护： 提供了与 Debian 软件包维护相关的工具，可以帮助跟踪和管理软件包的错误报告和修复。
- 文档生成： 提供了生成 Debian 软件包相关文档的工具。

一些 devscripts 中常用的工具包括：
- debuild: 用于构建 Debian 软件包的工具。
- debdiff: 用于比较 Debian 软件包之间的差异。
- uscan: 用于从网络上自动下载更新的工具。
- dch: 用于在 Debian 软件包的变更历史中记录修改信息。

总的来说，devscripts 是为了简化 Debian 软件包开发者的工作而创建的工具集，使他们能够更轻松地构建、维护和管理 Debian 软件包。

这些信息对于脚本、安装软件包、确保软件兼容性以及识别正在运行的操作系统非常有用。LSB是一种标准化的规范，可以帮助确保不同的 Linux 发行版在某些方面保持一致性，让开发者能够更轻松地编写跨发行版兼容的代码。

## dput
dput 是一个用于将本地打包好的 Debian 软件包上传到远程 Debian 软件仓库的工具。它允许开发者将他们构建好的软件包上传到合适的 Debian 软件仓库，使得其他用户可以访问并安装这些软件包。

一般来说，dput 可以实现以下功能：
- 上传软件包： 开发者通过 dput 命令将本地打包好的 Debian 软件包上传到远程的 Debian 软件仓库。
- 指定目标仓库： dput 允许指定上传软件包的目标仓库，这样开发者可以将软件包上传到正确的软件源，以便用户能够访问和安装这些软件包。
- 验证和授权： 在上传软件包之前，dput 会执行一些验证步骤，确保软件包满足特定的要求，并可能需要开发者提供相应的授权信息。
- 简化上传流程： 对于 Debian 软件包的开发者来说，dput 提供了一个简化的上传流程，使得他们可以更轻松地将软件包发布到合适的仓库中。

总的来说，dput 是 Debian 发行版中用于将本地构建好的软件包上传到远程仓库的方便工具，有助于软件包的分发和安装。

## nvi
nvi 是一个文本编辑器，它是 vi 编辑器的一个变种。vi 是一款经典的、强大的、使用广泛的文本编辑器，存在于许多类 Unix 系统中，而 nvi 是 vi 的改进版本。

nvi 的作用与 vi 类似，主要用于在命令行界面下编辑文本文件。它具有以下特点和功能：
- 模式编辑： 像 vi 一样，nvi 也是一个模式编辑器，具有不同的模式（命令模式、插入模式、可视模式等）。在不同模式下，可以执行不同的编辑操作。
- 强大的编辑功能： 支持文本搜索、替换、复制、粘贴等常见的编辑操作，以及一些高级的编辑功能。
- 轻量级和快速： nvi 在保持与 vi 兼容的同时，也注重效率和快速响应。
- 配置灵活： 用户可以根据自己的喜好进行配置和定制，通过设置编辑器的选项和配置文件来调整编辑器的行为。
- 遵循 POSIX 标准： nvi 遵循 POSIX 标准，因此在符合标准的系统上具有良好的可移植性。

总的来说，nvi 是一个面向命令行环境的文本编辑器，适合于对文本文件进行快速、高效的编辑操作，同时也保留了 vi 的许多特性和使用习惯。

## add-apt-repository -y ppa:tonimelisma/ppa 的作用

这个命令 add-apt-repository -y ppa:tonimelisma/ppa 在 Ubuntu 或基于 Debian 的系统中用于向软件源列表中添加一个 PPA（Personal Package Archive，个人软件包存档）。

具体来说：
- add-apt-repository 是一个用于添加软件源的工具。
- -y 参数表示自动确认，即在添加软件源时，不需要用户手动确认。
- ppa:tonimelisma/ppa 是要添加的 PPA 的地址。tonimelisma 是 PPA 的所有者或用户名，ppa 是个人软件包存档的名称。

当你运行这个命令时，它会将 ppa:tonimelisma/ppa 添加到你系统的软件源列表中。PPA 提供了一种在 Ubuntu 或基于 Debian 的系统上安装不在官方仓库中的软件的方法。添加 PPA 后，系统就能够从这个 PPA 中获取软件包并进行安装。

请注意，在使用任何 PPA 之前，确保你信任 PPA 的所有者，因为它们并非官方维护的软件源。添加不可信任或未经审核的 PPA 可能会导致系统安全问题或不稳定性。

## 示例
官方示例，有些方法（2023-12-19）已经废弃，手动替换即可
```golang
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	image1, err := vips.NewImageFromFile("input.jpg")
	checkError(err)

	// Rotate the picture upright and reset EXIF orientation tag
	err = image1.AutoRotate()
	checkError(err)

	ep := vips.NewDefaultJPEGExportParams()
	image1bytes, _, err := image1.Export(ep)
	checkError(err)
	err = ioutil.WriteFile("output.jpg", image1bytes, 0644)
	checkError(err)
}
```

```golang

package main

import (
	"fmt"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	image1, err := vips.NewImageFromFile("0001.png")
	checkError(err)

	// Rotate the picture upright and reset EXIF orientation tag
	err = image1.AutoRotate()
	checkError(err)

	image1bytes, _, err := image1.ExportNative()
	checkError(err)

	dirPath := "."
	fileName := fmt.Sprintf("resize_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)

	ok, err := WriteFile(dirPath, filePath, image1bytes, false)
	if !ok {
		fmt.Println("fileutil.WriteFile !ok")
		return
	}
	checkError(err)
	fmt.Println("successful, filepath:", filePath)
}
```

# imagick_v3 与 govips 对比
在生成速度、生成质量、占用空间上 govips 更优！

imagick_v3 的优势是很灵活，方便自定义一些属性！
