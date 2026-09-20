// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder

	if stageSet.Stage != nil {
		allOrdered := []*All{}
		for all := range stageSet.Stage.Alls {
			allOrdered = append(allOrdered, all)
		}
		sort.Slice(allOrdered, func(i, j int) bool {
			return stageSet.Stage.All_stagedOrder[allOrdered[i]] < stageSet.Stage.All_stagedOrder[allOrdered[j]]
		})
		for _, all := range allOrdered {
			allIdent := "__stage_0" + all.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.All{Name: %s}).Stage(stageSet.Stage)", allIdent, __gong__toRawStringLiteral(all.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allIdent, __gong__toRawStringLiteral(all.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", allIdent, __gong__toRawStringLiteral(all.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", allIdent, all.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", allIdent, all.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", allIdent, __gong__toRawStringLiteral(all.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", allIdent, __gong__toRawStringLiteral(all.MaxOccurs)))
			if all.Annotation != nil {
				targetIdent := "__stage_0" + all.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", allIdent, targetIdent))
			}
			for _, elem := range all.Sequences {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Alls {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Choices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", allIdent, allIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		annotationOrdered := []*Annotation{}
		for annotation := range stageSet.Stage.Annotations {
			annotationOrdered = append(annotationOrdered, annotation)
		}
		sort.Slice(annotationOrdered, func(i, j int) bool {
			return stageSet.Stage.Annotation_stagedOrder[annotationOrdered[i]] < stageSet.Stage.Annotation_stagedOrder[annotationOrdered[j]]
		})
		for _, annotation := range annotationOrdered {
			annotationIdent := "__stage_0" + annotation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Annotation{Name: %s}).Stage(stageSet.Stage)", annotationIdent, __gong__toRawStringLiteral(annotation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", annotationIdent, __gong__toRawStringLiteral(annotation.Name)))
			for _, elem := range annotation.Documentations {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Documentations = append(%s.Documentations, %s)", annotationIdent, annotationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		attributeOrdered := []*Attribute{}
		for attribute := range stageSet.Stage.Attributes {
			attributeOrdered = append(attributeOrdered, attribute)
		}
		sort.Slice(attributeOrdered, func(i, j int) bool {
			return stageSet.Stage.Attribute_stagedOrder[attributeOrdered[i]] < stageSet.Stage.Attribute_stagedOrder[attributeOrdered[j]]
		})
		for _, attribute := range attributeOrdered {
			attributeIdent := "__stage_0" + attribute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Attribute{Name: %s}).Stage(stageSet.Stage)", attributeIdent, __gong__toRawStringLiteral(attribute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", attributeIdent, __gong__toRawStringLiteral(attribute.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", attributeIdent, attribute.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", attributeIdent, __gong__toRawStringLiteral(attribute.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Default)))
			values.WriteString(fmt.Sprintf("\n\t%s.Use = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Use)))
			values.WriteString(fmt.Sprintf("\n\t%s.Form = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Form)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fixed = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Fixed)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", attributeIdent, __gong__toRawStringLiteral(attribute.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetNamespace = %s", attributeIdent, __gong__toRawStringLiteral(attribute.TargetNamespace)))
			values.WriteString(fmt.Sprintf("\n\t%s.SimpleType = %s", attributeIdent, __gong__toRawStringLiteral(attribute.SimpleType)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDXSD = %s", attributeIdent, __gong__toRawStringLiteral(attribute.IDXSD)))
			if attribute.Annotation != nil {
				targetIdent := "__stage_0" + attribute.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", attributeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		attributegroupOrdered := []*AttributeGroup{}
		for attributegroup := range stageSet.Stage.AttributeGroups {
			attributegroupOrdered = append(attributegroupOrdered, attributegroup)
		}
		sort.Slice(attributegroupOrdered, func(i, j int) bool {
			return stageSet.Stage.AttributeGroup_stagedOrder[attributegroupOrdered[i]] < stageSet.Stage.AttributeGroup_stagedOrder[attributegroupOrdered[j]]
		})
		for _, attributegroup := range attributegroupOrdered {
			attributegroupIdent := "__stage_0" + attributegroup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.AttributeGroup{Name: %s}).Stage(stageSet.Stage)", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", attributegroupIdent, attributegroup.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", attributegroupIdent, attributegroup.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", attributegroupIdent, attributegroup.Depth))
			if attributegroup.Annotation != nil {
				targetIdent := "__stage_0" + attributegroup.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", attributegroupIdent, targetIdent))
			}
			for _, elem := range attributegroup.AttributeGroups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", attributegroupIdent, attributegroupIdent, targetIdent))
			}
			for _, elem := range attributegroup.Attributes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", attributegroupIdent, attributegroupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		choiceOrdered := []*Choice{}
		for choice := range stageSet.Stage.Choices {
			choiceOrdered = append(choiceOrdered, choice)
		}
		sort.Slice(choiceOrdered, func(i, j int) bool {
			return stageSet.Stage.Choice_stagedOrder[choiceOrdered[i]] < stageSet.Stage.Choice_stagedOrder[choiceOrdered[j]]
		})
		for _, choice := range choiceOrdered {
			choiceIdent := "__stage_0" + choice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Choice{Name: %s}).Stage(stageSet.Stage)", choiceIdent, __gong__toRawStringLiteral(choice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", choiceIdent, __gong__toRawStringLiteral(choice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", choiceIdent, __gong__toRawStringLiteral(choice.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", choiceIdent, choice.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", choiceIdent, choice.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", choiceIdent, __gong__toRawStringLiteral(choice.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", choiceIdent, __gong__toRawStringLiteral(choice.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", choiceIdent, choice.IsDuplicatedInXSD))
			if choice.Annotation != nil {
				targetIdent := "__stage_0" + choice.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", choiceIdent, targetIdent))
			}
			for _, elem := range choice.Sequences {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Alls {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Choices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", choiceIdent, choiceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		complexcontentOrdered := []*ComplexContent{}
		for complexcontent := range stageSet.Stage.ComplexContents {
			complexcontentOrdered = append(complexcontentOrdered, complexcontent)
		}
		sort.Slice(complexcontentOrdered, func(i, j int) bool {
			return stageSet.Stage.ComplexContent_stagedOrder[complexcontentOrdered[i]] < stageSet.Stage.ComplexContent_stagedOrder[complexcontentOrdered[j]]
		})
		for _, complexcontent := range complexcontentOrdered {
			complexcontentIdent := "__stage_0" + complexcontent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ComplexContent{Name: %s}).Stage(stageSet.Stage)", complexcontentIdent, __gong__toRawStringLiteral(complexcontent.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complexcontentIdent, __gong__toRawStringLiteral(complexcontent.Name)))
		}
	}
	if stageSet.Stage != nil {
		complextypeOrdered := []*ComplexType{}
		for complextype := range stageSet.Stage.ComplexTypes {
			complextypeOrdered = append(complextypeOrdered, complextype)
		}
		sort.Slice(complextypeOrdered, func(i, j int) bool {
			return stageSet.Stage.ComplexType_stagedOrder[complextypeOrdered[i]] < stageSet.Stage.ComplexType_stagedOrder[complextypeOrdered[j]]
		})
		for _, complextype := range complextypeOrdered {
			complextypeIdent := "__stage_0" + complextype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ComplexType{Name: %s}).Stage(stageSet.Stage)", complextypeIdent, __gong__toRawStringLiteral(complextype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", complextypeIdent, complextype.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAnonymous = %t", complextypeIdent, complextype.IsAnonymous))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", complextypeIdent, complextype.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", complextypeIdent, complextype.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", complextypeIdent, __gong__toRawStringLiteral(complextype.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", complextypeIdent, complextype.IsDuplicatedInXSD))
			if complextype.OuterElement != nil {
				targetIdent := "__stage_0" + complextype.OuterElement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OuterElement = %s", complextypeIdent, targetIdent))
			}
			if complextype.Annotation != nil {
				targetIdent := "__stage_0" + complextype.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Sequences {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Alls {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Choices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			if complextype.Extension != nil {
				targetIdent := "__stage_0" + complextype.Extension.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = %s", complextypeIdent, targetIdent))
			}
			if complextype.SimpleContent != nil {
				targetIdent := "__stage_0" + complextype.SimpleContent.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleContent = %s", complextypeIdent, targetIdent))
			}
			if complextype.ComplexContent != nil {
				targetIdent := "__stage_0" + complextype.ComplexContent.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexContent = %s", complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Attributes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.AttributeGroups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		documentationOrdered := []*Documentation{}
		for documentation := range stageSet.Stage.Documentations {
			documentationOrdered = append(documentationOrdered, documentation)
		}
		sort.Slice(documentationOrdered, func(i, j int) bool {
			return stageSet.Stage.Documentation_stagedOrder[documentationOrdered[i]] < stageSet.Stage.Documentation_stagedOrder[documentationOrdered[j]]
		})
		for _, documentation := range documentationOrdered {
			documentationIdent := "__stage_0" + documentation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Documentation{Name: %s}).Stage(stageSet.Stage)", documentationIdent, __gong__toRawStringLiteral(documentation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Source)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Lang)))
		}
	}
	if stageSet.Stage != nil {
		elementOrdered := []*Element{}
		for element := range stageSet.Stage.Elements {
			elementOrdered = append(elementOrdered, element)
		}
		sort.Slice(elementOrdered, func(i, j int) bool {
			return stageSet.Stage.Element_stagedOrder[elementOrdered[i]] < stageSet.Stage.Element_stagedOrder[elementOrdered[j]]
		})
		for _, element := range elementOrdered {
			elementIdent := "__stage_0" + element.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Element{Name: %s}).Stage(stageSet.Stage)", elementIdent, __gong__toRawStringLiteral(element.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elementIdent, __gong__toRawStringLiteral(element.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", elementIdent, element.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", elementIdent, element.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", elementIdent, element.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", elementIdent, __gong__toRawStringLiteral(element.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", elementIdent, __gong__toRawStringLiteral(element.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", elementIdent, __gong__toRawStringLiteral(element.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", elementIdent, __gong__toRawStringLiteral(element.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", elementIdent, __gong__toRawStringLiteral(element.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default = %s", elementIdent, __gong__toRawStringLiteral(element.Default)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fixed = %s", elementIdent, __gong__toRawStringLiteral(element.Fixed)))
			values.WriteString(fmt.Sprintf("\n\t%s.Nillable = %s", elementIdent, __gong__toRawStringLiteral(element.Nillable)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", elementIdent, __gong__toRawStringLiteral(element.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Abstract = %s", elementIdent, __gong__toRawStringLiteral(element.Abstract)))
			values.WriteString(fmt.Sprintf("\n\t%s.Form = %s", elementIdent, __gong__toRawStringLiteral(element.Form)))
			values.WriteString(fmt.Sprintf("\n\t%s.Block = %s", elementIdent, __gong__toRawStringLiteral(element.Block)))
			values.WriteString(fmt.Sprintf("\n\t%s.Final = %s", elementIdent, __gong__toRawStringLiteral(element.Final)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", elementIdent, element.IsDuplicatedInXSD))
			if element.Annotation != nil {
				targetIdent := "__stage_0" + element.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", elementIdent, targetIdent))
			}
			if element.SimpleType != nil {
				targetIdent := "__stage_0" + element.SimpleType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleType = %s", elementIdent, targetIdent))
			}
			if element.ComplexType != nil {
				targetIdent := "__stage_0" + element.ComplexType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexType = %s", elementIdent, targetIdent))
			}
			for _, elem := range element.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", elementIdent, elementIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		enumerationOrdered := []*Enumeration{}
		for enumeration := range stageSet.Stage.Enumerations {
			enumerationOrdered = append(enumerationOrdered, enumeration)
		}
		sort.Slice(enumerationOrdered, func(i, j int) bool {
			return stageSet.Stage.Enumeration_stagedOrder[enumerationOrdered[i]] < stageSet.Stage.Enumeration_stagedOrder[enumerationOrdered[j]]
		})
		for _, enumeration := range enumerationOrdered {
			enumerationIdent := "__stage_0" + enumeration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Enumeration{Name: %s}).Stage(stageSet.Stage)", enumerationIdent, __gong__toRawStringLiteral(enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enumerationIdent, __gong__toRawStringLiteral(enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", enumerationIdent, __gong__toRawStringLiteral(enumeration.Value)))
			if enumeration.Annotation != nil {
				targetIdent := "__stage_0" + enumeration.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", enumerationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		extensionOrdered := []*Extension{}
		for extension := range stageSet.Stage.Extensions {
			extensionOrdered = append(extensionOrdered, extension)
		}
		sort.Slice(extensionOrdered, func(i, j int) bool {
			return stageSet.Stage.Extension_stagedOrder[extensionOrdered[i]] < stageSet.Stage.Extension_stagedOrder[extensionOrdered[j]]
		})
		for _, extension := range extensionOrdered {
			extensionIdent := "__stage_0" + extension.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Extension{Name: %s}).Stage(stageSet.Stage)", extensionIdent, __gong__toRawStringLiteral(extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extensionIdent, __gong__toRawStringLiteral(extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", extensionIdent, __gong__toRawStringLiteral(extension.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", extensionIdent, extension.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", extensionIdent, extension.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", extensionIdent, __gong__toRawStringLiteral(extension.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", extensionIdent, __gong__toRawStringLiteral(extension.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base = %s", extensionIdent, __gong__toRawStringLiteral(extension.Base)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", extensionIdent, __gong__toRawStringLiteral(extension.Ref)))
			for _, elem := range extension.Sequences {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Alls {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Choices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Attributes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.AttributeGroups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", extensionIdent, extensionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		groupOrdered := []*Group{}
		for group := range stageSet.Stage.Groups {
			groupOrdered = append(groupOrdered, group)
		}
		sort.Slice(groupOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_stagedOrder[groupOrdered[i]] < stageSet.Stage.Group_stagedOrder[groupOrdered[j]]
		})
		for _, group := range groupOrdered {
			groupIdent := "__stage_0" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupIdent, __gong__toRawStringLiteral(group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", groupIdent, __gong__toRawStringLiteral(group.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", groupIdent, __gong__toRawStringLiteral(group.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAnonymous = %t", groupIdent, group.IsAnonymous))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", groupIdent, group.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", groupIdent, __gong__toRawStringLiteral(group.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", groupIdent, __gong__toRawStringLiteral(group.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", groupIdent, group.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", groupIdent, group.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", groupIdent, __gong__toRawStringLiteral(group.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", groupIdent, __gong__toRawStringLiteral(group.MaxOccurs)))
			if group.Annotation != nil {
				targetIdent := "__stage_0" + group.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", groupIdent, targetIdent))
			}
			if group.OuterElement != nil {
				targetIdent := "__stage_0" + group.OuterElement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OuterElement = %s", groupIdent, targetIdent))
			}
			for _, elem := range group.Sequences {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Alls {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Choices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		lengthOrdered := []*Length{}
		for length := range stageSet.Stage.Lengths {
			lengthOrdered = append(lengthOrdered, length)
		}
		sort.Slice(lengthOrdered, func(i, j int) bool {
			return stageSet.Stage.Length_stagedOrder[lengthOrdered[i]] < stageSet.Stage.Length_stagedOrder[lengthOrdered[j]]
		})
		for _, length := range lengthOrdered {
			lengthIdent := "__stage_0" + length.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Length{Name: %s}).Stage(stageSet.Stage)", lengthIdent, __gong__toRawStringLiteral(length.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lengthIdent, __gong__toRawStringLiteral(length.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", lengthIdent, __gong__toRawStringLiteral(length.Value)))
			if length.Annotation != nil {
				targetIdent := "__stage_0" + length.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", lengthIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		maxinclusiveOrdered := []*MaxInclusive{}
		for maxinclusive := range stageSet.Stage.MaxInclusives {
			maxinclusiveOrdered = append(maxinclusiveOrdered, maxinclusive)
		}
		sort.Slice(maxinclusiveOrdered, func(i, j int) bool {
			return stageSet.Stage.MaxInclusive_stagedOrder[maxinclusiveOrdered[i]] < stageSet.Stage.MaxInclusive_stagedOrder[maxinclusiveOrdered[j]]
		})
		for _, maxinclusive := range maxinclusiveOrdered {
			maxinclusiveIdent := "__stage_0" + maxinclusive.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MaxInclusive{Name: %s}).Stage(stageSet.Stage)", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Value)))
			if maxinclusive.Annotation != nil {
				targetIdent := "__stage_0" + maxinclusive.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", maxinclusiveIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		maxlengthOrdered := []*MaxLength{}
		for maxlength := range stageSet.Stage.MaxLengths {
			maxlengthOrdered = append(maxlengthOrdered, maxlength)
		}
		sort.Slice(maxlengthOrdered, func(i, j int) bool {
			return stageSet.Stage.MaxLength_stagedOrder[maxlengthOrdered[i]] < stageSet.Stage.MaxLength_stagedOrder[maxlengthOrdered[j]]
		})
		for _, maxlength := range maxlengthOrdered {
			maxlengthIdent := "__stage_0" + maxlength.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MaxLength{Name: %s}).Stage(stageSet.Stage)", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Value)))
			if maxlength.Annotation != nil {
				targetIdent := "__stage_0" + maxlength.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", maxlengthIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		mininclusiveOrdered := []*MinInclusive{}
		for mininclusive := range stageSet.Stage.MinInclusives {
			mininclusiveOrdered = append(mininclusiveOrdered, mininclusive)
		}
		sort.Slice(mininclusiveOrdered, func(i, j int) bool {
			return stageSet.Stage.MinInclusive_stagedOrder[mininclusiveOrdered[i]] < stageSet.Stage.MinInclusive_stagedOrder[mininclusiveOrdered[j]]
		})
		for _, mininclusive := range mininclusiveOrdered {
			mininclusiveIdent := "__stage_0" + mininclusive.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MinInclusive{Name: %s}).Stage(stageSet.Stage)", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Value)))
			if mininclusive.Annotation != nil {
				targetIdent := "__stage_0" + mininclusive.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", mininclusiveIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		minlengthOrdered := []*MinLength{}
		for minlength := range stageSet.Stage.MinLengths {
			minlengthOrdered = append(minlengthOrdered, minlength)
		}
		sort.Slice(minlengthOrdered, func(i, j int) bool {
			return stageSet.Stage.MinLength_stagedOrder[minlengthOrdered[i]] < stageSet.Stage.MinLength_stagedOrder[minlengthOrdered[j]]
		})
		for _, minlength := range minlengthOrdered {
			minlengthIdent := "__stage_0" + minlength.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MinLength{Name: %s}).Stage(stageSet.Stage)", minlengthIdent, __gong__toRawStringLiteral(minlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", minlengthIdent, __gong__toRawStringLiteral(minlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", minlengthIdent, __gong__toRawStringLiteral(minlength.Value)))
			if minlength.Annotation != nil {
				targetIdent := "__stage_0" + minlength.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", minlengthIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		patternOrdered := []*Pattern{}
		for pattern := range stageSet.Stage.Patterns {
			patternOrdered = append(patternOrdered, pattern)
		}
		sort.Slice(patternOrdered, func(i, j int) bool {
			return stageSet.Stage.Pattern_stagedOrder[patternOrdered[i]] < stageSet.Stage.Pattern_stagedOrder[patternOrdered[j]]
		})
		for _, pattern := range patternOrdered {
			patternIdent := "__stage_0" + pattern.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Pattern{Name: %s}).Stage(stageSet.Stage)", patternIdent, __gong__toRawStringLiteral(pattern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", patternIdent, __gong__toRawStringLiteral(pattern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", patternIdent, __gong__toRawStringLiteral(pattern.Value)))
			if pattern.Annotation != nil {
				targetIdent := "__stage_0" + pattern.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", patternIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		restrictionOrdered := []*Restriction{}
		for restriction := range stageSet.Stage.Restrictions {
			restrictionOrdered = append(restrictionOrdered, restriction)
		}
		sort.Slice(restrictionOrdered, func(i, j int) bool {
			return stageSet.Stage.Restriction_stagedOrder[restrictionOrdered[i]] < stageSet.Stage.Restriction_stagedOrder[restrictionOrdered[j]]
		})
		for _, restriction := range restrictionOrdered {
			restrictionIdent := "__stage_0" + restriction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Restriction{Name: %s}).Stage(stageSet.Stage)", restrictionIdent, __gong__toRawStringLiteral(restriction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restrictionIdent, __gong__toRawStringLiteral(restriction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base = %s", restrictionIdent, __gong__toRawStringLiteral(restriction.Base)))
			if restriction.Annotation != nil {
				targetIdent := "__stage_0" + restriction.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", restrictionIdent, targetIdent))
			}
			for _, elem := range restriction.Enumerations {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Enumerations = append(%s.Enumerations, %s)", restrictionIdent, restrictionIdent, targetIdent))
			}
			if restriction.MinInclusive != nil {
				targetIdent := "__stage_0" + restriction.MinInclusive.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinInclusive = %s", restrictionIdent, targetIdent))
			}
			if restriction.MaxInclusive != nil {
				targetIdent := "__stage_0" + restriction.MaxInclusive.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaxInclusive = %s", restrictionIdent, targetIdent))
			}
			if restriction.Pattern != nil {
				targetIdent := "__stage_0" + restriction.Pattern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pattern = %s", restrictionIdent, targetIdent))
			}
			if restriction.WhiteSpace != nil {
				targetIdent := "__stage_0" + restriction.WhiteSpace.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = %s", restrictionIdent, targetIdent))
			}
			if restriction.MinLength != nil {
				targetIdent := "__stage_0" + restriction.MinLength.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinLength = %s", restrictionIdent, targetIdent))
			}
			if restriction.MaxLength != nil {
				targetIdent := "__stage_0" + restriction.MaxLength.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaxLength = %s", restrictionIdent, targetIdent))
			}
			if restriction.Length != nil {
				targetIdent := "__stage_0" + restriction.Length.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Length = %s", restrictionIdent, targetIdent))
			}
			if restriction.TotalDigit != nil {
				targetIdent := "__stage_0" + restriction.TotalDigit.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TotalDigit = %s", restrictionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		schemaOrdered := []*Schema{}
		for schema := range stageSet.Stage.Schemas {
			schemaOrdered = append(schemaOrdered, schema)
		}
		sort.Slice(schemaOrdered, func(i, j int) bool {
			return stageSet.Stage.Schema_stagedOrder[schemaOrdered[i]] < stageSet.Stage.Schema_stagedOrder[schemaOrdered[j]]
		})
		for _, schema := range schemaOrdered {
			schemaIdent := "__stage_0" + schema.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Schema{Name: %s}).Stage(stageSet.Stage)", schemaIdent, __gong__toRawStringLiteral(schema.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", schemaIdent, __gong__toRawStringLiteral(schema.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Xs = %s", schemaIdent, __gong__toRawStringLiteral(schema.Xs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", schemaIdent, schema.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", schemaIdent, schema.Depth))
			if schema.Annotation != nil {
				targetIdent := "__stage_0" + schema.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", schemaIdent, targetIdent))
			}
			for _, elem := range schema.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.SimpleTypes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleTypes = append(%s.SimpleTypes, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.ComplexTypes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexTypes = append(%s.ComplexTypes, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.AttributeGroups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", schemaIdent, schemaIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		sequenceOrdered := []*Sequence{}
		for sequence := range stageSet.Stage.Sequences {
			sequenceOrdered = append(sequenceOrdered, sequence)
		}
		sort.Slice(sequenceOrdered, func(i, j int) bool {
			return stageSet.Stage.Sequence_stagedOrder[sequenceOrdered[i]] < stageSet.Stage.Sequence_stagedOrder[sequenceOrdered[j]]
		})
		for _, sequence := range sequenceOrdered {
			sequenceIdent := "__stage_0" + sequence.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Sequence{Name: %s}).Stage(stageSet.Stage)", sequenceIdent, __gong__toRawStringLiteral(sequence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", sequenceIdent, sequence.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", sequenceIdent, sequence.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.MaxOccurs)))
			if sequence.Annotation != nil {
				targetIdent := "__stage_0" + sequence.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Sequences {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Alls {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Choices {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Groups {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Elements {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		simplecontentOrdered := []*SimpleContent{}
		for simplecontent := range stageSet.Stage.SimpleContents {
			simplecontentOrdered = append(simplecontentOrdered, simplecontent)
		}
		sort.Slice(simplecontentOrdered, func(i, j int) bool {
			return stageSet.Stage.SimpleContent_stagedOrder[simplecontentOrdered[i]] < stageSet.Stage.SimpleContent_stagedOrder[simplecontentOrdered[j]]
		})
		for _, simplecontent := range simplecontentOrdered {
			simplecontentIdent := "__stage_0" + simplecontent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SimpleContent{Name: %s}).Stage(stageSet.Stage)", simplecontentIdent, __gong__toRawStringLiteral(simplecontent.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simplecontentIdent, __gong__toRawStringLiteral(simplecontent.Name)))
			if simplecontent.Extension != nil {
				targetIdent := "__stage_0" + simplecontent.Extension.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = %s", simplecontentIdent, targetIdent))
			}
			if simplecontent.Restriction != nil {
				targetIdent := "__stage_0" + simplecontent.Restriction.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Restriction = %s", simplecontentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		simpletypeOrdered := []*SimpleType{}
		for simpletype := range stageSet.Stage.SimpleTypes {
			simpletypeOrdered = append(simpletypeOrdered, simpletype)
		}
		sort.Slice(simpletypeOrdered, func(i, j int) bool {
			return stageSet.Stage.SimpleType_stagedOrder[simpletypeOrdered[i]] < stageSet.Stage.SimpleType_stagedOrder[simpletypeOrdered[j]]
		})
		for _, simpletype := range simpletypeOrdered {
			simpletypeIdent := "__stage_0" + simpletype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.SimpleType{Name: %s}).Stage(stageSet.Stage)", simpletypeIdent, __gong__toRawStringLiteral(simpletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simpletypeIdent, __gong__toRawStringLiteral(simpletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", simpletypeIdent, __gong__toRawStringLiteral(simpletype.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", simpletypeIdent, simpletype.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", simpletypeIdent, simpletype.Depth))
			if simpletype.Annotation != nil {
				targetIdent := "__stage_0" + simpletype.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", simpletypeIdent, targetIdent))
			}
			if simpletype.Restriction != nil {
				targetIdent := "__stage_0" + simpletype.Restriction.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Restriction = %s", simpletypeIdent, targetIdent))
			}
			if simpletype.Union != nil {
				targetIdent := "__stage_0" + simpletype.Union.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Union = %s", simpletypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		totaldigitOrdered := []*TotalDigit{}
		for totaldigit := range stageSet.Stage.TotalDigits {
			totaldigitOrdered = append(totaldigitOrdered, totaldigit)
		}
		sort.Slice(totaldigitOrdered, func(i, j int) bool {
			return stageSet.Stage.TotalDigit_stagedOrder[totaldigitOrdered[i]] < stageSet.Stage.TotalDigit_stagedOrder[totaldigitOrdered[j]]
		})
		for _, totaldigit := range totaldigitOrdered {
			totaldigitIdent := "__stage_0" + totaldigit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TotalDigit{Name: %s}).Stage(stageSet.Stage)", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Value)))
			if totaldigit.Annotation != nil {
				targetIdent := "__stage_0" + totaldigit.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", totaldigitIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		unionOrdered := []*Union{}
		for union := range stageSet.Stage.Unions {
			unionOrdered = append(unionOrdered, union)
		}
		sort.Slice(unionOrdered, func(i, j int) bool {
			return stageSet.Stage.Union_stagedOrder[unionOrdered[i]] < stageSet.Stage.Union_stagedOrder[unionOrdered[j]]
		})
		for _, union := range unionOrdered {
			unionIdent := "__stage_0" + union.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Union{Name: %s}).Stage(stageSet.Stage)", unionIdent, __gong__toRawStringLiteral(union.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unionIdent, __gong__toRawStringLiteral(union.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MemberTypes = %s", unionIdent, __gong__toRawStringLiteral(union.MemberTypes)))
			if union.Annotation != nil {
				targetIdent := "__stage_0" + union.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", unionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		whitespaceOrdered := []*WhiteSpace{}
		for whitespace := range stageSet.Stage.WhiteSpaces {
			whitespaceOrdered = append(whitespaceOrdered, whitespace)
		}
		sort.Slice(whitespaceOrdered, func(i, j int) bool {
			return stageSet.Stage.WhiteSpace_stagedOrder[whitespaceOrdered[i]] < stageSet.Stage.WhiteSpace_stagedOrder[whitespaceOrdered[j]]
		})
		for _, whitespace := range whitespaceOrdered {
			whitespaceIdent := "__stage_0" + whitespace.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.WhiteSpace{Name: %s}).Stage(stageSet.Stage)", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Value)))
			if whitespace.Annotation != nil {
				targetIdent := "__stage_0" + whitespace.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", whitespaceIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/app/xsd/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "All":
					if !preserveOrder {
						inst := (&All{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(All)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Annotation":
					if !preserveOrder {
						inst := (&Annotation{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Annotation)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Attribute":
					if !preserveOrder {
						inst := (&Attribute{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Attribute)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "AttributeGroup":
					if !preserveOrder {
						inst := (&AttributeGroup{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AttributeGroup)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Choice":
					if !preserveOrder {
						inst := (&Choice{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Choice)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ComplexContent":
					if !preserveOrder {
						inst := (&ComplexContent{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ComplexContent)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ComplexType":
					if !preserveOrder {
						inst := (&ComplexType{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ComplexType)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Documentation":
					if !preserveOrder {
						inst := (&Documentation{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Documentation)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Element":
					if !preserveOrder {
						inst := (&Element{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Element)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Enumeration":
					if !preserveOrder {
						inst := (&Enumeration{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Enumeration)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Extension":
					if !preserveOrder {
						inst := (&Extension{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Extension)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group":
					if !preserveOrder {
						inst := (&Group{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Length":
					if !preserveOrder {
						inst := (&Length{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Length)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MaxInclusive":
					if !preserveOrder {
						inst := (&MaxInclusive{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MaxInclusive)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MaxLength":
					if !preserveOrder {
						inst := (&MaxLength{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MaxLength)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MinInclusive":
					if !preserveOrder {
						inst := (&MinInclusive{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MinInclusive)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MinLength":
					if !preserveOrder {
						inst := (&MinLength{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MinLength)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Pattern":
					if !preserveOrder {
						inst := (&Pattern{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Pattern)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Restriction":
					if !preserveOrder {
						inst := (&Restriction{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Restriction)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Schema":
					if !preserveOrder {
						inst := (&Schema{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Schema)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Sequence":
					if !preserveOrder {
						inst := (&Sequence{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Sequence)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SimpleContent":
					if !preserveOrder {
						inst := (&SimpleContent{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SimpleContent)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SimpleType":
					if !preserveOrder {
						inst := (&SimpleType{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SimpleType)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TotalDigit":
					if !preserveOrder {
						inst := (&TotalDigit{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TotalDigit)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Union":
					if !preserveOrder {
						inst := (&Union{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Union)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "WhiteSpace":
					if !preserveOrder {
						inst := (&WhiteSpace{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(WhiteSpace)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *All:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sequence); ok {
										inst.Sequences = append(inst.Sequences, typedTarget)
									}
								}
							}
						}
					case "Alls":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*All); ok {
										inst.Alls = append(inst.Alls, typedTarget)
									}
								}
							}
						}
					case "Choices":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Choice); ok {
										inst.Choices = append(inst.Choices, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					}
				case *Annotation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Documentations":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Documentation); ok {
										inst.Documentations = append(inst.Documentations, typedTarget)
									}
								}
							}
						}
					}
				case *Attribute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "Default":
						inst.Default = GongExtractString(rhs)
					case "Use":
						inst.Use = GongExtractString(rhs)
					case "Form":
						inst.Form = GongExtractString(rhs)
					case "Fixed":
						inst.Fixed = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "TargetNamespace":
						inst.TargetNamespace = GongExtractString(rhs)
					case "SimpleType":
						inst.SimpleType = GongExtractString(rhs)
					case "IDXSD":
						inst.IDXSD = GongExtractString(rhs)
					}
				case *AttributeGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "AttributeGroups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AttributeGroup); ok {
										inst.AttributeGroups = append(inst.AttributeGroups, typedTarget)
									}
								}
							}
						}
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Attributes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Attribute); ok {
										inst.Attributes = append(inst.Attributes, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					}
				case *Choice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sequence); ok {
										inst.Sequences = append(inst.Sequences, typedTarget)
									}
								}
							}
						}
					case "Alls":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*All); ok {
										inst.Alls = append(inst.Alls, typedTarget)
									}
								}
							}
						}
					case "Choices":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Choice); ok {
										inst.Choices = append(inst.Choices, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *ComplexContent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ComplexType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "IsAnonymous":
						inst.IsAnonymous = GongExtractBool(rhs)
					case "OuterElement":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
									inst.OuterElement = typedTarget
								}
							}
						}
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sequence); ok {
										inst.Sequences = append(inst.Sequences, typedTarget)
									}
								}
							}
						}
					case "Alls":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*All); ok {
										inst.Alls = append(inst.Alls, typedTarget)
									}
								}
							}
						}
					case "Choices":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Choice); ok {
										inst.Choices = append(inst.Choices, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Extension":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Extension); ok {
									inst.Extension = typedTarget
								}
							}
						}
					case "SimpleContent":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SimpleContent); ok {
									inst.SimpleContent = typedTarget
								}
							}
						}
					case "ComplexContent":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ComplexContent); ok {
									inst.ComplexContent = typedTarget
								}
							}
						}
					case "Attributes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Attribute); ok {
										inst.Attributes = append(inst.Attributes, typedTarget)
									}
								}
							}
						}
					case "AttributeGroups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AttributeGroup); ok {
										inst.AttributeGroups = append(inst.AttributeGroups, typedTarget)
									}
								}
							}
						}
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *Documentation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					}
				case *Element:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Default":
						inst.Default = GongExtractString(rhs)
					case "Fixed":
						inst.Fixed = GongExtractString(rhs)
					case "Nillable":
						inst.Nillable = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Abstract":
						inst.Abstract = GongExtractString(rhs)
					case "Form":
						inst.Form = GongExtractString(rhs)
					case "Block":
						inst.Block = GongExtractString(rhs)
					case "Final":
						inst.Final = GongExtractString(rhs)
					case "SimpleType":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SimpleType); ok {
									inst.SimpleType = typedTarget
								}
							}
						}
					case "ComplexType":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ComplexType); ok {
									inst.ComplexType = typedTarget
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *Enumeration:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sequence); ok {
										inst.Sequences = append(inst.Sequences, typedTarget)
									}
								}
							}
						}
					case "Alls":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*All); ok {
										inst.Alls = append(inst.Alls, typedTarget)
									}
								}
							}
						}
					case "Choices":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Choice); ok {
										inst.Choices = append(inst.Choices, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Base":
						inst.Base = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Attributes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Attribute); ok {
										inst.Attributes = append(inst.Attributes, typedTarget)
									}
								}
							}
						}
					case "AttributeGroups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AttributeGroup); ok {
										inst.AttributeGroups = append(inst.AttributeGroups, typedTarget)
									}
								}
							}
						}
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "IsAnonymous":
						inst.IsAnonymous = GongExtractBool(rhs)
					case "OuterElement":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
									inst.OuterElement = typedTarget
								}
							}
						}
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sequence); ok {
										inst.Sequences = append(inst.Sequences, typedTarget)
									}
								}
							}
						}
					case "Alls":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*All); ok {
										inst.Alls = append(inst.Alls, typedTarget)
									}
								}
							}
						}
					case "Choices":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Choice); ok {
										inst.Choices = append(inst.Choices, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					}
				case *Length:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MaxInclusive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MaxLength:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MinInclusive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MinLength:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Pattern:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Restriction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Base":
						inst.Base = GongExtractString(rhs)
					case "Enumerations":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Enumeration); ok {
										inst.Enumerations = append(inst.Enumerations, typedTarget)
									}
								}
							}
						}
					case "MinInclusive":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MinInclusive); ok {
									inst.MinInclusive = typedTarget
								}
							}
						}
					case "MaxInclusive":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MaxInclusive); ok {
									inst.MaxInclusive = typedTarget
								}
							}
						}
					case "Pattern":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Pattern); ok {
									inst.Pattern = typedTarget
								}
							}
						}
					case "WhiteSpace":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*WhiteSpace); ok {
									inst.WhiteSpace = typedTarget
								}
							}
						}
					case "MinLength":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MinLength); ok {
									inst.MinLength = typedTarget
								}
							}
						}
					case "MaxLength":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MaxLength); ok {
									inst.MaxLength = typedTarget
								}
							}
						}
					case "Length":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Length); ok {
									inst.Length = typedTarget
								}
							}
						}
					case "TotalDigit":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TotalDigit); ok {
									inst.TotalDigit = typedTarget
								}
							}
						}
					}
				case *Schema:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Xs":
						inst.Xs = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "SimpleTypes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SimpleType); ok {
										inst.SimpleTypes = append(inst.SimpleTypes, typedTarget)
									}
								}
							}
						}
					case "ComplexTypes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ComplexType); ok {
										inst.ComplexTypes = append(inst.ComplexTypes, typedTarget)
									}
								}
							}
						}
					case "AttributeGroups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AttributeGroup); ok {
										inst.AttributeGroups = append(inst.AttributeGroups, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					}
				case *Sequence:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sequence); ok {
										inst.Sequences = append(inst.Sequences, typedTarget)
									}
								}
							}
						}
					case "Alls":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*All); ok {
										inst.Alls = append(inst.Alls, typedTarget)
									}
								}
							}
						}
					case "Choices":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Choice); ok {
										inst.Choices = append(inst.Choices, typedTarget)
									}
								}
							}
						}
					case "Groups":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group); ok {
										inst.Groups = append(inst.Groups, typedTarget)
									}
								}
							}
						}
					case "Elements":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Element); ok {
										inst.Elements = append(inst.Elements, typedTarget)
									}
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					}
				case *SimpleContent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Extension":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Extension); ok {
									inst.Extension = typedTarget
								}
							}
						}
					case "Restriction":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Restriction); ok {
									inst.Restriction = typedTarget
								}
							}
						}
					}
				case *SimpleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Restriction":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Restriction); ok {
									inst.Restriction = typedTarget
								}
							}
						}
					case "Union":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Union); ok {
									inst.Union = typedTarget
								}
							}
						}
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					}
				case *TotalDigit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Union:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "MemberTypes":
						inst.MemberTypes = GongExtractString(rhs)
					}
				case *WhiteSpace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Annotation); ok {
									inst.Annotation = typedTarget
								}
							}
						}
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}
