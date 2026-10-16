// generated code by xsd2gong - do not edit
package common

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
	rePattern_2 = regexp.MustCompile("^([^:Z]*)$")
)

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on AccidentalText.
func (x *AccidentalText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("AccidentalText %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("AccidentalText %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("AccidentalText %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("AccidentalText %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Coda.
func (x *Coda) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Coda %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Dynamics.
func (x *Dynamics) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Dynamics %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("Dynamics %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("Dynamics %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("Dynamics %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Empty.
func (x *Empty) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyPlacement.
func (x *EmptyPlacement) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyPlacement %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyPlacementSmufl.
func (x *EmptyPlacementSmufl) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyPlacementSmufl %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyPrintStyle.
func (x *EmptyPrintStyle) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyPrintStyle %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyPrintStyleAlign.
func (x *EmptyPrintStyleAlign) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyPrintStyleAlign %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyPrintStyleAlignId.
func (x *EmptyPrintStyleAlignId) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyPrintStyleAlignId %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyPrintObjectStyleAlign.
func (x *EmptyPrintObjectStyleAlign) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyPrintObjectStyleAlign %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyTrillSound.
func (x *EmptyTrillSound) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("EmptyTrillSound %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on HorizontalTurn.
func (x *HorizontalTurn) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("HorizontalTurn %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Fermata.
func (x *Fermata) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Fermata %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Fingering.
func (x *Fingering) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Fingering %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FormattedSymbol.
func (x *FormattedSymbol) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("FormattedSymbol %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("FormattedSymbol %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("FormattedSymbol %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("FormattedSymbol %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FormattedSymbolId.
func (x *FormattedSymbolId) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("FormattedSymbolId %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("FormattedSymbolId %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("FormattedSymbolId %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("FormattedSymbolId %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FormattedText.
func (x *FormattedText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("FormattedText %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("FormattedText %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("FormattedText %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("FormattedText %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on FormattedTextId.
func (x *FormattedTextId) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("FormattedTextId %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	if x.Underline != 0 && (x.Underline < 0 || x.Underline > 3) {
		errs = append(errs, fmt.Errorf("FormattedTextId %q: field Underline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Underline))
	}
	if x.Overline != 0 && (x.Overline < 0 || x.Overline > 3) {
		errs = append(errs, fmt.Errorf("FormattedTextId %q: field Overline value %d violates XSD restriction bounds on number-of-lines", x.Name, x.Overline))
	}
	if x.LineThrough != 0 && (x.LineThrough < 0 || x.LineThrough > 3) {
		errs = append(errs, fmt.Errorf("FormattedTextId %q: field LineThrough value %d violates XSD restriction bounds on number-of-lines", x.Name, x.LineThrough))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Fret.
func (x *Fret) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Fret %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Level.
func (x *Level) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MidiDevice.
func (x *MidiDevice) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Port != 0 && (x.Port < 1 || x.Port > 16) {
		errs = append(errs, fmt.Errorf("MidiDevice %q: field Port value %d violates XSD restriction bounds on midi-16", x.Name, x.Port))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MidiInstrument.
func (x *MidiInstrument) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("MidiInstrument %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	if x.MidiChannel != 0 && (x.MidiChannel < 1 || x.MidiChannel > 16) {
		errs = append(errs, fmt.Errorf("MidiInstrument %q: field MidiChannel value %d violates XSD restriction bounds on midi-16", x.Name, x.MidiChannel))
	}
	if x.MidiBank != 0 && (x.MidiBank < 1 || x.MidiBank > 16384) {
		errs = append(errs, fmt.Errorf("MidiInstrument %q: field MidiBank value %d violates XSD restriction bounds on midi-16384", x.Name, x.MidiBank))
	}
	if x.MidiProgram != 0 && (x.MidiProgram < 1 || x.MidiProgram > 128) {
		errs = append(errs, fmt.Errorf("MidiInstrument %q: field MidiProgram value %d violates XSD restriction bounds on midi-128", x.Name, x.MidiProgram))
	}
	if x.MidiUnpitched != 0 && (x.MidiUnpitched < 1 || x.MidiUnpitched > 128) {
		errs = append(errs, fmt.Errorf("MidiInstrument %q: field MidiUnpitched value %d violates XSD restriction bounds on midi-128", x.Name, x.MidiUnpitched))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NameDisplay.
func (x *NameDisplay) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherPlay.
func (x *OtherPlay) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("OtherPlay %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Play.
func (x *Play) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Segno.
func (x *Segno) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Segno %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on StringType.
func (x *StringType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("StringType %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TypedText.
func (x *TypedText) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on WavyLine.
func (x *WavyLine) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number != 0 && (x.Number < 1 || x.Number > 16) {
		errs = append(errs, fmt.Errorf("WavyLine %q: field Number value %d violates XSD restriction bounds on number-level", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("WavyLine %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Encoding.
func (x *Encoding) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.EncodingDate != "" && !rePattern_2.MatchString(x.EncodingDate) {
		errs = append(errs, fmt.Errorf("Encoding %q: field EncodingDate value %q does not match XSD pattern", x.Name, x.EncodingDate))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Identification.
func (x *Identification) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Miscellaneous.
func (x *Miscellaneous) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MiscellaneousField.
func (x *MiscellaneousField) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.NameAttr == "" {
		errs = append(errs, fmt.Errorf("MiscellaneousField %q: required attribute NameAttr (name) is empty", x.Name, "NameAttr", "name"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Supports.
func (x *Supports) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Element == "" {
		errs = append(errs, fmt.Errorf("Supports %q: required attribute Element (element) is empty", x.Name, "Element", "element"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Appearance.
func (x *Appearance) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Distance.
func (x *Distance) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("Distance %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Glyph.
func (x *Glyph) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("Glyph %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on LineWidth.
func (x *LineWidth) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("LineWidth %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MeasureLayout.
func (x *MeasureLayout) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on NoteSize.
func (x *NoteSize) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherAppearance.
func (x *OtherAppearance) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Type == "" {
		errs = append(errs, fmt.Errorf("OtherAppearance %q: required attribute Type (type) is empty", x.Name, "Type", "type"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PageLayout.
func (x *PageLayout) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PageMargins.
func (x *PageMargins) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Scaling.
func (x *Scaling) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on StaffLayout.
func (x *StaffLayout) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on SystemDividers.
func (x *SystemDividers) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on SystemLayout.
func (x *SystemLayout) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on SystemMargins.
func (x *SystemMargins) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Bookmark.
func (x *Bookmark) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("Bookmark %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Link.
func (x *Link) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Href == "" {
		errs = append(errs, fmt.Errorf("Link %q: required attribute Href (href) is empty", x.Name, "Href", "href"))
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("Link %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on OtherText.
func (x *OtherText) ValidateSemantics() (errs []error) {
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
	for _, inst := range stage.GetInstancesSorted[*AccidentalText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Coda]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Dynamics]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Empty]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyPlacement]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyPlacementSmufl]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyPrintStyle]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyPrintStyleAlign]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyPrintStyleAlignId]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyPrintObjectStyleAlign]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyTrillSound]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*HorizontalTurn]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Fermata]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Fingering]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FormattedSymbol]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FormattedSymbolId]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FormattedText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*FormattedTextId]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Fret]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Level]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MidiDevice]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MidiInstrument]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NameDisplay]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherPlay]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Play]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Segno]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*StringType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TypedText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*WavyLine]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Encoding]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Identification]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Miscellaneous]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MiscellaneousField]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Supports]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Appearance]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Distance]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Glyph]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*LineWidth]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MeasureLayout]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*NoteSize]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherAppearance]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PageLayout]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PageMargins]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Scaling]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*StaffLayout]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*SystemDividers]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*SystemLayout]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*SystemMargins]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Bookmark]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Link]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*OtherText]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	return errs
}
