package question

import (
	"bytes"
	"image/color"

	"github.com/fogleman/gg"
)

func add() ([]byte, string, []int, error) {
	ctx := gg.NewContext(WIDTH, HEIGHT)

	ctx.SetColor(color.White)
	ctx.Clear()
	ctx.DrawCircle(300, 300, 300)
	ctx.SetColor(color.Black)
	ctx.Fill()

	buf := &bytes.Buffer{}

	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, "traffic lights", []int{0}, err
	}

	return buf.Bytes(), "traffic lights", []int{0}, nil
}
