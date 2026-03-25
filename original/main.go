package original

// import (
// 	"flag"
// 	"fmt"
// 	"math/rand"
// 	// rl "github.com/gen2brain/raylib-go/raylib"
// )

// var darkHeaderColor = rl.NewColor(20, 24, 26, 255)
// var clearColor = rl.NewColor(223, 224, 232, 255)
// var headerColor = rl.NewColor(58, 63, 94, 255)
// var buttonBackgroundColor = rl.NewColor(163, 167, 194, 255)
// var textColorLight = rl.NewColor(223, 224, 232, 255)
// var textColorDark = rl.NewColor(27, 21, 23, 255)
// var lineColor = rl.NewColor(27, 21, 23, 255)
// var coloredButtonTextColor = rl.NewColor(255, 238, 131, 255)

// const (
// 	fontSize           = 16
// 	fontSpacing        = 2
// 	noteButtonFontSize = 24

// 	unit                   = 30
// 	margin                 = 8
// 	screenWidth            = 270
// 	noteButtonSharpsOffset = 18
// 	noteButtonWidth        = 36
// 	buttonMargin           = 2
// 	topButtonMargin        = 15

// 	// scale = 4
// )

// func main() {

// 	fullScreenPtr := flag.Bool("fullscreen", false, "")
// 	scalePtr := flag.Bool("scale", false, "")
// 	flag.Parse()

// 	if *fullScreenPtr {
// 		rl.SetWindowState(rl.FlagFullscreenMode)
// 	}
// 	scale := 4
// 	windowWidth := int32(1080)
// 	windowHeight := int32(2408)
// 	if *scalePtr {
// 		scale = 1
// 		windowWidth = int32(270)
// 		windowHeight = int32(602)
// 	}

// 	// rl.InitWindow(270, 602, "note practice") // 1080, 2408
// 	rl.InitWindow(windowWidth, windowHeight, "note practice") // 1080, 2408

// 	rl.SetWindowState(rl.FlagVsyncHint)

// 	// rl.HideCursor()
// 	rl.SetTargetFPS(60)

// 	var camera rl.Camera2D
// 	camera.Zoom = float32(scale)

// 	assets := newAssets()
// 	session := newSession()
// 	buttons := newButtons()

// 	session.nextNote()
// 	buttons.update(session)

// 	var clickTimer float32 = 0

// 	pressedQuit := false
// 	for !rl.WindowShouldClose() && !pressedQuit {

// 		if rl.IsKeyPressed(rl.KeySpace) {
// 			session.nextNote()
// 		}
// 		if clickTimer > 0 {
// 			clickTimer = clickTimer - rl.GetFrameTime()
// 			if clickTimer < 0 {
// 				session.nextNote()
// 				buttons.update(session)
// 			}
// 		}
// 		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
// 			mpos := rl.GetMousePosition()
// 			clicked := false
// 			for _, b := range buttons.allButtons {
// 				if b.checkCollision(mpos) {
// 					clicked = true
// 					if b.note == session.currentNote && b.sharpFlat == session.sharpFlat {
// 						b.state = "correct"
// 						session.score = session.score + 1
// 					} else {
// 						b.state = "incorrect"
// 					}
// 				} else {
// 					b.state = "normal"
// 				}
// 			}
// 			if clicked {
// 				clickTimer = 1.0
// 			}
// 			quitRect := rl.Rectangle{X: 200, Y: 2, Width: 57, Height: 25}
// 			if rl.CheckCollisionPointRec(mpos, quitRect) {
// 				pressedQuit = true
// 			}
// 		}
// 		session.timer = session.timer + rl.GetFrameTime()

// 		rl.BeginDrawing()
// 		rl.ClearBackground(clearColor)
// 		rl.BeginMode2D(camera)

// 		// drawHeader
// 		rl.DrawRectangleV(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: screenWidth, Y: unit}, darkHeaderColor)
// 		rl.DrawTextEx(assets.font, "note practice", rl.Vector2{X: margin, Y: margin}, fontSize, fontSpacing, textColorLight)
// 		rl.DrawTextureEx(assets.quitButton, rl.Vector2{X: 200, Y: 2}, 0, 1, rl.White)
// 		rl.DrawTextEx(assets.font, "quit", rl.Vector2{X: 214, Y: 6}, fontSize, fontSpacing, textColorLight)

// 		// drawScore
// 		rl.DrawRectangleV(rl.Vector2{X: 0, Y: unit}, rl.Vector2{X: screenWidth, Y: unit}, headerColor)
// 		rl.DrawTextEx(assets.font, fmt.Sprintf("score: %d", session.score), rl.Vector2{X: margin, Y: margin + unit}, fontSize, fontSpacing, textColorLight)
// 		rl.DrawTextEx(assets.font, fmt.Sprintf("time: %d", int(session.timer)), rl.Vector2{X: 160, Y: margin + unit}, fontSize, fontSpacing, textColorLight)

// 		// treble
// 		drawStave(unit*4, assets)

// 		// bass
// 		drawStave(unit*10, assets)

