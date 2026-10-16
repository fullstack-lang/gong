// generated code by xsd2gong - do not edit
package barline

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
	rePattern_2 = regexp.MustCompile("^(([ ]*)|([1-9][0-9]*(, ?[1-9][0-9]*)*))$")
)

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on BarStyleColor.
func (x *BarStyleColor) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("BarStyleColor %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Barline.
func (x *Barline) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Ending.
func (x *Ending) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Number == "" {
		errs = append(errs, fmt.Errorf("Ending %q: required attribute Number (number) is empty", x.Name, "Number", "number"))
	}
	if x.Number != "" && !rePattern_2.MatchString(x.Number) {
		errs = append(errs, fmt.Errorf("Ending %q: field Number value %q does not match XSD pattern", x.Name, x.Number))
	}
	if x.Color != "" && !rePattern_1.MatchString(x.Color) {
		errs = append(errs, fmt.Errorf("Ending %q: field Color value %q does not match XSD pattern", x.Name, x.Color))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Repeat.
func (x *Repeat) ValidateSemantics() (errs []error) {
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
	for _, inst := range stage.GetInstancesSorted[*BarStyleColor]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Barline]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Ending]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Repeat]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	return errs
}
