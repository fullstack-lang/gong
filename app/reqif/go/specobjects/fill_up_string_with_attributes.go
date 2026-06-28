package specobjects

import (
	"log"
	"sort"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/app/reqif/go/xhtml"
)

type TargetEnum string

const (
	Title   TargetEnum = "Title"
	Subject TargetEnum = "Subject"
)

type AttributeWithRank struct {
	Rank  int
	Value string
}

// FillUpStringWithAttributes
func FillUpStringWithAttributes(stager *m.Stager, specObject *m.SPEC_OBJECT, target TargetEnum) (titleComplement string) {

	var attributesWithRank []AttributeWithRank

	attributesWithRank = collectAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_STRING, target, attributesWithRank)
	attributesWithRank = collectXHTMLAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_XHTML, target, attributesWithRank)
	attributesWithRank = collectAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_INTEGER, target, attributesWithRank)
	attributesWithRank = collectAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_REAL, target, attributesWithRank)
	attributesWithRank = collectAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_DATE, target, attributesWithRank)
	attributesWithRank = collectAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_ENUMERATION, target, attributesWithRank)

	sort.SliceStable(attributesWithRank, func(i, j int) bool {
		return attributesWithRank[i].Rank < attributesWithRank[j].Rank
	})

	for i, attr := range attributesWithRank {
		if i > 0 {
			titleComplement += " - "
		}
		titleComplement += attr.Value
	}

	return
}

func collectXHTMLAttributes(stager *m.Stager, attributes []*m.ATTRIBUTE_VALUE_XHTML, target TargetEnum, collected []AttributeWithRank) []AttributeWithRank {

	for _, attr := range attributes {
		rendering := GetAttributeDefinitionRendering(stager, attr)
		var isIn bool
		switch target {
		case Title:
			isIn = *rendering.GetShowInTitlePtr()
		case Subject:
			isIn = *rendering.GetShowInSubjectPtr()
		}

		if isIn {
			// Get the raw XHTML value
			rawXHTML := attr.GetValue()
			var value string

			// Convert the XHTML string to its Markdown equivalent
			markdown, err := xhtml.ToMarkdown(rawXHTML)
			if err != nil {
				// Log the error and fall back to the raw value to avoid crashing
				log.Printf("WARN: Could not convert XHTML to Markdown for attribute. Error: %v", err)
				value = rawXHTML // Fallback
			} else {
				value = markdown // Use the converted Markdown
			}

			rank := *rendering.GetRankPtr()
			collected = append(collected, AttributeWithRank{Rank: rank, Value: value})
		}
	}
	return collected
}

func collectAttributes[Attr m.Attribute](stager *m.Stager, attributes []Attr, target TargetEnum, collected []AttributeWithRank) []AttributeWithRank {

	for _, attr := range attributes {

		rendering := GetAttributeDefinitionRendering(stager, attr)
		var isIn bool

		// fetch the definition of the attribute, check if it is to be retained
		// if yes, append to attributesString
		switch target {
		case Title:
			isIn = *rendering.GetShowInTitlePtr()
		case Subject:
			isIn = *rendering.GetShowInSubjectPtr()
		}

		if isIn {
			rank := *rendering.GetRankPtr()
			collected = append(collected, AttributeWithRank{Rank: rank, Value: attr.GetValue()})
		}
	}

	return collected
}
