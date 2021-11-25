package images

import (
	"gart/config"
	"image/color"
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

type ContoursConfiguration struct {
	Width  int `form:"width,default=5000"`
	Height int `form:"height,default=5000"`
	Count  int `form:"count,default=500"`
}

func (c ContoursConfiguration) Parameters() []config.ImageParameter {
	return []config.ImageParameter{
		{
			Name:         "width",
			Type:         "range",
			DefaultValue: 5000,
			Min:          500,
			Max:          8000,
			Step:         100,
		},
		{
			Name:         "height",
			Type:         "range",
			DefaultValue: 5000,
			Min:          500,
			Max:          8000,
			Step:         100,
		},
		{
			Name:         "count",
			Type:         "range",
			DefaultValue: 500,
			Min:          1,
			Max:          1000,
		},
	}
}

func Contours(config ContoursConfiguration) string {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0x58, 0x18, 0x45, 0xFF},
		{0x90, 0x0C, 0x3F, 0xFF},
		{0xC7, 0x00, 0x39, 0xFF},
		{0xFF, 0x57, 0x33, 0xFF},
		{0xFF, 0xC3, 0x0F, 0xFF},
	}
	c := generativeart.NewCanva(config.Width, config.Height)
	c.SetBackground(color.RGBA{0x1a, 0x06, 0x33, 0xFF})
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewContourLine(config.Count))

	path := "assets/generated/contours.png"
	c.ToPNG(path)

	return path
}
