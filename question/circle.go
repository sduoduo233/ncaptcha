package question

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/fogleman/gg"
)

func circle() ([]byte, string, []int, error) {
	ans := []int{randint(0, 9)}
	ctx := gg.NewContext(WIDTH, HEIGHT)

	ctx.SetRGB(0, 0, 0)
	ctx.Clear()

	r := float64(randint(20, 80))

	ctx.SetRGB(1, 1, 1)
	ctx.SetLineWidth(5)
	for row := range 3 {
		for col := range 3 {
			x := float64(col*200 + 100)
			y := float64(row*200 + 100)
			if row*3+col != ans[0] {
				rr := r + float64(randint(5, 15))
				if randint(0, 100)%2 == 0 {
					rr = r - float64(randint(5, 15))
				}
				ctx.DrawCircle(x, y, rr)
			} else {
				ctx.DrawCircle(x, y, r)
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

	return buf.Bytes(), fmt.Sprintf("circles of %.0fpx raduis", r), ans, nil
}
