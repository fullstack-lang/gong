package models

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// VerySimpleCodeGenerator generates from elements of mdlPkg the file generatedFilePath with templateCode
func VerySimpleCodeGenerator(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	generatedFilePath string,
	templateCode string) {
	SimpleCodeGenerator(mdlPkg, pkgName, pkgGoPath, generatedFilePath, templateCode, map[string]string{})
}

func VerySimpleCodeGeneratorForGongStructWithNameField(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	generatedFilePath string,
	templateCode string) {
	SimpleCodeGeneratorForGongStructWithNameField(mdlPkg, pkgName, pkgGoPath, generatedFilePath, templateCode, map[string]string{})
}

// SimpleCodeGenerator generates from elements of mdlPkg the file generatedFilePath with templateCode and
// with subTemplateCode (code that accumulates per struct)
func SimpleCodeGenerator(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	generatedFilePath string,
	templateCode string,
	subTemplateCode map[string]string) {
	CodeGenerator(mdlPkg, pkgName, pkgGoPath, generatedFilePath, templateCode, subTemplateCode,
		map[string]string{}, map[string]string{},
		false,
		false)
}

// SimpleCodeGenerator generates from elements of mdlPkg the file generatedFilePath with templateCode and
// with subTemplateCode (code that accumulates per struct)
func SimpleCodeGeneratorForGongStructWithNameField(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	generatedFilePath string,
	templateCode string,
	subTemplateCode map[string]string) {
	CodeGenerator(mdlPkg, pkgName, pkgGoPath, generatedFilePath, templateCode, subTemplateCode,
		map[string]string{}, map[string]string{},
		true,
		false)
}

// SimpleCodeGenerator generates from elements of mdlPkg the file generatedFilePath with templateCode and
// with subTemplateCode (code that accumulates per struct)
// with subSubTemplateCode (code that accumulates per field of each struct)
func CodeGenerator(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	generatedFilePath string,
	templateCode string,
	subTemplateCode map[string]string,
	subSubTemplateCode map[string]string,
	// a sub sub template is generated within a sub template
	subSubToSubMap map[string]string,
	forGongStructWithNameFieldOnly bool,
	forGongStructWithHasOnUpdateSignatureOnly bool) {

	file, err := os.Create(generatedFilePath)

	if err != nil {
		file, err = os.Create(generatedFilePath)
	}

	if err != nil {
		log.Panic(err)
	}

	code := templateCode

	subCodes := make(map[string]string)
	for subTemplate := range subTemplateCode {
		subCodes[subTemplate] = ""
	}

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		if forGongStructWithNameFieldOnly && !_struct.HasNameField() {
			continue
		}
		if forGongStructWithHasOnUpdateSignatureOnly &&
			!_struct.HasOnAfterUpdateSignature {
			continue
		}
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		structName := strings.ToLower(_struct.Name)
		_ = structName

		subSubCodes := make(map[string]string)
		for subSubTemplate := range subSubTemplateCode {
			subSubCodes[subSubTemplate] = ""
		}

		// compute code from sub sub template
		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *PointerToGongStructField:

				fieldName := strings.ToLower(field.Name)
				assocStructName := strings.ToLower(field.GongStruct.Name)

				for subSubTemplate := range subSubTemplateCode {
					subSubCodes[subSubTemplate] += Replace6(subSubTemplateCode[subSubTemplate],
						"{{fieldName}}", fieldName,
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocStructName}}", assocStructName,
						"{{Structname}}", _struct.Name,
						"{{structname}}", structName)
				}
			}
		}

		// compute code from sub template
		for subTemplate := range subTemplateCode {
			subCodes[subTemplate] += Replace2(subTemplateCode[subTemplate],
				"{{Structname}}", _struct.Name,
				"{{structname}}", structName)

			// apply sub sub templates
			for subSubTemplate := range subSubTemplateCode {

				toReplace := "{{" + subSubTemplate + "}}"
				subCodes[subSubToSubMap[subSubTemplate]] = strings.ReplaceAll(
					subCodes[subSubToSubMap[subSubTemplate]],
					toReplace, subSubCodes[subSubTemplate])
			}
		}

	}
	for subTemplate := range subTemplateCode {
		toReplace := "{{" + string(subTemplate) + "}}"
		code = strings.ReplaceAll(code, toReplace, subCodes[subTemplate])
	}

	code = strings.ReplaceAll(code, "{{PkgName}}", pkgName)
	code = strings.ReplaceAll(code, "{{TitlePkgName}}", strings.Title(pkgName))
	code = strings.ReplaceAll(code, "{{pkgname}}", strings.ToLower(pkgName))
	code = strings.ReplaceAll(code, "{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))
	code = strings.ReplaceAll(code, "{{PkgPathAboveRoot}}", strings.ReplaceAll(pkgGoPath, "/go/models", ""))

	pkgPathRootWithoutSlashes := strings.ReplaceAll(pkgGoPath, "/models", "")
	pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "/", "_")
	pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "-", "_")
	pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, ".", "_")

	code = strings.ReplaceAll(code, "{{PkgPathRootWithoutSlashes}}", pkgPathRootWithoutSlashes)

	defer file.Close()
	fmt.Fprint(file, code)
}

