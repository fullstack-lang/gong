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

	map_Pkg_StructName_GongStruct := make(map[string]*gong.GongStruct)
	map_StructName_GongStruct := make(map[string]*gong.GongStruct)
	for gongStruct := range *stager.gongStage.GetInstancesSet[*gong.GongStruct]() {
		pkgName := "models"
		if gongStruct.ModelPkg != nil && gongStruct.ModelPkg.PkgGoName != "" {
			pkgName = gongStruct.ModelPkg.PkgGoName
		}
		map_Pkg_StructName_GongStruct[pkgName+"."+gongStruct.Name] = gongStruct
		if _, exists := map_StructName_GongStruct[gongStruct.Name]; !exists {
			map_StructName_GongStruct[gongStruct.Name] = gongStruct
		}
	}
	getGongStruct := func(pkgName, structName string) *gong.GongStruct {
		if pkgName != "" {
			if gs, ok := map_Pkg_StructName_GongStruct[pkgName+"."+structName]; ok {
				return gs
			}
		}
		return map_StructName_GongStruct[structName]
	}

	for gongStructShape := range *stager.stage.GetInstancesSet[*GongStructShape]() {
		pkgName, gongStructName := IdentifierMetaToPackageAndGongStructName(gongStructShape.IdentifierMeta)
		gongStruct := getGongStruct(pkgName, gongStructName)

		if gongStruct == nil {
			log.Println("doc removed shape", gongStructName)
			gongStructShape.Unstage(stager.stage)
			needCommit = true
			continue
		}
	}

	for fieldShape := range *stager.stage.GetInstancesSet[*AttributeShape]() {
		pkgName, structname, fieldShapeName := IdentifierMetaToPackageStructAndFieldName(fieldShape.IdentifierMeta)
		gongStruct := getGongStruct(pkgName, structname)

		if gongStruct == nil {
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
		pkgName, structname, fieldShapeName := IdentifierMetaToPackageStructAndFieldName(linkShape.IdentifierMeta)
		gongStruct := getGongStruct(pkgName, structname)

		if gongStruct == nil {
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
					targetGS := realField.GongStruct
					if targetGS == nil {
						break
					}
					targetStructName := targetGS.Name
					if _, targetOk := map_StructName_GongStruct[targetStructName]; !targetOk {
						break
					}
					expectedTargetMultiplicity := ZERO_ONE
					expectedSourceMultiplicity := MANY

					expectedFieldTypeIdentifierMeta := stager.computeExpectedFieldTypeIdentifierMeta(targetGS, linkShape.FieldTypeIdentifierMeta)

					if linkShape.TargetMultiplicity != expectedTargetMultiplicity {
						linkShape.TargetMultiplicity = expectedTargetMultiplicity
						needCommit = true
					}
					if linkShape.SourceMultiplicity != expectedSourceMultiplicity {
						linkShape.SourceMultiplicity = expectedSourceMultiplicity
						needCommit = true
					}
					if expectedFieldTypeIdentifierMeta != "" && linkShape.FieldTypeIdentifierMeta != expectedFieldTypeIdentifierMeta {
						linkShape.FieldTypeIdentifierMeta = expectedFieldTypeIdentifierMeta
						needCommit = true
					}
					fieldFound = true
				case *gong.SliceOfPointerToGongStructField:
					targetGS := realField.GongStruct
					if targetGS == nil {
						break
					}
					targetStructName := targetGS.Name
					if _, targetOk := map_StructName_GongStruct[targetStructName]; !targetOk {
						break
					}
					expectedTargetMultiplicity := MANY
					expectedSourceMultiplicity := MANY

					expectedFieldTypeIdentifierMeta := stager.computeExpectedFieldTypeIdentifierMeta(targetGS, linkShape.FieldTypeIdentifierMeta)

					if linkShape.TargetMultiplicity != expectedTargetMultiplicity {
						linkShape.TargetMultiplicity = expectedTargetMultiplicity
						needCommit = true
					}
					if linkShape.SourceMultiplicity != expectedSourceMultiplicity {
						linkShape.SourceMultiplicity = expectedSourceMultiplicity
						needCommit = true
					}
					if expectedFieldTypeIdentifierMeta != "" && linkShape.FieldTypeIdentifierMeta != expectedFieldTypeIdentifierMeta {
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

func (stager *Stager) computeExpectedFieldTypeIdentifierMeta(targetGS *gong.GongStruct, currentFieldTypeIdentifierMeta any) string {
	targetStructName := targetGS.Name
	targetPkgName := "models"
	if targetGS.ModelPkg != nil && targetGS.ModelPkg.PkgGoName != "" {
		targetPkgName = targetGS.ModelPkg.PkgGoName
	}

	// 1. If current FieldTypeIdentifierMeta is already valid and points to targetStructName with a compatible package, preserve it!
	if currentFieldTypeIdentifierMeta != nil {
		currentPkg, currentName := IdentifierMetaToPackageAndGongStructName(currentFieldTypeIdentifierMeta)
		if currentName == targetStructName {
			if currentPkg == targetPkgName {
				if str, ok := currentFieldTypeIdentifierMeta.(string); ok {
					return str
				}
			}
			for _, imp := range stager.stage.MetaPackageImports {
				if imp.Alias == RefPrefixReferencedPackage+currentPkg {
					cleanPath := strings.Trim(imp.Path, "\"")
					if (targetGS.ModelPkg != nil && targetGS.ModelPkg.PkgPath != "" && cleanPath == targetGS.ModelPkg.PkgPath) ||
						strings.HasSuffix(cleanPath, "/"+targetPkgName) ||
						(targetPkgName == "models" && (cleanPath == "models" || strings.HasSuffix(cleanPath, "/models"))) {
						if str, ok := currentFieldTypeIdentifierMeta.(string); ok {
							return str
						}
					}
				}
			}
		}
	}

	// 2. Check if a GongStructShape in the diagram has the target struct
	for gongStructShape := range *stager.stage.GetInstancesSet[*GongStructShape]() {
		pkg, name := IdentifierMetaToPackageAndGongStructName(gongStructShape.IdentifierMeta)
		if name == targetStructName && (pkg == targetPkgName || targetPkgName == "models" || pkg == "") {
			if str, ok := gongStructShape.IdentifierMeta.(string); ok {
				return str
			}
		}
	}

	// 3. Match via stager.stage.MetaPackageImports
	if len(stager.stage.MetaPackageImports) > 0 {
		for _, imp := range stager.stage.MetaPackageImports {
			cleanPath := strings.Trim(imp.Path, "\"")
			if targetGS.ModelPkg != nil && targetGS.ModelPkg.PkgPath != "" && cleanPath == targetGS.ModelPkg.PkgPath {
				return imp.Alias + "." + targetStructName + "{}"
			}
			if imp.Alias == RefPrefixReferencedPackage+targetPkgName {
				return imp.Alias + "." + targetStructName + "{}"
			}
			if strings.HasSuffix(cleanPath, "/"+targetPkgName) {
				return imp.Alias + "." + targetStructName + "{}"
			}
		}
	}

	// 4. Default to GongStructNameToIdentifierWithPackage
	return GongStructNameToIdentifierWithPackage(targetPkgName, targetStructName) + "{}"
}
