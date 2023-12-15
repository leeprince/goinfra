# pdf操作
---

# 使用“gopkg.in/gographics/imagick.v2/imagick”

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

解决操作：实际上原始的`ImageMagick-6`在复制到容器或者pod的`/etc/ImageMagick-6`后，
在上面操作过程中`/etc/ImageMagick-6/policy.xml`会被重写，
所以解决仅需重新将`ImageMagick-6/policy.xml`复制到`/etc/ImageMagick-6/policy.xml`。

# pkg-config
pkg-config是一个在编译过程中帮助程序找到库文件的工具。当你在编译一个程序时，如果它依赖于某些库文件，那么你需要告诉编译器这些库文件的位置。这就是pkg-config的作用。
具体来说，pkg-config可以做以下几件事：

- 提供库文件的路径：pkg-config可以告诉编译器库文件的位置，这样编译器就可以找到这些文件并将它们链接到你的程序中。
- 提供库文件的版本信息：pkg-config可以检查库文件的版本，确保你的程序使用的是正确的版本。
- 提供编译参数：pkg-config可以提供编译参数，这些参数可以告诉编译器如何编译你的程序。

在Linux系统中，pkg-config是一个非常重要的工具，它可以帮助你更容易地编译和安装程序。