// 		// draw clefs
// 		rl.DrawTextureEx(assets.trebleClef, rl.Vector2{X: margin, Y: 90}, 0, 1, rl.White)
// 		rl.DrawTextureEx(assets.bassClef, rl.Vector2{X: margin, Y: 292}, 0, 1, rl.White)

// 		// draw note(s)
// 		drawNote(assets, session)

// 		// draw note buttons
// 		rl.DrawRectangleV(rl.Vector2{X: 0, Y: unit * 16}, rl.Vector2{X: screenWidth, Y: unit * 5}, buttonBackgroundColor)

// 		for _, b := range buttons.allButtons {
// 			b.draw(assets)
// 		}

// 		rl.EndMode2D()
// 		rl.EndDrawing()
// 	}

// 	assets.cleanup()
// 	rl.CloseWindow()
// }

// func drawNote(assets *assets, session *session) {

// 	drawExtraLine := false
// 	ypos := 0
// 	if session.trebleBass == "treble" {
// 		ypos = 229 - (session.index * 12)
// 		drawExtraLine = session.index > 11 || session.index == 0
// 	} else {
// 		ypos = 397 + (12) - (session.index * 12)
// 		drawExtraLine = session.index > 11 || session.index == 0
// 	}

// 	rl.DrawTextureEx(assets.note, rl.Vector2{X: 180, Y: float32(ypos)}, 0, 1, rl.White)

// 	if drawExtraLine {
// 		rl.DrawLineV(rl.Vector2{X: 174, Y: float32(ypos + 11)}, rl.Vector2{X: 208, Y: float32(ypos + 11)}, lineColor)
// 		rl.DrawLineV(rl.Vector2{X: 174, Y: float32(ypos + 11 + 1)}, rl.Vector2{X: 208, Y: float32(ypos + 11 + 1)}, lineColor)
// 		rl.DrawLineV(rl.Vector2{X: 174, Y: float32(ypos + 11 + 2)}, rl.Vector2{X: 208, Y: float32(ypos + 11 + 2)}, lineColor)
// 	}

// 	switch session.sharpFlat {
// 	case "sharp":
// 		rl.DrawTextureEx(assets.sharp, rl.Vector2{X: 152, Y: float32(ypos - 11)}, 0, 1, rl.White)
// 	case "flat":
// 		rl.DrawTextureEx(assets.flat, rl.Vector2{X: 154, Y: float32(ypos) - 20}, 0, 1, rl.White)
// 	}
// }

// func drawStave(startY float32, assets *assets) {
// 	linePos := rl.Vector2{X: 0, Y: startY}
// 	for i := range 5 {
// 		var offset float32 = float32(i) * 24
// 		rl.DrawTextureEx(assets.line, rl.Vector2Add(linePos, rl.Vector2{Y: offset}), 0, 1, rl.White)
// 	}
// }

// func newButtonRow(pos rl.Vector2, notes []string) []*button {
// 	row := []*button{}
// 	drawPos := rl.Vector2{X: pos.X + buttonMargin, Y: pos.Y + buttonMargin}
// 	for i, note := range notes {
// 		if note == "" {
// 			continue
// 		}
// 		drawPos.X = pos.X + (float32(i) * noteButtonWidth) + buttonMargin
// 		row = append(row, newButton(drawPos, rl.Vector2{X: 32, Y: 56}, note, "nothing"))
// 	}
// 	return row
// }

// type buttons struct {
// 	allButtons []*button
// 	topRow     []*button
// }

// func newButtons() *buttons {
// 	topRow := newButtonRow(rl.Vector2{X: margin + noteButtonSharpsOffset, Y: unit * 16}, []string{"C", "D", "", "F", "G", "A"})
// 	bottomRow := newButtonRow(rl.Vector2{X: margin, Y: unit * 18}, []string{"C", "D", "E", "F", "G", "A", "B"})
// 	allButtons := append(topRow, bottomRow...)
// 	return &buttons{
// 		topRow:     topRow,
// 		allButtons: allButtons,
// 	}
// }

// func (b *buttons) update(session *session) {
// 	sharpFlat := session.sharpFlat
// 	if sharpFlat == "nothing" {
// 		sharpFlat = "sharp"
// 	}
// 	notes := []string{"C", "D", "F", "G", "A"}
// 	if session.sharpFlat == "flat" && (session.currentNote == "B" || session.currentNote == "E") {
// 		notes = []string{"D", "E", "G", "A", "B"}
// 	}

// 	for index, button := range b.topRow {
// 		if index != 2 {
// 			button.note = notes[index]
// 		}
// 		button.sharpFlat = sharpFlat
// 		var leftMargin float32 = 10
// 		if sharpFlat != "nothing" {
// 			leftMargin = 2
// 		}
// 		button.leftMargin = leftMargin
// 	}
// 	for _, b := range b.allButtons {
// 		b.state = "normal"
// 	}
// }

