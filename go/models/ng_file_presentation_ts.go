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
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

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

	// insertion point for declarations{{` + string(rune(NgPresentationTsInsertionPerStructDeclarations)) + `}}

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	{{structname}}: {{Structname}}DB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private {{structname}}Service: {{Structname}}Service,
		private frontRepoService: FrontRepoService,
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
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.{{structname}} = this.frontRepo.{{Structname}}s.get(id)

				// insertion point for recovery of durations{{` + string(rune(NgPresentationTsInsertionPerStructRecoveries)) + `}}
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				presentation: ["{{PkgPathRootWithoutSlashes}}-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				editor: ["{{PkgPathRootWithoutSlashes}}-" + "{{structname}}-detail", ID]
			}
		}]);
	}
}
`

// Insertion points
type NgPresentationTsInsertionPoint int

const (
	NgPresentationTsInsertionPerStructDeclarations NgPresentationTsInsertionPoint = iota
	NgPresentationTsInsertionPerStructRecoveries
	NgPresentationTsInsertionsNb
)

type NgPresentationSubTemplate int

const (
	NgPresentationTSTimeDurationDeclarations NgPresentationSubTemplate = iota
	NgPresentationTSTimeDurationRecoveries
)

var NgPresentationSubTemplateCode map[NgPresentationSubTemplate]string = map[NgPresentationSubTemplate]string{

	NgPresentationTSTimeDurationDeclarations: `
	// fields from {{FieldName}}
	{{FieldName}}_Hours: number
	{{FieldName}}_Minutes: number
	{{FieldName}}_Seconds: number`,
	NgPresentationTSTimeDurationRecoveries: `
				// computation of Hours, Minutes, Seconds for {{FieldName}}
				this.{{FieldName}}_Hours = Math.floor(this.{{structname}}.{{FieldName}} / (3600 * 1000 * 1000 * 1000))
				this.{{FieldName}}_Minutes = Math.floor(this.{{structname}}.{{FieldName}} % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.{{FieldName}}_Seconds = this.{{structname}}.{{FieldName}} % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)`,
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

		TSinsertions := make(map[NgPresentationTsInsertionPoint]string)
		for insertion := NgPresentationTsInsertionPoint(0); insertion < NgPresentationTsInsertionsNb; insertion++ {
			TSinsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				gongBasicField := field.(*GongBasicField)

				if gongBasicField.DeclaredType == "time.Duration" {
					TSinsertions[NgPresentationTsInsertionPerStructDeclarations] += Replace1(
						NgPresentationSubTemplateCode[NgPresentationTSTimeDurationDeclarations],
						"{{FieldName}}", gongBasicField.Name)

					TSinsertions[NgPresentationTsInsertionPerStructRecoveries] += Replace1(
						NgPresentationSubTemplateCode[NgPresentationTSTimeDurationRecoveries],
						"{{FieldName}}", gongBasicField.Name)
				}
			}
		}

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

		//
		// generate `per field` HTML code
		// cumulative sub template
		//
		subCodesHTML := ""
		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				gongBasicField := field.(*GongBasicField)

				// bais field (enum)
				if gongBasicField.GongEnum != nil {
					subCodesHTML += Replace2(NgPresentationEnumHtmlSubTemplateCode[string(NgPresentationHtmlEnum)],
						"{{FieldName}}", gongBasicField.Name,
						"{{EnumName}}", gongBasicField.GongEnum.Name)

				} else // basic field (not enum)
				{
					if gongBasicField.basicKind == types.Bool {
						subCodesHTML += Replace1(
							NgPresentationBoolHtmlSubTemplateCode[string(NgPresentationHtmlBool)],
							"{{FieldName}}", gongBasicField.Name)
					} else {

						if gongBasicField.DeclaredType != "time.Duration" {

							// conversion form go type to ts type
							TypeInput := ""
							switch gongBasicField.basicKind {
							case types.Int, types.Float64:
								TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\" "
							case types.String:
								TypeInput = "name=\"\" [ngModelOptions]=\"{standalone: true}\" 	"
							}
							subCodesHTML += Replace2(
								NgPresentationBasicFieldHtmlSubTemplateCode[string(NgPresentationHtmlBasicField)],
								"{{FieldName}}", gongBasicField.Name,
								"{{TypeInput}}", TypeInput)
						} else {
							subCodesHTML += Replace1(
								NgPresentationBasicFieldTimeDurationHtmlSubTemplateCode[string(NgPresentationHtmlBasicFieldTimeDuration)],
								"{{FieldName}}", gongBasicField.Name)
						}

					}
				}

			case *GongTimeField:
				gongTimeField := field.(*GongTimeField)
				subCodesHTML +=
					Replace1(NgPresentationTimeFieldHtmlSubTemplateCode[string(NgPresentationHtmlTimeField)],
						"{{FieldName}}", gongTimeField.Name)

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

		for insertion := NgPresentationTsInsertionPoint(0); insertion < NgPresentationTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TSinsertions[insertion])
		}

		// final replacement
		pkgPathRootWithoutSlashes := strings.ReplaceAll(pkgGoPath, "/models", "")
		pkgPathRootWithoutSlashes = strings.ReplaceAll(pkgPathRootWithoutSlashes, "/", ".")

		codeTS = Replace7(codeTS,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name),
			"{{PkgPathRootWithoutSlashes}}", pkgPathRootWithoutSlashes)
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
