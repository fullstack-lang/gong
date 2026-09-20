package models

import (
	"log"
	"strings"

	gong "github.com/fullstack-lang/gong/go/models"
)

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage

	pass := 0
	for {
		if stager.enforceSemanticOnePass(false, stage) {
			needCommit = true
			pass++
		} else {
			break
		}
	}

	if needCommit {
		log.Printf("Semantic enforcement performed in %d passes\n", pass)
		stage.CommitWithSuspendedCallbacks()
	}

	return
}

func (stager *Stager) enforceSemanticOnePass(needCommit bool, stage *Stage) bool {
	// VERY important because the probe only unstages objects
	// this is the Clean that delete them from slices and pointers that reference
	// them. If the checkout is not performed, the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	needCommit = stage.Clean() || needCommit

	needCommit = stager.enforceSemanticShapeWithCorrectMEtaIDentifiers() || needCommit

	return needCommit
}

func (stager *Stager) enforceSemanticShapeWithCorrectMEtaIDentifiers() (needCommit bool) {

	gongStructSet := stager.gongStage.GetInstancesMapByName[*gong.GongStruct]()
	// gongEnumSet := stager.gongStage.GetInstancesMapByName[*gong.GongEnum]()

	for gongStructShape := range *stager.stage.GetInstancesSet[*GongStructShape]() {
		gongStructName := IdentifierMetaToGongStructName(gongStructShape.IdentifierMeta)
		_, ok := gongStructSet[gongStructName]

		if !ok {
			log.Println("doc removed shape", gongStructName)
			gongStructShape.Unstage(stager.stage)
			needCommit = true
			continue
		}
	}

	for fieldShape := range *stager.stage.GetInstancesSet[*AttributeShape]() {
		structname, fieldShapeName := IdentifierMetaToStructAndFieldName(fieldShape.IdentifierMeta)

		gongStruct, ok := gongStructSet[structname]

		if !ok {
			log.Println("doc removed attribute shape", structname, fieldShapeName)
			fieldShape.Unstage(stager.stage)
			needCommit = true
			continue
		}

		var fieldFound bool
		for _, field := range gongStruct.Fields {
			if field.GetName() == fieldShapeName {
				switch realField := field.(type) {
				case *gong.GongBasicField:
					names := strings.Split(realField.DeclaredType, ".")
					fieldTypeName := names[len(names)-1]
					if fieldShape.Fieldtypename != fieldTypeName {
						fieldShape.Fieldtypename = fieldTypeName
						needCommit = true
					}
					fieldFound = true
				case *gong.GongTimeField:
					if fieldShape.Fieldtypename != "Time" {
						fieldShape.Fieldtypename = "Time"
						needCommit = true
					}
					fieldFound = true
				}
			}
		}
		if !fieldFound {
			log.Println("doc removed attribute shape", structname, fieldShapeName)
			fieldShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	for linkShape := range *stager.stage.GetInstancesSet[*LinkShape]() {
		structname, fieldShapeName := IdentifierMetaToStructAndFieldName(linkShape.IdentifierMeta)

		gongStruct, ok := gongStructSet[structname]

		if !ok {
			log.Println("doc removed link shape", structname, fieldShapeName)
			linkShape.Unstage(stager.stage)
			needCommit = true
			continue
		}

		var fieldFound bool
		for _, field := range gongStruct.Fields {
			if field.GetName() == fieldShapeName {
				switch realField := field.(type) {
				case *gong.PointerToGongStructField:
					targetStructName := realField.GongStruct.Name
					if _, targetOk := gongStructSet[targetStructName]; !targetOk {
						break
					}
					expectedTargetMultiplicity := ZERO_ONE
					expectedSourceMultiplicity := MANY
					targetPkgName := "models"
					if realField.GongStruct.ModelPkg != nil && realField.GongStruct.ModelPkg.PkgGoName != "" {
						targetPkgName = realField.GongStruct.ModelPkg.PkgGoName
					}
					expectedFieldTypeIdentifierMeta := GongStructNameToIdentifierWithPackage(targetPkgName, targetStructName) + "{}"

					if linkShape.TargetMultiplicity != expectedTargetMultiplicity {
						linkShape.TargetMultiplicity = expectedTargetMultiplicity
						needCommit = true
					}
					if linkShape.SourceMultiplicity != expectedSourceMultiplicity {
						linkShape.SourceMultiplicity = expectedSourceMultiplicity
						needCommit = true
					}
					if linkShape.FieldTypeIdentifierMeta != expectedFieldTypeIdentifierMeta {
						linkShape.FieldTypeIdentifierMeta = expectedFieldTypeIdentifierMeta
						needCommit = true
					}
					fieldFound = true
				case *gong.SliceOfPointerToGongStructField:
					targetStructName := realField.GongStruct.Name
					if _, targetOk := gongStructSet[targetStructName]; !targetOk {
						break
					}
					expectedTargetMultiplicity := MANY
					expectedSourceMultiplicity := MANY
					targetPkgName := "models"
					if realField.GongStruct.ModelPkg != nil && realField.GongStruct.ModelPkg.PkgGoName != "" {
						targetPkgName = realField.GongStruct.ModelPkg.PkgGoName
					}
					expectedFieldTypeIdentifierMeta := GongStructNameToIdentifierWithPackage(targetPkgName, targetStructName) + "{}"

					if linkShape.TargetMultiplicity != expectedTargetMultiplicity {
						linkShape.TargetMultiplicity = expectedTargetMultiplicity
						needCommit = true
					}
					if linkShape.SourceMultiplicity != expectedSourceMultiplicity {
						linkShape.SourceMultiplicity = expectedSourceMultiplicity
						needCommit = true
					}
					if linkShape.FieldTypeIdentifierMeta != expectedFieldTypeIdentifierMeta {
						linkShape.FieldTypeIdentifierMeta = expectedFieldTypeIdentifierMeta
						needCommit = true
					}
					fieldFound = true
				}
			}
		}
		if !fieldFound {
			log.Println("doc removed link shape", structname, fieldShapeName)
			linkShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	gongEnumSet := stager.gongStage.GetInstancesMapByName[*gong.GongEnum]()
	for gongEnumShape := range *stager.stage.GetInstancesSet[*GongEnumShape]() {

		gongEnumName := GongEnumIdentifierMetaToGongEnumName(gongEnumShape.IdentifierMeta)
		_, ok := gongEnumSet[gongEnumName]
		if !ok {
			log.Println("doc removed enum shape", gongEnumName)
			gongEnumShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	gongEnumValueSet := stager.gongStage.GetInstancesMapByName[*gong.GongEnumValue]()
	for gongEnumValueShape := range *stager.stage.GetInstancesSet[*GongEnumValueShape]() {

		gongEnumValueName := GongEnumValueShapeIdentifierMetaToValueName(gongEnumValueShape.IdentifierMeta)
		_, ok := gongEnumValueSet[gongEnumValueName]
		if !ok {
			log.Println("doc removed enum value shape", gongEnumValueName)
			gongEnumValueShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	return
}
