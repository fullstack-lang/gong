package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const NgLibFrontRepoServiceTemplate = `import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports {{` + string(NgLibFrontRepoServiceImports) + `}}

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template {{` + string(NgLibFrontRepoMapDecl) + `}}
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number
    Valid: boolean
}

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template {{` + string(NgLibFrontRepoServiceDecl) + `}}
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template {{` + string(NgLibFrontRepoObservableArrayType) + `}}
  ] = [ // insertion point sub template {{` + string(NgLibFrontRepoObservableRefs) + `}}
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations {{` + string(NgLibFrontRepoArraysDecls) + `}}
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting {{` + string(NgLibFrontRepoTypeCasting) + `}}

            // 
            // First Step: init map of instances
            // insertion point sub template for init {{` + string(NgLibFrontRepoInitMapInstances) + `}}

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem {{` + string(NgLibFrontRepoRedeemPointers) + `}}

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct {{` + string(NgLibFrontRepoPerStructPull) + `}}
}

// insertion point for get unique ID per struct {{` + string(NgLibFrontRepoPerStructGetUniqueID) + `}}
`

type NgLibFrontRepoServiceSubTemplate string

const (
	NgLibFrontRepoServiceImports       NgLibFrontRepoServiceSubTemplate = "ServiceImports"
	NgLibFrontRepoMapDecl              NgLibFrontRepoServiceSubTemplate = "MapDecl"
	NgLibFrontRepoObservableArrayType  NgLibFrontRepoServiceSubTemplate = "ObservableArrayType"
	NgLibFrontRepoTypeCasting          NgLibFrontRepoServiceSubTemplate = "TypeCasting"
	NgLibFrontRepoServiceDecl          NgLibFrontRepoServiceSubTemplate = "ServiceDecl"
	NgLibFrontRepoObservableRefs       NgLibFrontRepoServiceSubTemplate = "ObservableRefs"
	NgLibFrontRepoArraysDecls          NgLibFrontRepoServiceSubTemplate = "ArraysDecls"
	NgLibFrontRepoInitMapInstances     NgLibFrontRepoServiceSubTemplate = "InitMapInstances"
	NgLibFrontRepoRedeemPointers       NgLibFrontRepoServiceSubTemplate = "RedeemPointers"
	NgLibFrontRepoPerStructPull        NgLibFrontRepoServiceSubTemplate = "PerStructPull"
	NgLibFrontRepoPerStructGetUniqueID NgLibFrontRepoServiceSubTemplate = "PerStructGetUniqueID"
)

