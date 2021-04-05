package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "embed"
)

//go:embed ng_file_presentation.css
var NgFilePresentationCssTmpl string

const NgPresentationTemplateTS = `import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface {{structname}}DummyElement {
}

const ELEMENT_DATA: {{structname}}DummyElement[] = [
];

@Component({
	selector: 'app-{{structname}}-presentation',
	templateUrl: './{{structname}}-presentation.component.html',
  styleUrls: ['./{{structname}}-presentation.component.css'],
})
export class {{Structname}}PresentationComponent implements OnInit {

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	{{structname}}: {{Structname}}DB;
 
	constructor(
		private {{structname}}Service: {{Structname}}Service,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.get{{Structname}}();

		// observable for changes in 
		this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.get{{Structname}}()
				}
			}
		)
	}

	get{{Structname}}(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.{{structname}}Service.get{{Structname}}(id)
			.subscribe(
				{{structname}} => {
					this.{{structname}} = {{structname}}
				}
			);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				presentation: [structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				editor: ["{{structname}}-detail", ID]
			}
		}]);
	}
}
`

//
// Sub Templates
//
type NgPresentationEnumSubTemplate string

const (
	NgPresentationEnumImports         NgPresentationEnumSubTemplate = "NgPresentationEnumImports"
	NgPresentationEnumSelectListDecls NgPresentationEnumSubTemplate = "NgPresentationSelectListDecls"
	NgPresentationEnumSelectListInits NgPresentationEnumSubTemplate = "NgPresentationSelectListInits"
)

var NgPresentationEnumSubTemplateCode map[string]string = map[string]string{
	string(NgPresentationEnumImports): `
import { {{EnumName}}Select, {{EnumName}}List } from '../{{EnumName}}'`,

	string(NgPresentationEnumSelectListDecls): `
	{{EnumName}}List: {{EnumName}}Select[]`,

	string(NgPresentationEnumSelectListInits): `
		this.{{EnumName}}List = {{EnumName}}List`,
}

type NgPresentationPointerToStructSubTemplate string

const (
	NgPresentationPointerToStructSavingTS NgPresentationPointerToStructSubTemplate = "NgPresentationPointerToStructSavingTS"
)

var NgPresentationPointerToStructSubTemplateCode map[string]string = map[string]string{

	string(NgPresentationPointerToStructSavingTS): `
		if (this.{{structname}}.{{FieldName}} != undefined) {
			this.{{structname}}.{{FieldName}}ID = this.{{structname}}.{{FieldName}}.ID
			this.{{structname}}.{{FieldName}}Name = this.{{structname}}.{{FieldName}}.Name
		}`,
}

