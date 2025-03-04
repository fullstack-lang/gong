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

const NgClassTmpl = `// generated code - do not edit

import { {{Structname}}API } from './{{structname}}-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports{{` + string(rune(NgClassTsInsertionPerStructImports)) + `}}

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class {{Structname}} {

	static GONGSTRUCT_NAME = "{{Structname}}"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsDecl)) + `}}

	// insertion point for pointers and slices of pointers declarations{{` + string(rune(NgClassTsInsertionPerStructOtherDecls)) + `}}
}

export function Copy{{Structname}}To{{Structname}}API({{structname}}: {{Structname}}, {{structname}}API: {{Structname}}API) {

	{{structname}}API.CreatedAt = {{structname}}.CreatedAt
	{{structname}}API.DeletedAt = {{structname}}.DeletedAt
	{{structname}}API.ID = {{structname}}.ID

	// insertion point for basic fields copy operations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsCopyToAPI)) + `}}

	// insertion point for pointer fields encoding{{` + string(rune(NgClassTsInsertionPerStructPointerFieldsCopyToAPI)) + `}}

	// insertion point for slice of pointers fields encoding{{` + string(rune(NgClassTsInsertionPerStructSliceOfPointersFieldsCopyToAPI)) + `}}
}

// Copy{{Structname}}APITo{{Structname}} update basic, pointers and slice of pointers fields of {{structname}}
// from respectively the basic fields and encoded fields of pointers and slices of pointers of {{structname}}API
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function Copy{{Structname}}APITo{{Structname}}({{structname}}API: {{Structname}}API, {{structname}}: {{Structname}}, frontRepo: FrontRepo) {

	{{structname}}.CreatedAt = {{structname}}API.CreatedAt
	{{structname}}.DeletedAt = {{structname}}API.DeletedAt
	{{structname}}.ID = {{structname}}API.ID

	// insertion point for basic fields copy operations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsCopyFromAPI)) + `}}

	// insertion point for pointer fields encoding{{` + string(rune(NgClassTsInsertionPerStructPointerFieldsCopyFromAPI)) + `}}

	// insertion point for slice of pointers fields encoding{{` + string(rune(NgClassTsInsertionPerStructSliceOfPointersFieldsCopyFromAPI)) + `}}
}
`

// Insertion points
// insertion points in the main template
type NgClassTsInsertionPoint int

const (
	NgClassTsInsertionPerStructImports NgClassTsInsertionPoint = iota
	NgClassTsInsertionPerStructBasicFieldsDecl
	NgClassTsInsertionPerStructBasicFieldsCopyToAPI
	NgClassTsInsertionPerStructPointerFieldsCopyToAPI
	NgClassTsInsertionPerStructSliceOfPointersFieldsCopyToAPI
	NgClassTsInsertionPerStructBasicFieldsCopyFromAPI
	NgClassTsInsertionPerStructPointerFieldsCopyFromAPI
	NgClassTsInsertionPerStructSliceOfPointersFieldsCopyFromAPI
	NgClassTsInsertionPerStructOtherDecls
	NgClassTsInsertionsNb
)

type NgClassSubTemplate int

const (
	NgClassTSBasicFieldImports NgClassSubTemplate = iota

	NgClassTSBasicFieldDecls

	NgClassTSBasicFieldCopyToAPI

	NgClassTSPointerFieldCopyToAPI

	NgClassTSSliceOfPointersFieldCopyToAPI

	NgClassTSBasicFieldCopyFromAPI

	NgClassTSPointerFieldCopyFromAPI

	NgClassTSSliceOfPointersFieldCopyFromAPI

	NgClassTSTimeFieldDecls

	NgClassTSOtherDecls

	NgClassTSPointerToStructFieldsDecl

	NgClassTSSliceOfPtrToStructFieldsDecl

	NgClassTSOtherDeclsTimeDuration

	NgClassTSOtherDeclsEnumInt
)

var NgClassSubTemplateCode map[NgClassSubTemplate]string = map[NgClassSubTemplate]string{

	NgClassTSBasicFieldImports: `
import { {{AssocStructName}} } from './{{assocStructName}}'`,

	NgClassTSBasicFieldDecls: `
	{{FieldName}}: {{TypeInput}} = {{NullValue}}`,

	NgClassTSBasicFieldCopyToAPI: `
	{{structname}}API.{{FieldName}} = {{structname}}.{{FieldName}}`,

	NgClassTSPointerFieldCopyToAPI: `
	{{structname}}API.{{Structname}}PointersEncoding.{{FieldName}}ID.Valid = true
	if ({{structname}}.{{FieldName}} != undefined) {
		{{structname}}API.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64 = {{structname}}.{{FieldName}}.ID  
	} else {
		{{structname}}API.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64 = 0 		
	}
