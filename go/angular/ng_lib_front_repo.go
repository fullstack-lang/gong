package angular

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
)

const NgLibFrontRepoServiceTemplate = `import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports{{` + string(NgLibFrontRepoServiceImports) + `}}

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template{{` + string(NgLibFrontRepoMapDecl) + `}}

  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point{{` + string(NgLibFrontRepoSwitchGetArray) + `}}
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point{{` + string(NgLibFrontRepoSwitchGetMap) + `}}
      default:
        throw new Error("Type not recognized");
    }
  }
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template {{` + string(NgLibFrontRepoServiceDecl) + `}}
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [
    Observable<null>, // see below for the of(null) observable
    // insertion point sub template {{` + string(NgLibFrontRepoObservableArrayType) + `}}
  ] = [
      // Using "combineLatest" with a placeholder observable.
      //
      // This allows the typescript compiler to pass when no GongStruct is present in the front API
      //
      // The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
      // This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
      // expectation for a non-empty array of observables.
      of(null), // 
      // insertion point sub template{{` + string(NgLibFrontRepoObservableRefs) + `}}
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [
      of(null), // see above for justification
      // insertion point sub template{{` + string(NgLibFrontRepoObservableRefs) + `}}
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations {{` + string(NgLibFrontRepoArraysDecls) + `}}
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting {{` + string(NgLibFrontRepoTypeCasting) + `}}

            // 
            // First Step: init map of instances
            // insertion point sub template for init {{` + string(NgLibFrontRepoInitMapInstances) + `}}

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem {{` + string(NgLibFrontRepoRedeemPointers) + `}}

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem {{` + string(NgLibFrontRepoSortSlicesOfPointers) + `}}

            // hand over control flow to observer
            observer.next(this.frontRepo)
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
	NgLibFrontRepoSwitchGetArray       NgLibFrontRepoServiceSubTemplate = "SwitchGetArray"
	NgLibFrontRepoSwitchGetMap         NgLibFrontRepoServiceSubTemplate = "SwitchGetMap"
	NgLibFrontRepoObservableArrayType  NgLibFrontRepoServiceSubTemplate = "ObservableArrayType"
	NgLibFrontRepoTypeCasting          NgLibFrontRepoServiceSubTemplate = "TypeCasting"
	NgLibFrontRepoServiceDecl          NgLibFrontRepoServiceSubTemplate = "ServiceDecl"
	NgLibFrontRepoObservableRefs       NgLibFrontRepoServiceSubTemplate = "ObservableRefs"
	NgLibFrontRepoArraysDecls          NgLibFrontRepoServiceSubTemplate = "ArraysDecls"
	NgLibFrontRepoInitMapInstances     NgLibFrontRepoServiceSubTemplate = "InitMapInstances"
	NgLibFrontRepoRedeemPointers       NgLibFrontRepoServiceSubTemplate = "RedeemPointers"
	NgLibFrontRepoSortSlicesOfPointers NgLibFrontRepoServiceSubTemplate = "SortSlicesOfPointers"
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
  {{Structname}}s_array = new Array<{{Structname}}DB>() // array of repo instances
  {{Structname}}s = new Map<number, {{Structname}}DB>() // map of repo instances
  {{Structname}}s_batch = new Map<number, {{Structname}}DB>() // same but only in last GET (for finding repo instances to delete)
`,

	string(NgLibFrontRepoSwitchGetArray): `
      case '{{Structname}}':
        return this.{{Structname}}s_array as unknown as Array<Type>`,

	string(NgLibFrontRepoSwitchGetMap): `
      case '{{Structname}}':
        return this.{{Structname}}s_array as unknown as Map<number, Type>`,

	string(NgLibFrontRepoObservableArrayType): `
    Observable<{{Structname}}DB[]>,`,

	string(NgLibFrontRepoServiceDecl): `
    private {{structname}}Service: {{Structname}}Service,`,

	string(NgLibFrontRepoObservableRefs): `
      this.{{structname}}Service.get{{Structname}}s(this.GONG__StackPath),`,

	string(NgLibFrontRepoArraysDecls): `
            {{structname}}s_,`,

	string(NgLibFrontRepoTypeCasting): `
            var {{structname}}s: {{Structname}}DB[]
            {{structname}}s = {{structname}}s_ as {{Structname}}DB[]`,

	string(NgLibFrontRepoInitMapInstances): `
            // init the array
            this.frontRepo.{{Structname}}s_array = {{structname}}s

            // clear the map that counts {{Structname}} in the GET
            this.frontRepo.{{Structname}}s_batch.clear()

            {{structname}}s.forEach(
              {{structname}} => {
                this.frontRepo.{{Structname}}s.set({{structname}}.ID, {{structname}})
                this.frontRepo.{{Structname}}s_batch.set({{structname}}.ID, {{structname}})
              }
            )

            // clear {{structname}}s that are absent from the batch
            this.frontRepo.{{Structname}}s.forEach(
              {{structname}} => {
                if (this.frontRepo.{{Structname}}s_batch.get({{structname}}.ID) == undefined) {
                  this.frontRepo.{{Structname}}s.delete({{structname}}.ID)
                }
              }
            )

            // sort {{Structname}}s_array array
            this.frontRepo.{{Structname}}s_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
`,

	string(NgLibFrontRepoRedeemPointers): `
            {{structname}}s.forEach(
              {{structname}} => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming{{` + string(NgFrontRepoPtrToStructRedeeming) + `}}

                // insertion point for redeeming ONE-MANY associations{{` + string(NgFrontRepoSliceOfPointerRedeeming) + `}}
              }
            )`,

	string(NgLibFrontRepoSortSlicesOfPointers): `
            {{structname}}s.forEach(
              {{structname}} => {
                // insertion point for sorting{{` + string(NgFrontRepoSliceOfPointerSorting) + `}}
              }
            )`,

	string(NgLibFrontRepoPerStructPull): `

  // {{Structname}}Pull performs a GET on {{Structname}} of the stack and redeem association pointers 
  {{Structname}}Pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.{{structname}}Service.get{{Structname}}s(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            {{structname}}s,
          ]) => {
            // init the array
            this.frontRepo.{{Structname}}s_array = {{structname}}s

            // clear the map that counts {{Structname}} in the GET
            this.frontRepo.{{Structname}}s_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            {{structname}}s.forEach(
              {{structname}} => {
                this.frontRepo.{{Structname}}s.set({{structname}}.ID, {{structname}})
                this.frontRepo.{{Structname}}s_batch.set({{structname}}.ID, {{structname}})

                // insertion point for redeeming ONE/ZERO-ONE associations{{` + string(NgFrontRepoPtrToStructRedeeming) + `}}

                // insertion point for redeeming ONE-MANY associations{{` + string(NgFrontRepoSliceOfPointerRedeeming) + `}}
              }
            )

            // clear {{structname}}s that are absent from the GET
            this.frontRepo.{{Structname}}s.forEach(
              {{structname}} => {
                if (this.frontRepo.{{Structname}}s_batch.get({{structname}}.ID) == undefined) {
                  this.frontRepo.{{Structname}}s.delete({{structname}}.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
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
                  let _{{assocStructName}} = this.frontRepo.{{AssocStructName}}s.get({{structname}}.{{FieldName}}ID.Int64)
                  if (_{{assocStructName}}) {
                    {{structname}}.{{FieldName}} = _{{assocStructName}}
                  }
                }`,
}

