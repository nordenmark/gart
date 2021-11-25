package images

import (
	"gart/config"
	"image/color"
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

type ShapesConfiguration struct {
	Width  int `form:"width,default=5000"`
	Height int `form:"height,default=5000"`
	Count  int `form:"count,default=150"`
}

func (c ShapesConfiguration) Parameters() []config.ImageParameter {
	return []config.ImageParameter{
		{
			Name:         "width",
			Type:         "number",
			DefaultValue: 5000,
			Min:          500,
			Max:          8000,
		},
		{
			Name:         "height",
			Type:         "number",
			DefaultValue: 5000,
			Min:          500,
			Max:          8000,
		},
		{
			Name:         "count",
			Type:         "range",
			DefaultValue: 150,
			Min:          1,
			Max:          1000,
		},
	}
}

func Shapes(config ShapesConfiguration) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(config.Width, config.Height)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema([]color.RGBA{
		{0xCF, 0x2B, 0x34, 0xFF},
		{0xF0, 0x8F, 0x46, 0xFF},
		{0xF0, 0xC1, 0x29, 0xFF},
		{0x19, 0x6E, 0x94, 0xFF},
		{0x35, 0x3A, 0x57, 0xFF},
	})
	c.Draw(arts.NewRandomShape(config.Count))

	path := "assets/generated/shapes.png"
	c.ToPNG(path)

	return path
}
