package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/seanoneillcode/mobiletest/notepractice"
)

func main() {
	ebiten.SetWindowSize(270, 602) // 1080, 2480
	ebiten.SetWindowTitle("note practice")

	if err := ebiten.RunGame(notepractice.NewGame()); err != nil {
		log.Fatal(err)
	}
}
