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
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, all := range __gong__sortStageSetInstances(stageSet.Stage.Alls, stageSet.Stage.All_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			allIdent := "__models" + all.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.All{Name: %s}).Stage(stageSet.Stage)", allIdent, __gong__toRawStringLiteral(all.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", allIdent, __gong__toRawStringLiteral(all.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", allIdent, __gong__toRawStringLiteral(all.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", allIdent, all.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", allIdent, all.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", allIdent, __gong__toRawStringLiteral(all.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", allIdent, __gong__toRawStringLiteral(all.MaxOccurs)))
			if all.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + all.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", allIdent, targetIdent))
			}
			for _, elem := range all.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", allIdent, allIdent, targetIdent))
			}
			for _, elem := range all.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", allIdent, allIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, annotation := range __gong__sortStageSetInstances(stageSet.Stage.Annotations, stageSet.Stage.Annotation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			annotationIdent := "__models" + annotation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Annotation{Name: %s}).Stage(stageSet.Stage)", annotationIdent, __gong__toRawStringLiteral(annotation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", annotationIdent, __gong__toRawStringLiteral(annotation.Name)))
			for _, elem := range annotation.Documentations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Documentations = append(%s.Documentations, %s)", annotationIdent, annotationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attribute := range __gong__sortStageSetInstances(stageSet.Stage.Attributes, stageSet.Stage.Attribute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributeIdent := "__models" + attribute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Attribute{Name: %s}).Stage(stageSet.Stage)", attributeIdent, __gong__toRawStringLiteral(attribute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attribute.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", attributeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attributegroup := range __gong__sortStageSetInstances(stageSet.Stage.AttributeGroups, stageSet.Stage.AttributeGroup_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributegroupIdent := "__models" + attributegroup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AttributeGroup{Name: %s}).Stage(stageSet.Stage)", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasNameConflict = %t", attributegroupIdent, attributegroup.HasNameConflict))
			values.WriteString(fmt.Sprintf("\n\t%s.GoIdentifier = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.GoIdentifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", attributegroupIdent, __gong__toRawStringLiteral(attributegroup.Ref)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", attributegroupIdent, attributegroup.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", attributegroupIdent, attributegroup.Depth))
			if attributegroup.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attributegroup.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", attributegroupIdent, targetIdent))
			}
			for _, elem := range attributegroup.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", attributegroupIdent, attributegroupIdent, targetIdent))
			}
			for _, elem := range attributegroup.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", attributegroupIdent, attributegroupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, choice := range __gong__sortStageSetInstances(stageSet.Stage.Choices, stageSet.Stage.Choice_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			choiceIdent := "__models" + choice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Choice{Name: %s}).Stage(stageSet.Stage)", choiceIdent, __gong__toRawStringLiteral(choice.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", choiceIdent, __gong__toRawStringLiteral(choice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", choiceIdent, __gong__toRawStringLiteral(choice.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", choiceIdent, choice.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", choiceIdent, choice.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", choiceIdent, __gong__toRawStringLiteral(choice.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", choiceIdent, __gong__toRawStringLiteral(choice.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDuplicatedInXSD = %t", choiceIdent, choice.IsDuplicatedInXSD))
			if choice.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + choice.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", choiceIdent, targetIdent))
			}
			for _, elem := range choice.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", choiceIdent, choiceIdent, targetIdent))
			}
			for _, elem := range choice.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", choiceIdent, choiceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, complexcontent := range __gong__sortStageSetInstances(stageSet.Stage.ComplexContents, stageSet.Stage.ComplexContent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			complexcontentIdent := "__models" + complexcontent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ComplexContent{Name: %s}).Stage(stageSet.Stage)", complexcontentIdent, __gong__toRawStringLiteral(complexcontent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", complexcontentIdent, __gong__toRawStringLiteral(complexcontent.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, complextype := range __gong__sortStageSetInstances(stageSet.Stage.ComplexTypes, stageSet.Stage.ComplexType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			complextypeIdent := "__models" + complextype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ComplexType{Name: %s}).Stage(stageSet.Stage)", complextypeIdent, __gong__toRawStringLiteral(complextype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + complextype.OuterElement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OuterElement = %s", complextypeIdent, targetIdent))
			}
			if complextype.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + complextype.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			if complextype.Extension != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + complextype.Extension.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = %s", complextypeIdent, targetIdent))
			}
			if complextype.SimpleContent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + complextype.SimpleContent.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleContent = %s", complextypeIdent, targetIdent))
			}
			if complextype.ComplexContent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + complextype.ComplexContent.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexContent = %s", complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
			for _, elem := range complextype.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", complextypeIdent, complextypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, documentation := range __gong__sortStageSetInstances(stageSet.Stage.Documentations, stageSet.Stage.Documentation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentationIdent := "__models" + documentation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Documentation{Name: %s}).Stage(stageSet.Stage)", documentationIdent, __gong__toRawStringLiteral(documentation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Source)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", documentationIdent, __gong__toRawStringLiteral(documentation.Lang)))
		}
	}
	if stageSet.Stage != nil {
		for _, element := range __gong__sortStageSetInstances(stageSet.Stage.Elements, stageSet.Stage.Element_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elementIdent := "__models" + element.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Element{Name: %s}).Stage(stageSet.Stage)", elementIdent, __gong__toRawStringLiteral(element.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + element.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", elementIdent, targetIdent))
			}
			if element.SimpleType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + element.SimpleType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleType = %s", elementIdent, targetIdent))
			}
			if element.ComplexType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + element.ComplexType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexType = %s", elementIdent, targetIdent))
			}
			for _, elem := range element.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", elementIdent, elementIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, enumeration := range __gong__sortStageSetInstances(stageSet.Stage.Enumerations, stageSet.Stage.Enumeration_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			enumerationIdent := "__models" + enumeration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Enumeration{Name: %s}).Stage(stageSet.Stage)", enumerationIdent, __gong__toRawStringLiteral(enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enumerationIdent, __gong__toRawStringLiteral(enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", enumerationIdent, __gong__toRawStringLiteral(enumeration.Value)))
			if enumeration.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + enumeration.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", enumerationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, extension := range __gong__sortStageSetInstances(stageSet.Stage.Extensions, stageSet.Stage.Extension_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extensionIdent := "__models" + extension.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Extension{Name: %s}).Stage(stageSet.Stage)", extensionIdent, __gong__toRawStringLiteral(extension.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extensionIdent, __gong__toRawStringLiteral(extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", extensionIdent, __gong__toRawStringLiteral(extension.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", extensionIdent, extension.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", extensionIdent, extension.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", extensionIdent, __gong__toRawStringLiteral(extension.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", extensionIdent, __gong__toRawStringLiteral(extension.MaxOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base = %s", extensionIdent, __gong__toRawStringLiteral(extension.Base)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ref = %s", extensionIdent, __gong__toRawStringLiteral(extension.Ref)))
			for _, elem := range extension.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", extensionIdent, extensionIdent, targetIdent))
			}
			for _, elem := range extension.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", extensionIdent, extensionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, group := range __gong__sortStageSetInstances(stageSet.Stage.Groups, stageSet.Stage.Group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupIdent := "__models" + group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group{Name: %s}).Stage(stageSet.Stage)", groupIdent, __gong__toRawStringLiteral(group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
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
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + group.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", groupIdent, targetIdent))
			}
			if group.OuterElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + group.OuterElement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OuterElement = %s", groupIdent, targetIdent))
			}
			for _, elem := range group.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", groupIdent, groupIdent, targetIdent))
			}
			for _, elem := range group.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", groupIdent, groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, length := range __gong__sortStageSetInstances(stageSet.Stage.Lengths, stageSet.Stage.Length_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lengthIdent := "__models" + length.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Length{Name: %s}).Stage(stageSet.Stage)", lengthIdent, __gong__toRawStringLiteral(length.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lengthIdent, __gong__toRawStringLiteral(length.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", lengthIdent, __gong__toRawStringLiteral(length.Value)))
			if length.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + length.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", lengthIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, maxinclusive := range __gong__sortStageSetInstances(stageSet.Stage.MaxInclusives, stageSet.Stage.MaxInclusive_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			maxinclusiveIdent := "__models" + maxinclusive.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MaxInclusive{Name: %s}).Stage(stageSet.Stage)", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", maxinclusiveIdent, __gong__toRawStringLiteral(maxinclusive.Value)))
			if maxinclusive.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + maxinclusive.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", maxinclusiveIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, maxlength := range __gong__sortStageSetInstances(stageSet.Stage.MaxLengths, stageSet.Stage.MaxLength_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			maxlengthIdent := "__models" + maxlength.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MaxLength{Name: %s}).Stage(stageSet.Stage)", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", maxlengthIdent, __gong__toRawStringLiteral(maxlength.Value)))
			if maxlength.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + maxlength.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", maxlengthIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, mininclusive := range __gong__sortStageSetInstances(stageSet.Stage.MinInclusives, stageSet.Stage.MinInclusive_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mininclusiveIdent := "__models" + mininclusive.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MinInclusive{Name: %s}).Stage(stageSet.Stage)", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", mininclusiveIdent, __gong__toRawStringLiteral(mininclusive.Value)))
			if mininclusive.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + mininclusive.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", mininclusiveIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, minlength := range __gong__sortStageSetInstances(stageSet.Stage.MinLengths, stageSet.Stage.MinLength_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			minlengthIdent := "__models" + minlength.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MinLength{Name: %s}).Stage(stageSet.Stage)", minlengthIdent, __gong__toRawStringLiteral(minlength.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", minlengthIdent, __gong__toRawStringLiteral(minlength.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", minlengthIdent, __gong__toRawStringLiteral(minlength.Value)))
			if minlength.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + minlength.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", minlengthIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, pattern := range __gong__sortStageSetInstances(stageSet.Stage.Patterns, stageSet.Stage.Pattern_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			patternIdent := "__models" + pattern.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Pattern{Name: %s}).Stage(stageSet.Stage)", patternIdent, __gong__toRawStringLiteral(pattern.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", patternIdent, __gong__toRawStringLiteral(pattern.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", patternIdent, __gong__toRawStringLiteral(pattern.Value)))
			if pattern.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + pattern.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", patternIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, restriction := range __gong__sortStageSetInstances(stageSet.Stage.Restrictions, stageSet.Stage.Restriction_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			restrictionIdent := "__models" + restriction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Restriction{Name: %s}).Stage(stageSet.Stage)", restrictionIdent, __gong__toRawStringLiteral(restriction.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restrictionIdent, __gong__toRawStringLiteral(restriction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base = %s", restrictionIdent, __gong__toRawStringLiteral(restriction.Base)))
			if restriction.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", restrictionIdent, targetIdent))
			}
			for _, elem := range restriction.Enumerations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Enumerations = append(%s.Enumerations, %s)", restrictionIdent, restrictionIdent, targetIdent))
			}
			if restriction.MinInclusive != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.MinInclusive.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinInclusive = %s", restrictionIdent, targetIdent))
			}
			if restriction.MaxInclusive != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.MaxInclusive.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaxInclusive = %s", restrictionIdent, targetIdent))
			}
			if restriction.Pattern != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.Pattern.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pattern = %s", restrictionIdent, targetIdent))
			}
			if restriction.WhiteSpace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.WhiteSpace.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = %s", restrictionIdent, targetIdent))
			}
			if restriction.MinLength != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.MinLength.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MinLength = %s", restrictionIdent, targetIdent))
			}
			if restriction.MaxLength != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.MaxLength.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MaxLength = %s", restrictionIdent, targetIdent))
			}
			if restriction.Length != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.Length.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Length = %s", restrictionIdent, targetIdent))
			}
			if restriction.TotalDigit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + restriction.TotalDigit.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TotalDigit = %s", restrictionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, schema := range __gong__sortStageSetInstances(stageSet.Stage.Schemas, stageSet.Stage.Schema_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			schemaIdent := "__models" + schema.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Schema{Name: %s}).Stage(stageSet.Stage)", schemaIdent, __gong__toRawStringLiteral(schema.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", schemaIdent, __gong__toRawStringLiteral(schema.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Xs = %s", schemaIdent, __gong__toRawStringLiteral(schema.Xs)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", schemaIdent, schema.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", schemaIdent, schema.Depth))
			if schema.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + schema.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", schemaIdent, targetIdent))
			}
			for _, elem := range schema.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.SimpleTypes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SimpleTypes = append(%s.SimpleTypes, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.ComplexTypes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ComplexTypes = append(%s.ComplexTypes, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.AttributeGroups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeGroups = append(%s.AttributeGroups, %s)", schemaIdent, schemaIdent, targetIdent))
			}
			for _, elem := range schema.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", schemaIdent, schemaIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, sequence := range __gong__sortStageSetInstances(stageSet.Stage.Sequences, stageSet.Stage.Sequence_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sequenceIdent := "__models" + sequence.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Sequence{Name: %s}).Stage(stageSet.Stage)", sequenceIdent, __gong__toRawStringLiteral(sequence.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OuterElementName = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.OuterElementName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", sequenceIdent, sequence.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", sequenceIdent, sequence.Depth))
			values.WriteString(fmt.Sprintf("\n\t%s.MinOccurs = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.MinOccurs)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaxOccurs = %s", sequenceIdent, __gong__toRawStringLiteral(sequence.MaxOccurs)))
			if sequence.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + sequence.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Sequences {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sequences = append(%s.Sequences, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Alls {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Alls = append(%s.Alls, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Choices {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Choices = append(%s.Choices, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Groups {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Groups = append(%s.Groups, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
			for _, elem := range sequence.Elements {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elements = append(%s.Elements, %s)", sequenceIdent, sequenceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, simplecontent := range __gong__sortStageSetInstances(stageSet.Stage.SimpleContents, stageSet.Stage.SimpleContent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			simplecontentIdent := "__models" + simplecontent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SimpleContent{Name: %s}).Stage(stageSet.Stage)", simplecontentIdent, __gong__toRawStringLiteral(simplecontent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simplecontentIdent, __gong__toRawStringLiteral(simplecontent.Name)))
			if simplecontent.Extension != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + simplecontent.Extension.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = %s", simplecontentIdent, targetIdent))
			}
			if simplecontent.Restriction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + simplecontent.Restriction.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Restriction = %s", simplecontentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, simpletype := range __gong__sortStageSetInstances(stageSet.Stage.SimpleTypes, stageSet.Stage.SimpleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			simpletypeIdent := "__models" + simpletype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SimpleType{Name: %s}).Stage(stageSet.Stage)", simpletypeIdent, __gong__toRawStringLiteral(simpletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simpletypeIdent, __gong__toRawStringLiteral(simpletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", simpletypeIdent, __gong__toRawStringLiteral(simpletype.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Order = %d", simpletypeIdent, simpletype.Order))
			values.WriteString(fmt.Sprintf("\n\t%s.Depth = %d", simpletypeIdent, simpletype.Depth))
			if simpletype.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + simpletype.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", simpletypeIdent, targetIdent))
			}
			if simpletype.Restriction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + simpletype.Restriction.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Restriction = %s", simpletypeIdent, targetIdent))
			}
			if simpletype.Union != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + simpletype.Union.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Union = %s", simpletypeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, totaldigit := range __gong__sortStageSetInstances(stageSet.Stage.TotalDigits, stageSet.Stage.TotalDigit_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			totaldigitIdent := "__models" + totaldigit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TotalDigit{Name: %s}).Stage(stageSet.Stage)", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", totaldigitIdent, __gong__toRawStringLiteral(totaldigit.Value)))
			if totaldigit.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + totaldigit.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", totaldigitIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, union := range __gong__sortStageSetInstances(stageSet.Stage.Unions, stageSet.Stage.Union_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unionIdent := "__models" + union.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Union{Name: %s}).Stage(stageSet.Stage)", unionIdent, __gong__toRawStringLiteral(union.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unionIdent, __gong__toRawStringLiteral(union.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MemberTypes = %s", unionIdent, __gong__toRawStringLiteral(union.MemberTypes)))
			if union.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + union.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", unionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, whitespace := range __gong__sortStageSetInstances(stageSet.Stage.WhiteSpaces, stageSet.Stage.WhiteSpace_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			whitespaceIdent := "__models" + whitespace.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.WhiteSpace{Name: %s}).Stage(stageSet.Stage)", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", whitespaceIdent, __gong__toRawStringLiteral(whitespace.Value)))
			if whitespace.Annotation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + whitespace.Annotation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = %s", whitespaceIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

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

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/app/xsd/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

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

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "All":
					identifierMap[ident.Name] = __gong__stageSetInit(new(All), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Annotation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Annotation), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Attribute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Attribute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AttributeGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AttributeGroup), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Choice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Choice), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ComplexContent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ComplexContent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ComplexType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ComplexType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Documentation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Documentation), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Element":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Element), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Enumeration":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Enumeration), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Extension":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extension), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Length":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Length), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MaxInclusive":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MaxInclusive), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MaxLength":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MaxLength), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MinInclusive":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MinInclusive), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MinLength":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MinLength), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pattern":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pattern), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Restriction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Restriction), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Schema":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Schema), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Sequence":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sequence), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SimpleContent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SimpleContent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SimpleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SimpleType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TotalDigit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TotalDigit), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Union":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Union), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "WhiteSpace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WhiteSpace), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Documentations, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
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
						__gong__assignPointer(&inst.OuterElement, rhs, identifierMap)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "Order":
						inst.Order = GongExtractInt(rhs)
					case "Depth":
						inst.Depth = GongExtractInt(rhs)
					case "MinOccurs":
						inst.MinOccurs = GongExtractString(rhs)
					case "MaxOccurs":
						inst.MaxOccurs = GongExtractString(rhs)
					case "Extension":
						__gong__assignPointer(&inst.Extension, rhs, identifierMap)
					case "SimpleContent":
						__gong__assignPointer(&inst.SimpleContent, rhs, identifierMap)
					case "ComplexContent":
						__gong__assignPointer(&inst.ComplexContent, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
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
						__gong__assignPointer(&inst.SimpleType, rhs, identifierMap)
					case "ComplexType":
						__gong__assignPointer(&inst.ComplexType, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "IsDuplicatedInXSD":
						inst.IsDuplicatedInXSD = GongExtractBool(rhs)
					}
				case *Enumeration:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					}
				case *Group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Ref":
						inst.Ref = GongExtractString(rhs)
					case "IsAnonymous":
						inst.IsAnonymous = GongExtractBool(rhs)
					case "OuterElement":
						__gong__assignPointer(&inst.OuterElement, rhs, identifierMap)
					case "HasNameConflict":
						inst.HasNameConflict = GongExtractBool(rhs)
					case "GoIdentifier":
						inst.GoIdentifier = GongExtractString(rhs)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MaxInclusive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MaxLength:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MinInclusive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MinLength:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Pattern:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Restriction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Base":
						inst.Base = GongExtractString(rhs)
					case "Enumerations":
						__gong__assignSliceOfPointers(&inst.Enumerations, rhs, identifierMap)
					case "MinInclusive":
						__gong__assignPointer(&inst.MinInclusive, rhs, identifierMap)
					case "MaxInclusive":
						__gong__assignPointer(&inst.MaxInclusive, rhs, identifierMap)
					case "Pattern":
						__gong__assignPointer(&inst.Pattern, rhs, identifierMap)
					case "WhiteSpace":
						__gong__assignPointer(&inst.WhiteSpace, rhs, identifierMap)
					case "MinLength":
						__gong__assignPointer(&inst.MinLength, rhs, identifierMap)
					case "MaxLength":
						__gong__assignPointer(&inst.MaxLength, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "TotalDigit":
						__gong__assignPointer(&inst.TotalDigit, rhs, identifierMap)
					}
				case *Schema:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Xs":
						inst.Xs = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
					case "SimpleTypes":
						__gong__assignSliceOfPointers(&inst.SimpleTypes, rhs, identifierMap)
					case "ComplexTypes":
						__gong__assignSliceOfPointers(&inst.ComplexTypes, rhs, identifierMap)
					case "AttributeGroups":
						__gong__assignSliceOfPointers(&inst.AttributeGroups, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "OuterElementName":
						inst.OuterElementName = GongExtractString(rhs)
					case "Sequences":
						__gong__assignSliceOfPointers(&inst.Sequences, rhs, identifierMap)
					case "Alls":
						__gong__assignSliceOfPointers(&inst.Alls, rhs, identifierMap)
					case "Choices":
						__gong__assignSliceOfPointers(&inst.Choices, rhs, identifierMap)
					case "Groups":
						__gong__assignSliceOfPointers(&inst.Groups, rhs, identifierMap)
					case "Elements":
						__gong__assignSliceOfPointers(&inst.Elements, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Extension, rhs, identifierMap)
					case "Restriction":
						__gong__assignPointer(&inst.Restriction, rhs, identifierMap)
					}
				case *SimpleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Restriction":
						__gong__assignPointer(&inst.Restriction, rhs, identifierMap)
					case "Union":
						__gong__assignPointer(&inst.Union, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Union:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
					case "MemberTypes":
						inst.MemberTypes = GongExtractString(rhs)
					}
				case *WhiteSpace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Annotation":
						__gong__assignPointer(&inst.Annotation, rhs, identifierMap)
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

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}
