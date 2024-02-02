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

const NgClassAPITmpl = `// insertion point for imports{{` + string(rune(NgClassAPITsInsertionPerStructImports)) + `}}

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class {{Structname}}API {

	static GONGSTRUCT_NAME = "{{Structname}}"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations{{` + string(rune(NgClassAPITsInsertionPerStructBasicFieldsDecl)) + `}}

	// insertion point for other decls{{` + string(rune(NgClassAPITsInsertionPerStructOtherDecls)) + `}}

	{{Structname}}PointersEncoding: {{Structname}}PointersEncoding = new {{Structname}}PointersEncoding
}

export class {{Structname}}PointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields{{` + string(rune(NgClassAPITsInsertionPerStructPointersEncoding)) + `}}
}
`

// Insertion points
// insertion points in the main template
type NgClassAPITsInsertionPoint int

const (
	NgClassAPITsInsertionPerStructImports NgClassAPITsInsertionPoint = iota
	NgClassAPITsInsertionPerStructBasicFieldsDecl
	NgClassAPITsInsertionPerStructOtherDecls
	NgClassAPITsInsertionPerStructPointersEncoding
	NgClassAPITsInsertionsNb
)

type NgClassAPISubTemplate int

const (
	NgClassAPITSBasicFieldImports NgClassAPISubTemplate = iota
	NgClassAPITSBasicFieldDecls

	NgClassAPITSTimeFieldDecls

	NgClassAPITSOtherDecls

	NgClassAPITSPointerToStructFieldsEncodingDecl

	NgClassAPIPointersEncodingTSSliceOfPtrToStructFieldsDecl

	NgClassAPITSOtherDeclsTimeDuration

	NgClassAPITSOtherDeclsEnumInt
)

var NgClassAPISubTemplateCode map[NgClassAPISubTemplate]string = map[NgClassAPISubTemplate]string{

	NgClassAPITSBasicFieldImports: `
import { {{AssocStructName}}API } from './{{assocStructName}}-api'`,

	NgClassAPITSBasicFieldDecls: `
	{{FieldName}}: {{TypeInput}} = {{NullValue}}`,

	NgClassAPITSTimeFieldDecls: `
	{{FieldName}}: Date = new Date`,

	NgClassAPITSPointerToStructFieldsEncodingDecl: `
	{{FieldName}}ID: NullInt64 = new NullInt64 // if pointer is null, {{FieldName}}.ID = 0
`,

	NgClassAPIPointersEncodingTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}: number[] = []`,

	NgClassAPITSOtherDeclsTimeDuration: `
	{{FieldName}}_string?: string`,

	NgClassAPITSOtherDeclsEnumInt: `
	{{FieldName}}_string?: string`,
}

// MultiCodeGeneratorNgTable generates the code for the
// Detail components
func MultiCodeGeneratorNgClassAPI(modelPkg *models.ModelPkg) {

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

		TSinsertions := make(map[NgClassAPITsInsertionPoint]string)
		for insertion := NgClassAPITsInsertionPoint(0); insertion < NgClassAPITsInsertionsNb; insertion++ {
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

				TSinsertions[NgClassAPITsInsertionPerStructBasicFieldsDecl] += models.Replace3(
					NgClassAPISubTemplateCode[NgClassAPITSBasicFieldDecls],
					"{{FieldName}}", field.Name,
					"{{TypeInput}}", typeOfField,
					"{{NullValue}}", nullValue,
				)

				if field.DeclaredType == "time.Duration" {
					TSinsertions[NgClassAPITsInsertionPerStructOtherDecls] += models.Replace1(
						NgClassAPISubTemplateCode[NgClassAPITSOtherDeclsTimeDuration],
						"{{FieldName}}", field.Name)
				}

				if field.GongEnum != nil && field.GetBasicKind() == types.Int {

					TSinsertions[NgClassAPITsInsertionPerStructOtherDecls] += models.Replace1(
						NgClassAPISubTemplateCode[NgClassAPITSOtherDeclsEnumInt],
						"{{FieldName}}", field.Name)
				}
			case *models.GongTimeField:

				TSinsertions[NgClassAPITsInsertionPerStructBasicFieldsDecl] += models.Replace1(
					NgClassAPISubTemplateCode[NgClassAPITSTimeFieldDecls],
					"{{FieldName}}", field.Name)

			case *models.PointerToGongStructField:

				newImport := models.Replace2(NgClassAPISubTemplateCode[NgClassAPITSBasicFieldImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// for imports, duplicate imports or imports of the class itself are not allowed
				if !strings.Contains(TSinsertions[NgClassAPITsInsertionPerStructImports], newImport) &&
					_struct != field.GongStruct {
					TSinsertions[NgClassAPITsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassAPITsInsertionPerStructPointersEncoding] +=
					models.Replace2(NgClassAPISubTemplateCode[NgClassAPITSPointerToStructFieldsEncodingDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

			case *models.SliceOfPointerToGongStructField:

				newImport := models.Replace2(NgClassAPISubTemplateCode[NgClassAPITSBasicFieldImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// for imports, duplicate imports are not allowed
				if !strings.Contains(TSinsertions[NgClassAPITsInsertionPerStructImports], newImport) &&
					_struct != field.GongStruct {
					TSinsertions[NgClassAPITsInsertionPerStructImports] += newImport
				}

				TSinsertions[NgClassAPITsInsertionPerStructPointersEncoding] +=
					models.Replace2(NgClassAPISubTemplateCode[NgClassAPIPointersEncodingTSSliceOfPtrToStructFieldsDecl],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

			}
		}

		// generate the typescript file
		codeAPITS := NgClassAPITmpl

		for insertion := NgClassAPITsInsertionPoint(0); insertion < NgClassAPITsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeAPITS = strings.ReplaceAll(codeAPITS, toReplace, TSinsertions[insertion])
		}

		// final replacement
		{
			codeAPITS = models.Replace6(codeAPITS,
				"{{PkgName}}", PkgName,
				"{{TitlePkgName}}", strings.Title(PkgName),
				"{{pkgname}}", strings.ToLower(PkgName),
				"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			file, err := os.Create(filepath.Join(MatTargetPath, strings.ToLower(_struct.Name)+"-api.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeAPITS)
		}
	}
}
