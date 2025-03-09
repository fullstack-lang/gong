package models

import "fmt"

type Note struct {
	Name string

	Frequencies []*Freqency

	Start float64

	Duration float64

	Velocity float64

	// to help to decode
	Info string
}

type Freqency struct {
	Name string
}

const (
	A = "A"
	B = "B"
	C = "C"
	D = "D"
	E = "E"
	F = "F"
	G = "G"
)

var Letters = []string{C, D, E, F, G, A, B}

const (
	DoubleFlat  = "bb"
	Flat        = "b"
	Natural     = ""
	Sharp       = "#"
	DoubleSharp = "x"
)

var Accidentals = []string{DoubleFlat, Flat, Natural, Sharp, DoubleSharp}

const (
	OctaveMinus4 = -4
	OctaveMinus3 = -3
	OctaveMinus2 = -2
	OctaveMinus1 = -1
	Octave0      = 0
	Octave1      = 1
	Octave2      = 2
	Octave3      = 3
	Octave4      = 4
	Octave5      = 5
	Octave6      = 6
	Octave7      = 7
	Octave8      = 8
	Octave9      = 9
	Octave10     = 10
	Octave11     = 11
)

var Octaves = []int{OctaveMinus4, OctaveMinus3, OctaveMinus2, OctaveMinus1, Octave0, Octave1, Octave2, Octave3, Octave4, Octave5, Octave6, Octave7, Octave8, Octave9, Octave10, Octave11}

func GeneratePianoNotes() []string {
	var notes []string
	octave := 0
	for octave < 9 {
		for _, letter := range Letters {
			note := fmt.Sprintf("%s%s%d", letter, Natural, octave)
			notes = append(notes, note)
			switch letter {
			case C, D, F, G, A:
				note := fmt.Sprintf("%s%s%d", letter, Sharp, octave)
				notes = append(notes, note)

			}
		}
		octave++
	}
	return notes
}
