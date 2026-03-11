package models

import (
	"log"

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

	gongStructSet := *gong.GetGongstructInstancesMap[gong.GongStruct](stager.gongStage)
	// gongEnumSet := *gong.GetGongstructInstancesMap[gong.GongEnum](stager.gongStage)

	for gongStructShape := range *GetGongstructInstancesSetFromPointerType[*GongStructShape](stager.stage) {
		gongStructName := IdentifierMetaToGongStructName(gongStructShape.IdentifierMeta)
		_, ok := gongStructSet[gongStructName]

		if !ok {
			gongStructShape.Unstage(stager.stage)
			needCommit = true
			continue
		}
	}

	for fieldShape := range *GetGongstructInstancesSetFromPointerType[*AttributeShape](stager.stage) {
		structname, fieldShapeName := IdentifierMetaToStructAndFieldName(fieldShape.IdentifierMeta)

		_, ok := gongStructSet[structname]

		if !ok {
			fieldShape.Unstage(stager.stage)
			needCommit = true
			continue
		}

		var fieldFound bool
		for _, field := range gongStructSet[structname].Fields {
			if field.GetName() == fieldShapeName {
				fieldFound = true
			}
		}
		if !fieldFound {
			fieldShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	for linkShape := range *GetGongstructInstancesSetFromPointerType[*LinkShape](stager.stage) {
		structname, fieldShapeName := IdentifierMetaToStructAndFieldName(linkShape.IdentifierMeta)

		_, ok := gongStructSet[structname]

		if !ok {
			linkShape.Unstage(stager.stage)
			needCommit = true
			continue
		}

		var fieldFound bool
		for _, field := range gongStructSet[structname].Fields {
			if field.GetName() == fieldShapeName {
				fieldFound = true
			}
		}
		if !fieldFound {
			linkShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	gongEnumSet := *gong.GetGongstructInstancesMap[gong.GongEnum](stager.gongStage)
	for gongEnumShape := range *GetGongstructInstancesSetFromPointerType[*GongEnumShape](stager.stage) {

		gongEnumName := GongEnumIdentifierMetaToGongEnumName(gongEnumShape.IdentifierMeta)
		_, ok := gongEnumSet[gongEnumName]
		if !ok {
			gongEnumShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	gongEnumValueSet := *gong.GetGongstructInstancesMap[gong.GongEnumValue](stager.gongStage)
	for gongEnumValueShape := range *GetGongstructInstancesSetFromPointerType[*GongEnumValueShape](stager.stage) {

		gongEnumValueName := GongEnumValueShapeIdentifierMetaToValueName(gongEnumValueShape.IdentifierMeta)
		_, ok := gongEnumValueSet[gongEnumValueName]
		if !ok {
			gongEnumValueShape.Unstage(stager.stage)
			needCommit = true
		}
	}

	return
}
