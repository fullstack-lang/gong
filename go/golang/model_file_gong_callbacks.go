package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const ModelGongCallbacksFileTemplate = `package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case Tree:
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, target)
		}
	case Tree:
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, target)
		}
	case Tree:
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case Tree:
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point for generic get functions
	case Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	case Tree:
	}
}
`

// insertion points are places where the code is
// generated per gong struct
type ModelGongCallbacksStructInsertionId int

const (
	ModelGongCallbacksGenericGongstructTypes ModelGongCallbacksStructInsertionId = iota
	ModelGongCallbacksStructInsertionsNb
)

var ModelGongCallbacksStructSubTemplateCode map[ModelGongCallbacksStructInsertionId]string = // new line
map[ModelGongCallbacksStructInsertionId]string{
	ModelGongCallbacksGenericGongstructTypes: ` | *{{Structname}} | []*{{Structname}}`,
}

type ModelGongCallbacksFieldInsertionId int

func CodeGeneratorModelGongCallbacks(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongCallbacksFileTemplate

	subStructCodes := make(map[ModelGongCallbacksStructInsertionId]string)
	for subStructTemplate := range ModelGongCallbacksStructSubTemplateCode {
		subStructCodes[subStructTemplate] = ""
	}

	// sort gong structs per name (for reproductibility)
	gongStructs := []*models.GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		gongStructs = append(gongStructs, _struct)
	}
	sort.Slice(gongStructs[:], func(i, j int) bool {
		return gongStructs[i].Name < gongStructs[j].Name
	})

	for _, gongStruct := range gongStructs {

		if !gongStruct.HasNameField() {
			continue
		}

		// generatedCodeFromSubTemplate := models.Replace2(ModelGongCallbacksStructSubTemplateCode[subStructTemplate],
		// 	"{{structname}}", strings.ToLower(gongStruct.Name),
		// 	"{{Structname}}", gongStruct.Name)

		// subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate

	}

	for insertionPerStructId := ModelGongCallbacksStructInsertionId(0); insertionPerStructId < ModelGongCallbacksStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong_callbacks.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
