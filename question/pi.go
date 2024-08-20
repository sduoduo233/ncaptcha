package question

import (
	"bytes"
	"slices"

	"github.com/fogleman/gg"
)

func pi() ([]byte, string, []int, error) {
	PI := "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094"

	ans := []int{}

	ctx := gg.NewContext(WIDTH, HEIGHT)
	err := loadFontFace(ctx, 64)
	if err != nil {
		return nil, "", []int{}, err
	}

	ctx.SetRGB(0, 0, 0)
	ctx.Clear()

	_, h := ctx.MeasureString(PI)
	ctx.SetRGB(1, 1, 1)

	for row := range 12 {
		for col := range 3 {
			if randint(0, 15) == 1 {
				ans = append(ans, (row/4)*3+col)
				idx := row*18 + col*6 + randint(0, 5)
				PI = replaceAtIndex(PI, randomDigitRune(string(PI[idx])), idx)
			}
			ctx.DrawString(PI[row*18+col*6:row*18+(col+1)*6], float64((col%3)*200), float64((row/4)*200+(row%4)*int(h))+h)
		}
	}

	buf := &bytes.Buffer{}

	err = ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	return buf.Bytes(), "wrong digit of pi", ans, nil
}
