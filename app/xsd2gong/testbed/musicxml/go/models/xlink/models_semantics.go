// generated code by xsd2gong - do not edit
package xlink

import (
	"fmt"
	"regexp"
)

var (
	_ = fmt.Errorf
	_ = regexp.MustCompile
)

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Simple.
func (x *Simple) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("Simple %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	if x.Arcrole != "" && len(x.Arcrole) < 1 {
		errs = append(errs, fmt.Errorf("Simple %q: field Arcrole length %d < minLength 1", x.Name, len(x.Arcrole)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on Extended.
func (x *Extended) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("Extended %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on TitleEltType.
func (x *TitleEltType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ResourceType.
func (x *ResourceType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("ResourceType %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on LocatorType.
func (x *LocatorType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Href == "" {
		errs = append(errs, fmt.Errorf("LocatorType %q: required attribute Href (href) is empty", x.Name, "Href", "href"))
	}
	if x.Role != "" && len(x.Role) < 1 {
		errs = append(errs, fmt.Errorf("LocatorType %q: field Role length %d < minLength 1", x.Name, len(x.Role)))
	}
	return errs
}

// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on ArcType.
func (x *ArcType) ValidateSemantics() (errs []error) {
	if x == nil {
		return nil
	}
	if x.Arcrole != "" && len(x.Arcrole) < 1 {
		errs = append(errs, fmt.Errorf("ArcType %q: field Arcrole length %d < minLength 1", x.Name, len(x.Arcrole)))
	}
	return errs
}

// ValidateStageSemantics validates all staged instances in this package's Stage.
func ValidateStageSemantics(stage *Stage) (errs []error) {
	if stage == nil {
		return nil
	}
	for _, inst := range stage.GetInstancesSorted[*Simple]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*Extended]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*TitleEltType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ResourceType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*LocatorType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	for _, inst := range stage.GetInstancesSorted[*ArcType]() {
		errs = append(errs, inst.ValidateSemantics()...)
	}
	return errs
}
