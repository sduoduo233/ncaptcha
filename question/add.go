package question

import (
	"bytes"

	"github.com/fogleman/gg"
)

func add() ([]byte, []int, error) {
	ctx := gg.NewContext(WIDTH, HEIGHT)

	buf := &bytes.Buffer{}

	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, []int{0}, err
	}

	return buf.Bytes(), []int{0}, nil
}
