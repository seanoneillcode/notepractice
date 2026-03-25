package notepractice

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var darkHeaderColor = color.RGBA{20, 24, 26, 255}
var clearColor = color.RGBA{223, 224, 232, 255}
var headerColor = color.RGBA{58, 63, 94, 255}
var buttonBackgroundColor = color.RGBA{163, 167, 194, 255}
var textColorLight = color.RGBA{223, 224, 232, 255}
var textColorDark = color.RGBA{27, 21, 23, 255}
var lineColor = color.RGBA{27, 21, 23, 255}
var coloredButtonTextColor = color.RGBA{255, 238, 131, 255}

const (
	fontSize           = 16
	fontSpacing        = 2
	noteButtonFontSize = 24

	unit                   = 30
	margin                 = 8
	screenWidth            = 270
	noteButtonSharpsOffset = 18
	noteButtonWidth        = 36
	buttonMargin           = 2
	topButtonMargin        = 15
)

const scale = 1

type Game struct {
	images   map[string]*ebiten.Image
	font     *text.GoTextFace
	showNote bool

	inputHandler *inputHandler
	score        int
	timer        float64
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

		inputHandler: NewInputHandler(),
		score:        0,
	}
}

func (g *Game) Update() error {
	g.inputHandler.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(clearColor)

	if g.inputHandler.hasInput {
		g.drawImage(screen, "noteButtonCorrect", g.inputHandler.pos)
	}

	// drawHeader
	g.drawRect(screen, Vector2{}, Vector2{X: screenWidth, Y: unit}, darkHeaderColor)
	g.drawText(screen, "note practice", Vector2{X: margin, Y: margin}, textColorLight)
	// quit button
	g.drawImage(screen, "quitButton", Vector2{X: 200, Y: 2})
	g.drawText(screen, "quit", Vector2{X: 214, Y: 6}, textColorLight)
	// drawScore
	g.drawRect(screen, Vector2{X: 0, Y: unit}, Vector2{X: screenWidth, Y: unit}, headerColor)
	g.drawText(screen, fmt.Sprintf("score: %d", g.score), Vector2{X: margin, Y: margin + unit}, textColorLight)
	g.drawText(screen, fmt.Sprintf("time: %d", int(g.timer)), Vector2{X: 160, Y: margin + unit}, textColorLight)

	// treble
	g.drawStave(screen, unit*4)

	// bass
	g.drawStave(screen, unit*10)

	// draw clefs
	g.drawImage(screen, "trebleClef", Vector2{X: margin, Y: 90})
	g.drawImage(screen, "bassClef", Vector2{X: margin, Y: 292})

	// draw note(s)
	// g.drawNote(screen, session)

	// draw note buttons
	g.drawRect(screen, Vector2{X: 0, Y: unit * 16}, Vector2{X: screenWidth, Y: unit * 5}, buttonBackgroundColor)

	// for _, b := range buttons.allButtons {
	// 	b.draw(assets)
	// }
}

func (g *Game) drawRect(screen *ebiten.Image, pos Vector2, size Vector2, color color.Color) {
	vector.FillRect(screen, float32(pos.X), float32(pos.Y), float32(size.X), float32(size.Y), color, false)
}

func (g *Game) drawImage(screen *ebiten.Image, image string, pos Vector2) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(g.images[image], op)
}

func (g *Game) drawText(screen *ebiten.Image, value string, pos Vector2, color color.Color) {
	txtOp := &text.DrawOptions{}
	txtOp.GeoM.Translate(pos.X, pos.Y)
	txtOp.ColorScale.ScaleWithColor(color)
	text.Draw(screen, value, g.font, txtOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 270, 602
}

func (g *Game) drawStave(screen *ebiten.Image, startY float64) {
	linePos := Vector2{X: 0, Y: startY}
	for i := range 5 {
		var offset float64 = float64(i) * 24
		g.drawImage(screen, "line", Vector2{X: linePos.X, Y: linePos.Y + offset})
	}
}
