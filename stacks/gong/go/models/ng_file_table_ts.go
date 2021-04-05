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

//go:embed ng_file_table.css
var NgFileTableCssTmpl string

const NgTableTemplateTS = `// generated from NgTableTemplateTS
import { Component, OnInit, OnChanges, Input, Output, EventEmitter, Inject, Optional } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

// generated table component
@Component({
  selector: 'app-{{structname}}s-table',
  templateUrl: './{{structname}}s-table.component.html',
  styleUrls: ['./{{structname}}s-table.component.css'],
})
export class {{Structname}}sTableComponent implements OnInit {

  // used if the component is called as a selection component of {{Structname}} instances
  selection: SelectionModel<{{Structname}}DB>;
  initialSelection = new Array<{{Structname}}DB>();

  // the data source for the table
  {{structname}}s: {{Structname}}DB[];

  // front repo, that will be referenced by this.{{structname}}s
  frontRepo: FrontRepo

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  constructor(
    private {{structname}}Service: {{Structname}}Service,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of {{structname}} instances
    public dialogRef: MatDialogRef<{{Structname}}sTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    // https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
    // this is for routerLink on same component when only queryParameter changes
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };

    // observable for changes in structs
    this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.get{{Structname}}s()
        }
      }
    )
    if (dialogData == undefined) {
  	  this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display{{` + string(rune(NgTableTsInsertionPerStructColumns)) + `}}
      ]
    } else {
  	  this.displayedColumns = ['select', 'ID', // insertion point for columns to display{{` + string(rune(NgTableTsInsertionPerStructColumns)) + `}}
      ]
      this.selection = new SelectionModel<{{Structname}}DB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.get{{Structname}}s()
  }

  get{{Structname}}s(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo
        console.log("front repo pull returned")

        this.{{structname}}s = this.frontRepo.{{Structname}}s_array;
        // in case the component is called as a selection component
        if (this.dialogData != undefined) {
          this.{{structname}}s.forEach(
            {{structname}} => {
              let ID = this.dialogData.ID
              let revPointer = {{structname}}[this.dialogData.ReversePointer]
              if (revPointer.Int64 == ID) {
                this.initialSelection.push({{structname}})
              }
            }
          )
          this.selection = new SelectionModel<{{Structname}}DB>(allowMultiSelect, this.initialSelection);
        }
      }
    )
  }

  // new{{Structname}} initiate a new {{structname}}
  // create a new {{Structname}} objet
  new{{Structname}}() {
  }

  delete{{Structname}}({{structname}}ID: number, {{structname}}: {{Structname}}DB) {
    // list of {{structname}}s is truncated of {{structname}} before the delete
    this.{{structname}}s = this.{{structname}}s.filter(h => h !== {{structname}});

    this.{{structname}}Service.delete{{Structname}}({{structname}}ID).subscribe(
      {{structname}} => {
        this.{{structname}}Service.{{Structname}}ServiceChanged.next("delete")

        console.log("{{structname}} deleted")
      }
    );
  }

  edit{{Structname}}({{structname}}ID: number, {{structname}}: {{Structname}}DB) {

  }

  // display {{structname}} in router
  display{{Structname}}InRouter({{structname}}ID: number) {
    this.router.navigate( ["{{structname}}-display", {{structname}}ID])
  }

  // set editor outlet
  setEditorRouterOutlet({{structname}}ID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["{{structname}}-detail", {{structname}}ID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet({{structname}}ID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["{{structname}}-presentation", {{structname}}ID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.{{structname}}s.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.{{structname}}s.forEach(row => this.selection.select(row));
  }

  save() {

    let toUpdate = new Set<{{Structname}}DB>()

    // reset all initial selection of {{structname}} that belong to aclass through Anarrayofb
    this.initialSelection.forEach(
      {{structname}} => {
        {{structname}}[this.dialogData.ReversePointer].Int64 = 0
        {{structname}}[this.dialogData.ReversePointer].Valid = true
        toUpdate.add({{structname}})
      }
    )

    // from selection, set {{structname}} that belong to aclass through Anarrayofb
    this.selection.selected.forEach(
      {{structname}} => {
        console.log("selection ID " + {{structname}}.ID)
        let ID = +this.dialogData.ID
        {{structname}}[this.dialogData.ReversePointer].Int64 = ID
        {{structname}}[this.dialogData.ReversePointer].Valid = true
        toUpdate.add({{structname}})
      }
    )

    // update all {{structname}} (only update selection & initial selection)
    toUpdate.forEach(
      {{structname}} => {
        this.{{structname}}Service.update{{Structname}}({{structname}})
          .subscribe({{structname}} => {
            this.{{structname}}Service.{{Structname}}ServiceChanged.next("update")
            console.log("{{structname}} saved")
          });
      }
    )
    this.dialogRef.close('Pizza!');
  }
}
`

