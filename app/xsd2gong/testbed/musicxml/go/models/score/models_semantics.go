// generated code by xsd2gong - do not edit
package score

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
)

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Credit.
func (x *Credit) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Defaults.
func (x *Defaults) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on EmptyFont.
func (x *EmptyFont) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on GroupBarline.
func (x *GroupBarline) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("GroupBarline %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on GroupName.
func (x *GroupName) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("GroupName %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on GroupSymbol.
func (x *GroupSymbol) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("GroupSymbol %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on InstrumentLink.
func (x *InstrumentLink) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("InstrumentLink %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on LyricFont.
func (x *LyricFont) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on LyricLanguage.
func (x *LyricLanguage) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Lang == "" {
		errs = append(errs, fmt.Errorf("LyricLanguage %q: required attribute Lang (lang) is empty", x.Name, "Lang", "lang"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Opus.
func (x *Opus) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Href == "" {
		errs = append(errs, fmt.Errorf("Opus %q: required attribute Href (href) is empty", x.Name, "Href", "href"))
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("Opus %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PartGroup.
func (x *PartGroup) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PartLink.
func (x *PartLink) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Href == "" {
		errs = append(errs, fmt.Errorf("PartLink %q: required attribute Href (href) is empty", x.Name, "Href", "href"))
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("PartLink %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PartList.
func (x *PartList) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on PartName.
func (x *PartName) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("PartName %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Player.
func (x *Player) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("Player %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ScoreInstrument.
func (x *ScoreInstrument) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("ScoreInstrument %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ScorePart.
func (x *ScorePart) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("ScorePart %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Work.
func (x *Work) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ScorePartwise.
func (x *ScorePartwise) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ScoreTimewise.
func (x *ScoreTimewise) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Part.
func (x *Part) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("Part %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Measure.
func (x *Measure) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number == "" {
		errs = append(errs, fmt.Errorf("Measure %q: required attribute Number (number) is empty", x.Name, "Number", "number"))
	}
	if x.Text != "" && len(x.Text) < 1 {
		errs = append(errs, fmt.Errorf("Measure %q: field Text length %d < minLength 1", x.Name, len(x.Text)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on MeasureMusicDataChoice.
func (x *MeasureMusicDataChoice) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	choiceCount_1 := 0
	if x.Note != nil {
		choiceCount_1++
	}
	if x.Backup != nil {
		choiceCount_1++
	}
	if x.Forward != nil {
		choiceCount_1++
	}
	if x.Direction != nil {
		choiceCount_1++
	}
	if x.Attributes != nil {
		choiceCount_1++
	}
	if x.Harmony != nil {
		choiceCount_1++
	}
	if x.FiguredBass != nil {
		choiceCount_1++
	}
	if x.Print != nil {
		choiceCount_1++
	}
	if x.Sound != nil {
		choiceCount_1++
	}
	if x.Listening != nil {
		choiceCount_1++
	}
	if x.Barline != nil {
		choiceCount_1++
	}
	if x.Grouping != nil {
		choiceCount_1++
	}
	if x.Link != nil {
		choiceCount_1++
	}
	if x.Bookmark != nil {
		choiceCount_1++
	}
	if choiceCount_1 != 1 {
		errs = append(errs, fmt.Errorf("MeasureMusicDataChoice %q: xs:choice requires exactly 1 of [Note, Backup, Forward, Direction, Attributes, Harmony, FiguredBass, Print, Sound, Listening, Barline, Grouping, Link, Bookmark] (got %d)", x.Name, choiceCount_1))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ScoreTimewiseMeasure.
func (x *ScoreTimewiseMeasure) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number == "" {
		errs = append(errs, fmt.Errorf("ScoreTimewiseMeasure %q: required attribute Number (number) is empty", x.Name, "Number", "number"))
	}
	if x.Text != "" && len(x.Text) < 1 {
		errs = append(errs, fmt.Errorf("ScoreTimewiseMeasure %q: field Text length %d < minLength 1", x.Name, len(x.Text)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ScoreTimewiseMeasurePart.
func (x *ScoreTimewiseMeasurePart) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Id == "" {
		errs = append(errs, fmt.Errorf("ScoreTimewiseMeasurePart %q: required attribute Id (id) is empty", x.Name, "Id", "id"))
	}
	return errs
}

// ValidateStageSemantics validates all staged instances in this package's Stage.
func ValidateStageSemantics(stage *Stage) (errs []error) {
	if stage == nil {
		return nil
	}
	for _, inst := range stage.GetInstancesSorted[*Credit]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Defaults]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*EmptyFont]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*GroupBarline]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*GroupName]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*GroupSymbol]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*InstrumentLink]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*LyricFont]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*LyricLanguage]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Opus]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PartGroup]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PartLink]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PartList]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*PartName]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Player]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ScoreInstrument]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ScorePart]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Work]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ScorePartwise]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ScoreTimewise]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Part]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Measure]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*MeasureMusicDataChoice]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ScoreTimewiseMeasure]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ScoreTimewiseMeasurePart]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	return errs
}
