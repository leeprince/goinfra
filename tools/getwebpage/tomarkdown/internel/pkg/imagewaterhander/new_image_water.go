package imagewaterhander

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 16:15
 * @Desc:
 */

type ImageWaterHander struct {
	waterPosition WaterPosition
	waterText     WaterText
}

type WaterPosition struct {
	RectTopRightX    int
	RectTopRightY    int
	RectBottomRightX int
	RectBottomRightY int
}

type WaterText struct {
	TtfFilePath     string
	TextTopRight    string
	TextBottomRight string
	FontSize        float64
	Dpi             float64
}

func NewImageWaterHander(waterPosition WaterPosition, waterText WaterText) *ImageWaterHander {
	return &ImageWaterHander{
		waterPosition: waterPosition,
		waterText:     waterText,
	}
}

func (r *ImageWaterHander) checkInit() (err error) {
	return
}