const (
	NgFrontRepoSliceOfPointerRedeeming NgLibFrontRepoServiceSubSubTemplate = "NgFrontRepoSliceOfPointerRedeeming"
	NgFrontRepoSliceOfPointerSorting   NgLibFrontRepoServiceSubSubTemplate = "NgFrontRepoSliceOfPointerSorting"
)

var NgFrontRepoSliceOfPointerToStructTmplCode map[string]string = // new line
map[string]string{
	string(NgFrontRepoSliceOfPointerRedeeming): `
                // insertion point for slice of pointer field {{Structname}}.{{FieldName}} redeeming
                {
                  let _{{structname}} = this.frontRepo.{{Structname}}s.get({{assocStructName}}.{{Structname}}_{{FieldName}}DBID.Int64)
                  if (_{{structname}}) {
                    if (_{{structname}}.{{FieldName}} == undefined) {
                      _{{structname}}.{{FieldName}} = new Array<{{AssocStructName}}DB>()
                    }
                    _{{structname}}.{{FieldName}}.push({{assocStructName}})
                    if ({{assocStructName}}.{{Structname}}_{{FieldName}}_reverse == undefined) {
                      {{assocStructName}}.{{Structname}}_{{FieldName}}_reverse = _{{structname}}
                    }
                  }
                }`,
	string(NgFrontRepoSliceOfPointerSorting): `
                {{structname}}.{{FieldName}}?.sort((t1, t2) => {
                  if (t1.{{Structname}}_{{FieldName}}DBID_Index.Int64 > t2.{{Structname}}_{{FieldName}}DBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.{{Structname}}_{{FieldName}}DBID_Index.Int64 < t2.{{Structname}}_{{FieldName}}DBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })
`,
}

