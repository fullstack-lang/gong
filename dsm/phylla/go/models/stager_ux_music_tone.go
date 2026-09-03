package models

import (
	"fmt"
	"math"

	gongtone_models "github.com/fullstack-lang/gong/lib/tone/go/models"
)

func (stager *Stager) ux_tone_music() {
	stager.toneStage.Reset()

	plant := stager.GetCurrentPlant()
	if plant == nil || plant.PlantType != Music || plant.MusicAbstract == nil {
		stager.toneStage.Commit()
		return
	}

	keyboard := gongtone_models.GeneratePianoNotes()
	if len(keyboard) > 12 {
		keyboard = keyboard[12:] // start at C3
	}

	map_Freqs := make(map[string]*gongtone_models.Freqency)

	beatLength := stager.circumferenceLength / float64(plant.MusicAbstract.NbOfBeatsInTheme)
	if beatLength <= 0 {
		beatLength = 100
	}
	bps := plant.MusicAbstract.BeatsPerSecond
	if bps <= 0 {
		bps = 6.0
	}

	for idx, noteData := range stager.musicNotes {
		if !noteData.IsKept {
			continue
		}

		pitchIdx := noteData.Pitch
		if pitchIdx < 0 {
			pitchIdx = 0
		}
		if pitchIdx >= len(keyboard) {
			pitchIdx = len(keyboard) - 1
		}

		freqNotation := keyboard[pitchIdx]
		freq, ok := map_Freqs[freqNotation]
		if !ok {
			freq = (&gongtone_models.Freqency{Name: freqNotation}).Stage(stager.toneStage)
			map_Freqs[freqNotation] = freq
		}

		note := (&gongtone_models.Note{
			Name:     fmt.Sprintf("Note %d", idx),
			Start:    (noteData.CenterX / beatLength) / bps,
			Duration: 1.0 / bps,
			Velocity: plant.MusicAbstract.Level,
		}).Stage(stager.toneStage)
		note.Frequencies = append(note.Frequencies, freq)
	}

	player := new(gongtone_models.Player).Stage(stager.toneStage)
	player.OnDI = func(p *gongtone_models.Player) error {
		value := (p.Status == gongtone_models.PLAYING)
		if stager.cursor != nil {
			stager.cursor.PlayCursor(stager.cursorStage, value)
		}
		return nil
	}

	stager.toneStage.Commit()
}

func (stager *Stager) ux_cursor_music() {
	if stager.cursor == nil {
		return
	}

	plant := stager.GetCurrentPlant()
	if plant == nil || plant.PlantType != Music || plant.MusicAbstract == nil {
		stager.cursorStage.Commit()
		return
	}

	themeDuration := float64(plant.MusicAbstract.NbOfBeatsInTheme) / plant.MusicAbstract.BeatsPerSecond
	themeVisualLength := stager.circumferenceLength
	if themeVisualLength <= 0 {
		themeVisualLength = 100
	}

	originX := plant.MusicAbstract.OriginX
	originY := plant.MusicAbstract.OriginY
	if originY <= 0 {
		originY = 750.0
	}

	pitchSpacing := plant.MusicAbstract.PitchHeight * plant.RhombusSideLength
	if pitchSpacing <= 0 {
		pitchSpacing = 10
	}

	stager.cursor.Y1 = originY
	stager.cursor.Y2 = originY - float64(plant.MusicAbstract.NbPitchLines-1)*pitchSpacing

	maxX := -math.MaxFloat64
	hasKept := false
	for _, nd := range stager.musicNotes {
		if nd.IsKept {
			hasKept = true
			if nd.CenterX > maxX {
				maxX = nd.CenterX
			}
		}
	}
	if !hasKept {
		maxX = themeVisualLength
	}

	stager.cursor.StartX = originX
	beatWidth := themeVisualLength / float64(plant.MusicAbstract.NbOfBeatsInTheme)
	stager.cursor.EndX = originX + maxX + beatWidth

	if themeVisualLength > 0 && themeDuration > 0 {
		stager.cursor.DurationSeconds = ((stager.cursor.EndX - stager.cursor.StartX) / themeVisualLength) * themeDuration
	} else {
		stager.cursor.DurationSeconds = 10
	}

	stager.cursorStage.Commit()
}