// type assets struct {
// 	noteButton          rl.Texture2D
// 	noteButtonCorrect   rl.Texture2D
// 	noteButtonIncorrect rl.Texture2D
// 	note                rl.Texture2D
// 	sharp               rl.Texture2D
// 	flat                rl.Texture2D
// 	trebleClef          rl.Texture2D
// 	bassClef            rl.Texture2D
// 	line                rl.Texture2D
// 	quitButton          rl.Texture2D
// 	font                rl.Font
// }

// func newAssets() *assets {
// 	return &assets{
// 		noteButton:          rl.LoadTexture("res/note-button.png"),
// 		noteButtonCorrect:   rl.LoadTexture("res/note-button-correct.png"),
// 		noteButtonIncorrect: rl.LoadTexture("res/note-button-incorrect.png"),
// 		note:                rl.LoadTexture("res/note.png"),
// 		sharp:               rl.LoadTexture("res/sharp.png"),
// 		flat:                rl.LoadTexture("res/flat.png"),
// 		trebleClef:          rl.LoadTexture("res/treble-clef.png"),
// 		bassClef:            rl.LoadTexture("res/bass-clef.png"),
// 		line:                rl.LoadTexture("res/line.png"),
// 		quitButton:          rl.LoadTexture("res/quit-button.png"),
// 		font:                rl.LoadFont("res/FSEX302.ttf"),
// 	}
// }

// func (a *assets) cleanup() {
// 	rl.UnloadTexture(a.noteButton)
// 	rl.UnloadTexture(a.note)
// 	rl.UnloadTexture(a.sharp)
// 	rl.UnloadTexture(a.flat)
// 	rl.UnloadTexture(a.trebleClef)
// 	rl.UnloadTexture(a.bassClef)
// 	rl.UnloadTexture(a.line)
// 	rl.UnloadTexture(a.quitButton)
// 	rl.UnloadFont(a.font)
// }

// type session struct {
// 	state       string // reading, result
// 	timer       float32
// 	currentNote string // a,b,c,d,e,f,g
// 	sharpFlat   string // nothing, sharp, flat
// 	trebleBass  string // treble, bass
// 	index       int
// 	score       int
// }

// func newSession() *session {
// 	return &session{
// 		state: "reading",
// 	}
// }

// func (s *session) nextNote() {
// 	s.trebleBass = randString([]string{"treble", "bass"})
// 	s.sharpFlat = randString([]string{"nothing", "sharp", "flat"})
// 	notes := []string{"A", "B", "C", "D", "E", "F", "G"}

// 	s.index = rand.Intn(7 * 2) // low c, two octaves

// 	if s.trebleBass == "treble" {
// 		noteIndex := (s.index + 2) % 7
// 		s.currentNote = notes[noteIndex]
// 	} else {
// 		noteIndex := (s.index + 4) % 7
// 		s.currentNote = notes[noteIndex]
// 	}
// 	if s.sharpFlat == "sharp" && (s.currentNote == "E" || s.currentNote == "B") {
// 		s.sharpFlat = "nothing"
// 	}
// 	if s.sharpFlat == "flat" && (s.currentNote == "A" || s.currentNote == "F") {
// 		s.sharpFlat = "nothing"
// 	}
// }

// func randString(options []string) string {
// 	return options[rand.Intn(len(options))]
// }

// type button struct {
// 	pos        rl.Vector2
// 	size       rl.Vector2
// 	state      string // normal, correct, incorrect
// 	note       string
// 	leftMargin float32
// 	sharpFlat  string
// }

// func newButton(pos rl.Vector2, size rl.Vector2, note string, sharpFlat string) *button {
// 	var leftMargin float32 = 10
// 	if sharpFlat != "nothing" {
// 		leftMargin = 2
// 	}
// 	return &button{
// 		pos:        pos,
// 		size:       size,
// 		state:      "normal",
// 		note:       note,
// 		leftMargin: leftMargin,
// 		sharpFlat:  sharpFlat,
// 	}
// }

// func (b *button) draw(assets *assets) {
// 	image := assets.noteButton
// 	textColor := textColorDark
// 	if b.state == "correct" {
// 		image = assets.noteButtonCorrect
// 		textColor = coloredButtonTextColor
// 	}
// 	if b.state == "incorrect" {
// 		image = assets.noteButtonIncorrect
// 		textColor = coloredButtonTextColor
// 	}
// 	rl.DrawTextureEx(image, b.pos, 0, 1, rl.White)
// 	text := b.note
// 	if b.sharpFlat == "sharp" {
// 		text = text + "#"
// 	}
// 	if b.sharpFlat == "flat" {
// 		text = text + "b"
// 	}
// 	rl.DrawTextEx(assets.font, text, rl.Vector2Add(b.pos, rl.Vector2{X: b.leftMargin, Y: topButtonMargin}), noteButtonFontSize, fontSpacing, textColor)
// }

// func (b *button) checkCollision(mpos rl.Vector2) bool {
// 	rect := rl.Rectangle{X: b.pos.X, Y: b.pos.Y, Width: b.size.X, Height: b.size.Y}
// 	if rl.CheckCollisionPointRec(mpos, rect) {
// 		return true
// 	}
// 	return false
// }
