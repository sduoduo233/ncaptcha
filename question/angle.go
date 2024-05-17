package question

import (
	"bytes"
	"fmt"
	"math"
	"slices"

	"github.com/fogleman/gg"
)

func angle() ([]byte, string, []int, error) {
	ans := []int{randint(0, 9)}

	targetAngle := gg.Radians(float64(randint(10, 80)))

	ctx := gg.NewContext(WIDTH, HEIGHT)

	ctx.SetRGB(0, 0, 0)
	ctx.Clear()

	ctx.SetRGB(1, 1, 1)
	for row := range 3 {
		for col := range 3 {
			baseX := float64(col*200 + 20)
			baseY := float64(row*200 + 180)
			ctx.SetLineWidth(5)
			ctx.DrawLine(baseX, baseY, baseX+150, baseY)
			ctx.Stroke()
			if row*3+col != ans[0] {
				a := targetAngle - 2*math.Pi/360*float64(randint(5, 10))
				if randint(0, 100)%2 == 0 {
					a = targetAngle + 2*math.Pi/360*float64(randint(5, 10))
				}
				ctx.DrawLine(baseX, baseY, baseX+math.Cos(a)*float64(150), baseY-math.Sin(a)*float64(150))
			} else {
				ctx.DrawLine(baseX, baseY, baseX+math.Cos(targetAngle)*float64(150), baseY-math.Sin(targetAngle)*float64(150))
			}

			ctx.Stroke()
		}
	}

	buf := &bytes.Buffer{}
	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	return buf.Bytes(), fmt.Sprintf("%.1f° angles", gg.Degrees(targetAngle)), ans, nil
}
