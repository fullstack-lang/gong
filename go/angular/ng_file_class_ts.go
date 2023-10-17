package angular

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

const NgClassTmpl = `export class {{Structname}}API {
	// insertion point for basic fields declarations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsDecl)) + `}}
}
`
const NgClassDBTmpl = `// insertion point for imports{{` + string(rune(NgClassTsInsertionPerStructImports)) + `}}

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class {{Structname}}DB {

	static GONGSTRUCT_NAME = "{{Structname}}"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsDecl)) + `}}

	// insertion point for other declarations{{` + string(rune(NgClassTsInsertionPerStructOtherDecls)) + `}}

	{{Structname}}PointersEncoding: {{Structname}}PointersEncoding = new {{Structname}}PointersEncoding
}

export class {{Structname}}PointersEncoding {
	// insertion point for other declarations{{` + string(rune(NgClassTsInsertionPerStructPointersEncoding)) + `}}
}
`

// Insertion points
// insertion points in the main template
type NgClassTsInsertionPoint int

const (
	NgClassTsInsertionPerStructImports NgClassTsInsertionPoint = iota
	NgClassTsInsertionPerStructBasicFieldsDecl
	NgClassTsInsertionPerStructOtherDecls
	NgClassTsInsertionPerStructPointersEncoding
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

	NgClassPointersEncodingTSSliceOfPtrToStructFieldsDecl

	NgClassTSSliceOfPtrToGongStructReverseID

	NgClassTSOtherDeclsTimeDuration

	NgClassTSOtherDeclsEnumInt
)

var NgClassSubTemplateCode map[NgClassSubTemplate]string = map[NgClassSubTemplate]string{

	NgClassTSBasicFieldImports: `
import { {{AssocStructName}}DB } from './{{assocStructName}}-db'`,

	NgClassTSBasicFieldDecls: `
	{{FieldName}}: {{TypeInput}} = {{NullValue}}`,

	NgClassTSTimeFieldDecls: `
	{{FieldName}}: Date = new Date`,

	NgClassTSPointerToStructFieldsDecl: `
	{{FieldName}}?: {{TypeInput}}DB
	{{FieldName}}ID: NullInt64 = new NullInt64 // if pointer is null, {{FieldName}}.ID = 0
`,

	NgClassTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}: Array<{{TypeInput}}DB> = []`,

	NgClassPointersEncodingTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}: number[] = []`,

	NgClassTSSliceOfPtrToGongStructReverseID: `
	// reverse pointers encoding (to be removed)
	{{AssocStructName}}_{{FieldName}}DBID: NullInt64 = new NullInt64
	{{AssocStructName}}_{{FieldName}}DBID_Index: NullInt64  = new NullInt64 // store the index of the {{structname}} instance in {{AssocStructName}}.{{FieldName}}
	{{AssocStructName}}_{{FieldName}}_reverse?: {{AssocStructName}}DB 
`,

	NgClassTSOtherDeclsTimeDuration: `
	{{FieldName}}_string?: string`,

	NgClassTSOtherDeclsEnumInt: `
	{{FieldName}}_string?: string`,
}

// MultiCodeGeneratorNgTable generates the code for the
// Detail components
func MultiCodeGeneratorNgClass(
	mdlPkg *models.ModelPkg,
	PkgName,
	MatTargetPath,
	PkgGoPath string) {

	// have alphabetical order generation
	structList := []*models.GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		if !_struct.HasNameField() || _struct.IsIgnoredForFront {
			continue
		}

		// generate the typescript file
		codeDBTS := NgClassDBTmpl

		TSinsertions := make(map[NgClassTsInsertionPoint]string)
		for insertion := NgClassTsInsertionPoint(0); insertion < NgClassTsInsertionsNb; insertion++ {
			TSinsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *models.GongBasicField:

				// conversion form go type to ts type
				typeOfField := ""
				nullValue := ""

				// conversion form go type to ts type
				switch field.GetBasicKind() {
				case types.Int, types.Int64, types.Int32, types.Float64:
					typeOfField = "number"
					nullValue = "0"
				case types.String:
					typeOfField = "string"
					nullValue = "\"\""
				case types.Bool:
					typeOfField = "boolean"
					nullValue = "false"
				}

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsDecl] += models.Replace3(
					NgClassSubTemplateCode[NgClassTSBasicFieldDecls],
					"{{FieldName}}", field.Name,
					"{{TypeInput}}", typeOfField,
					"{{NullValue}}", nullValue,
				)

				if field.DeclaredType == "time.Duration" {
					TSinsertions[NgClassTsInsertionPerStructOtherDecls] += models.Replace1(
						NgClassSubTemplateCode[NgClassTSOtherDeclsTimeDuration],
						"{{FieldName}}", field.Name)
				}

				if field.GongEnum != nil && field.GetBasicKind() == types.Int {

					TSinsertions[NgClassTsInsertionPerStructOtherDecls] += models.Replace1(
						NgClassSubTemplateCode[NgClassTSOtherDeclsEnumInt],
						"{{FieldName}}", field.Name)
				}
			case *models.GongTimeField:

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsDecl] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSTimeFieldDecls],
					"{{FieldName}}", field.Name)

			case *models.PointerToGongStructField:

				newImport := models.Replace2(NgClassSubTemplateCode[NgClassTSBasicFieldImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// for imports, duplicate imports or imports of the class itself are not allowed
				if !strings.Contains(TSinsertions[NgClassTsInsertionPerStructImports], newImport) &&
					_struct != field.GongStruct {
					TSinsertions[NgClassTsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassTsInsertionPerStructOtherDecls] +=
					models.Replace2(NgClassSubTemplateCode[NgClassTSPointerToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

			case *models.SliceOfPointerToGongStructField:

				newImport := models.Replace2(NgClassSubTemplateCode[NgClassTSBasicFieldImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// for imports, duplicate imports are not allowed
				if !strings.Contains(TSinsertions[NgClassTsInsertionPerStructImports], newImport) &&
					_struct != field.GongStruct {
					TSinsertions[NgClassTsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassTsInsertionPerStructOtherDecls] +=
					models.Replace2(NgClassSubTemplateCode[NgClassTSSliceOfPtrToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

				TSinsertions[NgClassTsInsertionPerStructPointersEncoding] +=
					models.Replace2(NgClassSubTemplateCode[NgClassPointersEncodingTSSliceOfPtrToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field := field.(type) {
				case *models.SliceOfPointerToGongStructField:

					if field.GongStruct == _struct {

						newImport := models.Replace2(NgClassSubTemplateCode[NgClassTSBasicFieldImports],
							"{{AssocStructName}}", __struct.Name,
							"{{assocStructName}}", strings.ToLower(__struct.Name))

						// for imports, duplicate imports are not allowed
						if !strings.Contains(TSinsertions[NgClassTsInsertionPerStructImports], newImport) &&
							__struct != field.GongStruct {
							TSinsertions[NgClassTsInsertionPerStructImports] += newImport
						}

						TSinsertions[NgClassTsInsertionPerStructOtherDecls] +=
							models.Replace2(NgClassSubTemplateCode[NgClassTSSliceOfPtrToGongStructReverseID],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)
					}
				}
			}
		}

		for insertion := NgClassTsInsertionPoint(0); insertion < NgClassTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeDBTS = strings.ReplaceAll(codeDBTS, toReplace, TSinsertions[insertion])
		}

		// final replacement
		codeDBTS = models.Replace6(codeDBTS,
			"{{PkgName}}", PkgName,
			"{{TitlePkgName}}", strings.Title(PkgName),
			"{{pkgname}}", strings.ToLower(PkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
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
