package imagewaterhander

import (
	"errors"
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"golang.org/x/image/bmp"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 16:07
 * @Desc:
 */

type ImageFormat string

const (
	ImageFormatJpeg ImageFormat = "jpeg"
	ImageFormatPng  ImageFormat = "png"
	ImageFormatGif  ImageFormat = "gif"
	ImageFormatBmp  ImageFormat = "bmp"
	ImageFormatWebp ImageFormat = "webp"
)

// ImageWatermarkProcess 图片水印处理：处理右上角和右下角的水印，并用头部和底部的大部分像素作为覆盖颜色，并加上文案
func (r *ImageWaterHander) ImageWatermarkProcess(img image.Image, format ImageFormat, outputDir, fileName string) (err error) {
	if err = r.checkInit(); err != nil {
		if err != nil {
			fmt.Println("checkInit error:", err)
			return
		}
	}
	
	// 创建一个和原图一样大小的 RGBA 图片
	rgba := image.NewRGBA(img.Bounds())
	
	// 将原图复制到新的 RGBA 图片上
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	
	// 覆盖指定区域像素的方式
	err = r.overDominantColor(rgba) // 取指定区域中占大部分的颜色覆盖指定区域像素
	if err != nil {
		fmt.Println("overDominantColor error:", err)
		return
	}
	
	// 创建一个新的图片文件
	outFilePath := filepath.Join(outputDir, fileName)
	if _, ok := fileutil.CheckFileDirExist(outFilePath); !ok {
		// 创建目录
		err = os.MkdirAll(outputDir, os.ModePerm)
		if ok = os.IsNotExist(err); ok {
			err = errors.New("创建文件目录错误")
			return
		}
	}
	out, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println("ImageWatermarkProcess os.Create error:", err)
		return
	}
	defer out.Close()
	
	// 将 RGBA 图片编码为原始格式并写入新的图片文件
	switch format {
	case ImageFormatJpeg:
		var opt jpeg.Options
		opt.Quality = 80
		if err = jpeg.Encode(out, rgba, &opt); err != nil {
			fmt.Println("jpeg.Encode error:", err)
			return
		}
	case ImageFormatPng:
		if err = png.Encode(out, rgba); err != nil {
			fmt.Println("png.Encode error:", err)
			return
		}
	case ImageFormatGif:
		if err = gif.Encode(out, rgba, nil); err != nil {
			fmt.Println("gif.Encode error:", err)
			return
		}
	case ImageFormatBmp:
		if err = bmp.Encode(out, rgba); err != nil {
			fmt.Println("bmp.Encode error:", err)
			return
		}
	case ImageFormatWebp:
		// 对于 WEBP 图像，我们将其保存为 PNG 格式
		if err = png.Encode(out, rgba); err != nil {
			fmt.Println("webp png.Encode error:", err)
			return
		}
	default:
		err = errors.New("not match fromat")
		fmt.Println("format default err", err)
		return
	}
	return
}

