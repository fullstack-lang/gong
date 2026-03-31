package models

import (
	"cmp"
	"log"
	"slices"
	"sort"
	"strconv"

	"golang.org/x/exp/maps"
)

// FactorDuplicates performs a simplification of the xsd schema.
//
// For instance, in REQIF schema, the same "ALTERNATIVE-ID" element and its
// inner anonymous complex type are repeated 24 times.
// in order to avoid duplicated go code, one need to factor the xsd before
//
// This function is currently suited for REQIF schema and could be generalized later
func (schema *Schema) FactorDuplicates() {

	// Step 1. elements within choices
	map_NameXSD_Type_Element := make(map[string]map[string][]*Element)
	schema.extractMapOfElementsWithinChoice(map_NameXSD_Type_Element)
	schema.analyseMapOfElements(map_NameXSD_Type_Element)
	schema.factorElementsWithinChoices(map_NameXSD_Type_Element)

	// log.Println("")

	// Step 2 choices within complex types
	map_Choices := make(map[choiceId][]*Choice)
	schema.extractMapOfChoicesWithinComplexTypes(map_Choices)
	schema.analyseMapOfChoices(map_Choices)
	schema.factorChoicesWithinComplexType(map_Choices)

	// log.Println("")

	// Step 3 complex types within elements
	map_ComplexTypes := make(map[complexTypeId][]*ComplexType)
	schema.extractMapOfComplexTypesWithinElements(map_ComplexTypes)
	schema.analyseMapOfComplexTypes(map_ComplexTypes)
	schema.factorComplexTypesWithinElements(map_ComplexTypes)

	// log.Println("")

	// Step 4 : elements within alls
	map_ElementsWithinAlls := make(map[elementId][]*Element)
	schema.extractMapOfElementsWithinAll(map_ElementsWithinAlls)
	schema.analyseMapOfElementWithinAlls(map_ElementsWithinAlls)
	schema.factorElementsWithinAlls(map_ElementsWithinAlls)

	// log.Println("")

	map_ElementsWithinAlls = make(map[elementId][]*Element)
	schema.extractMapOfElementsWithinAll(map_ElementsWithinAlls)
	schema.analyseMapOfElementWithinAlls(map_ElementsWithinAlls)

	// log.Println("")
}

//
// element within all
//

type elementId struct {
	elementName       string
	minOccurs         int
	maxOccurs         int
	choiceElementName string
	choiceElementType string
}

func genElementId(element *Element) (res elementId) {

	choice := element.ComplexType.Choices[0]

	minOccurs, _ := strconv.Atoi(choice.MinOccurs)
	maxOccurs, _ := strconv.Atoi(choice.MaxOccurs)
	res = elementId{
		elementName:       element.NameXSD,
		minOccurs:         minOccurs,
		maxOccurs:         maxOccurs,
		choiceElementName: choice.Elements[0].NameXSD,
		choiceElementType: choice.Elements[0].Type,
	}
	return
}

func (schema *Schema) extractMapOfElementsWithinAll(map_Elements map[elementId][]*Element) {
	for _, complexType := range schema.ComplexTypes {
		for _, alls := range complexType.Alls {
			for _, element := range alls.Elements {
				if _complexType := element.ComplexType; _complexType != nil {
					// get rid of choices that are not with one element only
					if len(_complexType.Elements) > 0 ||
						len(_complexType.Sequences) > 0 ||
						len(_complexType.Alls) > 0 ||
						len(_complexType.Choices) != 1 {
						continue
					}
					elementId := genElementId(element)
					if _, ok := map_Elements[elementId]; ok {
						if !slices.Contains(map_Elements[elementId], element) {
							map_Elements[elementId] = append(map_Elements[elementId], element)
						}
					} else {
						map_Elements[elementId] = []*Element{element}
					}
				}
			}
		}
	}
}

func (*Schema) analyseMapOfElementWithinAlls(map_Elements map[elementId][]*Element) {
	elementIds := maps.Keys(map_Elements)
	slices.SortFunc(elementIds, func(a, b elementId) int {
		return cmp.Compare(a.choiceElementName, b.choiceElementName)
	})

	for _, elementId := range elementIds {
		elements, ok := map_Elements[elementId]
		if ok {
			log.Println("element", elements[0].NameXSD, elementId.choiceElementName, elementId.choiceElementType, len(map_Elements[elementId]))
		}
	}
}

