package config

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 01:56
 * @Desc:
 */

var C = &Config{}

type Config struct {
	WebPageUrlList      []string    `yaml:"WebPageUrlList"`
	FTP                 FTP         `yaml:"FTP"`
	SaveImagePathPrefix string      `yaml:"SaveImagePathPrefix"` // 保存本地、远程文件时，路径前缀
	SaveLocal           SaveLocal   `yaml:"SaveLocal"`           // 保存到本地
	SaveRemoter         SaveRemoter `yaml:"SaveRemoter"`         // 保存到远程 Wordpress
}

type FTP struct {
	Conf struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Conf"`
	AccessHost    string `yaml:"AccessHost"`
	TargetImgHost string `yaml:"TargetImgHost"` // 目标网站的图片主机；有些网站的图片不包含主机名，而是使用相对路径，所以需要指定一个默认的主机名
}

// SaveLocal 是否将图片和文章保存到本地
// 注意：只有开启了才会进行图片的水印处理
type SaveLocal struct {
	IsSave  bool   `yaml:"IsSave"`
	SaveDir string `yaml:"SaveDir"` // 保存路径
	Water   Water  `yaml:"Water"`   // 图片水印处理
}

type Water struct {
	WaterPosition    WaterPosition    `yaml:"WaterPosition"`    // 水印所在位置
	ConvertWaterText ConvertWaterText `yaml:"ConvertWaterText"` // 覆盖水印的文本
}

type WaterPosition struct {
	RectTopRightX    int `yaml:"RectTopRightX"`
	RectTopRightY    int `yaml:"RectTopRightY"`
	RectBottomRightX int `yaml:"RectBottomRightX"`
	RectBottomRightY int `yaml:"RectBottomRightY"`
}

type ConvertWaterText struct {
	TtfFilePath     string  `yaml:"TtfFilePath"` // 字体库位置
	TextTopRight    string  `yaml:"TextTopRight"`
	TextBottomRight string  `yaml:"TextBottomRight"`
	FontSize        float64 `yaml:"FontSize"`
	Dpi             float64 `yaml:"Dpi"`
}

type SaveRemoter struct {
	IsSave bool   `yaml:"IsSave"`
	Url    string `yaml:"Url"` // 保存的地址
}
