package question

import (
	"bytes"
	"slices"

	"github.com/fogleman/gg"
)

func color() ([]byte, string, []int, error) {
	ans := []int{randint(0, 9)}

	ctx := gg.NewContext(WIDTH, HEIGHT)

	r := randint(0, 255)
	g := randint(0, 255)
	b := randint(0, 255)

	for row := range 3 {
		for col := range 3 {
			if row*3+col == ans[0] {
				ctx.SetRGB255(r+randint(-5, 5), g+randint(-5, 5), b+randint(-5, 5))
			} else {
				ctx.SetRGB255(r, g, b)
			}
			ctx.DrawRectangle(float64(col*200), float64(row*200), 200, 200)
			ctx.Fill()
		}
	}

	buf := &bytes.Buffer{}
	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	return buf.Bytes(), "squares with different color", ans, nil
}
