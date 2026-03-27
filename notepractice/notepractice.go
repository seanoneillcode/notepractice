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
	images map[string]*ebiten.Image
	font   *text.GoTextFace

	inputHandler *inputHandler
	session      *session
	buttons      *buttons
	clickTimer   float64
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
			"extraLine":           LoadImage("res/extra-line.png"),
			"quitButton":          LoadImage("res/quit-button.png"),
		},
		font:         LoadFont("res/FSEX302.ttf"),
		inputHandler: NewInputHandler(),
		session:      NewSession(),
		buttons:      NewButtons(),
	}
}

func (g *Game) Update() error {
	g.inputHandler.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(clearColor)

	if g.clickTimer > 0 {
		g.clickTimer = g.clickTimer - 0.016
		if g.clickTimer < 0 {
			g.session.nextNote()
			g.buttons.update(g.session)
		}
	}

	if g.inputHandler.hasInput && g.inputHandler.releasedInput && g.clickTimer <= 0 {
		g.drawImage(screen, "noteButtonCorrect", g.inputHandler.pos)
		clicked := false
		for _, b := range g.buttons.allButtons {
			if b.checkCollision(g.inputHandler.pos) {
				g.session.canScore = false
				clicked = true
				if b.note == g.session.currentNote && b.sharpFlat == g.session.sharpFlat {
					b.state = "correct"
					g.session.score = g.session.score + 1
				} else {
					b.state = "incorrect"
				}
			} else {
				b.state = "normal"
			}
		}
		if clicked {
			g.clickTimer = 1.0
		}
	}
	if g.inputHandler.releasedInput {
		g.session.canScore = true
	}

	// drawHeader
	g.drawRect(screen, Vector2{}, Vector2{X: screenWidth, Y: unit}, darkHeaderColor)
	g.drawText(screen, "note practice", Vector2{X: margin, Y: margin}, textColorLight)

	// quit button
	g.drawImage(screen, "quitButton", Vector2{X: 200, Y: 2})
	g.drawText(screen, "quit", Vector2{X: 214, Y: 6}, textColorLight)

	// drawScore
	g.drawRect(screen, Vector2{X: 0, Y: unit}, Vector2{X: screenWidth, Y: unit}, headerColor)
	g.drawText(screen, fmt.Sprintf("score: %d", g.session.score), Vector2{X: margin, Y: margin + unit}, textColorLight)
	g.drawText(screen, fmt.Sprintf("time: %d", int(g.session.timer)), Vector2{X: 160, Y: margin + unit}, textColorLight)

	// treble
	g.drawStave(screen, unit*4)

	// bass
	g.drawStave(screen, unit*10)

	// draw clefs
	g.drawImage(screen, "trebleClef", Vector2{X: margin, Y: 90})
	g.drawImage(screen, "bassClef", Vector2{X: margin, Y: 292})

	// draw note(s)
	g.drawNote(screen, g.session)

	// draw note buttons
	g.drawRect(screen, Vector2{X: 0, Y: unit * 16}, Vector2{X: screenWidth, Y: unit * 5}, buttonBackgroundColor)
	for _, b := range g.buttons.allButtons {
		b.draw(screen, g)
	}
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

func (g *Game) drawNote(screen *ebiten.Image, session *session) {
	drawExtraLine := false
	ypos := 0
	if session.trebleBass == "treble" {
		ypos = 229 - (session.index * 12)
		drawExtraLine = session.index > 11 || session.index == 0
	} else {
		ypos = 397 + (12) - (session.index * 12)
		drawExtraLine = session.index > 11 || session.index == 0
	}

	g.drawImage(screen, "note", Vector2{X: 180, Y: float64(ypos)})

	if drawExtraLine {
		g.drawImage(screen, "extraLine", Vector2{X: 174, Y: float64(ypos + 11)})
	}

	switch session.sharpFlat {
	case "sharp":
		g.drawImage(screen, "sharp", Vector2{X: 152, Y: float64(ypos - 11)})
	case "flat":
		g.drawImage(screen, "flat", Vector2{X: 152, Y: float64(ypos - 20)})
	}
}
