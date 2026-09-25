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
		for _, gongbasicfield := range __gong__sortStageSetInstances(stageSet.Stage.GongBasicFields, stageSet.Stage.GongBasicField_stagedOrder) {
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
		for _, gongenum := range __gong__sortStageSetInstances(stageSet.Stage.GongEnums, stageSet.Stage.GongEnum_stagedOrder) {
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
		for _, gongenumvalue := range __gong__sortStageSetInstances(stageSet.Stage.GongEnumValues, stageSet.Stage.GongEnumValue_stagedOrder) {
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
		for _, gonglink := range __gong__sortStageSetInstances(stageSet.Stage.GongLinks, stageSet.Stage.GongLink_stagedOrder) {
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
		for _, gongnote := range __gong__sortStageSetInstances(stageSet.Stage.GongNotes, stageSet.Stage.GongNote_stagedOrder) {
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
		for _, gongstruct := range __gong__sortStageSetInstances(stageSet.Stage.GongStructs, stageSet.Stage.GongStruct_stagedOrder) {
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
		for _, gongtimefield := range __gong__sortStageSetInstances(stageSet.Stage.GongTimeFields, stageSet.Stage.GongTimeField_stagedOrder) {
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
		for _, metareference := range __gong__sortStageSetInstances(stageSet.Stage.MetaReferences, stageSet.Stage.MetaReference_stagedOrder) {
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
		for _, modelpkg := range __gong__sortStageSetInstances(stageSet.Stage.ModelPkgs, stageSet.Stage.ModelPkg_stagedOrder) {
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
		for _, pointertogongstructfield := range __gong__sortStageSetInstances(stageSet.Stage.PointerToGongStructFields, stageSet.Stage.PointerToGongStructField_stagedOrder) {
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
		for _, sliceofpointertogongstructfield := range __gong__sortStageSetInstances(stageSet.Stage.SliceOfPointerToGongStructFields, stageSet.Stage.SliceOfPointerToGongStructField_stagedOrder) {
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
		for _, stagesetfield := range __gong__sortStageSetInstances(stageSet.Stage.StageSetFields, stageSet.Stage.StageSetField_stagedOrder) {
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
		for _, stagesetmodel := range __gong__sortStageSetInstances(stageSet.Stage.StageSetModels, stageSet.Stage.StageSetModel_stagedOrder) {
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
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongBasicField), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GongEnum":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongEnum), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GongEnumValue":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongEnumValue), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GongLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongLink), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GongNote":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongNote), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GongStruct":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongStruct), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GongTimeField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GongTimeField), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MetaReference":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MetaReference), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ModelPkg":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ModelPkg), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PointerToGongStructField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PointerToGongStructField), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SliceOfPointerToGongStructField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SliceOfPointerToGongStructField), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StageSetField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StageSetField), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StageSetModel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StageSetModel), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
						__gong__assignPointer(&inst.GongEnum, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.GongEnumValues, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Links, rhs, identifierMap)
					}
				case *GongStruct:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "GongBasicFields":
						__gong__assignSliceOfPointers(&inst.GongBasicFields, rhs, identifierMap)
					case "GongTimeFields":
						__gong__assignSliceOfPointers(&inst.GongTimeFields, rhs, identifierMap)
					case "PointerToGongStructFields":
						__gong__assignSliceOfPointers(&inst.PointerToGongStructFields, rhs, identifierMap)
					case "SliceOfPointerToGongStructFields":
						__gong__assignSliceOfPointers(&inst.SliceOfPointerToGongStructFields, rhs, identifierMap)
					case "HasOnAfterUpdateSignature":
						inst.HasOnAfterUpdateSignature = GongExtractBool(rhs)
					case "IsIgnoredForFront":
						inst.IsIgnoredForFront = GongExtractBool(rhs)
					case "IsOmittedForMarshalling":
						inst.IsOmittedForMarshalling = GongExtractBool(rhs)
					case "ModelPkg":
						__gong__assignPointer(&inst.ModelPkg, rhs, identifierMap)
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
						__gong__assignPointer(&inst.StageSet, rhs, identifierMap)
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
						__gong__assignPointer(&inst.GongStruct, rhs, identifierMap)
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
						__gong__assignPointer(&inst.GongStruct, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.Fields, rhs, identifierMap)
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
