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
var buttonBackgroundColor = color.RGBA{163, 167, 194, 255}
var textColorLight = color.RGBA{223, 224, 232, 255}
var textColorDark = color.RGBA{27, 21, 23, 255}
var coloredButtonTextColor = color.RGBA{255, 238, 131, 255}

const (
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
	g := &Game{
		images: map[string]*ebiten.Image{
			"noteButton":          LoadImage("res/note-button.png"),
			"noteButtonCorrect":   LoadImage("res/note-button-correct.png"),
			"noteButtonIncorrect": LoadImage("res/note-button-incorrect.png"),
			"noteButtonActual":    LoadImage("res/note-button-actual.png"),
			"note":                LoadImage("res/note.png"),
			"sharp":               LoadImage("res/sharp.png"),
			"flat":                LoadImage("res/flat.png"),
			"trebleClef":          LoadImage("res/treble-clef.png"),
			"bassClef":            LoadImage("res/bass-clef.png"),
			"line":                LoadImage("res/line.png"),
			"extraLine":           LoadImage("res/extra-line.png"),
			"guideButton":         LoadImage("res/guide-button.png"),
			"noteA":               LoadImage("res/noteA.png"),
			"noteB":               LoadImage("res/noteB.png"),
			"noteC":               LoadImage("res/noteC.png"),
			"noteD":               LoadImage("res/noteD.png"),
			"noteE":               LoadImage("res/noteE.png"),
			"noteF":               LoadImage("res/noteF.png"),
			"noteG":               LoadImage("res/noteG.png"),
			"sharpmod":            LoadImage("res/sharpmod.png"),
			"flatmod":             LoadImage("res/flatmod.png"),
		},
		font:         LoadFont("res/FSEX302.ttf"),
		inputHandler: NewInputHandler(),
		session:      NewSession(),
		buttons:      NewButtons(),
	}

	g.session.nextNote()
	g.buttons.reset(g.session)

	return g
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
			g.buttons.reset(g.session)
		}
	}

	if g.inputHandler.hasInput && g.inputHandler.releasedInput && g.clickTimer <= 0 {
		g.drawImage(screen, "noteButtonCorrect", g.inputHandler.pos)
		clickedAnyButton := false
		for _, b := range g.buttons.allButtons {
			if b.checkCollision(g.inputHandler.pos) {
				clickedAnyButton = true
				g.session.canScore = false
				if b.note == g.session.currentNote && b.sharpFlat == g.session.sharpFlat {
					b.state = "correct"
					g.session.score = g.session.score + 1
				} else {
					b.state = "incorrect"
				}
			}
		}
		if clickedAnyButton {
			g.clickTimer = 1.0
			for _, b := range g.buttons.allButtons {
				if b.state == "normal" {
					if b.note == g.session.currentNote && b.sharpFlat == g.session.sharpFlat {
						b.state = "actual"
					}
				}
			}
		}
	}
	if g.inputHandler.releasedInput {
		g.session.canScore = true
	}

	// drawHeader
	g.drawRect(screen, Vector2{}, Vector2{X: screenWidth, Y: unit}, darkHeaderColor)
	g.drawImage(screen, "guideButton", Vector2{X: 232, Y: 2})

	// drawScore
	g.drawText(screen, fmt.Sprintf("score: %d", g.session.score), Vector2{X: margin, Y: margin}, textColorLight)

	// treble
	g.drawStave(screen, unit*3)

	// bass
	g.drawStave(screen, unit*9)

	// draw clefs
	g.drawImage(screen, "trebleClef", Vector2{X: margin, Y: 60})
	g.drawImage(screen, "bassClef", Vector2{X: margin, Y: 262})

	// draw note(s)
	g.drawNote(screen, g.session)

	// draw note buttons
	g.drawRect(screen, Vector2{X: 0, Y: unit * 15}, Vector2{X: screenWidth, Y: unit * 4}, buttonBackgroundColor)
	for _, b := range g.buttons.allButtons {
		b.draw(screen, g)
	}

	g.drawRect(screen, Vector2{Y: unit * 19}, Vector2{X: screenWidth, Y: unit * 2}, darkHeaderColor)
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

func (g *Game) drawImageWithColor(screen *ebiten.Image, image string, pos Vector2, color color.Color) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pos.X, pos.Y)
	op.GeoM.Scale(scale, scale)
	op.ColorScale.ScaleWithColor(color)
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
		ypos = 200 - (session.index * 12)
		drawExtraLine = session.index > 11 || session.index == 0
	} else {
		ypos = 380 - (session.index * 12)
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