// insertion points in the main template
type NgTableTsInsertionPoint int

const (
	NgTableTsInsertionPerStructColumns NgTableTsInsertionPoint = iota
	NgTableTsInsertionsNb
)

type NgTableSubTemplate int

const (
	NgTableTSPerStructColumn NgTableSubTemplate = iota
)

var NgTablelSubTemplateCode map[NgTableSubTemplate]string = map[NgTableSubTemplate]string{

	NgTableTSPerStructColumn: `
        "{{FieldName}}",`,
}

// MultiCodeGeneratorNgTable parses mdlPkg and generates the code for the
// Table component
func MultiCodeGeneratorNgTable(
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
		dirPath := filepath.Join(matTargetPath, strings.ToLower(_struct.Name)+"s-table")
		errd := os.Mkdir(dirPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + dirPath)
		}

    // generate the css file
		VerySimpleCodeGenerator(mdlPkg,
			pkgName,
			pkgGoPath,
			filepath.Join(dirPath, strings.ToLower(_struct.Name)+"s-table.component.css"),
			NgFileTableCssTmpl,
    )

		// generate the typescript file
		codeTS := NgTableTemplateTS

		TsInsertions := make(map[NgTableTsInsertionPoint]string)
		for insertion := NgTableTsInsertionPoint(0); insertion < NgTableTsInsertionsNb; insertion++ {
			TsInsertions[insertion] = ""
		}

		HtmlInsertions := make(map[NgTableHtmlInsertionPoint]string)
		for insertion := NgTableHtmlInsertionPoint(0); insertion < NgTableHtmlInsertionsNb; insertion++ {
			HtmlInsertions[insertion] = ""
		}

		codeHTML := NgTableTemplateHTML

		for _, field := range _struct.Fields {
			switch field.(type) {
			case *GongBasicField:
				modelBasicField := field.(*GongBasicField)

				// conversion form go type to ts type
				TypeInput := ""
				switch modelBasicField.basicKind {
				case types.Int, types.Float64:
					TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\" "
				case types.String:
					TypeInput = "name=\"\" [ngModelOptions]=\"{standalone: true}\" 	"
				}

				switch modelBasicField.basicKind {
				case types.Float64:
					HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTableHTMLBasicFloat64Field],
						"{{FieldName}}", modelBasicField.Name,
						"{{TypeInput}}", TypeInput)
				case types.Bool:
					HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTableHTMLBool],
						"{{FieldName}}", modelBasicField.Name,
						"{{TypeInput}}", TypeInput)
				default:
					HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTableHTMLBasicField],
						"{{FieldName}}", modelBasicField.Name,
						"{{TypeInput}}", TypeInput)
				}
				TsInsertions[NgTableTsInsertionPerStructColumns] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPerStructColumn],
						"{{FieldName}}", modelBasicField.Name)

			case *PointerToGongStructField:
				modelPointerToStructField := field.(*PointerToGongStructField)

				HtmlInsertions[NgTableHtmlInsertionColumn] += Replace3(NgTableHTMLSubTemplateCode[NgTablePointerToStructHTMLFormField],
					"{{FieldName}}", modelPointerToStructField.Name,
					"{{AssocStructName}}", modelPointerToStructField.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(modelPointerToStructField.GongStruct.Name))

				TsInsertions[NgTableTsInsertionPerStructColumns] +=
					Replace1(NgTablelSubTemplateCode[NgTableTSPerStructColumn],
						"{{FieldName}}", modelPointerToStructField.Name)
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
						HtmlInsertions[NgTableHtmlInsertionColumn] += Replace2(NgTableHTMLSubTemplateCode[NgTablePointerToSliceOfGongStructHTMLFormField],
							"{{FieldName}}", fieldSliceOfPointerToModel.Name,
							"{{AssocStructName}}", __struct.Name)

						TsInsertions[NgTableTsInsertionPerStructColumns] +=
							Replace1(NgTablelSubTemplateCode[NgTableTSPerStructColumn],
								"{{FieldName}}", fieldSliceOfPointerToModel.Name)
					}
				}
			}
		}

		for insertion := NgTableTsInsertionPoint(0); insertion < NgTableTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TsInsertions[insertion])
		}

		for insertion := NgTableHtmlInsertionPoint(0); insertion < NgTableHtmlInsertionsNb; insertion++ {
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
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"s-table.component.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
		{
			file, err := os.Create(filepath.Join(dirPath, strings.ToLower(_struct.Name)+"s-table.component.html"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeHTML)
		}

	}
}
