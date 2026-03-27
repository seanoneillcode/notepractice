package notepractice

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type touch struct {
	originX, originY int
	currX, currY     int
	duration         int
	wasPinch, isPan  bool
}

type inputHandler struct {
	touchIDs []ebiten.TouchID
	touches  map[ebiten.TouchID]*touch

	pos           Vector2
	hasInput      bool
	releasedInput bool
}

func NewInputHandler() *inputHandler {
	return &inputHandler{
		touches: map[ebiten.TouchID]*touch{},
	}
}

func (i *inputHandler) update() error {
	for id := range i.touches {
		if inpututil.IsTouchJustReleased(id) {
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

	// Update the current position and durations of any touches that have
	// neither begun nor ended in this frame.
	for _, id := range i.touchIDs {
		t := i.touches[id]
		t.duration = inpututil.TouchPressDuration(id)
		t.currX, t.currY = ebiten.TouchPosition(id)
	}

	i.hasInput = false
	if len(i.touches) > 0 {
		i.pos.X = float64(i.touches[0].currX)
		i.pos.Y = float64(i.touches[0].currY)
		i.hasInput = true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		i.hasInput = true
		posx, posy := ebiten.CursorPosition()
		i.pos.X, i.pos.Y = float64(posx), float64(posy)
	}
	if !i.hasInput {
		i.releasedInput = true
	}

	return nil
}
