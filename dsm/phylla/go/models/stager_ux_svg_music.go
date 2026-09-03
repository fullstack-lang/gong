package models

import (
	"fmt"
	"math"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type MusicNoteData struct {
	BeatNb    int
	Pitch     int
	CenterX   float64
	CenterY   float64
	IsKept    bool
	VoiceName string
}

type MusicBezierSegment struct {
	StartX, StartY                         float64
	ControlPointStartX, ControlPointStartY float64
	ControlPointEndX, ControlPointEndY     float64
	EndX, EndY                             float64
}

func (b *MusicBezierSegment) ComputeYFromX(x float64) float64 {
	const tolerance = 1e-6
	const maxIterations = 100

	bezierX := func(t float64) float64 {
		return math.Pow(1-t, 3)*b.StartX +
			3*math.Pow(1-t, 2)*t*b.ControlPointStartX +
			3*(1-t)*math.Pow(t, 2)*b.ControlPointEndX +
			math.Pow(t, 3)*b.EndX
	}

	bezierXPrime := func(t float64) float64 {
		return -3*math.Pow(1-t, 2)*b.StartX +
			3*(math.Pow(1-t, 2)-2*(1-t)*t)*b.ControlPointStartX +
			3*((1-t)*2*t-math.Pow(t, 2))*b.ControlPointEndX +
			3*math.Pow(t, 2)*b.EndX
	}

	t := 0.5
	for i := 0; i < maxIterations; i++ {
		xAtT := bezierX(t)
		xPrimeAtT := bezierXPrime(t)
		if math.Abs(xPrimeAtT) < 1e-12 {
			break
		}
		deltaT := (xAtT - x) / xPrimeAtT
		t -= deltaT
		if math.Abs(deltaT) < tolerance {
			break
		}
	}
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	return math.Pow(1-t, 3)*b.StartY +
		3*math.Pow(1-t, 2)*t*b.ControlPointStartY +
		3*(1-t)*math.Pow(t, 2)*b.ControlPointEndY +
		math.Pow(t, 3)*b.EndY
}

func quantizePitch(rawY float64, pitchHeight float64, isMinor bool) int {
	if pitchHeight <= 0 {
		pitchHeight = 10
	}
	ratio := rawY / pitchHeight
	pitch := int(math.Floor(ratio + 0.5))
	delta := ratio - float64(pitch)
	pitchAdjustment := 1
	if delta < 0 {
		pitchAdjustment = -1
	}

	relativePitch := pitch % 12
	if relativePitch < 0 {
		relativePitch += 12
	}

	if isMinor {
		switch relativePitch {
		case 1, 4, 6:
			pitch += pitchAdjustment
		case 9:
			pitch -= 1
		case 10:
			pitch += 1
		}
	} else {
		switch relativePitch {
		case 1, 3, 6, 8, 10:
			pitch += pitchAdjustment
		}
	}
	return pitch
}

func (stager *Stager) ux_svg_music() {
	stager.svgMusicStage.Reset()
	stager.musicNotes = nil

	plant := stager.selectedPlant
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}

	if plant == nil || plant.PlantType != Music || plant.MusicAbstract == nil {
		stager.svgMusicStage.Commit()
		return
	}

	ma := plant.MusicAbstract
	N := plant.N
	M := plant.M
	if N < 1 {
		N = 1
	}
	if M < 1 {
		M = 1
	}
	sideLength := plant.RhombusSideLength
	if sideLength <= 0 {
		sideLength = 100
	}
	insideAngle := plant.RhombusInsideAngle
	if insideAngle <= 0 {
		insideAngle = 120
	}

	// 1. Compute geometry
	insideAngleRad := insideAngle * math.Pi / 180
	sinHalfInsideAngle := math.Sin(insideAngleRad / 2)
	cosHalfInsideAngle := math.Cos(insideAngleRad / 2)

	y := float64(N)*sideLength*sinHalfInsideAngle + float64(M)*sideLength*sinHalfInsideAngle
	x := float64(N)*sideLength*cosHalfInsideAngle - float64(M)*sideLength*cosHalfInsideAngle
	initialAxisAngleDegree := math.Atan2(y, x) * 180 / math.Pi
	circumferenceLength := math.Sqrt(x*x + y*y)
	if circumferenceLength <= 0 {
		circumferenceLength = 100
	}
	stager.circumferenceLength = circumferenceLength

	rotRad := -initialAxisAngleDegree * math.Pi / 180
	cosRot := math.Cos(rotRad)
	sinRot := math.Sin(rotRad)

	type pt struct {
		x, y float64
	}

	var nextCircle pt
	minY := math.MaxFloat64
	for i := 0; i <= N; i++ {
		for j := 0; j <= M; j++ {
			cx := float64(i)*sideLength*cosHalfInsideAngle - float64(j)*sideLength*cosHalfInsideAngle
			cy := float64(i)*sideLength*sinHalfInsideAngle + float64(j)*sideLength*sinHalfInsideAngle
			rx := cx*cosRot - cy*sinRot
			ry := cx*sinRot + cy*cosRot
			if ry > 1e-4 && ry < minY {
				minY = ry
				nextCircle = pt{x: rx, y: ry}
			}
		}
	}
	if minY == math.MaxFloat64 {
		nextCircle = pt{x: sideLength, y: sideLength}
	}

	// 2. Compute construction circles
	Z := N + M + 1
	growingCircles := make([]pt, Z)
	for k := 0; k < Z; k++ {
		gx := float64(k) * nextCircle.x
		gy := float64(k) * nextCircle.y
		nbRotations := int(gx / circumferenceLength)
		if gx < 0 {
			nbRotations--
		}
		growingCircles[k] = pt{
			x: gx - float64(nbRotations)*circumferenceLength,
			y: gy,
		}
	}

	circleNPlusM := growingCircles[N+M]
	dx := circleNPlusM.x - circumferenceLength
	dy := circleNPlusM.y
	constAxisAngle := math.Atan2(dy, dx)
	constCircleX := dx / 2.0
	constCircleY := dy / 2.0

	// 3. Compute base growth curve Bezier segments
	controlRatio := ma.BezierControlLengthRatio
	if controlRatio <= 0 {
		controlRatio = 0.56
	}
	angleRad := constAxisAngle - math.Pi/2.0

	baseBeziers := make([]MusicBezierSegment, N+M)
	for k := 0; k < N+M; k++ {
		sX := growingCircles[k].x + constCircleX
		sY := growingCircles[k].y + constCircleY
		var eX, eY float64
		if k+1 < N+M {
			eX = growingCircles[k+1].x + constCircleX
			eY = growingCircles[k+1].y + constCircleY
		} else {
			eX = constCircleX + circumferenceLength
			eY = constCircleY
		}
		baseBeziers[k] = MusicBezierSegment{
			StartX:             sX,
			StartY:             sY,
			ControlPointStartX: sX + sideLength*controlRatio*math.Cos(angleRad),
			ControlPointStartY: sY + sideLength*controlRatio*math.Sin(angleRad),
			ControlPointEndX:   eX + sideLength*controlRatio*math.Cos(angleRad+math.Pi),
			ControlPointEndY:   eY + sideLength*controlRatio*math.Sin(angleRad+math.Pi),
			EndX:               eX,
			EndY:               eY,
		}
	}

	translateBeziers := func(src []MusicBezierSegment, tx, ty float64) []MusicBezierSegment {
		res := make([]MusicBezierSegment, len(src))
		for i, b := range src {
			res[i] = MusicBezierSegment{
				StartX:             b.StartX + tx,
				StartY:             b.StartY + ty,
				ControlPointStartX: b.ControlPointStartX + tx,
				ControlPointStartY: b.ControlPointStartY + ty,
				ControlPointEndX:   b.ControlPointEndX + tx,
				ControlPointEndY:   b.ControlPointEndY + ty,
				EndX:               b.EndX + tx,
				EndY:               b.EndY + ty,
			}
		}
		return res
	}

	// 4. Compute voices
	firstVoiceShiftX := ma.FirstVoiceShiftX * sideLength
	firstVoiceShiftY := ma.FirstVoiceShiftY * sideLength
	firstVoice := translateBeziers(baseBeziers, firstVoiceShiftX, firstVoiceShiftY)
	firstVoiceShiftedRight := translateBeziers(firstVoice, circumferenceLength, 0)

	pitchSpacing := ma.PitchHeight * sideLength
	if pitchSpacing <= 0 {
		pitchSpacing = 10
	}
	secondVoiceShiftY := nextCircle.y + float64(ma.PitchDifference)*pitchSpacing
	secondVoice := translateBeziers(firstVoice, nextCircle.x, secondVoiceShiftY)
	secondVoiceShiftedRight := translateBeziers(secondVoice, circumferenceLength, 0)

	// 5. SVG Canvas Setup
	svgObject := (&svg.SVG{Name: "Music SVG"}).Stage(stager.svgMusicStage)
	layer := (&svg.Layer{Name: "Music Layer"}).Stage(stager.svgMusicStage)
	svgObject.Layers = append(svgObject.Layers, layer)

	originX := ma.OriginX
	originY := ma.OriginY
	if originY <= 0 {
		originY = 750.0
	}

	// 6. Draw Grid: Pitch Lines (horizontal) & Beat Lines (vertical)
	nbBeats := ma.NbOfBeatsInTheme
	if nbBeats <= 0 {
		nbBeats = 16
	}
	beatWidth := circumferenceLength / float64(nbBeats)

	nbPitchLines := ma.NbPitchLines
	if nbPitchLines <= 0 {
		nbPitchLines = 50
	}
	nbBeatLines := ma.NbBeatLines
	if nbBeatLines <= 0 {
		nbBeatLines = 64
	}

	totalWidth := float64(nbBeatLines-1) * beatWidth
	totalHeight := float64(nbPitchLines-1) * pitchSpacing

	// Pitch lines
	for i := 0; i < nbPitchLines; i++ {
		relPitch := i % 12
		if ma.IsMinor {
			switch relPitch {
			case 1, 4, 6, 9, 10:
				continue
			}
		} else {
			switch relPitch {
			case 1, 3, 6, 8, 10:
				continue
			}
		}

		line := (&svg.Line{
			Name: fmt.Sprintf("PitchLine-%d", i),
			X1:   originX,
			Y1:   originY - float64(i)*pitchSpacing,
			X2:   originX + totalWidth,
			Y2:   originY - float64(i)*pitchSpacing,
		}).Stage(stager.svgMusicStage)
		line.Presentation.Stroke = "#dcdcdc"
		line.Presentation.StrokeWidth = 1.0
		line.Presentation.StrokeOpacity = 0.8
		if i%12 == 0 {
			line.Presentation.StrokeWidth = 2.0
			line.Presentation.Stroke = "#bfbfbf"
		}
		layer.Lines = append(layer.Lines, line)
	}

	// Beat lines
	for i := 0; i < nbBeatLines; i++ {
		line := (&svg.Line{
			Name: fmt.Sprintf("BeatLine-%d", i),
			X1:   originX + float64(i)*beatWidth,
			Y1:   originY,
			X2:   originX + float64(i)*beatWidth,
			Y2:   originY - totalHeight,
		}).Stage(stager.svgMusicStage)
		line.Presentation.Stroke = "#dcdcdc"
		line.Presentation.StrokeWidth = 1.0
		line.Presentation.StrokeOpacity = 0.8
		if i%nbBeats == 0 {
			line.Presentation.StrokeWidth = 2.0
			line.Presentation.Stroke = "#999999"
		}
		layer.Lines = append(layer.Lines, line)
	}

	// 7. Draw Voice Wave Curves
	drawVoicePath := func(beziers []MusicBezierSegment, name, strokeColor string, strokeWidth float64, opacity float64) {
		if len(beziers) == 0 {
			return
		}
		path := (&svg.Path{Name: name}).Stage(stager.svgMusicStage)
		path.Presentation.Stroke = strokeColor
		path.Presentation.StrokeWidth = strokeWidth
		path.Presentation.StrokeOpacity = opacity
		path.Presentation.Color = "none"
		path.Presentation.FillOpacity = 0.0

		def := fmt.Sprintf("M %0.2f %0.2f", originX+beziers[0].StartX, originY-beziers[0].StartY)
		for _, b := range beziers {
			def += fmt.Sprintf(" C %0.2f %0.2f, %0.2f %0.2f, %0.2f %0.2f",
				originX+b.ControlPointStartX, originY-b.ControlPointStartY,
				originX+b.ControlPointEndX, originY-b.ControlPointEndY,
				originX+b.EndX, originY-b.EndY,
			)
		}
		path.Definition = def
		layer.Paths = append(layer.Paths, path)
	}

	if ma.ShowFirstVoice {
		drawVoicePath(firstVoice, "First Voice", "#e74c3c", 2.5, 1.0)
	}
	if ma.ShowFirstVoiceShiftRight {
		drawVoicePath(firstVoiceShiftedRight, "First Voice Shift Right", "#888888", 2.0, 0.8)
	}
	if ma.ShowSecondVoice {
		drawVoicePath(secondVoice, "Second Voice", "#2ecc71", 2.5, 1.0)
	}
	if ma.ShowSecondVoiceShiftRight {
		drawVoicePath(secondVoiceShiftedRight, "Second Voice Shift Right", "#a8e6cf", 2.0, 0.8)
	}

	// 8. Compute Notes and Draw Note Rectangles
	computeAndDrawVoiceNotes := func(beziers []MusicBezierSegment, voiceName, strokeColor string, isDisplayed bool) int {
		if len(beziers) == 0 {
			return 0
		}
		rawDist := beziers[0].StartX / beatWidth
		nbMeasureJump := int(math.Floor(rawDist + 0.5))

		for beatNb := 0; beatNb < nbBeats; beatNb++ {
			noteX := float64(beatNb+nbMeasureJump) * beatWidth
			var noteY float64
			found := false
			for _, b := range beziers {
				if b.EndX >= noteX {
					noteY = b.ComputeYFromX(noteX)
					found = true
					break
				}
			}
			if !found && len(beziers) > 0 {
				noteY = beziers[len(beziers)-1].EndY
			}

			pitch := quantizePitch(noteY, pitchSpacing, ma.IsMinor)
			quantizedY := float64(pitch) * pitchSpacing
			isKept := ma.IsNotePlayed(beatNb)

			noteData := &MusicNoteData{
				BeatNb:    beatNb,
				Pitch:     pitch,
				CenterX:   noteX,
				CenterY:   quantizedY,
				IsKept:    isKept,
				VoiceName: voiceName,
			}
			stager.musicNotes = append(stager.musicNotes, noteData)

			if !isDisplayed {
				continue
			}

			rectWidth := 18.0
			rect := (&svg.Rect{
				Name:   fmt.Sprintf("%s-Note-%d", voiceName, beatNb),
				X:      originX + noteX - rectWidth/2.0,
				Y:      originY - quantizedY - rectWidth/2.0,
				Width:  rectWidth,
				Height: rectWidth,
				RX:     2.0,
			}).Stage(stager.svgMusicStage)

			rect.Presentation.Stroke = strokeColor
			if isKept {
				rect.Presentation.StrokeWidth = 2.0
				rect.Presentation.StrokeOpacity = 1.0
				rect.Presentation.Color = "white"
				rect.Presentation.FillOpacity = 0.9
			} else {
				rect.Presentation.Stroke = "#aaaaaa"
				rect.Presentation.StrokeWidth = 1.0
				rect.Presentation.StrokeOpacity = 0.4
				rect.Presentation.Color = "none"
				rect.Presentation.FillOpacity = 0.0
			}

			rect.IsSelectable = true
			capturedBeat := beatNb
			rect.OnSelect = func() {
				plant := stager.GetCurrentPlant()
				if plant != nil && plant.MusicAbstract != nil {
					plant.MusicAbstract.ToggleNotePlayed(capturedBeat)
					stager.stage.Commit()
				}
			}

			rect.HasToolTip = true
			rect.ToolTipText = fmt.Sprintf("%s | Beat: %d | Pitch: %d", voiceName, beatNb, pitch)

			layer.Rects = append(layer.Rects, rect)
		}
		return nbMeasureJump
	}

	jump1 := computeAndDrawVoiceNotes(firstVoice, "1st Voice", "#e74c3c", ma.ShowFirstVoiceNotes)
	computeAndDrawVoiceNotes(firstVoiceShiftedRight, "1st Voice Shift Right", "#888888", ma.ShowFirstVoiceNotesShiftRight)
	jump2 := computeAndDrawVoiceNotes(secondVoice, "2nd Voice", "#2ecc71", ma.ShowSecondVoiceNotes)
	computeAndDrawVoiceNotes(secondVoiceShiftedRight, "2nd Voice Shift Right", "#a8e6cf", ma.ShowSecondVoiceNotesShiftRight)

	ma.ActualBeatsTemporalShift = jump2 - jump1

	stager.svgMusicStage.Commit()
}
