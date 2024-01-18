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

const NgClassDBTmpl = `// insertion point for imports{{` + string(rune(NgClassDBTsInsertionPerStructImports)) + `}}

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class {{Structname}}DB {

	static GONGSTRUCT_NAME = "{{Structname}}"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations{{` + string(rune(NgClassDBTsInsertionPerStructBasicFieldsDecl)) + `}}

	// insertion point for pointers and slices of pointers declarations{{` + string(rune(NgClassDBTsInsertionPerStructOtherDecls)) + `}}

	{{Structname}}PointersEncoding: {{Structname}}PointersEncoding = new {{Structname}}PointersEncoding
}

export class {{Structname}}PointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields{{` + string(rune(NgClassDBTsInsertionPerStructPointersEncoding)) + `}}
}
`

// Insertion points
// insertion points in the main template
type NgClassDBTsInsertionPoint int

const (
	NgClassDBTsInsertionPerStructImports NgClassDBTsInsertionPoint = iota
	NgClassDBTsInsertionPerStructBasicFieldsDecl
	NgClassDBTsInsertionPerStructOtherDecls
	NgClassDBTsInsertionPerStructPointersEncoding
	NgClassDBTsInsertionsNb
)

type NgClassDBSubTemplate int

const (
	NgClassDBTSBasicFieldImports NgClassDBSubTemplate = iota
	NgClassDBTSBasicFieldDecls

	NgClassDBTSTimeFieldDecls

	NgClassDBTSOtherDecls

	NgClassDBTSPointerToStructFieldsDecl

	NgClassDBTSPointerToStructFieldsEncodingDecl

	NgClassDBTSSliceOfPtrToStructFieldsDecl

	NgClassDBPointersEncodingTSSliceOfPtrToStructFieldsDecl

	NgClassDBTSOtherDeclsTimeDuration

	NgClassDBTSOtherDeclsEnumInt
)

var NgClassDBSubTemplateCode map[NgClassDBSubTemplate]string = map[NgClassDBSubTemplate]string{

	NgClassDBTSBasicFieldImports: `
import { {{AssocStructName}}DB } from './{{assocStructName}}-db'`,

	NgClassDBTSBasicFieldDecls: `
	{{FieldName}}: {{TypeInput}} = {{NullValue}}`,

	NgClassDBTSTimeFieldDecls: `
	{{FieldName}}: Date = new Date`,

	NgClassDBTSPointerToStructFieldsDecl: `
	{{FieldName}}?: {{TypeInput}}DB
`,

	NgClassDBTSPointerToStructFieldsEncodingDecl: `
	{{FieldName}}ID: NullInt64 = new NullInt64 // if pointer is null, {{FieldName}}.ID = 0
`,

	NgClassDBTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}: Array<{{TypeInput}}DB> = []`,

	NgClassDBPointersEncodingTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}: number[] = []`,

	NgClassDBTSOtherDeclsTimeDuration: `
	{{FieldName}}_string?: string`,

	NgClassDBTSOtherDeclsEnumInt: `
	{{FieldName}}_string?: string`,
}

// MultiCodeGeneratorNgTable generates the code for the
// Detail components
func MultiCodeGeneratorNgClassDB(modelPkg *models.ModelPkg) {

	PkgName := modelPkg.Name
	PkgGoPath := modelPkg.PkgPath
	MatTargetPath := modelPkg.NgDataLibrarySourceCodeDirectory

	// have alphabetical order generation
	structList := []*models.GongStruct{}
	for _, _struct := range modelPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		if !_struct.HasNameField() || _struct.IsIgnoredForFront {
			continue
		}

		TSinsertions := make(map[NgClassDBTsInsertionPoint]string)
		for insertion := NgClassDBTsInsertionPoint(0); insertion < NgClassDBTsInsertionsNb; insertion++ {
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

				TSinsertions[NgClassDBTsInsertionPerStructBasicFieldsDecl] += models.Replace3(
					NgClassDBSubTemplateCode[NgClassDBTSBasicFieldDecls],
					"{{FieldName}}", field.Name,
					"{{TypeInput}}", typeOfField,
					"{{NullValue}}", nullValue,
				)

				if field.DeclaredType == "time.Duration" {
					TSinsertions[NgClassDBTsInsertionPerStructOtherDecls] += models.Replace1(
						NgClassDBSubTemplateCode[NgClassDBTSOtherDeclsTimeDuration],
						"{{FieldName}}", field.Name)
				}

				if field.GongEnum != nil && field.GetBasicKind() == types.Int {

					TSinsertions[NgClassDBTsInsertionPerStructOtherDecls] += models.Replace1(
						NgClassDBSubTemplateCode[NgClassDBTSOtherDeclsEnumInt],
						"{{FieldName}}", field.Name)
				}
			case *models.GongTimeField:

				TSinsertions[NgClassDBTsInsertionPerStructBasicFieldsDecl] += models.Replace1(
					NgClassDBSubTemplateCode[NgClassDBTSTimeFieldDecls],
					"{{FieldName}}", field.Name)

			case *models.PointerToGongStructField:

				newImport := models.Replace2(NgClassDBSubTemplateCode[NgClassDBTSBasicFieldImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// for imports, duplicate imports or imports of the class itself are not allowed
				if !strings.Contains(TSinsertions[NgClassDBTsInsertionPerStructImports], newImport) &&
					_struct != field.GongStruct {
					TSinsertions[NgClassDBTsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassDBTsInsertionPerStructOtherDecls] +=
					models.Replace2(NgClassDBSubTemplateCode[NgClassDBTSPointerToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

				TSinsertions[NgClassDBTsInsertionPerStructPointersEncoding] +=
					models.Replace2(NgClassDBSubTemplateCode[NgClassDBTSPointerToStructFieldsEncodingDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

			case *models.SliceOfPointerToGongStructField:

				newImport := models.Replace2(NgClassDBSubTemplateCode[NgClassDBTSBasicFieldImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// for imports, duplicate imports are not allowed
				if !strings.Contains(TSinsertions[NgClassDBTsInsertionPerStructImports], newImport) &&
					_struct != field.GongStruct {
					TSinsertions[NgClassDBTsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassDBTsInsertionPerStructOtherDecls] +=
					models.Replace2(NgClassDBSubTemplateCode[NgClassDBTSSliceOfPtrToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

				TSinsertions[NgClassDBTsInsertionPerStructPointersEncoding] +=
					models.Replace2(NgClassDBSubTemplateCode[NgClassDBPointersEncodingTSSliceOfPtrToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

			}
		}

		// generate the typescript file
		codeDBTS := NgClassDBTmpl

		for insertion := NgClassDBTsInsertionPoint(0); insertion < NgClassDBTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeDBTS = strings.ReplaceAll(codeDBTS, toReplace, TSinsertions[insertion])
		}

		// final replacement
		{
			codeDBTS = models.Replace6(codeDBTS,
				"{{PkgName}}", PkgName,
				"{{TitlePkgName}}", strings.Title(PkgName),
				"{{pkgname}}", strings.ToLower(PkgName),
				"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			file, err := os.Create(filepath.Join(MatTargetPath, strings.ToLower(_struct.Name)+"-db.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeDBTS)
		}
	}
}
