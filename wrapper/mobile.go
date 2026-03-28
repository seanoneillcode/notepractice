package mobiletest

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/seanoneillcode/mobiletest/notepractice"
)

func init() {
	mobile.SetGame(notepractice.NewGame())
}

func Dummy() {}