// MultiCodeGeneratorNgPresentation parses mdlPkg and generates the code for the
// Presentation components
func MultiCodeGeneratorNgPresentation(
	mdlPkg *ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	for _, _struct := range mdlPkg.GongStructs {

		if !_struct.HasNameField() {
			continue
		}

		// create the component directory
		dirPath := filepath.Join(matTargetPath, strings.ToLower(_struct.Name)+"-presentation")
		errd := os.Mkdir(dirPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + dirPath)
		}

		// generate the css file
		VerySimpleCodeGenerator(mdlPkg,
			pkgName,
			pkgGoPath,
			filepath.Join(dirPath, strings.ToLower(_struct.Name)+"-presentation.component.css"),
			NgFilePresentationCssTmpl,
		)

		// generate the typescript file
		codeTS := NgPresentationTemplateTS
		codeHTML := NgPresentationTemplateHTML

		//
		// for typescript, enums
		// generate list of fields with type enum in order to generate EnumList stuff
		//

		// parse all fields. When a field is of type BasicField with a EnumType
		// then retain the enum type
		modelEnums := make(map[string]*GongEnum)
		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				modelBasicField := field.(*GongBasicField)

				if modelBasicField.GongEnum != nil {
					modelEnums[modelBasicField.GongEnum.Name] = modelBasicField.GongEnum
				}
			}
		}

		// sort ModelEnums
		sorted := make([]string, 0)
		for k := range modelEnums {
			sorted = append(sorted, k)
		}
		sort.Strings(sorted)

		// cumulative sub template
		subCodesEnumTS := make(map[string]string)
		for subTemplate := range NgPresentationEnumSubTemplateCode {
			subCodesEnumTS[subTemplate] = ""
		}

		// generate sub codes from enum lists
		for _, _enum := range sorted {
			for subTemplate := range NgPresentationEnumSubTemplateCode {
				subCodesEnumTS[subTemplate] += Replace1(NgPresentationEnumSubTemplateCode[subTemplate],
					"{{EnumName}}", _enum)
			}
		}

		// insert sub codes in main template for enum lists
		for subTemplate := range NgPresentationEnumSubTemplateCode {
			toReplace := "{{" + string(subTemplate) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, subCodesEnumTS[subTemplate])
		}

		//
		// for typescript, generate pointer to struct code
		//
		subCodesPointerToStructTS := make(map[string]string)
		for subTemplate := range NgPresentationPointerToStructSubTemplateCode {
			subCodesPointerToStructTS[subTemplate] = ""
		}
		for _, field := range _struct.Fields {
			switch field.(type) {
			case *PointerToGongStructField:
				modelPointerToStruct := field.(*PointerToGongStructField)
				for subTemplate := range NgPresentationPointerToStructSubTemplateCode {
					subCodesPointerToStructTS[subTemplate] += Replace1(
						NgPresentationPointerToStructSubTemplateCode[subTemplate],
						"{{FieldName}}", modelPointerToStruct.Name)
				}
			}
		}

		// insert sub codes in main template for pointer to bools
		for subTemplate := range NgPresentationPointerToStructSubTemplateCode {
			toReplace := "{{" + string(subTemplate) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, subCodesPointerToStructTS[subTemplate])
		}

		//
		// generate `per field` HTML code
		// cumulative sub template
		//
		subCodesHTML := ""
		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				modelBasicField := field.(*GongBasicField)

				// bais field (enum)
				if modelBasicField.GongEnum != nil {
					subCodesHTML += Replace2(NgPresentationEnumHtmlSubTemplateCode[string(NgPresentationHtmlEnum)],
						"{{FieldName}}", modelBasicField.Name,
						"{{EnumName}}", modelBasicField.GongEnum.Name)

				} else // basic field (not enum)
				{
					if modelBasicField.basicKind == types.Bool {
						subCodesHTML += Replace1(
							NgPresentationBoolHtmlSubTemplateCode[string(NgPresentationHtmlBool)],
							"{{FieldName}}", modelBasicField.Name)
					} else {
						// conversion form go type to ts type
						TypeInput := ""
						switch modelBasicField.basicKind {
						case types.Int, types.Float64:
							TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\" "
						case types.String:
							TypeInput = "name=\"\" [ngModelOptions]=\"{standalone: true}\" 	"
						}
						subCodesHTML += Replace2(
							NgPresentationBasicFieldHtmlSubTemplateCode[string(NgPresentationHtmlBasicField)],
							"{{FieldName}}", modelBasicField.Name,
							"{{TypeInput}}", TypeInput)
					}
				}

			case *PointerToGongStructField:
				modelPointerToStructField := field.(*PointerToGongStructField)
				subCodesHTML += Replace3(
					NgPresentationPointerToStructHtmlSubTemplateCode[string(NgPresentationPointerToStructHtmlFormField)],
					"{{FieldName}}", modelPointerToStructField.Name,
					"{{AssocStructName}}", modelPointerToStructField.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(modelPointerToStructField.GongStruct.Name))
			}
		}

		{
			toReplace := "{{" + string(NgPresentationHtmlField) + "}}"
			codeHTML = strings.ReplaceAll(codeHTML, toReplace, subCodesHTML)
		}

		// final replacement
		codeTS = Replace6(codeTS,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
		codeHTML = Replace6(codeHTML,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
		{
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"-presentation.component.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
		{
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"-presentation.component.html"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeHTML)
		}

	}
}
