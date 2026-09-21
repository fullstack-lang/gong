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
		gongbasicfieldOrdered := []*GongBasicField{}
		for gongbasicfield := range stageSet.Stage.GongBasicFields {
			gongbasicfieldOrdered = append(gongbasicfieldOrdered, gongbasicfield)
		}
		sort.Slice(gongbasicfieldOrdered, func(i, j int) bool {
			return stageSet.Stage.GongBasicField_stagedOrder[gongbasicfieldOrdered[i]] < stageSet.Stage.GongBasicField_stagedOrder[gongbasicfieldOrdered[j]]
		})
		for _, gongbasicfield := range gongbasicfieldOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongbasicfieldIdent := "__models" + gongbasicfield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongBasicField{Name: %s}).Stage(stageSet.Stage)", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.BasicKindName = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.BasicKindName)))
			values.WriteString(fmt.Sprintf("\n\t%s.DeclaredType = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.DeclaredType)))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", gongbasicfieldIdent, gongbasicfield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", gongbasicfieldIdent, __gong__toRawStringLiteral(gongbasicfield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", gongbasicfieldIdent, gongbasicfield.IsAccordionEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", gongbasicfieldIdent, gongbasicfield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTextArea = %t", gongbasicfieldIdent, gongbasicfield.IsTextArea))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokeWidth = %t", gongbasicfieldIdent, gongbasicfield.IsBespokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeWidth = %d", gongbasicfieldIdent, gongbasicfield.BespokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokeHeight = %t", gongbasicfieldIdent, gongbasicfield.IsBespokeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeHeight = %d", gongbasicfieldIdent, gongbasicfield.BespokeHeight))
			if gongbasicfield.GongEnum != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + gongbasicfield.GongEnum.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnum = %s", gongbasicfieldIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongenumOrdered := []*GongEnum{}
		for gongenum := range stageSet.Stage.GongEnums {
			gongenumOrdered = append(gongenumOrdered, gongenum)
		}
		sort.Slice(gongenumOrdered, func(i, j int) bool {
			return stageSet.Stage.GongEnum_stagedOrder[gongenumOrdered[i]] < stageSet.Stage.GongEnum_stagedOrder[gongenumOrdered[j]]
		})
		for _, gongenum := range gongenumOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumIdent := "__models" + gongenum.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnum{Name: %s}).Stage(stageSet.Stage)", gongenumIdent, __gong__toRawStringLiteral(gongenum.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumIdent, __gong__toRawStringLiteral(gongenum.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %d", gongenumIdent, int(gongenum.Type)))
			for _, elem := range gongenum.GongEnumValues {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnumValues = append(%s.GongEnumValues, %s)", gongenumIdent, gongenumIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongenumvalueOrdered := []*GongEnumValue{}
		for gongenumvalue := range stageSet.Stage.GongEnumValues {
			gongenumvalueOrdered = append(gongenumvalueOrdered, gongenumvalue)
		}
		sort.Slice(gongenumvalueOrdered, func(i, j int) bool {
			return stageSet.Stage.GongEnumValue_stagedOrder[gongenumvalueOrdered[i]] < stageSet.Stage.GongEnumValue_stagedOrder[gongenumvalueOrdered[j]]
		})
		for _, gongenumvalue := range gongenumvalueOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumvalueIdent := "__models" + gongenumvalue.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnumValue{Name: %s}).Stage(stageSet.Stage)", gongenumvalueIdent, __gong__toRawStringLiteral(gongenumvalue.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumvalueIdent, __gong__toRawStringLiteral(gongenumvalue.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", gongenumvalueIdent, __gong__toRawStringLiteral(gongenumvalue.Value)))
		}
	}
	if stageSet.Stage != nil {
		gonglinkOrdered := []*GongLink{}
		for gonglink := range stageSet.Stage.GongLinks {
			gonglinkOrdered = append(gonglinkOrdered, gonglink)
		}
		sort.Slice(gonglinkOrdered, func(i, j int) bool {
			return stageSet.Stage.GongLink_stagedOrder[gonglinkOrdered[i]] < stageSet.Stage.GongLink_stagedOrder[gonglinkOrdered[j]]
		})
		for _, gonglink := range gonglinkOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gonglinkIdent := "__models" + gonglink.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongLink{Name: %s}).Stage(stageSet.Stage)", gonglinkIdent, __gong__toRawStringLiteral(gonglink.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gonglinkIdent, __gong__toRawStringLiteral(gonglink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Recv = %s", gonglinkIdent, __gong__toRawStringLiteral(gonglink.Recv)))
			values.WriteString(fmt.Sprintf("\n\t%s.ImportPath = %s", gonglinkIdent, __gong__toRawStringLiteral(gonglink.ImportPath)))
		}
	}
	if stageSet.Stage != nil {
		gongnoteOrdered := []*GongNote{}
		for gongnote := range stageSet.Stage.GongNotes {
			gongnoteOrdered = append(gongnoteOrdered, gongnote)
		}
		sort.Slice(gongnoteOrdered, func(i, j int) bool {
			return stageSet.Stage.GongNote_stagedOrder[gongnoteOrdered[i]] < stageSet.Stage.GongNote_stagedOrder[gongnoteOrdered[j]]
		})
		for _, gongnote := range gongnoteOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongnoteIdent := "__models" + gongnote.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongNote{Name: %s}).Stage(stageSet.Stage)", gongnoteIdent, __gong__toRawStringLiteral(gongnote.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongnoteIdent, __gong__toRawStringLiteral(gongnote.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Body = %s", gongnoteIdent, __gong__toRawStringLiteral(gongnote.Body)))
			values.WriteString(fmt.Sprintf("\n\t%s.BodyHTML = %s", gongnoteIdent, __gong__toRawStringLiteral(gongnote.BodyHTML)))
			for _, elem := range gongnote.Links {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Links = append(%s.Links, %s)", gongnoteIdent, gongnoteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongstructOrdered := []*GongStruct{}
		for gongstruct := range stageSet.Stage.GongStructs {
			gongstructOrdered = append(gongstructOrdered, gongstruct)
		}
		sort.Slice(gongstructOrdered, func(i, j int) bool {
			return stageSet.Stage.GongStruct_stagedOrder[gongstructOrdered[i]] < stageSet.Stage.GongStruct_stagedOrder[gongstructOrdered[j]]
		})
		for _, gongstruct := range gongstructOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongstructIdent := "__models" + gongstruct.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongStruct{Name: %s}).Stage(stageSet.Stage)", gongstructIdent, __gong__toRawStringLiteral(gongstruct.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongstructIdent, __gong__toRawStringLiteral(gongstruct.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.HasOnAfterUpdateSignature = %t", gongstructIdent, gongstruct.HasOnAfterUpdateSignature))
			values.WriteString(fmt.Sprintf("\n\t%s.IsIgnoredForFront = %t", gongstructIdent, gongstruct.IsIgnoredForFront))
			values.WriteString(fmt.Sprintf("\n\t%s.IsOmittedForMarshalling = %t", gongstructIdent, gongstruct.IsOmittedForMarshalling))
			for _, elem := range gongstruct.GongBasicFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongBasicFields = append(%s.GongBasicFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			for _, elem := range gongstruct.GongTimeFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongTimeFields = append(%s.GongTimeFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			for _, elem := range gongstruct.PointerToGongStructFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointerToGongStructFields = append(%s.PointerToGongStructFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			for _, elem := range gongstruct.SliceOfPointerToGongStructFields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SliceOfPointerToGongStructFields = append(%s.SliceOfPointerToGongStructFields, %s)", gongstructIdent, gongstructIdent, targetIdent))
			}
			if gongstruct.ModelPkg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + gongstruct.ModelPkg.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ModelPkg = %s", gongstructIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongtimefieldOrdered := []*GongTimeField{}
		for gongtimefield := range stageSet.Stage.GongTimeFields {
			gongtimefieldOrdered = append(gongtimefieldOrdered, gongtimefield)
		}
		sort.Slice(gongtimefieldOrdered, func(i, j int) bool {
			return stageSet.Stage.GongTimeField_stagedOrder[gongtimefieldOrdered[i]] < stageSet.Stage.GongTimeField_stagedOrder[gongtimefieldOrdered[j]]
		})
		for _, gongtimefield := range gongtimefieldOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongtimefieldIdent := "__models" + gongtimefield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongTimeField{Name: %s}).Stage(stageSet.Stage)", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", gongtimefieldIdent, gongtimefield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", gongtimefieldIdent, gongtimefield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", gongtimefieldIdent, gongtimefield.IsAccordionEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeTimeFormat = %s", gongtimefieldIdent, __gong__toRawStringLiteral(gongtimefield.BespokeTimeFormat)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeFormOnly = %t", gongtimefieldIdent, gongtimefield.TimeFormOnly))
		}
	}
	if stageSet.Stage != nil {
		metareferenceOrdered := []*MetaReference{}
		for metareference := range stageSet.Stage.MetaReferences {
			metareferenceOrdered = append(metareferenceOrdered, metareference)
		}
		sort.Slice(metareferenceOrdered, func(i, j int) bool {
			return stageSet.Stage.MetaReference_stagedOrder[metareferenceOrdered[i]] < stageSet.Stage.MetaReference_stagedOrder[metareferenceOrdered[j]]
		})
		for _, metareference := range metareferenceOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metareferenceIdent := "__models" + metareference.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MetaReference{Name: %s}).Stage(stageSet.Stage)", metareferenceIdent, __gong__toRawStringLiteral(metareference.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metareferenceIdent, __gong__toRawStringLiteral(metareference.Name)))
		}
	}
	if stageSet.Stage != nil {
		modelpkgOrdered := []*ModelPkg{}
		for modelpkg := range stageSet.Stage.ModelPkgs {
			modelpkgOrdered = append(modelpkgOrdered, modelpkg)
		}
		sort.Slice(modelpkgOrdered, func(i, j int) bool {
			return stageSet.Stage.ModelPkg_stagedOrder[modelpkgOrdered[i]] < stageSet.Stage.ModelPkg_stagedOrder[modelpkgOrdered[j]]
		})
		for _, modelpkg := range modelpkgOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			modelpkgIdent := "__models" + modelpkg.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ModelPkg{Name: %s}).Stage(stageSet.Stage)", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PkgGoName = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.PkgGoName)))
			values.WriteString(fmt.Sprintf("\n\t%s.PkgPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.PkgPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.PathToGoSubDirectory = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.PathToGoSubDirectory)))
			values.WriteString(fmt.Sprintf("\n\t%s.OrmPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.OrmPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.DbOrmPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.DbOrmPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.DbLiteOrmPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.DbLiteOrmPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.DbPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.DbPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.ControllersPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.ControllersPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.FullstackPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.FullstackPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.StackPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.Level1StackPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.Level1StackPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.StaticPkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.StaticPkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.ProbePkgGenPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.ProbePkgGenPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgWorkspacePath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgWorkspacePath)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgWorkspaceName = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgWorkspaceName)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgDataLibrarySourceCodeDirectory = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgDataLibrarySourceCodeDirectory)))
			values.WriteString(fmt.Sprintf("\n\t%s.NgSpecificLibrarySourceCodeDirectory = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.NgSpecificLibrarySourceCodeDirectory)))
			values.WriteString(fmt.Sprintf("\n\t%s.MaterialLibDatamodelTargetPath = %s", modelpkgIdent, __gong__toRawStringLiteral(modelpkg.MaterialLibDatamodelTargetPath)))
			if modelpkg.StageSet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + modelpkg.StageSet.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StageSet = %s", modelpkgIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		pointertogongstructfieldOrdered := []*PointerToGongStructField{}
		for pointertogongstructfield := range stageSet.Stage.PointerToGongStructFields {
			pointertogongstructfieldOrdered = append(pointertogongstructfieldOrdered, pointertogongstructfield)
		}
		sort.Slice(pointertogongstructfieldOrdered, func(i, j int) bool {
			return stageSet.Stage.PointerToGongStructField_stagedOrder[pointertogongstructfieldOrdered[i]] < stageSet.Stage.PointerToGongStructField_stagedOrder[pointertogongstructfieldOrdered[j]]
		})
		for _, pointertogongstructfield := range pointertogongstructfieldOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointertogongstructfieldIdent := "__models" + pointertogongstructfield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PointerToGongStructField{Name: %s}).Stage(stageSet.Stage)", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", pointertogongstructfieldIdent, pointertogongstructfield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", pointertogongstructfieldIdent, pointertogongstructfield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", pointertogongstructfieldIdent, __gong__toRawStringLiteral(pointertogongstructfield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", pointertogongstructfieldIdent, pointertogongstructfield.IsAccordionEnd))
			values.WriteString(fmt.Sprintf("\n\t%s.IsType = %t", pointertogongstructfieldIdent, pointertogongstructfield.IsType))
			if pointertogongstructfield.GongStruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + pointertogongstructfield.GongStruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongStruct = %s", pointertogongstructfieldIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		sliceofpointertogongstructfieldOrdered := []*SliceOfPointerToGongStructField{}
		for sliceofpointertogongstructfield := range stageSet.Stage.SliceOfPointerToGongStructFields {
			sliceofpointertogongstructfieldOrdered = append(sliceofpointertogongstructfieldOrdered, sliceofpointertogongstructfield)
		}
		sort.Slice(sliceofpointertogongstructfieldOrdered, func(i, j int) bool {
			return stageSet.Stage.SliceOfPointerToGongStructField_stagedOrder[sliceofpointertogongstructfieldOrdered[i]] < stageSet.Stage.SliceOfPointerToGongStructField_stagedOrder[sliceofpointertogongstructfieldOrdered[j]]
		})
		for _, sliceofpointertogongstructfield := range sliceofpointertogongstructfieldOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sliceofpointertogongstructfieldIdent := "__models" + sliceofpointertogongstructfield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SliceOfPointerToGongStructField{Name: %s}).Stage(stageSet.Stage)", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Index = %d", sliceofpointertogongstructfieldIdent, sliceofpointertogongstructfield.Index))
			values.WriteString(fmt.Sprintf("\n\t%s.CompositeStructName = %s", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.CompositeStructName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionStart = %t", sliceofpointertogongstructfieldIdent, sliceofpointertogongstructfield.IsAccordionStart))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionName = %s", sliceofpointertogongstructfieldIdent, __gong__toRawStringLiteral(sliceofpointertogongstructfield.AccordionName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAccordionEnd = %t", sliceofpointertogongstructfieldIdent, sliceofpointertogongstructfield.IsAccordionEnd))
			if sliceofpointertogongstructfield.GongStruct != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + sliceofpointertogongstructfield.GongStruct.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongStruct = %s", sliceofpointertogongstructfieldIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		stagesetfieldOrdered := []*StageSetField{}
		for stagesetfield := range stageSet.Stage.StageSetFields {
			stagesetfieldOrdered = append(stagesetfieldOrdered, stagesetfield)
		}
		sort.Slice(stagesetfieldOrdered, func(i, j int) bool {
			return stageSet.Stage.StageSetField_stagedOrder[stagesetfieldOrdered[i]] < stageSet.Stage.StageSetField_stagedOrder[stagesetfieldOrdered[j]]
		})
		for _, stagesetfield := range stagesetfieldOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stagesetfieldIdent := "__models" + stagesetfield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StageSetField{Name: %s}).Stage(stageSet.Stage)", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PackageName = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.PackageName)))
			values.WriteString(fmt.Sprintf("\n\t%s.PackagePath = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.PackagePath)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsLocal = %t", stagesetfieldIdent, stagesetfield.IsLocal))
			values.WriteString(fmt.Sprintf("\n\t%s.ImportAlias = %s", stagesetfieldIdent, __gong__toRawStringLiteral(stagesetfield.ImportAlias)))
		}
	}
	if stageSet.Stage != nil {
		stagesetmodelOrdered := []*StageSetModel{}
		for stagesetmodel := range stageSet.Stage.StageSetModels {
			stagesetmodelOrdered = append(stagesetmodelOrdered, stagesetmodel)
		}
		sort.Slice(stagesetmodelOrdered, func(i, j int) bool {
			return stageSet.Stage.StageSetModel_stagedOrder[stagesetmodelOrdered[i]] < stageSet.Stage.StageSetModel_stagedOrder[stagesetmodelOrdered[j]]
		})
		for _, stagesetmodel := range stagesetmodelOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stagesetmodelIdent := "__models" + stagesetmodel.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StageSetModel{Name: %s}).Stage(stageSet.Stage)", stagesetmodelIdent, __gong__toRawStringLiteral(stagesetmodel.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stagesetmodelIdent, __gong__toRawStringLiteral(stagesetmodel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsManual = %t", stagesetmodelIdent, stagesetmodel.IsManual))
			for _, elem := range stagesetmodel.Fields {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fields = append(%s.Fields, %s)", stagesetmodelIdent, stagesetmodelIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/go/models"
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
		case "github.com/fullstack-lang/gong/go/models":
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
				case "GongBasicField":
					if !preserveOrder {
						inst := (&GongBasicField{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongBasicField)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongEnum":
					if !preserveOrder {
						inst := (&GongEnum{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongEnum)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongEnumValue":
					if !preserveOrder {
						inst := (&GongEnumValue{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongEnumValue)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongLink":
					if !preserveOrder {
						inst := (&GongLink{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongLink)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongNote":
					if !preserveOrder {
						inst := (&GongNote{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongNote)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongStruct":
					if !preserveOrder {
						inst := (&GongStruct{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongStruct)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongTimeField":
					if !preserveOrder {
						inst := (&GongTimeField{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongTimeField)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MetaReference":
					if !preserveOrder {
						inst := (&MetaReference{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MetaReference)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ModelPkg":
					if !preserveOrder {
						inst := (&ModelPkg{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ModelPkg)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PointerToGongStructField":
					if !preserveOrder {
						inst := (&PointerToGongStructField{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PointerToGongStructField)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SliceOfPointerToGongStructField":
					if !preserveOrder {
						inst := (&SliceOfPointerToGongStructField{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SliceOfPointerToGongStructField)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StageSetField":
					if !preserveOrder {
						inst := (&StageSetField{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StageSetField)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StageSetModel":
					if !preserveOrder {
						inst := (&StageSetModel{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StageSetModel)
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
				case *GongBasicField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "BasicKindName":
						inst.BasicKindName = GongExtractString(rhs)
					case "GongEnum":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GongEnum); ok {
									inst.GongEnum = typedTarget
								}
							}
						}
					case "DeclaredType":
						inst.DeclaredType = GongExtractString(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "IsTextArea":
						inst.IsTextArea = GongExtractBool(rhs)
					case "IsBespokeWidth":
						inst.IsBespokeWidth = GongExtractBool(rhs)
					case "BespokeWidth":
						inst.BespokeWidth = GongExtractInt(rhs)
					case "IsBespokeHeight":
						inst.IsBespokeHeight = GongExtractBool(rhs)
					case "BespokeHeight":
						inst.BespokeHeight = GongExtractInt(rhs)
					}
				case *GongEnum:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongEnumType(GongExtractInt(rhs))
					case "GongEnumValues":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongEnumValue); ok {
										inst.GongEnumValues = append(inst.GongEnumValues, typedTarget)
									}
								}
							}
						}
					}
				case *GongEnumValue:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *GongLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Recv":
						inst.Recv = GongExtractString(rhs)
					case "ImportPath":
						inst.ImportPath = GongExtractString(rhs)
					}
				case *GongNote:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Body":
						inst.Body = GongExtractString(rhs)
					case "BodyHTML":
						inst.BodyHTML = GongExtractString(rhs)
					case "Links":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongLink); ok {
										inst.Links = append(inst.Links, typedTarget)
									}
								}
							}
						}
					}
				case *GongStruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongBasicFields":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongBasicField); ok {
										inst.GongBasicFields = append(inst.GongBasicFields, typedTarget)
									}
								}
							}
						}
					case "GongTimeFields":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongTimeField); ok {
										inst.GongTimeFields = append(inst.GongTimeFields, typedTarget)
									}
								}
							}
						}
					case "PointerToGongStructFields":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*PointerToGongStructField); ok {
										inst.PointerToGongStructFields = append(inst.PointerToGongStructFields, typedTarget)
									}
								}
							}
						}
					case "SliceOfPointerToGongStructFields":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*SliceOfPointerToGongStructField); ok {
										inst.SliceOfPointerToGongStructFields = append(inst.SliceOfPointerToGongStructFields, typedTarget)
									}
								}
							}
						}
					case "HasOnAfterUpdateSignature":
						inst.HasOnAfterUpdateSignature = GongExtractBool(rhs)
					case "IsIgnoredForFront":
						inst.IsIgnoredForFront = GongExtractBool(rhs)
					case "IsOmittedForMarshalling":
						inst.IsOmittedForMarshalling = GongExtractBool(rhs)
					case "ModelPkg":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ModelPkg); ok {
									inst.ModelPkg = typedTarget
								}
							}
						}
					}
				case *GongTimeField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					case "BespokeTimeFormat":
						inst.BespokeTimeFormat = GongExtractString(rhs)
					case "TimeFormOnly":
						inst.TimeFormOnly = GongExtractBool(rhs)
					}
				case *MetaReference:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ModelPkg:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PkgGoName":
						inst.PkgGoName = GongExtractString(rhs)
					case "PkgPath":
						inst.PkgPath = GongExtractString(rhs)
					case "StageSet":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StageSetModel); ok {
									inst.StageSet = typedTarget
								}
							}
						}
					case "PathToGoSubDirectory":
						inst.PathToGoSubDirectory = GongExtractString(rhs)
					case "OrmPkgGenPath":
						inst.OrmPkgGenPath = GongExtractString(rhs)
					case "DbOrmPkgGenPath":
						inst.DbOrmPkgGenPath = GongExtractString(rhs)
					case "DbLiteOrmPkgGenPath":
						inst.DbLiteOrmPkgGenPath = GongExtractString(rhs)
					case "DbPkgGenPath":
						inst.DbPkgGenPath = GongExtractString(rhs)
					case "ControllersPkgGenPath":
						inst.ControllersPkgGenPath = GongExtractString(rhs)
					case "FullstackPkgGenPath":
						inst.FullstackPkgGenPath = GongExtractString(rhs)
					case "StackPkgGenPath":
						inst.StackPkgGenPath = GongExtractString(rhs)
					case "Level1StackPkgGenPath":
						inst.Level1StackPkgGenPath = GongExtractString(rhs)
					case "StaticPkgGenPath":
						inst.StaticPkgGenPath = GongExtractString(rhs)
					case "ProbePkgGenPath":
						inst.ProbePkgGenPath = GongExtractString(rhs)
					case "NgWorkspacePath":
						inst.NgWorkspacePath = GongExtractString(rhs)
					case "NgWorkspaceName":
						inst.NgWorkspaceName = GongExtractString(rhs)
					case "NgDataLibrarySourceCodeDirectory":
						inst.NgDataLibrarySourceCodeDirectory = GongExtractString(rhs)
					case "NgSpecificLibrarySourceCodeDirectory":
						inst.NgSpecificLibrarySourceCodeDirectory = GongExtractString(rhs)
					case "MaterialLibDatamodelTargetPath":
						inst.MaterialLibDatamodelTargetPath = GongExtractString(rhs)
					}
				case *PointerToGongStructField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongStruct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GongStruct); ok {
									inst.GongStruct = typedTarget
								}
							}
						}
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					case "IsType":
						inst.IsType = GongExtractBool(rhs)
					}
				case *SliceOfPointerToGongStructField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongStruct":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GongStruct); ok {
									inst.GongStruct = typedTarget
								}
							}
						}
					case "Index":
						inst.Index = GongExtractInt(rhs)
					case "CompositeStructName":
						inst.CompositeStructName = GongExtractString(rhs)
					case "IsAccordionStart":
						inst.IsAccordionStart = GongExtractBool(rhs)
					case "AccordionName":
						inst.AccordionName = GongExtractString(rhs)
					case "IsAccordionEnd":
						inst.IsAccordionEnd = GongExtractBool(rhs)
					}
				case *StageSetField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PackageName":
						inst.PackageName = GongExtractString(rhs)
					case "PackagePath":
						inst.PackagePath = GongExtractString(rhs)
					case "IsLocal":
						inst.IsLocal = GongExtractBool(rhs)
					case "ImportAlias":
						inst.ImportAlias = GongExtractString(rhs)
					}
				case *StageSetModel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Fields":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*StageSetField); ok {
										inst.Fields = append(inst.Fields, typedTarget)
									}
								}
							}
						}
					case "IsManual":
						inst.IsManual = GongExtractBool(rhs)
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
