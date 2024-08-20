package question

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"

	_ "embed"
)

//go:embed unifont-15.0.06.ttf
var unifont []byte

// randomDigitRune returns a random digit other than n
func randomDigitRune(n string) rune {
	nn, err := strconv.Atoi(n)
	if err != nil {
		return []rune(n)[0]
	}
	for {
		i := randint(0, 9)
		if i == nn {
			continue
		}
		return []rune(strconv.Itoa(i))[0]
	}

}

func randint(min, max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}

	return min + int(n.Int64())
}

func randomOffset(m int) int {
	if randint(0, 2) == 0 {
		return -1 * randint(1, m)
	} else {
		return randint(1, m)
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func loadFontFace(ctx *gg.Context, points float64) error {
	font, err := truetype.Parse(unifont)
	if err != nil {
		return fmt.Errorf("parse font: %w", err)
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size: points,
	})

	ctx.SetFontFace(face)

	// calculation of dc.fontHeight is different in ctx.SetFontFace and
	// ctx.LoadFontFace for some reason
	// so we have to set dc.fontHeight to not to break existing code
	v := reflect.ValueOf(ctx).Elem()
	field := v.FieldByName("fontHeight")
	setUnexportedField(field, points*72/96)

	return nil
}

func setUnexportedField(field reflect.Value, value interface{}) {
	// copied from https://stackoverflow.com/a/60598827/17863092
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}