func CodeGeneratorNgFrontRepo(
	mdlPkg *models.ModelPkg,
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
	structList := []*models.GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		if !_struct.HasNameField() || _struct.IsIgnoredForFront {
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
		fieldSliceOfPtrSortingToStructCodes := ""

		// compute code per field
		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *models.PointerToGongStructField:

				fieldName := strings.ToLower(field.Name)
				assocStructName := strings.ToLower(field.GongStruct.Name)

				fieldPointerToStructCodes += models.Replace6(
					NgFrontRepoPtrToStructTmplCodes[string(NgFrontRepoPtrToStructRedeeming)],
					"{{fieldName}}", fieldName,
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", assocStructName,
					"{{Structname}}", _struct.Name,
					"{{structname}}", structName)

			case *models.SliceOfPointerToGongStructField:

				fieldName := strings.ToLower(field.Name)
				assocStructName := strings.ToLower(field.GongStruct.Name)

				fieldSliceOfPtrSortingToStructCodes += models.Replace6(
					NgFrontRepoSliceOfPointerToStructTmplCode[string(NgFrontRepoSliceOfPointerSorting)],
					"{{fieldName}}", fieldName,
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", assocStructName,
					"{{Structname}}", _struct.Name,
					"{{structname}}", structName)
			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field := field.(type) {
				case *models.SliceOfPointerToGongStructField:

					if field.GongStruct == _struct {

						fieldName := strings.ToLower(field.Name)
						assocStructName := strings.ToLower(field.GongStruct.Name)
						__structName := strings.ToLower(__struct.Name)

						fieldSliceOfPtrToStructCodes += models.Replace6(
							NgFrontRepoSliceOfPointerToStructTmplCode[string(NgFrontRepoSliceOfPointerRedeeming)],
							"{{fieldName}}", fieldName,
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocStructName}}", assocStructName,
							"{{Structname}}", __struct.Name,
							"{{structname}}", __structName)
					}
				}
			}
		}

		// compute code from sub template
		for subTemplate := range NgFrontRepoPerStructTmplCodes {
			perStructCodes[subTemplate] += models.Replace3(NgFrontRepoPerStructTmplCodes[subTemplate],
				"{{Structname}}", _struct.Name,
				"{{structname}}", structName,
				"{{Prime}}", strconv.Itoa(models.Primes[structIndex+10]))

			toReplace := "{{" + string(NgFrontRepoPtrToStructRedeeming) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldPointerToStructCodes)

			toReplace = "{{" + string(NgFrontRepoSliceOfPointerRedeeming) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldSliceOfPtrToStructCodes)

			toReplace = "{{" + string(NgFrontRepoSliceOfPointerSorting) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldSliceOfPtrSortingToStructCodes)
		}
	}

	//
	// final per struct replacement
	//
	for perStructTmplCode := range NgFrontRepoPerStructTmplCodes {
		toReplace := "{{" + string(perStructTmplCode) + "}}"
		code = strings.ReplaceAll(code, toReplace, perStructCodes[perStructTmplCode])
	}

	code = models.Replace4(code,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", strings.Title(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""))

	defer file.Close()
	fmt.Fprint(file, code)
}
