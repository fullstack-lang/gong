package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	_ "embed"
)

//go:embed ng_file_detail.css
var NgFileDetailCssTmpl string

const NgDetailTemplateTS = `// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports{{` + string(rune(NgDetailTsInsertionPerStructImports)) + `}}

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-{{structname}}-detail',
	templateUrl: './{{structname}}-detail.component.html',
	styleUrls: ['./{{structname}}-detail.component.css'],
})
export class {{Structname}}DetailComponent implements OnInit {

	// insertion point for declarations{{` + string(rune(NgDetailTsInsertionPerStructDeclarations)) + `}}

	// the {{Structname}}DB of interest
	{{structname}}: {{Structname}}DB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private {{structname}}Service: {{Structname}}Service,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.get{{Structname}}()

		// observable for changes in structs
		this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.get{{Structname}}()
				}
			}
		)

		// insertion point for initialisation of enums list{{` + string(rune(NgDetailTsInsertionPerStructInits)) + `}}
	}

	get{{Structname}}(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo {{Structname}}Pull returned")

				if (id != 0 && association == undefined) {
					this.{{structname}} = frontRepo.{{Structname}}s.get(id)
				} else {
					this.{{structname}} = new ({{Structname}}DB)
				}

				// insertion point for recovery of form controls value for bool fields{{` + string(rune(NgDetailTsInsertionPerStructRecoveries)) + `}}
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// insertion point for saving value of form controls of boolean fields{{` + string(rune(NgDetailTsInsertionPerStructSaves)) + `}}

		if (id != 0 && association == undefined) {
			// insertion point for saving value of reverse pointers{{` + string(rune(NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate)) + `}}

			this.{{structname}}Service.update{{Structname}}(this.{{structname}})
				.subscribe({{structname}} => {
					this.{{structname}}Service.{{Structname}}ServiceChanged.next("update")

					console.log("{{structname}} saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer{{` + string(rune(NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner)) + `}}
			}
			this.{{structname}}Service.post{{Structname}}(this.{{structname}}).subscribe({{structname}} => {

				this.{{structname}}Service.{{Structname}}ServiceChanged.next("post")

				this.{{structname}} = {} // reset fields
				console.log("{{structname}} added")
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.{{structname}}.ID,
			ReversePointer: reverseField,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}
}
`

// Insertion points
// insertion points in the main template
type NgDetailTsInsertionPoint int

const (
	NgDetailTsInsertionPerStructImports NgDetailTsInsertionPoint = iota
	NgDetailTsInsertionPerStructDeclarations
	NgDetailTsInsertionPerStructInits
	NgDetailTsInsertionPerStructRecoveries
	NgDetailTsInsertionPerStructSaves
	NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate
	NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner
	NgDetailTsInsertionsNb
)

type NgDetailSubTemplate int

const (
	NgDetailTSEnumImports NgDetailSubTemplate = iota
	NgDetailTSEnumDeclarations
	NgDetailTSEnumInits

	NgDetailTSBooleanDeclarations
	NgDetailTSBooleanRecoveries
	NgDetailTSBooleanSaves

	NgDetailTSPointerToGongStructSaves

	NgDetailTSReversePointerToSliceOfGongStructSavesWhenCreateFromOwner
	NgDetailTSReversePointerToSliceOfGongStructSavesWhenUpdate
)