`,

	NgClassTSSliceOfPointersFieldCopyToAPI: `
	{{structname}}API.{{Structname}}PointersEncoding.{{FieldName}} = []
	for (let _{{assocStructName}} of {{structname}}.{{FieldName}}) {
		{{structname}}API.{{Structname}}PointersEncoding.{{FieldName}}.push(_{{assocStructName}}.ID)
	}
`,

	NgClassTSBasicFieldCopyFromAPI: `
	{{structname}}.{{FieldName}} = {{structname}}API.{{FieldName}}`,

	NgClassTSPointerFieldCopyFromAPI: `
	{{structname}}.{{FieldName}} = frontRepo.map_ID_{{AssocStructName}}.get({{structname}}API.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64)`,

	NgClassTSSliceOfPointersFieldCopyFromAPI: `
	if (!Array.isArray({{structname}}API.{{Structname}}PointersEncoding.{{FieldName}})) {
		console.error('Rects is not an array:', {{structname}}API.{{Structname}}PointersEncoding.{{FieldName}});
		return;
	}

	{{structname}}.{{FieldName}} = new Array<{{AssocStructName}}>()
	for (let _id of {{structname}}API.{{Structname}}PointersEncoding.{{FieldName}}) {
		let _{{assocStructName}} = frontRepo.map_ID_{{AssocStructName}}.get(_id)
		if (_{{assocStructName}} != undefined) {
			{{structname}}.{{FieldName}}.push(_{{assocStructName}}!)
		}
	}`,

	NgClassTSTimeFieldDecls: `
	{{FieldName}}: Date = new Date`,

	NgClassTSPointerToStructFieldsDecl: `
	{{FieldName}}?: {{TypeInput}}
`,

	NgClassTSSliceOfPtrToStructFieldsDecl: `
	{{FieldName}}: Array<{{TypeInput}}> = []`,

	NgClassTSOtherDeclsTimeDuration: `
	{{FieldName}}_string?: string`,

	NgClassTSOtherDeclsEnumInt: `
	{{FieldName}}_string?: string`,
}

// MultiCodeGeneratorNgTable generates the code for the
// Detail components
func MultiCodeGeneratorNgClass(modelPkg *models.ModelPkg) {

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

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsCopyToAPI] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSBasicFieldCopyToAPI],
					"{{FieldName}}", field.Name,
				)

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsCopyFromAPI] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSBasicFieldCopyFromAPI],
					"{{FieldName}}", field.Name,
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

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsCopyToAPI] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSBasicFieldCopyToAPI],
					"{{FieldName}}", field.Name,
				)

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsCopyFromAPI] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSBasicFieldCopyFromAPI],
					"{{FieldName}}", field.Name,
				)

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

				TSinsertions[NgClassTsInsertionPerStructPointerFieldsCopyToAPI] +=
					models.Replace2(NgClassSubTemplateCode[NgClassTSPointerFieldCopyToAPI],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

				TSinsertions[NgClassTsInsertionPerStructPointerFieldsCopyFromAPI] +=
					models.Replace3(NgClassSubTemplateCode[NgClassTSPointerFieldCopyFromAPI],
						"{{AssocStructName}}", field.GongStruct.Name,
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

				TSinsertions[NgClassTsInsertionPerStructSliceOfPointersFieldsCopyToAPI] +=
					models.Replace3(NgClassSubTemplateCode[NgClassTSSliceOfPointersFieldCopyToAPI],
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocStructName}}", strings.ToLower(field.GongStruct.Name),
						"{{FieldName}}", field.Name)

				TSinsertions[NgClassTsInsertionPerStructSliceOfPointersFieldsCopyFromAPI] +=
					models.Replace3(NgClassSubTemplateCode[NgClassTSSliceOfPointersFieldCopyFromAPI],
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocStructName}}", strings.ToLower(field.GongStruct.Name),
						"{{FieldName}}", field.Name)
			}
		}

		// generate the typescript file
		codeTS := NgClassTmpl

		for insertion := NgClassTsInsertionPoint(0); insertion < NgClassTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TSinsertions[insertion])
		}

		{
			codeTS = models.Replace6(codeTS,
				"{{PkgName}}", PkgName,
				"{{TitlePkgName}}", strings.Title(PkgName),
				"{{pkgname}}", strings.ToLower(PkgName),
				"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
				"{{Structname}}", _struct.Name,
				"{{structname}}", strings.ToLower(_struct.Name))

			file, err := os.Create(filepath.Join(MatTargetPath, strings.ToLower(_struct.Name)+".ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
	}
}
