package question

import (
	"bytes"
	"slices"

	"github.com/fogleman/gg"
)

type aabb struct {
	minX, minY, maxX, maxY int
}

func (a *aabb) collides(b *aabb) bool {
	return a.minX <= b.maxX && a.maxX >= b.minX && a.minY <= b.maxY && a.maxY >= b.minY
}

func resistor() ([]byte, string, []int, error) {
	ans := []int{}

	ctx := gg.NewContext(WIDTH, HEIGHT)

	ctx.SetHexColor("#008c4a")
	ctx.DrawRectangle(0, 0, WIDTH, HEIGHT)
	ctx.Fill()

	aabbs := make([]aabb, 0)

	tile0 := aabb{0, 0, 200, 200}
	tile1 := aabb{200, 0, 400, 200}
	tile2 := aabb{400, 0, 600, 200}
	tile3 := aabb{0, 200, 200, 400}
	tile4 := aabb{200, 200, 400, 400}
	tile5 := aabb{400, 200, 600, 400}
	tile6 := aabb{0, 400, 200, 600}
	tile7 := aabb{200, 400, 400, 600}
	tile8 := aabb{400, 400, 600, 600}

	n := randint(100, 150)
	for len(aabbs) < n {
		rotated := randint(0, 2) == 1
		w := 50
		h := 20
		if rotated {
			w, h = h, w
		}
		x := randint(0, WIDTH-w)
		y := randint(0, HEIGHT-h)

		collided := false
		for _, a := range aabbs {
			if a.collides(&aabb{minX: x, minY: y, maxX: x + w, maxY: y + h}) {
				collided = true
				break
			}
		}

		if !collided {
			ctx.DrawRectangle(float64(x), float64(y), float64(w), float64(h))
			ctx.SetHexColor("#ceb58c")
			ctx.Fill()

			black := "#000000"
			brown := "#7e4b26"
			red := "#fd0018"
			orange := "#fd9822"
			yellow := "#fefe32"
			green := "#11b053"
			blue := "#1050ce"
			violet := "#9911fc"
			gray := "#a6a6a6"
			white := "#ffffff"

			colors := []string{black, brown, red, orange, yellow, green, blue, violet, gray, white}

			band1 := randint(0, 10)
			band2 := randint(0, 10)
			r := randint(0, 6)
			multiplier := 1
			multiplierColor := black
			switch r {
			case 1:
				multiplier = 10
				multiplierColor = brown
			case 2:
				multiplier = 100
				multiplierColor = red
			case 3:
				multiplier = 1000
				multiplierColor = orange
			case 4:
				multiplier = 10000
				multiplierColor = yellow
			case 5:
				multiplier = 100000
				multiplierColor = green
			}

			if randint(0, 100) < 2 {
				band1 = 2
				band2 = 2
				multiplier = 10
				multiplierColor = brown
			}

			if !rotated {
				ctx.SetHexColor(colors[band1])
				ctx.DrawRectangle(float64(x+4+4), float64(y), 4, float64(h))
				ctx.Fill()
				ctx.SetHexColor(colors[band2])
				ctx.DrawRectangle(float64(x+12+4), float64(y), 4, float64(h))
				ctx.Fill()
				ctx.SetHexColor(multiplierColor)
				ctx.DrawRectangle(float64(x+20+4), float64(y), 4, float64(h))
				ctx.Fill()
				if randint(0, 100)%2 == 0 {
					ctx.SetHexColor(red)
				} else {
					ctx.SetHexColor(green)
				}
				ctx.DrawRectangle(float64(x+40), float64(y), 4, float64(h))
				ctx.Fill()
			} else {
				ctx.SetHexColor(colors[band1])
				ctx.DrawRectangle(float64(x), float64(y+4+4), float64(w), 4)
				ctx.Fill()
				ctx.SetHexColor(colors[band2])
				ctx.DrawRectangle(float64(x), float64(y+12+4), float64(w), 4)
				ctx.Fill()
				ctx.SetHexColor(multiplierColor)
				ctx.DrawRectangle(float64(x), float64(y+20+4), float64(w), 4)
				ctx.Fill()
				if randint(0, 100)%2 == 0 {
					ctx.SetHexColor(red)
				} else {
					ctx.SetHexColor(green)
				}
				ctx.DrawRectangle(float64(x), float64(y+40), float64(w), 4)
				ctx.Fill()
			}

			value := (band1*10 + band2) * multiplier
			if value == 220 {
				thisAabb := aabb{
					minX: x,
					minY: y,
					maxX: x + w,
					maxY: y + h,
				}
				if tile0.collides(&thisAabb) {
					ans = append(ans, 0)
				}
				if tile1.collides(&thisAabb) {
					ans = append(ans, 1)
				}
				if tile2.collides(&thisAabb) {
					ans = append(ans, 2)
				}
				if tile3.collides(&thisAabb) {
					ans = append(ans, 3)
				}
				if tile4.collides(&thisAabb) {
					ans = append(ans, 4)
				}
				if tile5.collides(&thisAabb) {
					ans = append(ans, 5)
				}
				if tile6.collides(&thisAabb) {
					ans = append(ans, 6)
				}
				if tile7.collides(&thisAabb) {
					ans = append(ans, 7)
				}
				if tile8.collides(&thisAabb) {
					ans = append(ans, 8)
				}
			}

			aabbs = append(aabbs, aabb{
				minX: x,
				minY: y,
				maxX: x + w,
				maxY: y + h,
			})
		}

	}

	buf := &bytes.Buffer{}
	err := ctx.EncodePNG(buf)
	if err != nil {
		return nil, "", []int{}, err
	}

	slices.Sort(ans)
	ans = slices.Compact(ans)

	return buf.Bytes(), "220Ω resistor", ans, nil
}