func MultiCodeGenerator(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	// if path contains {{Structname}}, that means a file has to be generated per struct
	path string,
	templateCode string,
	subTemplateCode map[string]string,
	subSubTemplateCode map[string]string,
	// a sub sub template is generated within a sub template
	subSubToSubMap map[string]string,
) {

	var file *os.File
	var err error

	for _, _struct := range mdlPkg.GongStructs {

		code := templateCode

		subCodes := make(map[string]string)
		for subTemplate := range subTemplateCode {
			subCodes[subTemplate] = ""
		}

		file, err = os.Create(strings.ReplaceAll(path, "{{Structname}}", _struct.Name))
		if err != nil {
			log.Panic(err)
		}

		structName := strings.ToLower(_struct.Name)
		_ = structName

		subSubCodes := make(map[string]string)
		for subSubTemplate := range subSubTemplateCode {
			subSubCodes[subSubTemplate] = ""
		}

		// compute code from sub sub template
		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *PointerToGongStructField:

				fieldName := strings.ToLower(field.Name)
				assocStructName := strings.ToLower(field.GongStruct.Name)

				for subSubTemplate := range subSubTemplateCode {
					subSubCodes[subSubTemplate] += Replace6(subSubTemplateCode[subSubTemplate],
						"{{fieldName}}", fieldName,
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocStructName}}", assocStructName,
						"{{Structname}}", _struct.Name,
						"{{structname}}", structName)
				}
			}
		}

		// compute sub codes from sub templates
		for subTemplate := range subTemplateCode {
			subCodes[subTemplate] += Replace2(subTemplateCode[subTemplate],
				"{{Structname}}", _struct.Name,
				"{{structname}}", structName)

			// apply sub sub templates
			for subSubTemplate := range subSubTemplateCode {

				toReplace := "{{" + subSubTemplate + "}}"
				subCodes[subSubToSubMap[subSubTemplate]] = strings.ReplaceAll(
					subCodes[subSubToSubMap[subSubTemplate]],
					toReplace, subSubCodes[subSubTemplate])
			}
		}

		for subTemplate := range subTemplateCode {
			toReplace := "{{" + string(subTemplate) + "}}"
			code = strings.ReplaceAll(code, toReplace, subCodes[subTemplate])
		}

		code = strings.ReplaceAll(code, "{{PkgName}}", pkgName)
		code = strings.ReplaceAll(code, "{{TitlePkgName}}", strings.Title(pkgName))
		code = strings.ReplaceAll(code, "{{pkgname}}", strings.ToLower(pkgName))
		code = strings.ReplaceAll(code, "{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))

		defer file.Close()
		fmt.Fprint(file, code)
	}
}
