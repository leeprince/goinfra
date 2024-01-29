package main

import (
	"fmt"
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
	"strings"
	
	"golang.org/x/image/bmp"
)

// input.png 图片水印格式参数
// const (
// 	imagePath        = "./input/input.png"
// 	rectTopRightX    = 300
// 	rectTopRightY    = 60
// 	rectBottomRightX = 550
// 	rectBottomRightY = 60
// )

// 1699934865230.png 图片水印格式参数
const (
	imagePath        = "./input/1699934865230.png"
	rectTopRightX    = 300
	rectTopRightY    = 35
	rectBottomRightX = 535
	rectBottomRightY = 36
)

const (
	textTopRight    = "itgogogo.cn"
	textBottomRight = "itgogogo.cn IT go go go 编程资料站"
)

var (
	// 输出的文件夹，结尾必须带上/否则会连在一起
	outputFileDir = "./output/"
)

func main() {
	// 单个图片水印处理
	// SingleImageWatermarkProcess()
	
	// 应用：遍历文件夹
	WalkDirImageWatermarkProcess()
}

func SingleImageWatermarkProcess() {
	// 打开图片文件
	file, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	// 解码图片
	img, format, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	err = ImageWatermarkProcess(img, imagePath, format)
	if err != nil {
		fmt.Println("ImageWatermarkProcess", err)
		return
	}
	fmt.Println("ImageWatermarkProcess success:", file.Name())
}

// WalkDirImageWatermarkProcess 遍历文件夹并进行单个图片水印处理
func WalkDirImageWatermarkProcess() {
	// 输入的文件夹，结尾必须带上/否则会连在一起
	intputFileDir := "./tmpintput/"
	// 输出的文件夹，结尾必须带上/否则会连在一起
	outputFileDir = "./tmpoutput/"
	
	// 遍历文件夹取出图片列表
	files, err := os.ReadDir(intputFileDir)
	if err != nil {
		fmt.Printf("Error reading directory %v: %v\n", "path/to/your/directory", err)
		return
	}
	
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("filepath.Walk err err:", err)
			continue
		}
		
		path := filepath.Join(intputFileDir, file.Name())
		f, ferr := os.Open(path)
		if ferr != nil {
			fmt.Printf("Error opening file %v: %v\n", path, ferr)
			err = ferr
			continue
		}
		
		img, format, derr := image.Decode(f)
		f.Close()
		if derr != nil {
			fmt.Printf("Error image.Decode %v\n", derr)
			continue
		}
		
		err = ImageWatermarkProcess(img, path, format)
		if err != nil {
			fmt.Println("ImageWatermarkProcess", err)
			return
		}
		
		fmt.Println("ImageWatermarkProcess success:", file.Name())
	}
	return
}

// ImageWatermarkProcess 图片水印处理：处理右上角和右下角的水印，并用头部和底部的大部分像素作为覆盖颜色，并加上文案
func ImageWatermarkProcess(img image.Image, filePath, format string) (err error) {
	// 创建一个和原图一样大小的 RGBA 图片
	rgba := image.NewRGBA(img.Bounds())
	
	// 将原图复制到新的 RGBA 图片上
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)
	
	// 覆盖指定区域像素的方式
	// OverWhiteColor 取白色像素覆盖指定区域像素
	// OverWhiteColor(rgba)
	// OverAverageColor 取指定区域平均颜色像素覆盖指定区域像素
	// OverAverageColor(rgba)
	// OverDominantColor 取指定区域平均颜色像素覆盖指定区域像素
	err = OverDominantColor(rgba)
	if err != nil {
		fmt.Println("OverDominantColor error:", err)
		return
	}
	
	// 创建一个新的图片文件
	outFilePath := fmt.Sprintf("%s%s.%s", outputFileDir, GetFilePathOfName(filePath), format)
	out, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println("os.Create error:", err)
		return
	}
	defer out.Close()
	
	// 将 RGBA 图片编码为原始格式并写入新的图片文件
	switch format {
	case "jpeg":
		var opt jpeg.Options
		opt.Quality = 80
		if err = jpeg.Encode(out, rgba, &opt); err != nil {
			fmt.Println("jpeg.Encode error:", err)
			return
		}
	case "png":
		if err = png.Encode(out, rgba); err != nil {
			fmt.Println("png.Encode error:", err)
			return
		}
	case "gif":
		if err = gif.Encode(out, rgba, nil); err != nil {
			fmt.Println("gif.Encode error:", err)
			return
		}
	case "bmp":
		if err = bmp.Encode(out, rgba); err != nil {
			fmt.Println("bmp.Encode error:", err)
			return
		}
	case "webp":
		// 对于 WEBP 图像，我们将其保存为 PNG 格式
		if err = png.Encode(out, rgba); err != nil {
			fmt.Println("webp png.Encode error:", err)
			return
		}
	default:
		fmt.Println("jpeg.Encode error:", err)
		return
	}
	return
}

// OverWhiteColor 取白色像素覆盖指定区域像素
func OverWhiteColor(rgba *image.RGBA) {
	// 定义要覆盖的区域，这里我们覆盖右上角和右下角的部分
	rectTopRight := image.Rect(rgba.Bounds().Max.X-rectTopRightX, 0, rgba.Bounds().Max.X, rectTopRightY)
	rectBottomRight := image.Rect(rgba.Bounds().Max.X-rectBottomRightX, rgba.Bounds().Max.Y-rectBottomRightY, rgba.Bounds().Max.X, rgba.Bounds().Max.Y)
	
	// 创建一个白色的矩形，用于覆盖水印
	whiteColorTopRight := image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	whiteColorBottomRight := image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	
	// 在 RGBA 图片上绘制白色矩形
	draw.Draw(rgba, rectTopRight, whiteColorTopRight, image.Point{}, draw.Src)
	draw.Draw(rgba, rectBottomRight, whiteColorBottomRight, image.Point{}, draw.Src)
	
	return
}

