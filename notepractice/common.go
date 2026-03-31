package notepractice

import (
	"bytes"
	"embed"
	"image"
	"log"

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