func (*Schema) factorElementsWithinAlls(map_ElementsWithinAlls map[elementId][]*Element) {

	elements := maps.Keys(map_ElementsWithinAlls)
	slices.SortFunc(elements, func(a, b elementId) int {
		return cmp.Compare(a.choiceElementName, b.choiceElementName)
	})

	for _, elementId := range elements {

		sliceElements := map_ElementsWithinAlls[elementId]
		elementToKeep := sliceElements[0]

		// log.Println(len(sliceElements), "elements for element name", elementType, "type", elementType)
		for _, elementToRemoveFromAll := range sliceElements[1:] {
			elementToRemoveFromAll.IsDuplicatedInXSD = true
			if outerAll, ok := elementToRemoveFromAll.OuterParticle.(*All); ok {
				var newElemWithinAll []*Element
				for _, _element := range outerAll.Elements {
					if _element == elementToRemoveFromAll {
						newElemWithinAll = append(newElemWithinAll, elementToKeep)
					} else {
						newElemWithinAll = append(newElemWithinAll, _element)
					}
				}
				outerAll.Elements = newElemWithinAll
			}
		}
	}
}

//
// complex type within element
//

type complexTypeId struct {
	minOccurs         int
	maxOccurs         int
	choiceElementName string
	choiceElementType string
}

func genComplexTypeId(complexType *ComplexType) (res complexTypeId) {

	choice := complexType.Choices[0]

	minOccurs, _ := strconv.Atoi(choice.MinOccurs)
	maxOccurs, _ := strconv.Atoi(choice.MaxOccurs)
	res = complexTypeId{
		minOccurs:         minOccurs,
		maxOccurs:         maxOccurs,
		choiceElementName: choice.Elements[0].NameXSD,
		choiceElementType: choice.Elements[0].Type,
	}
	return
}

func (schema *Schema) extractMapOfComplexTypesWithinElements(map_ComplexTypes map[complexTypeId][]*ComplexType) {
	for _, complexType := range schema.ComplexTypes {
		for _, alls := range complexType.Alls {
			for _, element := range alls.Elements {
				if _complexType := element.ComplexType; _complexType != nil {
					// get rid of choices that are not with one element only
					if len(_complexType.Elements) > 0 ||
						len(_complexType.Sequences) > 0 ||
						len(_complexType.Alls) > 0 ||
						len(_complexType.Choices) != 1 {
						continue
					}
					complexTypeId := genComplexTypeId(_complexType)
					if _, ok := map_ComplexTypes[complexTypeId]; ok {
						if !slices.Contains(map_ComplexTypes[complexTypeId], _complexType) {
							map_ComplexTypes[complexTypeId] = append(map_ComplexTypes[complexTypeId], _complexType)
						}
					} else {
						map_ComplexTypes[complexTypeId] = []*ComplexType{_complexType}
					}
				}
			}
		}
	}
}

func (*Schema) analyseMapOfComplexTypes(map_ComplexTypes map[complexTypeId][]*ComplexType) {
	complextypes := maps.Keys(map_ComplexTypes)
	slices.SortFunc(complextypes, func(a, b complexTypeId) int {
		return cmp.Compare(a.choiceElementName, b.choiceElementName)
	})

	for _, complextypeId := range complextypes {
		_ = complextypeId
		// log.Println("complex type", complextypeId.elementName, complextypeId.elementType, len(map_ComplexTypes[complextypeId]))
	}
}

func (*Schema) factorComplexTypesWithinElements(map_ComplexTypes map[complexTypeId][]*ComplexType) {

	complextypes := maps.Keys(map_ComplexTypes)
	slices.SortFunc(complextypes, func(a, b complexTypeId) int {
		return cmp.Compare(a.choiceElementName, b.choiceElementName)
	})

	for _, complextypeId := range complextypes {

		sliceElements := map_ComplexTypes[complextypeId]
		complexTypeToKeep := sliceElements[0]

		// log.Println(len(sliceElements), "elements for element name", elementType, "type", elementType)
		for _, complexTypeToRemoveFromElement := range sliceElements[1:] {
			complexTypeToRemoveFromElement.IsDuplicatedInXSD = true
			if outerElement, ok := complexTypeToRemoveFromElement.OuterParticle.(*Element); ok {
				if _complexType := outerElement.ComplexType; _complexType != nil {
					if _complexType == complexTypeToRemoveFromElement {
						outerElement.ComplexType = complexTypeToKeep
					}
				}
			}
		}
	}
}

//
// choice within complex type
//

// if choice is with only those elements, one can factor them
type choiceId struct {
	minOccurs         int
	maxOccurs         int
	choiceElementName string
	choiceElementType string
}

// choices within complex type

func genChoiceId(choice *Choice) (res choiceId) {
	minOccurs, _ := strconv.Atoi(choice.MinOccurs)
	maxOccurs, _ := strconv.Atoi(choice.MaxOccurs)
	res = choiceId{
		minOccurs:         minOccurs,
		maxOccurs:         maxOccurs,
		choiceElementName: choice.Elements[0].NameXSD,
		choiceElementType: choice.Elements[0].Type,
	}
	return
}