// OverAverageColor 取指定区域平均颜色像素覆盖指定区域像素
func OverAverageColor(rgba *image.RGBA) {
	// 定义要覆盖的区域，这里我们覆盖右上角和右下角的部分
	rectTopRight := image.Rect(rgba.Bounds().Max.X-rectTopRightX, 0, rgba.Bounds().Max.X, rectTopRightY)
	rectBottomRight := image.Rect(rgba.Bounds().Max.X-rectBottomRightX, rgba.Bounds().Max.Y-rectBottomRightY, rgba.Bounds().Max.X, rgba.Bounds().Max.Y)
	
	// 计算区域的平均像素颜色
	avgColorTopRight := averageColor(rgba, rectTopRight)
	avgColorBottomRight := averageColor(rgba, rectBottomRight)
	
	// 创建一个平均颜色的矩形，用于覆盖水印
	avgTopRight := image.NewUniform(avgColorTopRight)
	avgBottomRight := image.NewUniform(avgColorBottomRight)
	
	// 在 RGBA 图片上绘制平均颜色的矩形
	draw.Draw(rgba, rectTopRight, avgTopRight, image.Point{}, draw.Src)
	draw.Draw(rgba, rectBottomRight, avgBottomRight, image.Point{}, draw.Src)
	
	return
}

// OverDominantColor 取指定区域平均颜色像素覆盖指定区域像素
func OverDominantColor(rgba *image.RGBA) (err error) {
	// 定义要覆盖的区域，这里我们覆盖右上角和右下角的部分
	rectTopRight := image.Rect(rgba.Bounds().Max.X-rectTopRightX, 0, rgba.Bounds().Max.X, rectTopRightY)
	rectBottomRight := image.Rect(rgba.Bounds().Max.X-rectBottomRightX, rgba.Bounds().Max.Y-rectBottomRightY, rgba.Bounds().Max.X, rgba.Bounds().Max.Y)
	
	// 计算区域的平均像素颜色：存在一个问题是字体较多时使用到了字体颜色
	// dominantColorTopRight := dominantColor(rgba, rectTopRight)
	// dominantColorBottomRight := dominantColor(rgba, rectBottomRight)
	// 计算区域的平均像素颜色：优化取整个高度的头部/底部作为右上角的区域判断
	dominantColorTopRight := dominantColor(rgba, image.Rect(0, 0, rgba.Bounds().Max.X, rectTopRightY))
	dominantColorBottomRight := dominantColor(rgba, image.Rect(0, rgba.Bounds().Max.Y-rectBottomRightY, rgba.Bounds().Max.X, rgba.Bounds().Max.Y))
	
	// 创建一个平均颜色的矩形，用于覆盖水印
	dominantTopRight := image.NewUniform(dominantColorTopRight)
	dominantBottomRight := image.NewUniform(dominantColorBottomRight)
	
	// 在 RGBA 图片上绘制平均颜色的矩形
	draw.Draw(rgba, rectTopRight, dominantTopRight, image.Point{}, draw.Src)
	draw.Draw(rgba, rectBottomRight, dominantBottomRight, image.Point{}, draw.Src)
	
	// 在右上角和右下角添加红色的文本
	err = addLabel(rgba, textTopRight, rectTopRight, color.RGBA{R: 255, A: 255})
	if err != nil {
		fmt.Println("addLabel textTopRight", err)
		return
	}
	err = addLabel(rgba, textBottomRight, rectBottomRight, color.RGBA{R: 255, A: 255})
	if err != nil {
		fmt.Println("addLabel textBottomRight", err)
		return
	}
	
	return
}

// averageColor 计算图像指定区域的平均像素颜色
func averageColor(img *image.RGBA, rect image.Rectangle) color.RGBA {
	var r, g, b, a uint32
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			rr, gg, bb, aa := img.At(x, y).RGBA()
			r += rr
			g += gg
			b += bb
			a += aa
		}
	}
	totalPixels := uint32(rect.Dx() * rect.Dy())
	return color.RGBA{
		R: uint8(r / totalPixels),
		G: uint8(g / totalPixels),
		B: uint8(b / totalPixels),
		A: uint8(a / totalPixels),
	}
}

// dominantColor 计算图像指定区域中占大部分的颜色
func dominantColor(img *image.RGBA, rect image.Rectangle) color.RGBA {
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

// addLabel 在图像的指定位置添加文本标签
func addLabel(img *image.RGBA, label string, rect image.Rectangle, col color.RGBA) (err error) {
	// 字体
	var fontFace font.Face
	// 默认字体
	fontFace = basicfont.Face7x13
	
	// 加载字体
	fontBytes, err := os.ReadFile("./Arial Unicode.ttf")
	if err != nil {
		fmt.Println("os.ReadFile ttf", err)
		return
	}
	opentypeFont, err := opentype.Parse(fontBytes)
	if err != nil {
		fmt.Println("oopentypes.Parse", err)
		return
	}
	fontFace, err = opentype.NewFace(opentypeFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
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

// GetFilePathOfName 获取URL 的文件路径和名称："./0001.pdf 结果是：0001
func GetFilePathOfName(filePath string) string {
	sArr := strings.Split(filePath, "/")
	nameArr := strings.Split(sArr[len(sArr)-1], ".")
	realNameArr := nameArr[:len(nameArr)-1]
	s := strings.Join(realNameArr, ".")
	return s
}