var NgDetailSubTemplateCode map[NgDetailSubTemplate]string = map[NgDetailSubTemplate]string{

	NgDetailTSEnumImports: `
import { {{EnumName}}Select, {{EnumName}}List } from '../{{EnumName}}'`,
	NgDetailTSEnumDeclarations: `
	{{EnumName}}List: {{EnumName}}Select[]`,
	NgDetailTSEnumInits: `
		this.{{EnumName}}List = {{EnumName}}List`,

	NgDetailTSBooleanDeclarations: `
	{{FieldName}}FormControl = new FormControl(false);`,
	NgDetailTSBooleanRecoveries: `
				this.{{FieldName}}FormControl.setValue(this.{{structname}}.{{FieldName}})`,
	NgDetailTSBooleanSaves: `
		this.{{structname}}.{{FieldName}} = this.{{FieldName}}FormControl.value`,

	NgDetailTSPointerToGongStructSaves: `
		if (this.{{structname}}.{{FieldName}}ID == undefined) {
			this.{{structname}}.{{FieldName}}ID = new NullInt64
		}
		if (this.{{structname}}.{{FieldName}} != undefined) {
			this.{{structname}}.{{FieldName}}ID.Int64 = this.{{structname}}.{{FieldName}}.ID
			this.{{structname}}.{{FieldName}}ID.Valid = true
			this.{{structname}}.{{FieldName}}Name = this.{{structname}}.{{FieldName}}.Name
		} else {
			this.{{structname}}.{{FieldName}}ID.Int64 = 0
			this.{{structname}}.{{FieldName}}ID.Valid = true
			this.{{structname}}.{{FieldName}}Name = ""
		}`,

	NgDetailTSReversePointerToSliceOfGongStructSavesWhenCreateFromOwner: `
				case "{{AssocStructName}}_{{FieldName}}":
					this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID = new NullInt64
					this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Int64 = id
					this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Valid = true
					break`,

	NgDetailTSReversePointerToSliceOfGongStructSavesWhenUpdate: `
			if (this.{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse != undefined) {
				this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID = new NullInt64
				this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Int64 = this.{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse.ID
				this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Valid = true
				this.{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse = undefined // very important, otherwise, circular JSON
			}`,
}