func (schema *Schema) extractMapOfChoicesWithinComplexTypes(map_Choices map[choiceId][]*Choice) {
	for _, complexType := range schema.ComplexTypes {
		for _, alls := range complexType.Alls {
			for _, element := range alls.Elements {
				if _complexType := element.ComplexType; _complexType != nil {
					for _, choice := range _complexType.Choices {
						// get rid of choices that are not with one element only
						if len(choice.Alls) > 0 ||
							len(choice.Sequences) > 0 ||
							len(choice.Choices) > 0 ||
							len(choice.Elements) != 1 {
							continue
						}
						choiceId := genChoiceId(choice)
						if _, ok := map_Choices[choiceId]; ok {
							if !slices.Contains(map_Choices[choiceId], choice) {
								map_Choices[choiceId] = append(map_Choices[choiceId], choice)
							}
						} else {
							map_Choices[choiceId] = []*Choice{choice}
						}
					}
				}
			}
		}
	}
}

func (*Schema) analyseMapOfChoices(map_Choices map[choiceId][]*Choice) {
	choices := maps.Keys(map_Choices)
	slices.SortFunc(choices, func(a, b choiceId) int {
		return cmp.Compare(a.choiceElementName, b.choiceElementName)
	})

	for _, choiceId := range choices {
		_ = choiceId
		// log.Println("choice", choiceId.elementName, choiceId.elementType, len(map_Choices[choiceId]))
	}
}

func (*Schema) factorChoicesWithinComplexType(map_Choices map[choiceId][]*Choice) {
	for _, sliceElements := range map_Choices {

		firstOfTypeElement := sliceElements[0]

		// log.Println(len(sliceElements), "elements for element name", elementType, "type", elementType)
		for _, choiceToRemoveFromComplexType := range sliceElements[1:] {
			choiceToRemoveFromComplexType.IsDuplicatedInXSD = true
			if complexType, ok := choiceToRemoveFromComplexType.OuterParticle.(*Choice); ok {
				var newChoicesWithinElement []*Choice
				for _, _elem := range complexType.Choices {
					if _elem == choiceToRemoveFromComplexType {
						newChoicesWithinElement = append(newChoicesWithinElement, firstOfTypeElement)
					} else {
						newChoicesWithinElement = append(newChoicesWithinElement, _elem)
					}
				}
				complexType.Choices = newChoicesWithinElement
			}
		}
	}
}

// elements within choices

func (schema *Schema) extractMapOfElementsWithinChoice(map_Elements map[string]map[string][]*Element) {
	for _, namedComplexType := range schema.ComplexTypes {
		for _, all := range namedComplexType.Alls {
			for _, element := range all.Elements {
				element.OuterParticle = all
				if anonymousComplexType := element.ComplexType; anonymousComplexType != nil {
					anonymousComplexType.OuterParticle = element
					for _, choice := range anonymousComplexType.Choices {
						choice.OuterParticle = namedComplexType
						for _, _e := range choice.Elements {
							_e.OuterParticle = choice
							if _, ok := map_Elements[_e.NameXSD]; !ok {
								map_Elements[_e.NameXSD] = make(map[string][]*Element)
							}

							if _, ok := map_Elements[_e.NameXSD][_e.Type]; ok {

								// check that _element is not in the slice already
								if !slices.Contains(map_Elements[_e.NameXSD][_e.Type], _e) {
									map_Elements[_e.NameXSD][_e.Type] = append(
										map_Elements[_e.NameXSD][_e.Type], _e)
								}
							} else {
								map_Elements[_e.NameXSD][_e.Type] = []*Element{_e}
							}
						}
					}
				}
			}
		}
	}
}

func (*Schema) analyseMapOfElements(map_NameXSD_Type_Element map[string]map[string][]*Element) {
	elementsNames := maps.Keys(map_NameXSD_Type_Element)
	sort.Strings(elementsNames)

	for _, elementName := range elementsNames {

		for elementType := range map_NameXSD_Type_Element[elementName] {
			_ = elementType
			// log.Println(elementName, elementType, len(map_NameXSD_Type_Element[elementName][elementType]))
		}
	}
}

func (*Schema) factorElementsWithinChoices(map_NameXSD_Type_Element map[string]map[string][]*Element) {
	for elementNameXSD := range map_NameXSD_Type_Element {
		for elementType := range map_NameXSD_Type_Element[elementNameXSD] {
			if sliceElements, ok := map_NameXSD_Type_Element[elementNameXSD][elementType]; ok {
				firstOfTypeElement := sliceElements[0]

				// log.Println(len(sliceElements), "elements for element name", elementType, "type", elementType)
				for _, elementToRemoveFromChoice := range sliceElements[1:] {
					elementToRemoveFromChoice.IsDuplicatedInXSD = true
					if choice, ok := elementToRemoveFromChoice.OuterParticle.(*Choice); ok {
						var newElementsWithinChoice []*Element
						for _, _elem := range choice.Elements {
							if _elem == elementToRemoveFromChoice {
								newElementsWithinChoice = append(newElementsWithinChoice, firstOfTypeElement)
							} else {
								newElementsWithinChoice = append(newElementsWithinChoice, _elem)
							}
						}
						choice.Elements = newElementsWithinChoice
					}
				}
			}
		}
	}
}
