package question

import (
	"bytes"
	"crypto/rand"
	"slices"

	"github.com/fogleman/gg"
)

func prime() ([]byte, string, []int, error) {
	ans := []int{randint(0, 9)}
	ctx := gg.NewContext(WIDTH, HEIGHT)

	ctx.SetRGB(0, 0, 0)
	ctx.Clear()

	ctx.LoadFontFace("./unifont-15.0.06.ttf", 32)

	_, h := ctx.MeasureString("0123456789")

	ctx.SetRGB(1, 1, 1)
	for row := range 3 {
		for col := range 3 {
			if randint(0, 5) == 1 {
				p, _ := rand.Prime(rand.Reader, 32)
				ctx.DrawString(p.String(), float64(col*200), float64(row*200)+h+100)
			} else {
				p1, _ := rand.Prime(rand.Reader, 16)
				p2, _ := rand.Prime(rand.Reader, 16)
				ctx.DrawString(p1.Mul(p1, p2).String(), float64(col*200), float64(row*200)+h+100)
			}
		}
	}

	buf := &bytes.Buffer{}
	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	return buf.Bytes(), "prime number", ans, nil
}
