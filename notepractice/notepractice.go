package notepractice

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const scale = 1

type Game struct {
	images   map[string]*ebiten.Image
	font     *text.GoTextFace
	showNote bool

	touchIDs []ebiten.TouchID
	touches  map[ebiten.TouchID]*touch
}

type touch struct {
	originX, originY int
	currX, currY     int
	duration         int
	wasPinch, isPan  bool
}

func NewGame() *Game {
	return &Game{
		images: map[string]*ebiten.Image{
			"noteButton":          LoadImage("res/note-button.png"),
			"noteButtonCorrect":   LoadImage("res/note-button-correct.png"),
			"noteButtonIncorrect": LoadImage("res/note-button-incorrect.png"),
			"note":                LoadImage("res/note.png"),
			"sharp":               LoadImage("res/sharp.png"),
			"flat":                LoadImage("res/flat.png"),
			"trebleClef":          LoadImage("res/treble-clef.png"),
			"bassClef":            LoadImage("res/bass-clef.png"),
			"line":                LoadImage("res/line.png"),
			"quitButton":          LoadImage("res/quit-button.png"),
		},
		font:     LoadFont("res/FSEX302.ttf"),
		showNote: false,
		touches:  map[ebiten.TouchID]*touch{},
	}
}

func (g *Game) Update() error {
	// What touches have just ended?
	for id := range g.touches {
		if inpututil.IsTouchJustReleased(id) {
			delete(g.touches, id)
		}
	}

	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	for _, id := range g.touchIDs {
		x, y := ebiten.TouchPosition(id)
		g.touches[id] = &touch{
			originX: x, originY: y,
			currX: x, currY: y,
		}
	}

	// Update the current position and durations of any touches that have
	// neither begun nor ended in this frame.
	for _, id := range g.touchIDs {
		t := g.touches[id]
		t.duration = inpututil.TouchPressDuration(id)
		t.currX, t.currY = ebiten.TouchPosition(id)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!\n touches: "+fmt.Sprintf("%d", len(g.touches)))

	g.drawImage(screen, "quitButton", Vector2{X: 20, Y: 40})

	g.drawText(screen, "quit", Vector2{X: 20 + 2, Y: 40 + 2})

	if len(g.touches) > 0 {
		g.drawImage(screen, "noteButton", Vector2{X: 50, Y: 140})

		for _, t := range g.touches {
			g.drawImage(screen, "noteButtonCorrect", Vector2{X: float64(t.currX), Y: float64(t.currY)})
		}
	}
}

func (g *Game) drawImage(screen *ebiten.Image, image string, pos Vector2) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(g.images[image], op)
}

func (g *Game) drawText(screen *ebiten.Image, value string, pos Vector2) {
	// Draw variable-width text onto the screen.
	txtOp := &text.DrawOptions{}
	// Start drawing at the top center of the screen.
	txtOp.GeoM.Translate(pos.X, pos.Y)
	// By default, the text is white. We can call ScaleWithColor to specify a different color.
	//colorGreen := color.RGBA{0, 255, 0, 255}
	//txtOp.ColorScale.ScaleWithColor(colorGreen)
	text.Draw(screen, value, g.font, txtOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 270, 602
}
