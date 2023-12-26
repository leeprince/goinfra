# pdf操作

---

# 实操“gopkg.in/gographics/imagick.v2/imagick” 库将 pdf 转成图片的过程

`gopkg.in/gographics/imagick.v2/imagick` 是一个 Go 语言的库，它提供了对 ImageMagick 的 MagickWand C API 的绑定。ImageMagick 是一个强大的图像处理库，可以用来创建、编辑、组合或转换各种格式的图像。使我们可以在 Go 语言中方便地使用 ImageMagick 的功能，例如：

- 读取和写入各种图像格式，包括 GIF、JPEG、PNG、PDF、WebP 等。
- 调整图像的大小、裁剪、旋转、模糊、锐化等。
- 提取图像的元数据，例如 EXIF 信息。
- 对图像进行色彩管理和色彩空间转换。
- 创建和处理图像序列（例如 GIF 动画）。

可以看到库很强大，要想使用它，掌握各种方法是一方面，另外一方便是搭建好库本身依赖的各种包



> 解决所有问题后打包好的镜像：`docker pull leeprince/golang1_19:imagick_v2`



`gopkg.in/gographics/imagick.v2/imagick` 是一个 Go 语言的库，它提供了对 ImageMagick 的 MagickWand C API 的绑定。ImageMagick 是一个强大的图像处理库，可以用来创建、编辑、组合或转换各种格式的图像。使我们可以在 Go 语言中方便地使用 ImageMagick 的功能，例如：

- 读取和写入各种图像格式，包括 GIF、JPEG、PNG、PDF、WebP 等。
- 调整图像的大小、裁剪、旋转、模糊、锐化等。
- 提取图像的元数据，例如 EXIF 信息。
- 对图像进行色彩管理和色彩空间转换。
- 创建和处理图像序列（例如 GIF 动画）。

可以看到库很强大，要想使用它，掌握各种方法是一方面，另外一方便是搭建好库本身依赖的各种包



## 一、安装 ImageMagick、ghostscript、libmagickwand-dev【必选】

```
apt-get update
apt-get remove imagemagick
apt-get install imagemagick
apt-get install ghostscript


# 解决报错：Package MagickWand was not found in the pkg-config search path.
apt-get install libmagickwand-dev
```



## 二、下载并编译安装 leptonica，用于处理图片【可选】

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



## 三、安装图片处理工具 wkhtmltopdf、wkhtmltox【可选】

