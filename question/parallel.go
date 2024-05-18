package question

import (
	"bytes"
	"slices"

	"github.com/fogleman/gg"
)

func parallel() ([]byte, string, []int, error) {
	ans := []int{}

	ctx := gg.NewContext(WIDTH, HEIGHT)

	ctx.SetRGB(0, 0, 0)
	ctx.Clear()

	ctx.SetRGB(1, 1, 1)

	for row := range 3 {
		for col := range 3 {
			baseX := float64(col*200 + 20)
			baseY := float64(row*200 + 180)
			ctx.SetLineWidth(3)
			ctx.DrawLine(baseX, baseY, baseX+150, baseY)
			ctx.Stroke()

			offsetX := randomOffset(2)
			offsetY := randomOffset(2)
			if randint(0, 5) == 1 {
				ans = append(ans, row*3+col)
				offsetX = 0
				offsetY = 0
			}

			ctx.DrawLine(baseX, baseY-50, baseX+150+float64(offsetX), baseY-50+float64(offsetY))
			ctx.Stroke()
		}
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	buf := &bytes.Buffer{}
	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	return buf.Bytes(), "parallel lines", ans, nil
}
