package notepractice

import (
	"bytes"
	"embed"
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed res/*.png
var folder embed.FS

func LoadImage(imageFileName string) *ebiten.Image {
	return loadImage(imageFileName)
}

func loadImage(imageFileName string) *ebiten.Image {
	b, err := folder.ReadFile(imageFileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func LoadFont(fileName string) *text.GoTextFace {
	b, err := folder.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(b))
	if err != nil {
		log.Panic(err)
	}
	fontFace := &text.GoTextFace{
		Source: fontSource,
		Size:   16,
	}
	return fontFace
}

type Rect struct {
	pos  Vector2
	size Vector2
}

type Vector2 struct {
	X float64
	Y float64
}

func isPointInRect(point Vector2, rx, ry, sx, sy float64) bool {
	if point.X > rx+sx || point.X < rx {
		return false
	}
	if point.Y > ry+sy || point.Y < ry {
		return false
	}
	return true
}

func hsvToRGB(h float64, s float64, v float64) (r, g, b uint8) {
	C := v * s
	X := C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := v - C
	var Rnot, Gnot, Bnot float64
	switch {
	case 0 <= h && h < 60:
		Rnot, Gnot, Bnot = C, X, 0
	case 60 <= h && h < 120:
		Rnot, Gnot, Bnot = X, C, 0
	case 120 <= h && h < 180:
		Rnot, Gnot, Bnot = 0, C, X
	case 180 <= h && h < 240:
		Rnot, Gnot, Bnot = 0, X, C
	case 240 <= h && h < 300:
		Rnot, Gnot, Bnot = X, 0, C
	case 300 <= h && h < 360:
		Rnot, Gnot, Bnot = C, 0, X
	}
	r = uint8(math.Round((Rnot + m) * 255))
	g = uint8(math.Round((Gnot + m) * 255))
	b = uint8(math.Round((Bnot + m) * 255))
	return r, g, b
}