debian 操作系统安装方式参考：[服务器 Ubuntu18.4 安装 wkhtmltopdf（debian操作系统都适用）](http://mp.weixin.qq.com/s?__biz=MzkyMzYyNjQxOQ==&mid=2247484108&idx=1&sn=8adc24474699d6ee1d23d1e827f47bff&chksm=c1e37c10f694f506864f6e824871d0cf401592ba0807d53e276652a5a1997aeb824978fc9879&scene=21#wechat_redirect)



## 五、遇到的其他问题

### 关于 authorized 的问题

```
ImageMagick "not authorized" PDF errors
```

```
ERROR_POLICY: attempt to perform an operation not allowed by the security policy `PDF' @ error/constitute.c/IsCoderAuthorized/421
```



解决操作：

> 参考资料：https://cromwell-intl.com/open-source/pdf-not-authorized.html
> 参考资料：https://stackoverflow.com/questions/52998331/imagemagick-security-policy-pdf-blocking-conversion

实际上原始的`ImageMagick-6`在复制到容器或者pod的`/etc/ImageMagick-6`后，
在上面操作过程中`/etc/ImageMagick-6/policy.xml`会被重写，
所以解决仅需重新将`ImageMagick-6/policy.xml`复制到`/etc/ImageMagick-6/policy.xml`。

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE policymap [
  <!ELEMENT policymap (policy)*>
  <!ATTLIST policymap xmlns CDATA #FIXED ''>
  <!ELEMENT policy EMPTY>
  <!ATTLIST policy xmlns CDATA #FIXED '' domain NMTOKEN #REQUIRED
    name NMTOKEN #IMPLIED pattern CDATA #IMPLIED rights NMTOKEN #IMPLIED
    stealth NMTOKEN #IMPLIED value CDATA #IMPLIED>
]>
<!--
  Configure ImageMagick policies.

  Domains include system, delegate, coder, filter, path, or resource.

  Rights include none, read, write, execute and all.  Use | to combine them,
  for example: "read | write" to permit read from, or write to, a path.

  Use a glob expression as a pattern.

  Suppose we do not want users to process MPEG video images:

    <policy domain="delegate" rights="none" pattern="mpeg:decode" />

  Here we do not want users reading images from HTTP:

    <policy domain="coder" rights="none" pattern="HTTP" />

  The /repository file system is restricted to read only.  We use a glob
  expression to match all paths that start with /repository:

    <policy domain="path" rights="read" pattern="/repository/*" />

  Lets prevent users from executing any image filters:

    <policy domain="filter" rights="none" pattern="*" />

  Any large image is cached to disk rather than memory:

    <policy domain="resource" name="area" value="1GP"/>

  Use the default system font unless overwridden by the application:

    <policy domain="system" name="font" value="/usr/share/fonts/favorite.ttf"/>

  Define arguments for the memory, map, area, width, height and disk resources
  with SI prefixes (.e.g 100MB).  In addition, resource policies are maximums
  for each instance of ImageMagick (e.g. policy memory limit 1GB, -limit 2GB
  exceeds policy maximum so memory limit is 1GB).

  Rules are processed in order.  Here we want to restrict ImageMagick to only
  read or write a small subset of proven web-safe image types:

    <policy domain="delegate" rights="none" pattern="*" />
    <policy domain="filter" rights="none" pattern="*" />
    <policy domain="coder" rights="none" pattern="*" />
    <policy domain="coder" rights="read|write" pattern="{GIF,JPEG,PNG,WEBP}" />
-->
<policymap>
  <policy domain="delegate" rights="none" pattern="URL" />
  <policy domain="delegate" rights="none" pattern="HTTPS" />
  <policy domain="delegate" rights="none" pattern="HTTP" />
  <policy domain="path" rights="none" pattern="@*"/>
</policymap>
```



## 六、相关概念

### （一）ImageMagick

ImageMagick 是一个强大的图像处理库，它可以用于创建、编辑、组合或转换各种格式的位图图像。它可以读取、转换和写入多达200种以上的图像格式，包括PNG、JPEG、JPEG-2000、GIF、WebP、Postscript、PDF等。
ImageMagick的主要功能包括：

- 格式转换：可以将图像从一种格式转换为另一种格式。
- 图像编辑：可以调整图像的大小、模糊、裁剪、去除斑点、旋转、应用各种特效等。
- 图像绘制：可以在图像上添加线条、文字、多边形、椭圆、曲线等。
- 图像计算：可以进行图像的数学运算，如加、减、乘、除等。
- 动画处理：可以创建、编辑和处理各种动画图像。
- 处理大图像：可以处理大于内存大小的图像。

ImageMagick是开源的，可以在各种操作系统上使用，包括Linux、Windows、Mac OS X等。它提供了命令行工具和开发库，可以通过编程语言（如C、C++、Python、Ruby等）调用其功能。



### （二）ImageMagick 和 MagickWand 作用是什么，关系是什么

ImageMagick 是一个开源的软件套件，用于创建、编辑、合成和转换位图图像。它支持众多图片格式，并提供了丰富的命令行工具，可以用来处理图像，比如改变尺寸、旋转、裁剪、添加特效等等。
MagickWand 是 ImageMagick 的一个编程接口，允许开发者使用 C、C++、Python、PHP 等编程语言调用 ImageMagick 的功能，实现图像处理、编辑和生成。
MagickWand 可以被视为 ImageMagick 的一个库或接口，使得开发者能够在自己的应用程序中方便地利用 ImageMagick 的强大功能。



### （三）leptonica、ImageMagick、ghostscript作用是什么，关系是什么

根据参考资料，我了解到：

- leptonica是一个开源的图像处理库，它提供了一系列用于图像处理和分析的函数和工具。它可以用于图像的读取、写入、转换、缩放、旋转、滤波、分割等操作，以及文字识别和图像分析等任务。
- ImageMagick是一个强大的图像处理库，它可以用于创建、编辑、组合或转换各种格式的位图图像。它支持多种图像格式，并提供了丰富的功能，包括格式转换、图像编辑、图像绘制、图像计算、动画处理等。
- ghostscript是一个用于解释和渲染PostScript和PDF文件的开源引擎。它可以将PostScript和PDF文件转换为图像文件，也可以进行一些图像处理操作，如合并、分割、压缩等。

关于它们的关系，leptonica、ImageMagick和ghostscript是三个独立的图像处理库，它们有一些共同的功能，但也有一些区别。leptonica主要关注图像处理和分析，ImageMagick提供了更全面的图像处理功能，而ghostscript专注于解释和渲染PostScript和PDF文件。
在某些情况下，这三个库可以结合使用。

例如，可以使用leptonica读取和处理图像，然后使用ImageMagick进行格式转换或进一步的图像处理，最后使用ghostscript将结果转换为PDF或其他格式。

总结起来，leptonica、ImageMagick和ghostscript都是用于图像处理的库，它们各自有不同的功能和特点，可以根据具体需求选择合适的库进行使用。



### （四）pkg-config

pkg-config是一个在编译过程中帮助程序找到库文件的工具。当你在编译一个程序时，如果它依赖于某些库文件，那么你需要告诉编译器这些库文件的位置。这就是pkg-config的作用。
具体来说，pkg-config可以做以下几件事：

- 提供库文件的路径：pkg-config可以告诉编译器库文件的位置，这样编译器就可以找到这些文件并将它们链接到你的程序中。
- 提供库文件的版本信息：pkg-config可以检查库文件的版本，确保你的程序使用的是正确的版本。
- 提供编译参数：pkg-config可以提供编译参数，这些参数可以告诉编译器如何编译你的程序。

在Linux系统中，pkg-config是一个非常重要的工具，它可以帮助你更容易地编译和安装程序。



### （五）build-essential

build-essential 是针对 Debian 和 Ubuntu 等基于 Debian 的 Linux 发行版的软件包。它实际上是一个包含了用于编译软件的基本工具集合。

主要组成部分包括：

- GNU Compiler Collection (GCC)：这是一个非常重要的编译器集合，支持多种编程语言，如C、C++、Objective-C、Fortran等。它是开发C语言和许多其他语言程序的主要工具。
- GNU Make：是一个构建自动化工具，用于管理源代码的编译过程。它读取一个名为 Makefile 的文件，其中包含了关于如何编译和链接代码的规则。
- dpkg-dev：包含了一些工具和库，用于构建和管理Debian软件包（.deb文件）。

这些工具的集合使得在Linux系统上进行软件开发更加方便，因为它们提供了编译和构建程序所需的基本组件。使用这些工具，开发者可以编译源代码、创建可执行文件，并将软件打包成易于分发和安装的软件包。



## 七、测试代码

```
package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v2/imagick"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/15 14:12
 * @Desc:
 */

func main() {
	pdfBytes, err := ReadFile(".", "0001.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}

	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)

	imageByte, ok := CreateImage(pdfBytes, "jpeg")
	if !ok {
		fmt.Println("CreateImage !ok")
		return
	}

	ok, err = WriteFile(dirPath, filePath, imageByte, false)
	if !ok {
		fmt.Println("fileutil.WriteFile !ok")
		return
	}
	checkError(err)
	fmt.Println("successful, filepath:", filePath)
}

func CreateImage(data []byte, toImageType string) ([]byte, bool) {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	//sourceImagePath := getSourceImageForCover(filepath.Dir(pathNoExtension))
	mw := imagick.NewMagickWand()

	err = mw.ReadImageBlob(data)
	if err != nil {
		fmt.Println("[CreateImage] ReadImageBlob err:", err)
		return nil, false
	}


	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)
	mw.SetImageFormat(toImageType)


	content := mw.GetImageBlob()

	return content, true
}


func ReadFileBytesByUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Errorf("resp.StatusCode != http.StatusOK")
		return nil, err
	}

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "0" {
		err = errors.Errorf("contentLength == 0")
		return nil, err
	}

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll")
		return nil, err
	}
	if fileLen := len(fileBytes); fileLen <= 0 {
		err = errors.Errorf("fileBytes len 0")
		return nil, err
	}

	return fileBytes, nil
}

// 写入数据到文件
func WriteFile(dirPath, filename string, data []byte, isAppend bool) (ok bool, err error) {
	filePath := filepath.Join(dirPath, filename)
	if _, ok = CheckFileDirExist(filePath); !ok {
		// 创建目录
		err = os.MkdirAll(dirPath, os.ModePerm)
		if ok = os.IsNotExist(err); ok {
			err = errors.New("创建文件目录错误")
			return
		}
	}

	flag := os.O_CREATE | os.O_RDWR
	if isAppend {
		flag = flag | os.O_APPEND
	}
	fs, fErr := os.OpenFile(filePath, flag, 0666)
	if fErr != nil {
		err = fErr
		return
	}
	defer fs.Close()

	// 创建带有缓冲区的Writer对象
	writer := bufio.NewWriter(fs)
	// 写入数据
	if _, err = writer.Write(data); err != nil {
		return
	}
	// 自动添加换行符
	if isAppend {
		if _, err = writer.Write([]byte("\n")); err != nil {
			return
		}
	}

	// 刷新缓冲区
	writer.Flush()

	ok = true
	return
}

// 检查文件/目录是否存在
func CheckFileDirExist(filePath string) (os.FileInfo, bool) {
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, false
	}
	return finfo, true
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

var (
	FileNoExistErr = errors.New("file not exist")
)

// 读取文件
func ReadFile(filePath, filename string) (data []byte, err error) {
	fileSrc := filepath.Join(filePath, filename)
	if _, ok := CheckFileDirExist(fileSrc); !ok {
		return nil, FileNoExistErr
	}
	data, err = os.ReadFile(fileSrc)
	return
}

```



## 八、基于镜像构建并运行

1. cd 通过上面的<测试代码>构建的项目中
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

5. 运行代码：运行成功后，则可以将 pdf 转为图片

```
go run main.go
```



# 实操“gopkg.in/gographics/imagick.v/imagick” 库将 pdf 转成图片的过程

> 官方文档：https://pkg.go.dev/gopkg.in/gographics/imagick.v3@v3.5.0#section-readme
> github仓库：https://github.com/gographics/imagick/tree/v3.5.0
> 官方示例：https://github.com/gographics/imagick/tree/v3.5.0/examples



> 解决所有问题后打包好的镜像：`docker pull leeprince/golang1_19:imagick_v3_7.1.1-23`



## 一、须知
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



## 二、注意

1. 官方示例：`https://github.com/gographics/imagick/tree/v3.5.0/examples/docker` 的示例 `Dockerfile` 文件中的<IMAGEMAGICK_VERSION>参数一定要匹配上 <## 须知> 的版本。
> 官方原版
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

> 修改后

```
FROM golang:1.19

...

ENV IMAGEMAGICK_VERSION=7.1.1-23

...

WORKDIR /www

VOLUME ["/www"]

EXPOSE 80

CMD tail -f /dev/null
```

## 三、问题


### （一）undefined: imagick.Initialize

问题

```
root@file-center-format-ctl-66f8df46c7-tw54b:/work# go run format-ctl/cmd/main.go -conf=format-ctl/configs/conf.yaml
format-ctl/internal/domain/formatdomain/image.go:17:10: undefined: imagick.Initialize
format-ctl/internal/domain/formatdomain/image.go:18:16: undefined: imagick.Terminate
format-ctl/internal/domain/formatdomain/image.go:20:16: undefined: imagick.NewMagickWand
format-ctl/internal/domain/formatdomain/image.go:34:40: undefined: imagick.ALPHA_CHANNEL_REMOVE
```

解决：开启 CGO_ENABLED

```
CGO_ENABLED=1 go run format-ctl/cmd/main.go -conf=format-ctl/configs/conf.yaml
```



### （二）ImageMagick 的依赖丢失

在本地直接基于自己构建的镜像创建容器，或者在构建容器之上通过 Dockerfile 重新构建都正常，但是使用在构建容器之上通过 Dockerfile 重新构建并在 k8s 中创建 pod 时发现下载的 ImageMagick 的依赖丢失

本地运行正常的情况下，在容器内`/usr/local/lib`下可以看到
```
ImageMagick-7.1.1           libMagickCore-7.Q16HDRI.so         libMagickWand-7.Q16HDRI.a   libMagickWand-7.Q16HDRI.so.10      python3.11
libMagickCore-7.Q16HDRI.a   libMagickCore-7.Q16HDRI.so.10      libMagickWand-7.Q16HDRI.la  libMagickWand-7.Q16HDRI.so.10.0.1
libMagickCore-7.Q16HDRI.la  libMagickCore-7.Q16HDRI.so.10.0.1  libMagickWand-7.Q16HDRI.so  pkgconfig
```

但是k8s构建之后发现没有了。

原因：安装 ImageMagick 时所在的目录发生变更，导致通过 k8s 运行时，出现丢失的情况。问题在于使用了`cd`

解决：

```
RUN cd && \
	wget https://github.com/ImageMagick/ImageMagick/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
	
---改后
RUN set -ex && \
	wget https://github.com/ImageMagick/ImageMagick/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
```



### （三）创建 pod 时报错 

在本地直接基于自己构建的镜像创建容器，或者在构建容器之上通过 Dockerfile 重新构建都正常，但是使用在构建容器之上通过 Dockerfile 重新构建并在 k8s 中创建 pod 时报错 

> 观察
```
PS F:file-center> kubectl get pods -n test | findstr file
format-ctl-6ff6db9c77-sb264                        1/2     CrashLoopBackOff   9          26m
```

查看错误信息： kubectl describe -n test format-ctl-6ff6db9c77-sb264

```
Events:
  Type     Reason     Age                  From               Message
  ----     ------     ----                 ----               -------
  Normal   Scheduled  18m                  default-scheduler  Successfully assigned test/file-center-format-ctl-6ff6db9c77-sb264 to cloud-k8s-node04-test
  Normal   Pulled     17m                  kubelet            Successfully pulled image "ccr.ccs.tencentyun.com/golden-cloud/file-center-format-ctl-test:v31" in 46.322607238s
  Normal   Pulled     17m                  kubelet            Successfully pulled image "ccr.ccs.tencentyun.com/golden-cloud/filebeat:7.10.0" in 261.117147ms
  Normal   Created    17m                  kubelet            Created container filebeat
  Warning  Failed     17m                  kubelet            Error: failed to start container "backend-apps": Error response from daemon: OCI runtime create failed: invalid mount {Destination:www Type:bind Source:/var/lib/docker/volumes/8a967d82b837f36c7097fc3af8db4454b3e749d0174b6972ac88aa9f6805da3d/_data Options:[rbind]}: mount destination www not absolute: unknown
  Normal   Pulling    17m                  kubelet            Pulling image "ccr.ccs.tencentyun.com/golden-cloud/filebeat:7.10.0"
  Normal   Started    17m                  kubelet            Started container filebeat
  Normal   Pulled     17m                  kubelet            Successfully pulled image "ccr.ccs.tencentyun.com/golden-cloud/file-center-format-ctl-test:v31" in 192.673352ms
  Warning  Failed     17m                  kubelet            Error: failed to start container "backend-apps": Error response from daemon: OCI runtime create failed: invalid mount {Destination:www Type:bind Source:/var/lib/docker/volumes/4b6903e9005905063bad27ab588d66e2ed13c83ef83e1418afa5ddd12542608e/_data Options:[rbind]}: mount destination www not absolute: unknown
  Warning  Failed     16m                  kubelet            Error: failed to start container "backend-apps": Error response from daemon: OCI runtime create failed: invalid mount {Destination:www Type:bind Source:/var/lib/docker/volumes/5e779c0f01b1ebf3669fa4a66281d792039a630901164663b4b9f2b4249a8608/_data Options:[rbind]}: mount destination www not absolute: unknown
  Normal   Pulled     16m                  kubelet            Successfully pulled image "ccr.ccs.tencentyun.com/golden-cloud/file-center-format-ctl-test:v31" in 198.731528ms
  Normal   Pulling    16m (x4 over 18m)    kubelet            Pulling image "ccr.ccs.tencentyun.com/golden-cloud/file-center-format-ctl-test:v31"
  Normal   Created    16m (x4 over 17m)    kubelet            Created container backend-apps
  Normal   Pulled     16m                  kubelet            Successfully pulled image "ccr.ccs.tencentyun.com/golden-cloud/file-center-format-ctl-test:v31" in 264.153464ms
  Warning  Failed     16m                  kubelet            Error: failed to start container "backend-apps": Error response from daemon: OCI runtime create failed: invalid mount {Destination:www Type:bind Source:/var/lib/docker/volumes/2026f1ec46f7891dff25146bc43377ed65a249c8725d05fdcbcff011542494dd/_data Options:[rbind]}: mount destination www not absolute: unknown
```

解决：使用 VOLUME 指令允许你在容器中创建挂载点时，使用绝对路径

```
VOLUME ["www"]

--- 改后
VOLUME ["/www"]
```



#### CGO_ENABLED 参数说明

CGO_ENABLED 是一个用于控制 Go 编译器行为的环境变量。它主要用于决定是否启用 Go 的 CGO（C Go）支持，这个支持允许 Go 代码与 C 代码进行交互。
当 CGO_ENABLED 被设置为 1 时，Go 编译器会启用 CGO。这意味着你可以在 Go 代码中使用 cgo 包，将 Go 代码与 C 代码进行混合编程，调用 C 函数，并在 Go 中使用 C 库。

例如，在使用外部 C 库或者需要与 C 代码交互的情况下，你可能需要设置 CGO_ENABLED=1。
但是，当 CGO_ENABLED 被设置为 0 时，Go 编译器将禁用 CGO。这意味着你的 Go 代码只能使用纯 Go 实现，无法调用 C 函数或使用 C 库。

在一些特定场景下，禁用 CGO 可能有好处，比如在跨平台编译、构建静态二进制文件或者在一些特定环境下提高性能等方面。
一般来说，在不需要与 C 代码交互的情况下，禁用 CGO 可以提高构建速度并简化部署。但在需要使用 C 代码的场景下，需要将 CGO_ENABLED 设置为 1。
设置 CGO_ENABLED 的方式可以是环境变量的方式，也可以在编译时通过命令行参数来指定。

例如，在命令行中编译一个 Go 程序并指定 CGO_ENABLED：

CGO_ENABLED=1 go build main.go   # 启用 CGO
CGO_ENABLED=0 go build main.go   # 禁用 CGO

这样你可以根据需要灵活地控制是否启用 CGO。

## 四、测试代码

```
package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"os"
	"path/filepath"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/23 21:32
 * @Desc:
 */

func main() {
	fileBytes, err := ReadFile(".", "0001.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}
	toImageType := "jpg"
	
	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)
	
	// --- imagick v3 --------------------------------
	// 使用的时候都放到外面去初始化了
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	
	// sourceImagePath := getSourceImageForCover(filepath.Dir(pathNoExtension))
	mw := imagick.NewMagickWand()
	// defer clearImagickWand(mw)
	
	// mw.SetResolution(192, 192)
	// mw.SetResolution(350, 350)
	// mw.SetImageResolution(350, 350)
	// mw.SetImageCompressionQuality(100)
	
	err = mw.ReadImageBlob(fileBytes)
	if err != nil {
		fmt.Println("[CreateImage] ReadImageBlob err:", err)
		return
	}
	
	if toImageType == "jpg" {
		toImageType = "jpeg"
	}
	
	/*width := mw.GetImageWidth()
	height := mw.GetImageHeight()
	
	// Calculate half the size
	hWidth := uint(width * 2)
	hHeight := uint(height * 2)
	
	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		fmt.Println("ResizeImage failed：", err)
		return
	}*/
	
	// length := mw.GetImageIterations()
	// fmt.Println("length", length)
	// fmt.Println("width", mw.GetImageWidth())
	// fmt.Println("height", mw.GetImageHeight())
	
	// pix := imagick.NewPixelWand()
	// pix.SetColor("white")
	// mw.SetBackgroundColor(pix)
	
	// mw.GetImageDepth()
	// mw.SetImageDepth(16)
	
	// 激活、停用、重置或设置alpha通道。
	err = mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)
	if err != nil {
		fmt.Println("SetImageAlphaChannel failed:", err)
		return
	}
	
	// 设置需要转化为的图片格式
	err = mw.SetImageFormat(toImageType)
	if err != nil {
		fmt.Println("SetImageFormat failed:", err)
		return
	}
	
	// 转化后的图片字节流
	imageByte := mw.GetImageBlob()
	
	// 可选：是否在本地保存为图片
	err = mw.WriteImage(filePath)
	if err != nil {
		fmt.Println("[CreateImage] WriteImage failed:", err)
		return
	}
	
	// --- imagick v3 --------------------------------
	
	fmt.Println("successful, filepath:", filePath)
	
	imageBase64 := base64.StdEncoding.EncodeToString(imageByte)
	fmt.Println("imageBase64:", imageBase64)
	
	fmt.Println("successful, filepath:", filePath)
	
}

var (
	FileNoExistErr = errors.New("file not exist")
)

// 读取文件
func ReadFile(filePath, filename string) (data []byte, err error) {
	fileSrc := filepath.Join(filePath, filename)
	if _, ok := CheckFileDirExist(fileSrc); !ok {
		return nil, FileNoExistErr
	}
	data, err = os.ReadFile(fileSrc)
	return
}

// 检查文件/目录是否存在
func CheckFileDirExist(filePath string) (os.FileInfo, bool) {
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, false
	}
	return finfo, true
}

```



## 五、基于镜像构建并运行

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

5. 运行代码：运行成功后，则可以将 pdf 转为图片

```
go run main.go
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

VOLUME ["/www"]

EXPOSE 80

CMD tail -f /dev/null
```



# 实操 "github.com/davidbyttow/govips/v2/vips" 库将 pdf 转成图片的过程

> 官方文档：https://github.com/davidbyttow/govips
> 关于libvips：https://github.com/libvips/libvips



该包通过将所有图像操作公开在 Go 中的第一类类型上，包装了 libvips 图像处理库的核心功能。Libvips 通常比其他图形处理器（如 GraphicsMagick 和 ImageMagick）快 4-8 倍。检查基准测试：速度和内存使用。

这样做的目的是使开发人员能够在 Go 中构建速度极快的图像处理器，这非常适合并发请求。



> 解决所有问题后打包好的镜像：`docker pull leeprince/golang1_19:govips_v2`



## 一、完整 Dockerfile

> 后来才发现官方有基于ubuntu构建的环境（https://github.com/davidbyttow/govips/tree/master/build），发现原来 `libvips` 可以直接通过 `apt-get install -y libvips-dev`，并包含其他依赖，如：imagemagick、libmagickwand，但是同样也遇到一些问题，在下面给出了解决办法

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

VOLUME ["/www"]

EXPOSE 80

CMD tail -f /dev/null
```



## 二、问题

### 1）libjpeg-turbo8-dev
`Package libjpeg-turbo8-dev is not available, but is referred to by another package.`

解决：
```
libjpeg-dev 代替 libjpeg-turbo8-dev
```



### 2）meson

安装 libvips 时需要使用 `meson`, 所以需要提前安装
```
apt-get install -y meson
```



### 3）error: cannot find builtin CJK font

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



> 语言包
>
> - fonts-noto-cjk: Noto 字体系列，包括中文、日文、韩文等各种语言的字体。
> - fonts-arphic-ukai: AR PL UKai 中文字体。
> - fonts-arphic-uming: AR PL UMing 中文字体。



## 三、测试代码

官方示例，有些方法（2023-12-19）已经废弃，手动替换即可

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
	
  // 导出图像
  // 替换前
  /*ep := vips.NewDefaultJPEGExportParams()
	image1bytes, _, err := image1.Export(ep)*/
  // 替换后
	image1bytes, _, err := image1.ExportNative()
	checkError(err)

  // 写入本地文件
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



## 四、基于镜像构建并运行

1. cd 通过上面的<测试代码>构建的项目中
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

5. 运行代码：运行成功后，则可以将 pdf 转为图片

```
go run main.go
```



# imagick_v3 与 govips 对比

## 功能测试（一次性功能测试）
- 在生成速度、生成质量、占用空间上 govips_v2 更优！
- 多页PDF时，imagick_v3 处理最后一页；govips_v2 处理第一页
- imagick_v3 的优势是很灵活，方便自定义一些属性！



## 性能测试

> 基准测试与压力测试不同之处：
- 目的不同： 基准测试旨在评估代码或算法在标准条件下的性能，而压力测试则旨在测试系统在高负载或极限条件下的稳定性和表现。
- 应用场景不同： 基准测试通常用于评估和改进代码、算法等实现，而压力测试主要用于验证系统在负载增加时的行为和稳定性。

### 1、并发测试

#### 1）无并发
> govips_v2 更优

```
imagick_v3            ||| govips_v2
...wg.wait...         ||| ...wg.wait...
cost mill time: 134ms ||| cost mill time: 117ms
```

#### 2）并发2
> govips_v2 更优

```
imagick_v3            ||| govips_v2
...wg.wait...         ||| ...wg.wait...
cost mill time: 135ms ||| cost mill time: 130ms
cost mill time: 135ms ||| cost mill time: 130ms
```

#### 3）并发4
> 基本持平

```
imagick_v3            ||| govips_v2
...wg.wait...         ||| ...wg.wait...
cost mill time: 155ms ||| cost mill time: 173ms
cost mill time: 158ms ||| cost mill time: 174ms
cost mill time: 169ms ||| cost mill time: 177ms
cost mill time: 176ms ||| cost mill time: 177ms
```

#### 4）并发8
> imagick_v3 更优

```
imagick_v3            ||| govips_v2
...wg.wait...         ||| ...wg.wait...
cost mill time: 207ms ||| cost mill time: 321ms
cost mill time: 227ms ||| cost mill time: 322ms
cost mill time: 229ms ||| cost mill time: 324ms
cost mill time: 232ms ||| cost mill time: 323ms
cost mill time: 234ms ||| cost mill time: 324ms
cost mill time: 238ms ||| cost mill time: 323ms
cost mill time: 240ms ||| cost mill time: 323ms
cost mill time: 246ms ||| cost mill time: 325ms
```

#### 5）并发16
> imagick_v3 更优

```
imagick_v3            ||| govips_v2
...wg.wait...         ||| ...wg.wait...
cost mill time: 338ms ||| cost mill time: 540ms
cost mill time: 381ms ||| cost mill time: 542ms
cost mill time: 384ms ||| cost mill time: 544ms
cost mill time: 397ms ||| cost mill time: 546ms
cost mill time: 399ms ||| cost mill time: 547ms
cost mill time: 399ms ||| cost mill time: 544ms
cost mill time: 414ms ||| cost mill time: 545ms
cost mill time: 437ms ||| cost mill time: 548ms
cost mill time: 442ms ||| cost mill time: 550ms
cost mill time: 443ms ||| cost mill time: 550ms
cost mill time: 446ms ||| cost mill time: 553ms
cost mill time: 446ms ||| cost mill time: 659ms
cost mill time: 449ms ||| cost mill time: 659ms
cost mill time: 451ms ||| cost mill time: 660ms
cost mill time: 452ms ||| cost mill time: 661ms
cost mill time: 452ms ||| cost mill time: 662ms
```

#### 总结
无并发情况下，govips_v2 的性能更优，
并发在4个左右
小于4个：govips_v2 的性能更优
等于4个：imagick_v3 与 govips_v2 的处理速度基本持平
大于4个：govips_v2 的性能更优



### 2、基准测试

#### 1）基于时间1s
> govips_v2 更优

imagick_v3
```
root@b9de6110f73d:/www# go test -bench=. -run=none -benchtime=1s
goos: linux
goarch: amd64
pkg: leeprince/imagick_v3_test
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByImagickV3-16             8         137313025 ns/op
PASS
ok      leeprince/imagick_v3_test       1.234s
```

govips_v2
```
root@ec8637f12a8b:/www# go test -bench=. -run=none -benchtime=1s
goos: linux
goarch: amd64
pkg: leeprince/govips
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByGovips-16               14          77429864 ns/op
PASS
ok      leeprince/govips        1.884s
```

#### 2）基于时间2s
> govips_v2 更优

imagick_v3
```
root@b9de6110f73d:/www# go test -bench=. -run=none -benchtime=2s
goos: linux
goarch: amd64
pkg: leeprince/imagick_v3_test
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByImagickV3-16            18         125225556 ns/op
PASS
ok      leeprince/imagick_v3_test       2.390s
```

govips_v2
```
root@ec8637f12a8b:/www# go test -bench=. -run=none -benchtime=2s
goos: linux
goarch: amd64
pkg: leeprince/govips
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByGovips-16               30          80373650 ns/op
PASS
ok      leeprince/govips        3.981s
```

#### 3）基于时间4s
> govips_v2 更优

imagick_v3
```
root@b9de6110f73d:/www# go test -bench=. -run=none -benchtime=4s
goos: linux
goarch: amd64
pkg: leeprince/imagick_v3_test
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByImagickV3-16            30         141284520 ns/op
PASS
ok      leeprince/imagick_v3_test       4.404s
```

govips_v2
```
root@ec8637f12a8b:/www# go test -bench=. -run=none -benchtime=4s
goos: linux
goarch: amd64
pkg: leeprince/govips
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByGovips-16               61          84797231 ns/op
PASS
ok      leeprince/govips        8.200s
```


#### 4）基于次数1次
> govips_v2 更优

imagick_v3
```
root@b9de6110f73d:/www# go test -bench=. -run=none -count=1
goos: linux
goarch: amd64
pkg: leeprince/imagick_v3_test
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByImagickV3-16             9         130929433 ns/op
PASS
ok      leeprince/imagick_v3_test       2.330s
```

govips_v2
```
root@ec8637f12a8b:/www# go test -bench=. -run=none -count=1
goos: linux
goarch: amd64
pkg: leeprince/govips
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByGovips-16               15          80620080 ns/op
PASS
ok      leeprince/govips        1.949s
```

#### 5）基于次数2次
> govips_v2 更优

imagick_v3
```
root@b9de6110f73d:/www# go test -bench=. -run=none -count=2
goos: linux
goarch: amd64
pkg: leeprince/imagick_v3_test
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByImagickV3-16             9         124839256 ns/op
BenchmarkCustomerPdfToImagesByImagickV3-16             8         129674675 ns/op
PASS
ok      leeprince/imagick_v3_test       3.425s
```

govips_v2
```
root@ec8637f12a8b:/www# go test -bench=. -run=none -count=2
goos: linux
goarch: amd64
pkg: leeprince/govips
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCustomerPdfToImagesByGovips-16               15          82978773 ns/op
BenchmarkCustomerPdfToImagesByGovips-16               14          92253700 ns/op
PASS
ok      leeprince/govips        4.355s
```







