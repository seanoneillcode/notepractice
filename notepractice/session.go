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