var NgFrontRepoPerStructTmplCodes map[string]string = // new line
map[string]string{

	string(NgLibFrontRepoServiceImports): `
import { {{Structname}}DB } from './{{structname}}-db'
import { {{Structname}}Service } from './{{structname}}.service'
`,

	string(NgLibFrontRepoMapDecl): `
  {{Structname}}s_array = new Array<{{Structname}}DB>(); // array of repo instances
  {{Structname}}s = new Map<number, {{Structname}}DB>(); // map of repo instances
  {{Structname}}s_batch = new Map<number, {{Structname}}DB>(); // same but only in last GET (for finding repo instances to delete)`,

	string(NgLibFrontRepoObservableArrayType): `
    Observable<{{Structname}}DB[]>,`,

	string(NgLibFrontRepoServiceDecl): `
    private {{structname}}Service: {{Structname}}Service,`,

	string(NgLibFrontRepoObservableRefs): `
      this.{{structname}}Service.get{{Structname}}s(),`,

	string(NgLibFrontRepoArraysDecls): `
            {{structname}}s_,`,

	string(NgLibFrontRepoTypeCasting): `
            var {{structname}}s: {{Structname}}DB[]
            {{structname}}s = {{structname}}s_`,

	string(NgLibFrontRepoInitMapInstances): `
            // init the array
            FrontRepoSingloton.{{Structname}}s_array = {{structname}}s

            // clear the map that counts {{Structname}} in the GET
            FrontRepoSingloton.{{Structname}}s_batch.clear()
            
            {{structname}}s.forEach(
              {{structname}} => {
                FrontRepoSingloton.{{Structname}}s.set({{structname}}.ID, {{structname}})
                FrontRepoSingloton.{{Structname}}s_batch.set({{structname}}.ID, {{structname}})
              }
            )
            
            // clear {{structname}}s that are absent from the batch
            FrontRepoSingloton.{{Structname}}s.forEach(
              {{structname}} => {
                if (FrontRepoSingloton.{{Structname}}s_batch.get({{structname}}.ID) == undefined) {
                  FrontRepoSingloton.{{Structname}}s.delete({{structname}}.ID)
                }
              }
            )`,

	string(NgLibFrontRepoRedeemPointers): `
            {{structname}}s.forEach(
              {{structname}} => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming{{` + string(NgFrontRepoPtrToStructRedeeming) + `}}

                // insertion point for redeeming ONE-MANY associations{{` + string(NgFrontRepoSliceOfPointerRedeeming) + `}}
              }
            )`,

	string(NgLibFrontRepoPerStructPull): `

  // {{Structname}}Pull performs a GET on {{Structname}} of the stack and redeem association pointers 
  {{Structname}}Pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.{{structname}}Service.get{{Structname}}s()
        ]).subscribe(
          ([ // insertion point sub template 
            {{structname}}s,
          ]) => {
            // init the array
            FrontRepoSingloton.{{Structname}}s_array = {{structname}}s

            // clear the map that counts {{Structname}} in the GET
            FrontRepoSingloton.{{Structname}}s_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            {{structname}}s.forEach(
              {{structname}} => {
                FrontRepoSingloton.{{Structname}}s.set({{structname}}.ID, {{structname}})
                FrontRepoSingloton.{{Structname}}s_batch.set({{structname}}.ID, {{structname}})

                // insertion point for redeeming ONE/ZERO-ONE associations {{` + string(NgFrontRepoPtrToStructRedeeming) + `}}

                // insertion point for redeeming ONE-MANY associations {{` + string(NgFrontRepoSliceOfPointerRedeeming) + `}}
              }
            )

            // clear {{structname}}s that are absent from the GET
            FrontRepoSingloton.{{Structname}}s.forEach(
              {{structname}} => {
                if (FrontRepoSingloton.{{Structname}}s_batch.get({{structname}}.ID) == undefined) {
                  FrontRepoSingloton.{{Structname}}s.delete({{structname}}.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }`,

	string(NgLibFrontRepoPerStructGetUniqueID): `
export function get{{Structname}}UniqueID(id: number): number {
  return {{Prime}} * id
}`,
}

type NgLibFrontRepoServiceSubSubTemplate string

const (
	NgFrontRepoPtrToStructRedeeming NgLibFrontRepoServiceSubSubTemplate = "NgFrontRepoPtrToStructRedeeming"
)

// for each sub sub template, what sub template it relates to
var NgLibFrontRepoPointerToStructSubSubToSubMap map[string]string = //
map[string]string{
	string(NgFrontRepoPtrToStructRedeeming): string(NgLibFrontRepoRedeemPointers),
}

var NgFrontRepoPtrToStructTmplCodes map[string]string = // new line
map[string]string{
	string(NgFrontRepoPtrToStructRedeeming): `
                // insertion point for pointer field {{FieldName}} redeeming
                {
                  let _{{assocStructName}} = FrontRepoSingloton.{{AssocStructName}}s.get({{structname}}.{{FieldName}}ID.Int64)
                  if (_{{assocStructName}}) {
                    {{structname}}.{{FieldName}} = _{{assocStructName}}
                  }
                }`,
}

const (
	NgFrontRepoSliceOfPointerRedeeming NgLibFrontRepoServiceSubSubTemplate = "NgFrontRepoSliceOfPointerRedeeming"
)

