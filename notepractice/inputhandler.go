package notepractice

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type touch struct {
	originX, originY int
	currX, currY     int
}

type inputHandler struct {
	touchIDs []ebiten.TouchID
	touches  map[ebiten.TouchID]*touch

	pos           Vector2
	releasedInput bool
}

func NewInputHandler() *inputHandler {
	return &inputHandler{
		touches: map[ebiten.TouchID]*touch{},
	}
}

func (i *inputHandler) update() error {
	i.releasedInput = false
	for id := range i.touches {
		if inpututil.IsTouchJustReleased(id) {
			i.releasedInput = true
			delete(i.touches, id)
		}
	}

	i.touchIDs = inpututil.AppendJustPressedTouchIDs(i.touchIDs[:0])
	for _, id := range i.touchIDs {
		x, y := ebiten.TouchPosition(id)
		i.touches[id] = &touch{
			originX: x, originY: y,
			currX: x, currY: y,
		}
	}

	// Update the current position of touches
	for id, t := range i.touches {
		t.currX, t.currY = ebiten.TouchPosition(id)
	}

	if len(i.touches) > 0 {
		i.pos.X = float64(i.touches[0].currX)
		i.pos.Y = float64(i.touches[0].currY)
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		posx, posy := ebiten.CursorPosition()
		i.pos.X, i.pos.Y = float64(posx), float64(posy)
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		i.releasedInput = true
	}

	return nil
}
