package mobiletest

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/mobiletest/notepractice"
)

func main() {
	ebiten.SetWindowSize(1080, 2480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(notepractice.NewGame()); err != nil {
		log.Fatal(err)
	}
}
