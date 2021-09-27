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

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports{{` + string(rune(NgDetailTsInsertionPerStructImports)) + `}}

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// {{Structname}}DetailComponent is initilizaed from different routes
// {{Structname}}DetailComponentState detail different cases 
enum {{Structname}}DetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state{{` + string(rune(NgDetailTsInsertionPerStructEnumFieldDeclarations)) + `}}
}

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

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: {{Structname}}DetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private {{structname}}Service: {{Structname}}Service,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id');
		this.originStruct = this.route.snapshot.paramMap.get('originStruct');
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName');

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = {{Structname}}DetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = {{Structname}}DetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation{{` + string(rune(NgDetailTsInsertionPerStructCaseInitFieldDeclarations)) + `}}
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

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

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case {{Structname}}DetailComponentState.CREATE_INSTANCE:
						this.{{structname}} = new ({{Structname}}DB)
						break;
					case {{Structname}}DetailComponentState.UPDATE_INSTANCE:
						this.{{structname}} = frontRepo.{{Structname}}s.get(this.id)
						break;
					// insertion point for init of association field{{` + string(rune(NgDetailTsInsertionPerStructCaseSetField)) + `}}
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields{{` + string(rune(NgDetailTsInsertionPerStructRecoveries)) + `}}
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field{{` + string(rune(NgDetailTsInsertionPerStructSaves)) + `}}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers{{` + string(rune(NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate)) + `}}

		switch (this.state) {
			case {{Structname}}DetailComponentState.UPDATE_INSTANCE:
				this.{{structname}}Service.update{{Structname}}(this.{{structname}})
					.subscribe({{structname}} => {
						this.{{structname}}Service.{{Structname}}ServiceChanged.next("update")
					});
				break;
			default:
				this.{{structname}}Service.post{{Structname}}(this.{{structname}}).subscribe({{structname}} => {
					this.{{structname}}Service.{{Structname}}ServiceChanged.next("post")
					this.{{structname}} = {} // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string ) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.{{structname}}.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.{{structname}}.ID
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "{{Structname}}"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.{{structname}}.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event) {
		if (this.{{structname}}.Name == undefined) {
			this.{{structname}}.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
`

// Insertion points
// insertion points in the main template
type NgDetailTsInsertionPoint int

const (
	NgDetailTsInsertionPerStructImports NgDetailTsInsertionPoint = iota
	NgDetailTsInsertionPerStructEnumFieldDeclarations
	NgDetailTsInsertionPerStructDeclarations
	NgDetailTsInsertionPerStructCaseInitFieldDeclarations
	NgDetailTsInsertionPerStructCaseSetField
	NgDetailTsInsertionPerStructInits
	NgDetailTsInsertionPerStructSorting
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

	NgDetailTSTimeDurationDeclarations
	NgDetailTSTimeDurationRecoveries
	NgDetailTSTimeDurationSaves

	NgDetailTSPointerToGongStructSaves

	NgDetailTSReversePointerToSliceOfGongStructStateEnumDeclaration
	NgDetailTSReversePointerToSliceOfGongStructStateCaseComputation
	NgDetailTSReversePointerToSliceOfGongStructStateCaseSetField
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

	NgDetailTSTimeDurationDeclarations: `
	{{FieldName}}_Hours: number
	{{FieldName}}_Minutes: number
	{{FieldName}}_Seconds: number`,

	NgDetailTSTimeDurationRecoveries: `
				// computation of Hours, Minutes, Seconds for {{FieldName}}
				this.{{FieldName}}_Hours = Math.floor(this.{{structname}}.{{FieldName}} / (3600 * 1000 * 1000 * 1000))
				this.{{FieldName}}_Minutes = Math.floor(this.{{structname}}.{{FieldName}} % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.{{FieldName}}_Seconds = this.{{structname}}.{{FieldName}} % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)`,

	NgDetailTSTimeDurationSaves: `
		this.{{structname}}.{{FieldName}} =
			this.{{FieldName}}_Hours * (3600 * 1000 * 1000 * 1000) +
			this.{{FieldName}}_Minutes * (60 * 1000 * 1000 * 1000) +
			this.{{FieldName}}_Seconds * (1000 * 1000 * 1000)`,

	NgDetailTSPointerToGongStructSaves: `
		if (this.{{structname}}.{{FieldName}}ID == undefined) {
			this.{{structname}}.{{FieldName}}ID = new NullInt64
		}
		if (this.{{structname}}.{{FieldName}} != undefined) {
			this.{{structname}}.{{FieldName}}ID.Int64 = this.{{structname}}.{{FieldName}}.ID
			this.{{structname}}.{{FieldName}}ID.Valid = true
		} else {
			this.{{structname}}.{{FieldName}}ID.Int64 = 0
			this.{{structname}}.{{FieldName}}ID.Valid = true
		}`,

	NgDetailTSReversePointerToSliceOfGongStructStateEnumDeclaration: `
	CREATE_INSTANCE_WITH_ASSOCIATION_{{AssocStructName}}_{{FieldName}}_SET,`,

	NgDetailTSReversePointerToSliceOfGongStructStateCaseComputation: `
					case "{{FieldName}}":
						console.log("{{Structname}}" + " is instanciated with back pointer to instance " + this.id + " {{AssocStructName}} association {{FieldName}}")
						this.state = {{Structname}}DetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_{{AssocStructName}}_{{FieldName}}_SET
						break;`,

	NgDetailTSReversePointerToSliceOfGongStructStateCaseSetField: `
					case {{Structname}}DetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_{{AssocStructName}}_{{FieldName}}_SET:
						this.{{structname}} = new ({{Structname}}DB)
						this.{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse = frontRepo.{{AssocStructName}}s.get(this.id)
						break;`,

	NgDetailTSReversePointerToSliceOfGongStructSavesWhenUpdate: `
		if (this.{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse != undefined) {
			if (this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID == undefined) {
				this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID = new NullInt64
			}
			this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Int64 = this.{{structname}}.{{AssocStructName}}_{{FieldName}}_reverse.ID
			this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID.Valid = true
			if (this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID_Index == undefined) {
				this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID_Index = new NullInt64
			}
			this.{{structname}}.{{AssocStructName}}_{{FieldName}}DBID_Index.Valid = true
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
			switch field := field.(type) {
			case *GongBasicField:

				if field.GongEnum != nil {

					var importToInsert = Replace1(
						NgDetailSubTemplateCode[NgDetailTSEnumImports],
						"{{EnumName}}", field.GongEnum.Name)

					// cannot insert twice the same import
					if !strings.Contains(TSinsertions[NgDetailTsInsertionPerStructImports], importToInsert) {
						TSinsertions[NgDetailTsInsertionPerStructImports] += importToInsert
					}

					var declarationToInsert = Replace1(
						NgDetailSubTemplateCode[NgDetailTSEnumDeclarations],
						"{{EnumName}}", field.GongEnum.Name)

					if !strings.Contains(TSinsertions[NgDetailTsInsertionPerStructDeclarations], declarationToInsert) {
						TSinsertions[NgDetailTsInsertionPerStructDeclarations] += declarationToInsert
					}

					var initsToInsert = Replace1(
						NgDetailSubTemplateCode[NgDetailTSEnumInits],
						"{{EnumName}}", field.GongEnum.Name)

					if !strings.Contains(TSinsertions[NgDetailTsInsertionPerStructInits], initsToInsert) {
						TSinsertions[NgDetailTsInsertionPerStructInits] += initsToInsert
					}
				}

				if field.basicKind == types.Bool {
					TSinsertions[NgDetailTsInsertionPerStructDeclarations] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSBooleanDeclarations],
						"{{FieldName}}", field.Name)

					TSinsertions[NgDetailTsInsertionPerStructRecoveries] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSBooleanRecoveries],
						"{{FieldName}}", field.Name)

					TSinsertions[NgDetailTsInsertionPerStructSaves] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSBooleanSaves],
						"{{FieldName}}", field.Name)
				}

				if field.DeclaredType == "time.Duration" {
					TSinsertions[NgDetailTsInsertionPerStructDeclarations] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSTimeDurationDeclarations],
						"{{FieldName}}", field.Name)

					TSinsertions[NgDetailTsInsertionPerStructRecoveries] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSTimeDurationRecoveries],
						"{{FieldName}}", field.Name)

					TSinsertions[NgDetailTsInsertionPerStructSaves] += Replace1(
						NgDetailSubTemplateCode[NgDetailTSTimeDurationSaves],
						"{{FieldName}}", field.Name)
				}
			case *PointerToGongStructField:

				TSinsertions[NgDetailTsInsertionPerStructSaves] += Replace1(
					NgDetailSubTemplateCode[NgDetailTSPointerToGongStructSaves],
					"{{FieldName}}", field.Name)
			}
		}

		codeHTML := NgDetailTemplateHTML

		HtmlInsertions := make(map[NgDetailHtmlInsertionPoint]string)
		for insertion := NgDetailHtmlInsertionPoint(0); insertion < NgDetailHtmlInsertionsNb; insertion++ {
			HtmlInsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *GongBasicField:

				// bais field (enum)
				if field.GongEnum != nil {

					HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
						Replace2(NgDetailHtmlSubTemplateCode[NgDetailHtmlEnum],
							"{{FieldName}}", field.Name,
							"{{EnumName}}", field.GongEnum.Name)

				} else // basic field (not enum)
				{
					if field.basicKind == types.Bool {
						HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
							Replace1(NgDetailHtmlSubTemplateCode[NgDetailHtmlBool],
								"{{FieldName}}", field.Name)
					} else if field.DeclaredType != "time.Duration" {

						// conversion form go type to ts type
						TypeInput := ""
						switch field.basicKind {
						case types.Int, types.Int64, types.Float64:
							TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\""
							HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
								Replace2(NgDetailHtmlSubTemplateCode[NgDetailHtmlBasicField],
									"{{FieldName}}", field.Name,
									"{{TypeInput}}", TypeInput)
						case types.String:
							TypeInput = "name=\"\" [ngModelOptions]=\"{standalone: true}\""
							HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
								Replace2(NgDetailHtmlSubTemplateCode[NgDetailHtmlBasicStringField],
									"{{FieldName}}", field.Name,
									"{{TypeInput}}", TypeInput)
						}

					} else {
						HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
							Replace1(NgDetailHtmlSubTemplateCode[NgDetailHtmlTimeDuration],
								"{{FieldName}}", field.Name)
					}
				}
			case *GongTimeField:

				HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
					Replace1(NgDetailHtmlSubTemplateCode[NgDetailHtmlTimeField],
						"{{FieldName}}", field.Name)

			case *PointerToGongStructField:

				HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
					Replace3(NgDetailHtmlSubTemplateCode[NgDetailPointerToStructHtmlFormField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))
			case *SliceOfPointerToGongStructField:

				htmlCodeForField := Replace3(NgDetailHtmlSubTemplateCode[NgDetailSliceOfPointerToStructHtml],
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// check if this is a field of a MANY MANY association
				if strings.HasSuffix(field.Name, "Use") {

					intermediateStruct := field.GongStruct

					// a "Use" struct should have exactly 2 fields (Name and the pointer to the next struct)
					if len(intermediateStruct.Fields) != 2 {
						log.Panicf("%s struct is a Use struct. Expected 2 fields, got %d",
							field.Name, len(intermediateStruct.Fields))
					}

					intermediateStructField := intermediateStruct.Fields[1]
					_ = intermediateStructField

					// get the end type. first  (cast on ta PointerToStructField)
					switch intermediateStructField := intermediateStructField.(type) {
					case *PointerToGongStructField:
						nextStruct := (intermediateStructField).GongStruct

						addedHtmlCode := Replace5(NgDetailHtmlSubTemplateCode[NgDetailSliceOfPointerToStructManyManyHtml],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocStructName}}", strings.ToLower(field.GongStruct.Name),
							"{{IntermediateStructField}}", intermediateStructField.GetName(),
							"{{NextAssociatedStruct}}", nextStruct.Name)
						toReplace := "{{" + string(rune(NgDetailHtmlInsertionPerStructFieldsManyMany)) + "}}"
						htmlCodeForField = strings.ReplaceAll(htmlCodeForField, toReplace, addedHtmlCode)
					default:
					}

				}
				HtmlInsertions[NgDetailHtmlInsertionPerStructFields] += htmlCodeForField

				// HtmlInsertions[NgDetailHtmlInsertionPerStructFieldsManyMany] +=
				// 	Replace3(NgDetailHtmlSubTemplateCode[NgDetailSliceOfPointerToStructManyManyHtml],
				// 		"{{FieldName}}", field.Name,
				// 		"{{AssocStructName}}", field.GongStruct.Name,
				// 		"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))
			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field := field.(type) {
				case *SliceOfPointerToGongStructField:

					if field.GongStruct == _struct {
						TSinsertions[NgDetailTsInsertionPerStructEnumFieldDeclarations] +=
							Replace2(NgDetailSubTemplateCode[NgDetailTSReversePointerToSliceOfGongStructStateEnumDeclaration],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						TSinsertions[NgDetailTsInsertionPerStructCaseInitFieldDeclarations] +=
							Replace2(NgDetailSubTemplateCode[NgDetailTSReversePointerToSliceOfGongStructStateCaseComputation],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						TSinsertions[NgDetailTsInsertionPerStructCaseSetField] +=
							Replace2(NgDetailSubTemplateCode[NgDetailTSReversePointerToSliceOfGongStructStateCaseSetField],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						TSinsertions[NgDetailTsInsertionPerStructReversePointerSaveWhenUpdate] +=
							Replace2(NgDetailSubTemplateCode[NgDetailTSReversePointerToSliceOfGongStructSavesWhenUpdate],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						HtmlInsertions[NgDetailHtmlInsertionPerStructFields] +=
							Replace3(NgDetailHtmlSubTemplateCode[NgDetailSliceOfPointerToStructReverseHtml],
								"{{FieldName}}", field.Name,
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
