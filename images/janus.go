package images

import (
	"gart/config"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/common"
)

type JanusConfiguration struct {
	Width  int     `form:"width,default=5000"`
	Height int     `form:"height,default=5000"`
	Count  int     `form:"count,default=10"`
	Decay  float64 `form:"decay,default=0.2"`
}

func (c JanusConfiguration) Parameters() []config.ImageParameter {
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
			DefaultValue: 10,
			Min:          1,
			Max:          10,
		},
		{
			Name:         "decay",
			Type:         "range",
			DefaultValue: 0.2,
			Step:         0.1,
			Min:          0.1,
			Max:          10,
		},
	}
}

func draw(count int, decay float64, c *generativeart.Canva) {
	ctex := gg.NewContextForRGBA(c.Img())
	s := 5020.0
	r := 0.3

	for i := 0; i < count; i++ {
		k := i % len(c.Opts().ColorSchema())
		ctex.Push()
		ctex.Translate(float64(c.Width()/2), float64(c.Height()/2))

		//theta += rand.Float64()*math.Pi/2
		theta := common.RandomRangeFloat64(math.Pi/4, 3*math.Pi/4)
		x1, y1 := math.Cos(theta)*r, math.Sin(theta)*r
		x2, y2 := -x1, -y1

		noise := common.RandomRangeFloat64(-math.Abs(y1), math.Abs(y1))
		y1 += noise
		y2 += noise

		//r = r - r*j.decay
		s = s * 0.836
		ctex.Scale(s, s)
		//r = r * 0.836
		ctex.DrawArc(x1, y1, 1.0, math.Pi*3/2+theta, math.Pi*5/2+theta)
		ctex.SetColor(c.Opts().ColorSchema()[k])
		ctex.Fill()
		ctex.DrawArc(x2, y2, 1.0, math.Pi/2+theta, math.Pi*3/2+theta)
		ctex.SetColor(c.Opts().ColorSchema()[k])
		ctex.Fill()
		ctex.Pop()
	}
}

func Janus(config JanusConfiguration) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(config.Width, config.Height)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(common.DarkRed)
	c.SetForeground(common.LightPink)

	draw(config.Count, config.Decay, c)
	// c.Draw(arts.NewJanus(config.Count, config.Decay))

	path := "assets/generated/janus.png"
	c.ToPNG(path)

	return path
}
