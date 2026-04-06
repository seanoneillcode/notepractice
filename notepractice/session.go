package notepractice

import (
	"math/rand"
)

type session struct {
	state       string // reading, result
	timer       float32
	currentNote string // a,b,c,d,e,f,g
	sharpFlat   string // nothing, sharp, flat
	trebleBass  string // treble, bass
	index       int
	score       int
	canScore    bool
}

func NewSession() *session {
	return &session{
		state: "reading",
	}
}

func (s *session) update() {
	s.timer = s.timer - 0.01666666 // ebiten is fixed update
}

func (s *session) reset(timer float32) {
	s.state = "reading"
	s.timer = timer
	s.score = 0
	s.canScore = true
	s.nextNote()
}

func getTime(timeOption int) float32 {
	switch timeOption {
	case 0:
		return 60
	case 1:
		return 60 * 2
	case 2:
		return 60 * 4
	case 3:
		return 60 * 8
	}
	return 999
}

func (s *session) nextNote() {
	s.trebleBass = randString([]string{"treble", "bass"})
	s.sharpFlat = randString([]string{"nothing", "sharp", "flat"})
	notes := []string{"A", "B", "C", "D", "E", "F", "G"}

	s.index = rand.Intn(7 * 2) // low c, two octaves

	if s.trebleBass == "treble" {
		noteIndex := (s.index + 2) % 7
		s.currentNote = notes[noteIndex]
	} else {
		noteIndex := (s.index + 4) % 7
		s.currentNote = notes[noteIndex]
	}
	if s.sharpFlat == "sharp" && (s.currentNote == "E" || s.currentNote == "B") {
		s.sharpFlat = "nothing"
	}
	if s.sharpFlat == "flat" && (s.currentNote == "A" || s.currentNote == "F") {
		s.sharpFlat = "nothing"
	}
}

func randString(options []string) string {
	return options[rand.Intn(len(options))]
}
