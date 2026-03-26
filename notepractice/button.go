package notepractice

import "github.com/hajimehoshi/ebiten/v2"

type buttons struct {
	allButtons []*button
	topRow     []*button
}

func NewButtons() *buttons {
	topRow := newButtonRow(Vector2{X: margin + noteButtonSharpsOffset, Y: unit * 16}, []string{"C", "D", "", "F", "G", "A"})
	bottomRow := newButtonRow(Vector2{X: margin, Y: unit * 18}, []string{"C", "D", "E", "F", "G", "A", "B"})
	allButtons := append(topRow, bottomRow...)
	return &buttons{
		topRow:     topRow,
		allButtons: allButtons,
	}
}

func (b *buttons) update(session *session) {
	sharpFlat := session.sharpFlat
	if sharpFlat == "nothing" {
		sharpFlat = "sharp"
	}
	notes := []string{"C", "D", "F", "G", "A"}
	if session.sharpFlat == "flat" && (session.currentNote == "B" || session.currentNote == "E") {
		notes = []string{"D", "E", "G", "A", "B"}
	}

	for index, button := range b.topRow {
		if index != 2 {
			button.note = notes[index]
		}
		button.sharpFlat = sharpFlat
		var leftMargin float64 = 10
		if sharpFlat != "nothing" {
			leftMargin = 2
		}
		button.leftMargin = leftMargin
	}
	for _, b := range b.allButtons {
		b.state = "normal"
	}
}

func newButtonRow(pos Vector2, notes []string) []*button {
	row := []*button{}
	drawPos := Vector2{X: pos.X + buttonMargin, Y: pos.Y + buttonMargin}
	for i, note := range notes {
		if note == "" {
			continue
		}
		drawPos.X = pos.X + (float64(i) * noteButtonWidth) + buttonMargin
		row = append(row, newButton(drawPos, Vector2{X: 32, Y: 56}, note, "nothing"))
	}
	return row
}

type button struct {
	pos        Vector2
	size       Vector2
	state      string // normal, correct, incorrect
	note       string
	leftMargin float64
	sharpFlat  string
}

func newButton(pos Vector2, size Vector2, note string, sharpFlat string) *button {
	var leftMargin float64 = 10
	if sharpFlat != "nothing" {
		leftMargin = 2
	}
	return &button{
		pos:        pos,
		size:       size,
		state:      "normal",
		note:       note,
		leftMargin: leftMargin,
		sharpFlat:  sharpFlat,
	}
}

func (b *button) draw(screen *ebiten.Image, g *Game) {
	image := "noteButton"
	textColor := textColorDark
	if b.state == "correct" {
		image = "noteButtonCorrect"
		textColor = coloredButtonTextColor
	}
	if b.state == "incorrect" {
		image = "noteButtonIncorrect"
		textColor = coloredButtonTextColor
	}

	g.drawImage(screen, image, b.pos)
	text := b.note
	if b.sharpFlat == "sharp" {
		text = text + "#"
	}
	if b.sharpFlat == "flat" {
		text = text + "b"
	}
	// hurr
	old := g.font.Size
	g.font.Size = noteButtonFontSize
	g.drawText(screen, text, Vector2{X: b.pos.X + b.leftMargin, Y: b.pos.Y + topButtonMargin}, textColor)
	g.font.Size = old
}

func (b *button) checkCollision(mpos Vector2) bool {
	rx, ry, sx, sy := b.pos.X, b.pos.Y, b.size.X, b.size.Y
	if mpos.X > rx+sx || mpos.X < rx {
		return false
	}
	if mpos.Y > ry+sy || mpos.Y < ry {
		return false
	}
	return true
}
