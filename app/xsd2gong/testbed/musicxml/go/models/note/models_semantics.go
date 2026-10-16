// generated code by xsd2gong - do not edit
package note

import (
	"fmt"
	"regexp"
)

var (
	_ = fmt.Errorf
	_ = regexp.MustCompile
)

var (
	rePattern_1 = regexp.MustCompile("^(#[\\dA-F]{6}([\\dA-F][\\dA-F])?)$")
	rePattern_2 = regexp.MustCompile("^([1-9][0-9]*(, ?[1-9][0-9]*)*)$")
)

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Accidental.
func (x *Accidental) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Accidental %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on AccidentalMark.
func (x *AccidentalMark) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("AccidentalMark %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Arpeggiate.
func (x *Arpeggiate) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Arpeggiate %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Arpeggiate %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Articulations.
func (x *Articulations) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Arrow.
func (x *Arrow) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Arrow %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Assess.
func (x *Assess) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Assess %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Backup.
func (x *Backup) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Beam.
func (x *Beam) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 8) {
		errs = append(errs, fmt.Errorf("Beam %q: field Number value %d violates XSD restriction bounds on beam-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Beam %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Bend.
func (x *Bend) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Bend %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on BreathMark.
func (x *BreathMark) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("BreathMark %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Caesura.
func (x *Caesura) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Caesura %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Elision.
func (x *Elision) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Elision %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyLine.
func (x *EmptyLine) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyLine %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Extend.
func (x *Extend) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Extend %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Figure.
func (x *Figure) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FiguredBass.
func (x *FiguredBass) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("FiguredBass %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Forward.
func (x *Forward) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Glissando.
func (x *Glissando) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Glissando %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Glissando %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Grace.
func (x *Grace) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HammerOnPullOff.
func (x *HammerOnPullOff) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("HammerOnPullOff %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("HammerOnPullOff %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Handbell.
func (x *Handbell) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Handbell %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HarmonClosed.
func (x *HarmonClosed) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HarmonMute.
func (x *HarmonMute) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("HarmonMute %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Harmonic.
func (x *Harmonic) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Harmonic %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HeelToe.
func (x *HeelToe) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("HeelToe %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Hole.
func (x *Hole) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Hole %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HoleClosed.
func (x *HoleClosed) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Instrument.
func (x *Instrument) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("Instrument %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Listen.
func (x *Listen) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Lyric.
func (x *Lyric) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Lyric %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Lyric %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Mordent.
func (x *Mordent) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Mordent %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NonArpeggiate.
func (x *NonArpeggiate) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("NonArpeggiate %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("NonArpeggiate %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Notations.
func (x *Notations) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Note.
func (x *Note) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Note %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Note %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	choiceCount_1 := 0
	if x.Pitch != nil {
		choiceCount_1++
	}
	if x.Unpitched != nil {
		choiceCount_1++
	}
	if x.Rest != nil {
		choiceCount_1++
	}
	if choiceCount_1 != 1 {
		errs = append(errs, fmt.Errorf("Note %q: xs:choice requires exactly 1 of [Pitch, Unpitched, Rest] (got %d)", x.Name, choiceCount_1))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NoteType.
func (x *NoteType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Notehead.
func (x *Notehead) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Notehead %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NoteheadText.
func (x *NoteheadText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Ornaments.
func (x *Ornaments) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherNotation.
func (x *OtherNotation) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("OtherNotation %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("OtherNotation %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherPlacementText.
func (x *OtherPlacementText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("OtherPlacementText %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Pitch.
func (x *Pitch) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if true && (x.Octave < 0 || x.Octave > 9) {
		errs = append(errs, fmt.Errorf("Pitch %q: field Octave value %d violates XSD restriction bounds on octave", x.Name, x.Octave))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PlacementText.
func (x *PlacementText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("PlacementText %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Release.
func (x *Release) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Rest.
func (x *Rest) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.DisplayOctave != 0 && (x.DisplayOctave < 0 || x.DisplayOctave > 9) {
		errs = append(errs, fmt.Errorf("Rest %q: field DisplayOctave value %d violates XSD restriction bounds on octave", x.Name, x.DisplayOctave))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Slide.
func (x *Slide) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Slide %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Slide %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Slur.
func (x *Slur) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Slur %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Slur %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Stem.
func (x *Stem) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Stem %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on StrongAccent.
func (x *StrongAccent) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("StrongAccent %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Tap.
func (x *Tap) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Tap %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Technical.
func (x *Technical) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TextElementData.
func (x *TextElementData) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("TextElementData %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("TextElementData %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("TextElementData %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("TextElementData %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Tie.
func (x *Tie) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Tie %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Tied.
func (x *Tied) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Tied %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Tied %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Tremolo.
func (x *Tremolo) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Tremolo %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if true && (x.Value < 0 || x.Value > 8) {
		errs = append(errs, fmt.Errorf("Tremolo %q: field Value value %d violates XSD restriction bounds on tremolo-marks", x.Name, x.Value))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Tuplet.
func (x *Tuplet) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Tuplet %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TupletDot.
func (x *TupletDot) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("TupletDot %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TupletNumber.
func (x *TupletNumber) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("TupletNumber %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TupletPortion.
func (x *TupletPortion) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TupletType.
func (x *TupletType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("TupletType %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Unpitched.
func (x *Unpitched) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.DisplayOctave != 0 && (x.DisplayOctave < 0 || x.DisplayOctave > 9) {
		errs = append(errs, fmt.Errorf("Unpitched %q: field DisplayOctave value %d violates XSD restriction bounds on octave", x.Name, x.DisplayOctave))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Wait.
func (x *Wait) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Wait %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	return errs
}

// ValidateStageSemantics validates all staged instances in this package's Stage.
func ValidateStageSemantics(stage *Stage) (errs []error) {
	if stage == nil {
		return nil
	}
	for _, inst := range stage.GetInstancesSorted[*Accidental]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*AccidentalMark]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Arpeggiate]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Articulations]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Arrow]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Assess]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Backup]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Beam]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Bend]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*BreathMark]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Caesura]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Elision]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyLine]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Extend]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Figure]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FiguredBass]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Forward]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Glissando]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Grace]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HammerOnPullOff]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Handbell]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HarmonClosed]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HarmonMute]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Harmonic]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HeelToe]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Hole]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HoleClosed]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Instrument]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Listen]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Lyric]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Mordent]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NonArpeggiate]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Notations]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Note]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NoteType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Notehead]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NoteheadText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Ornaments]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherNotation]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherPlacementText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Pitch]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PlacementText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Release]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Rest]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Slide]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Slur]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Stem]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*StrongAccent]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Tap]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Technical]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TextElementData]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Tie]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Tied]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Tremolo]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Tuplet]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TupletDot]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TupletNumber]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TupletPortion]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TupletType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Unpitched]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Wait]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	return errs
}
