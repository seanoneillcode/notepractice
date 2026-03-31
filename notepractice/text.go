package notepractice

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	textCharacterImages = map[rune]*ebiten.Image{}
)

func (g *Game) drawText(screen *ebiten.Image, str string, pos Vector2, color color.Color, scale float64) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(color)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(pos.X, pos.Y)

	y := 0
	const (
		cw = 10
		ch = 12
	)
	for _, c := range str {
		if c == '\n' {
			y = ch
			continue
		}
		s, ok := textCharacterImages[c]
		if !ok {
			cval := int(c)
			index := -1
			if cval > 96 && cval < 123 {
				index = int(c) - 97
			}
			if cval > 47 && cval < 59 {
				index = int(c) - 48 + 26 // the width of the preceding letters
			}
			if c == ',' {
				index = 36
			}
			if c == '.' {
				index = 37
			}
			if c == '!' {
				index = 38
			}
			if c == '?' {
				index = 39
			}
			if c == ':' {
				index = 40
			}
			if c == ' ' {
				op.GeoM.Translate(float64((cw-5))*scale, 0)
			}
			if index != -1 {
				sx := index * cw
				rect := image.Rect(sx, 0, sx+cw-1, ch-1)
				s = g.images["text-source"].SubImage(rect).(*ebiten.Image)
				textCharacterImages[c] = s
			}
		}
		if s != nil {
			op.GeoM.Translate(float64((cw-3))*scale, float64(y)*scale)
			screen.DrawImage(s, op)
		}
	}
}