// overDominantColor 取指定区域中占大部分的颜色覆盖指定区域像素
func (r *ImageWaterHander) overDominantColor(rgba *image.RGBA) (err error) {
	// 定义要覆盖的区域，这里我们覆盖右上角和右下角的部分
	var rectTopRight, rectBottomRight image.Rectangle
	
	// 计算图像指定区域中占大部分的颜色：优化取整个高度的头部/底部作为右上角的区域判断
	var dominantColorTopRight, dominantColorBottomRight color.RGBA
	
	// 创建一个平均颜色的矩形，用于覆盖水印
	
	if r.waterPosition.RectTopRightX > 0 && r.waterPosition.RectTopRightY > 0 {
		// 定义要覆盖的区域，这里我们覆盖右上角的部分
		rectTopRight = image.Rect(rgba.Bounds().Max.X-r.waterPosition.RectTopRightX, 0, rgba.Bounds().Max.X, r.waterPosition.RectTopRightY)
		// 计算图像指定区域中占大部分的颜色：优化取整个高度的头部作为右上角的区域判断
		dominantColorTopRight = r.dominantColor(rgba, image.Rect(0, 0, rgba.Bounds().Max.X, r.waterPosition.RectTopRightY))
		// 创建一个平均颜色的矩形，用于覆盖水印
		dominantTopRight := image.NewUniform(dominantColorTopRight)
		// 在 RGBA 图片上绘制平均颜色的矩形
		draw.Draw(rgba, rectTopRight, dominantTopRight, image.Point{}, draw.Src)
		// 在右上角添加红色的文本
		err = r.addLabel(rgba, r.waterText.TextTopRight, rectTopRight, color.RGBA{R: 255, A: 255})
		if err != nil {
			fmt.Println("addLabel textTopRight", err)
			return
		}
	}
	if r.waterPosition.RectBottomRightX > 0 && r.waterPosition.RectBottomRightY > 0 {
		// 定义要覆盖的区域，这里我们覆盖右下角的部分
		rectBottomRight = image.Rect(rgba.Bounds().Max.X-r.waterPosition.RectBottomRightX, rgba.Bounds().Max.Y-r.waterPosition.RectBottomRightY, rgba.Bounds().Max.X, rgba.Bounds().Max.Y)
		// 计算图像指定区域中占大部分的颜色：优化取整个高度的头部作为右下角的区域判断
		dominantColorBottomRight = r.dominantColor(rgba, image.Rect(0, rgba.Bounds().Max.Y-r.waterPosition.RectBottomRightY, rgba.Bounds().Max.X, rgba.Bounds().Max.Y))
		// 创建一个平均颜色的矩形，用于覆盖水印
		dominantBottomRight := image.NewUniform(dominantColorBottomRight)
		// 在 RGBA 图片上绘制平均颜色的矩形
		draw.Draw(rgba, rectBottomRight, dominantBottomRight, image.Point{}, draw.Src)
		// 在右下角添加红色的文本
		err = r.addLabel(rgba, r.waterText.TextBottomRight, rectBottomRight, color.RGBA{R: 255, A: 255})
		if err != nil {
			fmt.Println("addLabel textBottomRight", err)
			return
		}
	}
	
	return
}

// dominantColor 计算图像指定区域中占大部分的颜色
func (r *ImageWaterHander) dominantColor(img *image.RGBA, rect image.Rectangle) color.RGBA {
	colorFreq := make(map[color.RGBA]uint32)
	var dominantColor color.RGBA
	var maxFreq uint32
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			colorFreq[c]++
			if colorFreq[c] > maxFreq {
				maxFreq = colorFreq[c]
				dominantColor = c
			}
		}
	}
	return dominantColor
}

var opentypeFont *opentype.Font
var addLabelOnce sync.Once

// addLabel 在图像的指定位置添加文本标签
func (r *ImageWaterHander) addLabel(img *image.RGBA, label string, rect image.Rectangle, col color.RGBA) (err error) {
	// 字体
	var fontFace font.Face
	// 默认字体
	fontFace = basicfont.Face7x13
	
	// 加载字体
	addLabelOnce.Do(func() {
		fontBytes, fonterr := os.ReadFile(r.waterText.TtfFilePath)
		if fonterr != nil {
			err = fonterr
			fmt.Println("os.ReadFile ttf err:", err)
			return
		}
		opentypeFont, err = opentype.Parse(fontBytes)
		if err != nil {
			fmt.Println("oopentypes.Parse", err)
			return
		}
	})
	if err != nil {
		fmt.Println("addLabelOnce err:", err)
		return
	}
	if opentypeFont == nil {
		fmt.Println("opentypeFont is nil")
		err = errors.New("opentypeFont is nil")
		return
	}
	
	fontFace, err = opentype.NewFace(opentypeFont, &opentype.FaceOptions{
		Size:    r.waterText.FontSize, // 字体大小
		DPI:     r.waterText.Dpi,      // 字体分辨率
		Hinting: font.HintingFull,
	})
	if err != nil {
		fmt.Println("opentype.NewFace", err)
		return
	}
	
	// 计算文本显示位置
	// 计算文本的宽度和高度
	var width fixed.Int26_6
	for _, r := range label {
		width += font.MeasureString(fontFace, string(r))
	}
	height := fontFace.Metrics().Height
	// 计算文本的起始位置
	x := rect.Min.X + (rect.Dx()-int(width.Round()))/2
	y := rect.Min.Y + (rect.Dy()-int(height.Round()))/2 + int(fontFace.Metrics().Ascent.Round())
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}
	
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: fontFace,
		Dot:  point,
	}
	d.DrawString(label)
	return
}
