package question

import (
	"bytes"
	"slices"

	"github.com/fogleman/gg"
)

func e() ([]byte, string, []int, error) {
	E := "2.7182818284590452353602874713526624977572470936999595749669676277240766303535475945713821785251664274274663919320030599218174135966290435729003342952605956307381323286279434907632338298807531952510190115738341879307021540891499348841675092447614606680822648001684774118537423454424371075390777449920695517027618386062613313845830007520449338265602976067371132007093287091274437470472306969772093101416"

	ans := []int{}

	ctx := gg.NewContext(WIDTH, HEIGHT)
	err := loadFontFace(ctx, 64)
	if err != nil {
		return nil, "", []int{}, err
	}

	ctx.SetRGB(0, 0, 0)
	ctx.Clear()

	_, h := ctx.MeasureString(E)
	ctx.SetRGB(1, 1, 1)

	for row := range 12 {
		for col := range 3 {
			if randint(0, 15) == 1 {
				ans = append(ans, (row/4)*3+col)
				idx := row*18 + col*6 + randint(0, 5)
				E = replaceAtIndex(E, randomDigitRune(string(E[idx])), idx)
			}
			ctx.DrawString(E[row*18+col*6:row*18+(col+1)*6], float64((col%3)*200), float64((row/4)*200+(row%4)*int(h))+h)
		}
	}

	buf := &bytes.Buffer{}

	err = ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	return buf.Bytes(), "wrong digit of e", ans, nil
}