var NgFrontRepoSliceOfPointerToStructTmplCode map[string]string = // new line
map[string]string{
	string(NgFrontRepoSliceOfPointerRedeeming): `
                // insertion point for slice of pointer field {{Structname}}.{{FieldName}} redeeming
                {
                  let _{{structname}} = FrontRepoSingloton.{{Structname}}s.get({{assocStructName}}.{{Structname}}_{{FieldName}}DBID.Int64)
                  if (_{{structname}}) {
                    if (_{{structname}}.{{FieldName}} == undefined) {
                      _{{structname}}.{{FieldName}} = new Array<{{Structname}}DB>()
                    }
                    _{{structname}}.{{FieldName}}.push({{assocStructName}})
                    if ({{assocStructName}}.{{Structname}}_{{FieldName}}_reverse == undefined) {
                      {{assocStructName}}.{{Structname}}_{{FieldName}}_reverse = _{{structname}}
                    }
                  }
                }`,
}

func CodeGeneratorNgFrontRepo(
	mdlPkg *ModelPkg,
	pkgName string,
	matTargetPath string,
	pkgGoPath string) {

	file, err := os.Create(filepath.Join(matTargetPath, "front-repo.service.ts"))
	if err != nil {
		log.Panic(err)
	}

	code := NgLibFrontRepoServiceTemplate

	perStructCodes := make(map[string]string)
	for subTemplate := range NgFrontRepoPerStructTmplCodes {
		perStructCodes[subTemplate] = ""
	}

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		if !_struct.HasNameField() {
			continue
		}
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for structIndex, _struct := range structList {

		structName := strings.ToLower(_struct.Name)
		_ = structName

		fieldPointerToStructCodes := ""
		fieldSliceOfPtrToStructCodes := ""

		// compute code per field
		for _, field := range _struct.Fields {
			switch field.(type) {
			case *PointerToGongStructField:
				fieldPointerToModel := field.(*PointerToGongStructField)

				fieldName := strings.ToLower(fieldPointerToModel.Name)
				assocStructName := strings.ToLower(fieldPointerToModel.GongStruct.Name)

				fieldPointerToStructCodes += Replace6(
					NgFrontRepoPtrToStructTmplCodes[string(NgFrontRepoPtrToStructRedeeming)],
					"{{fieldName}}", fieldName,
					"{{FieldName}}", fieldPointerToModel.Name,
					"{{AssocStructName}}", fieldPointerToModel.GongStruct.Name,
					"{{assocStructName}}", assocStructName,
					"{{Structname}}", _struct.Name,
					"{{structname}}", structName)

			case *SliceOfPointerToGongStructField:
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

						fieldName := strings.ToLower(fieldSliceOfPointerToModel.Name)
						assocStructName := strings.ToLower(fieldSliceOfPointerToModel.GongStruct.Name)
						__structName := strings.ToLower(__struct.Name)

						fieldSliceOfPtrToStructCodes += Replace6(
							NgFrontRepoSliceOfPointerToStructTmplCode[string(NgFrontRepoSliceOfPointerRedeeming)],
							"{{fieldName}}", fieldName,
							"{{FieldName}}", fieldSliceOfPointerToModel.Name,
							"{{AssocStructName}}", fieldSliceOfPointerToModel.GongStruct.Name,
							"{{assocStructName}}", assocStructName,
							"{{Structname}}", __struct.Name,
							"{{structname}}", __structName)
					}
				}
			}
		}

		// compute code from sub template
		for subTemplate := range NgFrontRepoPerStructTmplCodes {
			perStructCodes[subTemplate] += Replace3(NgFrontRepoPerStructTmplCodes[subTemplate],
				"{{Structname}}", _struct.Name,
				"{{structname}}", structName,
				"{{Prime}}", strconv.Itoa(Primes[structIndex+10]))

			toReplace := "{{" + string(NgFrontRepoPtrToStructRedeeming) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldPointerToStructCodes)

			toReplace = "{{" + string(NgFrontRepoSliceOfPointerRedeeming) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldSliceOfPtrToStructCodes)
		}
	}

	//
	// final per struct replacement
	//
	for perStructTmplCode := range NgFrontRepoPerStructTmplCodes {
		toReplace := "{{" + string(perStructTmplCode) + "}}"
		code = strings.ReplaceAll(code, toReplace, perStructCodes[perStructTmplCode])
	}

	code = Replace4(code,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))

	defer file.Close()
	fmt.Fprint(file, code)
}