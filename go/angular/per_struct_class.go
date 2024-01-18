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

import { {{Structname}}DB } from './{{structname}}-db'
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

export function Copy{{Structname}}To{{Structname}}DB({{structname}}: {{Structname}}, {{structname}}DB: {{Structname}}DB) {

	// insertion point for basic fields copy operations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsCopyToDB)) + `}}

	// insertion point for pointer fields encoding{{` + string(rune(NgClassTsInsertionPerStructPointerFieldsCopyToDB)) + `}}

	// insertion point for slice of pointers fields encoding{{` + string(rune(NgClassTsInsertionPerStructSliceOfPointersFieldsCopyToDB)) + `}}
}

export function Copy{{Structname}}DBTo{{Structname}}({{structname}}DB: {{Structname}}DB, {{structname}}: {{Structname}}, frontRepo: FrontRepo) {

	// insertion point for basic fields copy operations{{` + string(rune(NgClassTsInsertionPerStructBasicFieldsCopyFromDB)) + `}}

	// insertion point for pointer fields encoding{{` + string(rune(NgClassTsInsertionPerStructPointerFieldsCopyFromDB)) + `}}

	// insertion point for slice of pointers fields encoding{{` + string(rune(NgClassTsInsertionPerStructSliceOfPointersFieldsCopyFromDB)) + `}}
}
`

// Insertion points
// insertion points in the main template
type NgClassTsInsertionPoint int

const (
	NgClassTsInsertionPerStructImports NgClassTsInsertionPoint = iota
	NgClassTsInsertionPerStructBasicFieldsDecl
	NgClassTsInsertionPerStructBasicFieldsCopyToDB
	NgClassTsInsertionPerStructPointerFieldsCopyToDB
	NgClassTsInsertionPerStructSliceOfPointersFieldsCopyToDB
	NgClassTsInsertionPerStructBasicFieldsCopyFromDB
	NgClassTsInsertionPerStructPointerFieldsCopyFromDB
	NgClassTsInsertionPerStructSliceOfPointersFieldsCopyFromDB
	NgClassTsInsertionPerStructOtherDecls
	NgClassTsInsertionsNb
)

type NgClassSubTemplate int

const (
	NgClassTSBasicFieldImports NgClassSubTemplate = iota

	NgClassTSBasicFieldDecls

	NgClassTSBasicFieldCopyToDB

	NgClassTSPointerFieldCopyToDB

	NgClassTSSliceOfPointersFieldCopyToDB

	NgClassTSBasicFieldCopyFromDB

	NgClassTSPointerFieldCopyFromDB

	NgClassTSSliceOfPointersFieldCopyFromDB

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

	NgClassTSBasicFieldCopyToDB: `
	{{structname}}DB.{{FieldName}} = {{structname}}.{{FieldName}}`,

	NgClassTSPointerFieldCopyToDB: `
    {{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}}ID.Valid = true
	if ({{structname}}.{{FieldName}} != undefined) {
		{{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64 = {{structname}}.{{FieldName}}.ID  
    } else {
		{{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64 = 0 		
	}
`,

	NgClassTSSliceOfPointersFieldCopyToDB: `
	{{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}} = []
    for (let _{{assocStructName}} of {{structname}}.{{FieldName}}) {
		{{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}}.push(_{{assocStructName}}.ID)
    }
	`,

	NgClassTSBasicFieldCopyFromDB: `
	{{structname}}.{{FieldName}} = {{structname}}DB.{{FieldName}}`,

	NgClassTSPointerFieldCopyFromDB: `
	{{structname}}.{{FieldName}} = frontRepo.{{AssocStructName}}s.get({{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64)`,

	NgClassTSSliceOfPointersFieldCopyFromDB: `
	{{structname}}.{{FieldName}} = new Array<{{AssocStructName}}>()
	for (let _id of {{structname}}DB.{{Structname}}PointersEncoding.{{FieldName}}) {
	  let _{{assocStructName}} = frontRepo.{{AssocStructName}}s.get(_id)
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

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsCopyToDB] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSBasicFieldCopyToDB],
					"{{FieldName}}", field.Name,
				)

				TSinsertions[NgClassTsInsertionPerStructBasicFieldsCopyFromDB] += models.Replace1(
					NgClassSubTemplateCode[NgClassTSBasicFieldCopyFromDB],
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

				TSinsertions[NgClassTsInsertionPerStructPointerFieldsCopyToDB] +=
					models.Replace2(NgClassSubTemplateCode[NgClassTSPointerFieldCopyToDB],
						"{{FieldName}}", field.Name,
						"{{TypeInput}}", field.GongStruct.Name)

				TSinsertions[NgClassTsInsertionPerStructPointerFieldsCopyFromDB] +=
					models.Replace3(NgClassSubTemplateCode[NgClassTSPointerFieldCopyFromDB],
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

				TSinsertions[NgClassTsInsertionPerStructSliceOfPointersFieldsCopyToDB] +=
					models.Replace3(NgClassSubTemplateCode[NgClassTSSliceOfPointersFieldCopyToDB],
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocStructName}}", strings.ToLower(field.GongStruct.Name),
						"{{FieldName}}", field.Name)

				TSinsertions[NgClassTsInsertionPerStructSliceOfPointersFieldsCopyFromDB] +=
					models.Replace3(NgClassSubTemplateCode[NgClassTSSliceOfPointersFieldCopyFromDB],
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
