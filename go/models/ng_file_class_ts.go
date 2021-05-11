package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const NgClassTmpl = `export class {{Structname}}API {
	// insertion point for basic fields declarations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsDecl)) + `}}
}
`
const NgClassDBTmpl = `// insertion point for imports{{` + string(rune(NgClassTsInsertionPerStructImports)) + `}}

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class {{Structname}}DB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsDecl)) + `}}

	// insertion point for other declarations{{` + string(rune(NgClassTsInsertionPerStructOtherDecls)) + `}}
}
`

// Insertion points
// insertion points in the main template
type NgClassTsInsertionPoint int

const (
	NgClassTsInsertionPerStructImports NgClassTsInsertionPoint = iota
	NgClassTsInsertionPerStructBasicFieldsDecl
	NgClassTsInsertionPerStructOtherDecls
	NgClassTsInsertionsNb
)

type NgClassSubTemplate int

const (
	NgClassTSBasicFieldImports NgClassSubTemplate = iota
	NgClassTSBasicFieldDecls

	NgClassTSTimeFieldDecls

	NgClassTSOtherDecls

	NgClassTSPointerToStructFieldsDecl

	NgClassTSSliceOfPtrToStructFieldsDecl

	NgClassTSSliceOfPtrToGongStructReverseID

	NgClassTSOtherDeclsTimeDuration
)

var NgClassSubTemplateCode map[NgClassSubTemplate]string = map[NgClassSubTemplate]string{

	NgClassTSBasicFieldImports: `
import { {{AssocStructName}}DB } from './{{assocStructName}}-db'`,

	NgClassTSBasicFieldDecls: `
	{{FieldName}}?: {{TypeInput}}`,

	NgClassTSTimeFieldDecls: `
	{{FieldName}}?: Date`,

	NgClassTSPointerToStructFieldsDecl: `
	{{FieldName}}?: {{TypeInput}}DB
	{{FieldName}}ID?: NullInt64
	{{FieldName}}Name?: string
`,

	NgClassTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}?: Array<{{TypeInput}}DB>`,

	NgClassTSSliceOfPtrToGongStructReverseID: `
	{{AssocStructName}}_{{FieldName}}DBID?: NullInt64
	{{AssocStructName}}_{{FieldName}}DBID_Index?: NullInt64 // store the index of the {{structname}} instance in {{AssocStructName}}.{{FieldName}}
	{{AssocStructName}}_{{FieldName}}_reverse?: {{AssocStructName}}DB
