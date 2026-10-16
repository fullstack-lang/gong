// generated code by xsd2gong - do not edit
package direction

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

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Accord.
func (x *Accord) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if true && (x.TuningOctave < 0 || x.TuningOctave > 9) {
		errs = append(errs, fmt.Errorf("Accord %q: field TuningOctave value %d violates XSD restriction bounds on octave", x.Name, x.TuningOctave))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on AccordionRegistration.
func (x *AccordionRegistration) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("AccordionRegistration %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.AccordionMiddle != 0 && (x.AccordionMiddle < 1 || x.AccordionMiddle > 3) {
		errs = append(errs, fmt.Errorf("AccordionRegistration %q: field AccordionMiddle value %d violates XSD restriction bounds on accordion-middle", x.Name, x.AccordionMiddle))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Barre.
func (x *Barre) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Barre %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Bass.
func (x *Bass) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HarmonyAlter.
func (x *HarmonyAlter) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("HarmonyAlter %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on BassStep.
func (x *BassStep) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("BassStep %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Beater.
func (x *Beater) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on BeatUnitTied.
func (x *BeatUnitTied) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Bracket.
func (x *Bracket) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Bracket %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Bracket %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Dashes.
func (x *Dashes) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Dashes %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Dashes %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Degree.
func (x *Degree) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on DegreeAlter.
func (x *DegreeAlter) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("DegreeAlter %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on DegreeType.
func (x *DegreeType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("DegreeType %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on DegreeValue.
func (x *DegreeValue) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("DegreeValue %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Direction.
func (x *Direction) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on DirectionType.
func (x *DirectionType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Effect.
func (x *Effect) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Feature.
func (x *Feature) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FirstFret.
func (x *FirstFret) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Frame.
func (x *Frame) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Frame %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FrameNote.
func (x *FrameNote) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Glass.
func (x *Glass) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Grouping.
func (x *Grouping) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Harmony.
func (x *Harmony) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Harmony %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HarpPedals.
func (x *HarpPedals) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("HarpPedals %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Image.
func (x *Image) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Source == "" {
		errs = append(errs, fmt.Errorf("Image %q: required attribute Source (source) is empty", x.Name, "Source", "source"))
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("Image %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on InstrumentChange.
func (x *InstrumentChange) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("InstrumentChange %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Inversion.
func (x *Inversion) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Inversion %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Kind.
func (x *Kind) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Kind %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Listening.
func (x *Listening) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MeasureNumbering.
func (x *MeasureNumbering) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("MeasureNumbering %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Membrane.
func (x *Membrane) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Metal.
func (x *Metal) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Metronome.
func (x *Metronome) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Metronome %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MetronomeBeam.
func (x *MetronomeBeam) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 8) {
		errs = append(errs, fmt.Errorf("MetronomeBeam %q: field Number value %d violates XSD restriction bounds on beam-level", x.Name, x.Number))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MetronomeNote.
func (x *MetronomeNote) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MetronomeTied.
func (x *MetronomeTied) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MetronomeTuplet.
func (x *MetronomeTuplet) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Numeral.
func (x *Numeral) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NumeralKey.
func (x *NumeralKey) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NumeralRoot.
func (x *NumeralRoot) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("NumeralRoot %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if true && (x.Value < 1 || x.Value > 7) {
		errs = append(errs, fmt.Errorf("NumeralRoot %q: field Value value %d violates XSD restriction bounds on numeral-value", x.Name, x.Value))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OctaveShift.
func (x *OctaveShift) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("OctaveShift %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("OctaveShift %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Offset.
func (x *Offset) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherDirection.
func (x *OtherDirection) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("OtherDirection %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherListening.
func (x *OtherListening) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("OtherListening %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("OtherListening %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Pedal.
func (x *Pedal) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Pedal %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Pedal %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PedalTuning.
func (x *PedalTuning) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PerMinute.
func (x *PerMinute) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Percussion.
func (x *Percussion) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Percussion %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Pitched.
func (x *Pitched) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PrincipalVoice.
func (x *PrincipalVoice) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("PrincipalVoice %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Print.
func (x *Print) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Root.
func (x *Root) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on RootStep.
func (x *RootStep) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("RootStep %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Scordatura.
func (x *Scordatura) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Sound.
func (x *Sound) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Sound %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on StaffDivide.
func (x *StaffDivide) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("StaffDivide %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Stick.
func (x *Stick) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on StringMute.
func (x *StringMute) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("StringMute %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Swing.
func (x *Swing) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Sync.
func (x *Sync) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.TimeOnly != "" && !rePattern_2.MatchString(x.TimeOnly) {
		errs = append(errs, fmt.Errorf("Sync %q: field TimeOnly value %q does not match XSD pattern", x.Name, x.TimeOnly))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Timpani.
func (x *Timpani) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Wedge.
func (x *Wedge) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("Wedge %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Wedge %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Wood.
func (x *Wood) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on StyleText.
func (x *StyleText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("StyleText %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TimeModification.
func (x *TimeModification) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on VirtualInstrument.
func (x *VirtualInstrument) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateStageSemantics validates all staged instances in this package's Stage.
func ValidateStageSemantics(stage *Stage) (errs []error) {
	if stage == nil {
		return nil
	}
	for _, inst := range stage.GetInstancesSorted[*Accord]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*AccordionRegistration]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Barre]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Bass]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HarmonyAlter]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*BassStep]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Beater]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*BeatUnitTied]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Bracket]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Dashes]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Degree]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*DegreeAlter]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*DegreeType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*DegreeValue]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Direction]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*DirectionType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Effect]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Feature]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FirstFret]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Frame]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FrameNote]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Glass]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Grouping]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Harmony]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HarpPedals]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Image]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*InstrumentChange]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Inversion]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Kind]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Listening]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MeasureNumbering]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Membrane]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Metal]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Metronome]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MetronomeBeam]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MetronomeNote]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MetronomeTied]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MetronomeTuplet]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Numeral]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NumeralKey]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NumeralRoot]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OctaveShift]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Offset]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherDirection]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherListening]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Pedal]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PedalTuning]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PerMinute]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Percussion]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Pitched]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PrincipalVoice]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Print]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Root]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*RootStep]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Scordatura]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Sound]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*StaffDivide]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Stick]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*StringMute]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Swing]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Sync]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Timpani]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Wedge]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Wood]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*StyleText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TimeModification]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*VirtualInstrument]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	return errs
}
