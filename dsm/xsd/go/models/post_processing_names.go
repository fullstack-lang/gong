package models

import (
	"cmp"
	"slices"
)

func prefix(s string) string {
	return s + "_Inlined"
}

func PostProcessingNames(stage *Stage) {

	// map of embedded complex struct within elements
	map_EmbeddedComplexType := make(map[*ComplexType]*Element)
	map_EmbeddedGroup := make(map[*Group]*Element)
	setOfGoIdentifiers := make(map[string]any)

	for _, x := range GetGongstrucsSorted[*ComplexType](stage) {

		if x.IsAnonymous {
			continue
		}
		x.Name = x.NameXSD

		computeGoIdentifier(x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)

		if _x, ok := map_EmbeddedComplexType[x]; ok {
			x.Name = prefix(_x.Name)
		}

		x.ModelGroup.nameRecursively(x.Name)

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
		for _, ag := range x.AttributeGroups {
			ag.Name = prefix(x.Name)
		}
		if x.SimpleContent != nil {
			x.SimpleContent.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Group](stage) {

		if x.Ref != "" {
			continue
		}
		x.Name = x.NameXSD

		computeGoIdentifier("Group_"+x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)

		if _x, ok := map_EmbeddedGroup[x]; ok {
			x.Name = prefix(_x.Name)
		}

		x.ModelGroup.nameRecursively(x.Name)

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*SimpleType](stage) {
		x.Name = x.NameXSD

		if x.Restriction != nil {
			x.Restriction.Name = prefix(x.Name)
		}
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Schema](stage) {
		x.Name = "Schema"
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Restriction](stage) {
		for _, e := range x.Enumerations {
			e.Name = prefix(x.Name)
		}
		if x.MinInclusive != nil {
			x.MinInclusive.Name = prefix(x.Name)
		}
		if x.MaxInclusive != nil {
			x.MaxInclusive.Name = prefix(x.Name)
		}
		if x.Pattern != nil {
			x.Pattern.Name = prefix(x.Name)
		}
		if x.WhiteSpace != nil {
			x.WhiteSpace.Name = prefix(x.Name)
		}
		if x.MinLength != nil {
			x.MinLength.Name = prefix(x.Name)
		}
		if x.MaxLength != nil {
			x.MaxLength.Name = prefix(x.Name)
		}
		if x.Length != nil {
			x.Length.Name = prefix(x.Name)
		}
		if x.TotalDigit != nil {
			x.TotalDigit.Name = prefix(x.Name)
		}

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	//
	// Restrictions
	//
	for _, x := range GetGongstrucsSorted[*Enumeration](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Enumeration](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*MinInclusive](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*MaxInclusive](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Pattern](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*WhiteSpace](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*MinLength](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*MaxLength](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Length](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*TotalDigit](stage) {
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	for _, x := range GetGongstrucsSorted[*Annotation](stage) {
		for _, d := range x.Documentations {
			d.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*Attribute](stage) {
		x.Name = x.NameXSD
		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}
	for _, x := range GetGongstrucsSorted[*AttributeGroup](stage) {
		x.Name = x.NameXSD

		computeGoIdentifier("AttributeGroup_"+x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

	for _, x := range GetGongstrucsSorted[*SimpleContent](stage) {

		if x.Extension != nil {
			x.Extension.Name = prefix(x.Name)
		}
		if x.Restriction != nil {
			x.Restriction.Name = prefix(x.Name)
		}
	}

	elems := GetGongstrucsSorted[*Element](stage)
	slices.SortFunc(elems,
		func(a, b *Element) int {
			return cmp.Compare(a.GetOrder(), b.GetOrder())
		})
	for _, x := range elems {
		x.Name = x.NameXSD
		x.GoIdentifier = xsdNameToGoIdentifier(x.Name)

		if x.ComplexType != nil {
			map_EmbeddedComplexType[x.ComplexType] = x

			if x.ComplexType.Name == "" {
				x.ComplexType.Name = "A_" + x.Name
			}
			computeGoIdentifier(x.ComplexType.Name, &x.ComplexType.WithGoIdentifier, setOfGoIdentifiers)

			setOfGoIdentifiers := make(map[string]any)
			computeGoIdentifier(x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)
		}
		for _, group := range x.Groups {
			map_EmbeddedGroup[group] = x
			computeGoIdentifier(x.Name, &x.WithGoIdentifier, setOfGoIdentifiers)
		}

		if x.Annotation != nil {
			x.Annotation.Name = prefix(x.Name)
		}
	}

}