// MultiCodeGeneratorNgDetail parses mdlPkg and generates the code for the
// Detail components
func MultiCodeGeneratorNgDetail(
	mdlPkg *ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range mdlPkg.GongStructs {

		if !_struct.HasNameField() {
			continue
		}

		// create the component directory
		dirPath := filepath.Join(matTargetPath, strings.ToLower(_struct.Name)+"-detail")
		errd := os.Mkdir(dirPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + dirPath)

			// add a tempo because on windows, the later creation of file can fails
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)

		// generate the css file
		VerySimpleCodeGenerator(mdlPkg,
			pkgName,
			pkgGoPath,
			filepath.Join(dirPath, strings.ToLower(_struct.Name)+"-detail.component.css"),
			NgFileDetailCssTmpl,
		)

		// generate the typescript file
		codeTS := NgDetailTemplateTS

		TSinsertions := make(map[NgDetailTsInsertionPoint]string)
		for insertion := NgDetailTsInsertionPoint(0); insertion < NgDetailTsInsertionsNb; insertion++ {
			TSinsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				modelBasicField := field.(*GongBasicField)

				if modelBasicField.GongEnum != nil {

					var importToInsert = Replace1(
						NgDetailSubTemplateCode[NgDetailTSEnumImports],
						"{{EnumName}}", modelBasicField.GongEnum.Name)

					// cannot insert twice the same import
					if !strings.Contains(TSinsertions[NgDetailTsInsertionPerStructImports], importToInsert) {
						TSinsertions[NgDetailTsInsertionPerStructImports] += importToInsert
					}

					var declarationToInsert = Replace1(
						NgDetailSubTemplateCode[NgDetailTSEnumDeclarations],
						"{{EnumName}}", modelBasicField.GongEnum.Name)

					if !strings.Contains(TSinsertions[NgDetailTsInsertionPerStructDeclarations], declarationToInsert) {
						TSinsertions[NgDetailTsInsertionPerStructDeclarations] += declarationToInsert
					}

					var initsToInsert = Replace1(
						NgDetailSubTemplateCode[NgDetailTSEnumInits],
						"{{EnumName}}", modelBasicField.GongEnum.Name)

					if !strings.Contains(TSinsertions[NgDetailTsInsertionPerStructInits], initsToInsert) {
						TSinsertions[NgDetailTsInsertionPerStructInits] += initsToInsert
					}
				}

				if modelBasicField.basicKind == types.Bool {
					TSinsertions[NgDetailTsInsertionPerStructDeclarations] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSBooleanDeclarations],
						"{{FieldName}}", modelBasicField.Name)

					TSinsertions[NgDetailTsInsertionPerStructRecoveries] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSBooleanRecoveries],
						"{{FieldName}}", modelBasicField.Name)

					TSinsertions[NgDetailTsInsertionPerStructSaves] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSBooleanSaves],
						"{{FieldName}}", modelBasicField.Name)
				}
			case *PointerToGongStructField:
				modelPointerToStruct := field.(*PointerToGongStructField)
				TSinsertions[NgDetailTsInsertionPerStructSaves] += Replace1(
					NgDetailSubTemplateCode[NgDetailTSPointerToGongStructSaves],
					"{{FieldName}}", modelPointerToStruct.Name)
			}
		}

		codeHTML := NgDetailTemplateHTML

		HtmlInsertions := make(map[NgDetailHtmlInsertionPoint]string)
		for insertion := NgDetailHtmlInsertionPoint(0); insertion < NgDetailHtmlInsertionsNb; insertion++ {
			HtmlInsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				modelBasicField := field.(*GongBasicField)

				// bais field (enum)
				if modelBasicField.GongEnum != nil {

					HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
						Replace2(NgDetailHtmlSubTemplateCode[NgDetailHtmlEnum],
							"{{FieldName}}", modelBasicField.Name,
							"{{EnumName}}", modelBasicField.GongEnum.Name)

				} else // basic field (not enum)
				{
					if modelBasicField.basicKind == types.Bool {
						HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
							Replace1(NgDetailHtmlSubTemplateCode[NgDetailHtmlBool],
								"{{FieldName}}", modelBasicField.Name)
					} else {
						// conversion form go type to ts type
						TypeInput := ""
						switch modelBasicField.basicKind {
						case types.Int, types.Float64:
							TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\""
						case types.String:
							TypeInput = "name=\"\" [ngModelOptions]=\"{standalone: true}\""
						}
						HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
							Replace2(NgDetailHtmlSubTemplateCode[NgDetailHtmlBasicField],
								"{{FieldName}}", modelBasicField.Name,
								"{{TypeInput}}", TypeInput)
					}
				}

			case *PointerToGongStructField:
				modelPointerToStructField := field.(*PointerToGongStructField)
				HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
					Replace3(NgDetailHtmlSubTemplateCode[NgDetailPointerToStructHtmlFormField],
						"{{FieldName}}", modelPointerToStructField.Name,
						"{{AssocStructName}}", modelPointerToStructField.GongStruct.Name,
						"{{assocStructName}}", strings.ToLower(modelPointerToStructField.GongStruct.Name))
			case *SliceOfPointerToGongStructField:
				modelSliceOfPointerToStructField := field.(*SliceOfPointerToGongStructField)
				HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
					Replace3(NgDetailHtmlSubTemplateCode[NgDetailSliceOfPointerToStructHtml],
						"{{FieldName}}", modelSliceOfPointerToStructField.Name,
						"{{AssocStructName}}", modelSliceOfPointerToStructField.GongStruct.Name,
						"{{assocStructName}}", strings.ToLower(modelSliceOfPointerToStructField.GongStruct.Name))
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
						TSinsertions[NgDetailTsInsertionPerStructReversePointerSaveWhenCreateFromOwner] +=
							Replace2(NgDetailSubTemplateCode[NgDetailTSReversePointerToSliceOfGongStructSavesWhenCreateFromOwner],
								"{{FieldName}}", fieldSliceOfPointerToModel.Name,
								"{{AssocStructName}}", __struct.Name)

						TSinsertions[NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate] +=
							Replace2(NgDetailSubTemplateCode[NgDetailTSReversePointerToSliceOfGongStructSavesWhenUpdate],
								"{{FieldName}}", fieldSliceOfPointerToModel.Name,
								"{{AssocStructName}}", __struct.Name)

						HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
							Replace3(NgDetailHtmlSubTemplateCode[NgDetailSliceOfPointerToStructReverseHtml],
								"{{FieldName}}", fieldSliceOfPointerToModel.Name,
								"{{AssocStructName}}", __struct.Name,
								"{{assocStructName}}", strings.ToLower(__struct.Name))
					}
				}
			}
		}

		for insertion := NgDetailTsInsertionPoint(0); insertion < NgDetailTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TSinsertions[insertion])
		}

		for insertion := NgDetailHtmlInsertionPoint(0); insertion < NgDetailHtmlInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeHTML = strings.ReplaceAll(codeHTML, toReplace, HtmlInsertions[insertion])
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
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"-detail.component.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
		{
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"-detail.component.html"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeHTML)
		}

	}
}