`,

	NgClassTSOtherDeclsTimeDuration: `
	{{FieldName}}_string?: string`,
}

// MultiCodeGeneratorNgTable generates the code for the
// Detail components
func MultiCodeGeneratorNgClass(
	mdlPkg *ModelPkg,
	PkgName,
	MatTargetPath,
	PkgGoPath string) {

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		// generate the typescript file
		codeTS := NgClassTmpl
		codeDBTS := NgClassDBTmpl

		TSinsertions := make(map[NgClassTsInsertionPoint]string)
		for insertion := NgClassTsInsertionPoint(0); insertion < NgClassTsInsertionsNb; insertion++ {
			TSinsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				gongBasicField := field.(*GongBasicField)

				// conversion form go type to ts type
				typeOfField := ""

				// conversion form go type to ts type
				switch gongBasicField.basicKind {
				case types.Int, types.Int64, types.Int32, types.Float64:
					typeOfField = "number"
				case types.String, types.Bool:
					typeOfField = "string"
				}

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsDecl] += Replace2(
					NgClassSubTemplateCode[NgClassTSBasicFieldDecls],
					"{{FieldName}}", gongBasicField.Name,
					"{{TypeInput}}", typeOfField)

				if gongBasicField.DeclaredType == "time.Duration" {
					TSinsertions[NgClassTsInsertionPerStructOtherDecls] += Replace1(
						NgClassSubTemplateCode[NgClassTSOtherDeclsTimeDuration],
						"{{FieldName}}", gongBasicField.Name)
				}
			case *GongTimeField:
				gongBasicField := field.(*GongTimeField)

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsDecl] += Replace1(
					NgClassSubTemplateCode[NgClassTSTimeFieldDecls],
					"{{FieldName}}", gongBasicField.Name)

			case *PointerToGongStructField:
				modelPointerToStructField := field.(*PointerToGongStructField)
				_ = modelPointerToStructField

				newImport := Replace2(NgClassSubTemplateCode[NgClassTSBasicFieldImports],
					"{{AssocStructName}}", modelPointerToStructField.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(modelPointerToStructField.GongStruct.Name))

				// for imports, duplicate imports or imports of the class itself are not allowed
				if !strings.Contains(TSinsertions[NgClassTsInsertionPerStructImports], newImport) &&
					_struct != modelPointerToStructField.GongStruct {
					TSinsertions[NgClassTsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassTsInsertionPerStructOtherDecls] +=
					Replace2(NgClassSubTemplateCode[NgClassTSPointerToStructFieldsDecl],
						"{{FieldName}}", modelPointerToStructField.Name,
						"{{TypeInput}}", modelPointerToStructField.GongStruct.Name)

			case *SliceOfPointerToGongStructField:
				fieldSliceOfPointerToModel := field.(*SliceOfPointerToGongStructField)
				_ = fieldSliceOfPointerToModel

				newImport := Replace2(NgClassSubTemplateCode[NgClassTSBasicFieldImports],
					"{{AssocStructName}}", fieldSliceOfPointerToModel.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(fieldSliceOfPointerToModel.GongStruct.Name))

				// for imports, duplicate imports are not allowed
				if !strings.Contains(TSinsertions[NgClassTsInsertionPerStructImports], newImport) &&
					_struct != fieldSliceOfPointerToModel.GongStruct {
					TSinsertions[NgClassTsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassTsInsertionPerStructOtherDecls] +=
					Replace2(NgClassSubTemplateCode[NgClassTSSliceOfPtrToStructFieldsDecl],
						"{{FieldName}}", fieldSliceOfPointerToModel.Name,
						"{{TypeInput}}", fieldSliceOfPointerToModel.GongStruct.Name)

			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field.(type) {
				case *SliceOfPointerToGongStructField:
					fieldSliceOfPointerToModel := field.(*SliceOfPointerToGongStructField)

					if fieldSliceOfPointerToModel.GongStruct == _struct {

						newImport := Replace2(NgClassSubTemplateCode[NgClassTSBasicFieldImports],
							"{{AssocStructName}}", __struct.Name,
							"{{assocStructName}}", strings.ToLower(__struct.Name))

						// for imports, duplicate imports are not allowed
						if !strings.Contains(TSinsertions[NgClassTsInsertionPerStructImports], newImport) &&
							__struct != fieldSliceOfPointerToModel.GongStruct {
							TSinsertions[NgClassTsInsertionPerStructImports] += newImport
						}

						TSinsertions[NgClassTsInsertionPerStructOtherDecls] +=
							Replace2(NgClassSubTemplateCode[NgClassTSSliceOfPtrToGongStructReverseID],
								"{{FieldName}}", fieldSliceOfPointerToModel.Name,
								"{{AssocStructName}}", __struct.Name)
					}
				}
			}
		}

		for insertion := NgClassTsInsertionPoint(0); insertion < NgClassTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TSinsertions[insertion])
			codeDBTS = strings.ReplaceAll(codeDBTS, toReplace, TSinsertions[insertion])
		}

		// final replacement
		codeTS = Replace6(codeTS,
			"{{PkgName}}", PkgName,
			"{{TitlePkgName}}", strings.Title(PkgName),
			"{{pkgname}}", strings.ToLower(PkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
		codeDBTS = Replace6(codeDBTS,
			"{{PkgName}}", PkgName,
			"{{TitlePkgName}}", strings.Title(PkgName),
			"{{pkgname}}", strings.ToLower(PkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
		{
			file, err := os.Create(filepath.Join(MatTargetPath, strings.ToLower(_struct.Name)+"-api.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
		{
			file, err := os.Create(filepath.Join(MatTargetPath, strings.ToLower(_struct.Name)+"-db.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeDBTS)
		}
	}
}
