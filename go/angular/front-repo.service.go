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
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports{{` + string(rune(NgLibFrontRepoServiceImports)) + `}}

import { BackRepoData } from './back-repo-data'

export const StackType = "{{PkgPathRoot}}/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template{{` + string(rune(NgLibFrontRepoMapDecl)) + `}}

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point{{` + string(rune(NgLibFrontRepoSwitchGetFrontArray)) + `}}
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point{{` + string(rune(NgLibFrontRepoSwitchGetFrontMap)) + `}}
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
	SourceStruct: string = ""	// The "Aclass"
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
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template {{` + string(rune(NgLibFrontRepoServiceDecl)) + `}}
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
		// insertion point sub template {{` + string(rune(NgLibFrontRepoObservableArrayType)) + `}}
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template{{` + string(rune(NgLibFrontRepoObservableRefs)) + `}}
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
			// insertion point sub template{{` + string(rune(NgLibFrontRepoObservableRefs)) + `}}
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations {{` + string(rune(NgLibFrontRepoArraysDecls)) + `}}
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting {{` + string(rune(NgLibFrontRepoTypeCasting)) + `}}

						// 
						// First Step: init map of instances
						// insertion point sub template for init {{` + string(rune(NgLibFrontRepoInitMapInstances)) + `}}

						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem {{` + string(rune(NgLibFrontRepoInitFrontObjects)) + `}}

						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/{{PkgPathRoot}}/v1/ws/stage'
		let paramString = params.toString()
		let url = ` + "`" + "${basePath}?${paramString}" + "`" + `
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init {{` + string(rune(NgLibFrontRepoInitMapInstancesFromWebSocket)) + `}}

				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem {{` + string(rune(NgLibFrontRepoInitFrontObjectsFromWebSocket)) + `}}


				observer.next(this.frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct {{` + string(rune(NgLibFrontRepoPerStructGetUniqueID)) + `}}
`

type FrontRepoInsertionPointId int

const (
	NgLibFrontRepoServiceImports FrontRepoInsertionPointId = iota
	NgLibFrontRepoMapDecl
	NgLibFrontRepoSwitchGetFrontArray
	NgLibFrontRepoSwitchGetFrontMap
	NgLibFrontRepoObservableArrayType
	NgLibFrontRepoTypeCasting
	NgLibFrontRepoServiceDecl
	NgLibFrontRepoObservableRefs
	NgLibFrontRepoArraysDecls
	NgLibFrontRepoInitMapInstances
	NgLibFrontRepoInitMapInstancesFromWebSocket
	NgLibFrontRepoRedeemPointers
	NgLibFrontRepoInitFrontObjects
	NgLibFrontRepoInitFrontObjectsFromWebSocket
	NgLibFrontRepoSlicesOfPointersDecode
	NgLibFrontRepoPerStructGetUniqueID
)

var NgFrontRepoPerStructTmplCodes map[FrontRepoInsertionPointId]string = // new line
map[FrontRepoInsertionPointId]string{

	NgLibFrontRepoServiceImports: `
import { {{Structname}}API } from './{{structname}}-api'
import { {{Structname}}, Copy{{Structname}}APITo{{Structname}} } from './{{structname}}'
import { {{Structname}}Service } from './{{structname}}.service'
`,

	NgLibFrontRepoMapDecl: `
	array_{{Structname}}s = new Array<{{Structname}}>() // array of front instances
	map_ID_{{Structname}} = new Map<number, {{Structname}}>() // map of front instances
`,

	NgLibFrontRepoSwitchGetFrontArray: `
			case '{{Structname}}':
				return this.array_{{Structname}}s as unknown as Array<Type>`,

	NgLibFrontRepoSwitchGetFrontMap: `
			case '{{Structname}}':
				return this.map_ID_{{Structname}} as unknown as Map<number, Type>`,

	NgLibFrontRepoObservableArrayType: `
		Observable<{{Structname}}API[]>,`,

	NgLibFrontRepoServiceDecl: `
		private {{structname}}Service: {{Structname}}Service,`,

	NgLibFrontRepoObservableRefs: `
			this.{{structname}}Service.get{{Structname}}s(this.GONG__StackPath, this.frontRepo),`,

	NgLibFrontRepoArraysDecls: `
						{{structname}}s_,`,

	NgLibFrontRepoTypeCasting: `
						var {{structname}}s: {{Structname}}API[]
						{{structname}}s = {{structname}}s_ as {{Structname}}API[]`,

	NgLibFrontRepoInitMapInstances: `
						// init the arrays
						this.frontRepo.array_{{Structname}}s = []
						this.frontRepo.map_ID_{{Structname}}.clear()

						{{structname}}s.forEach(
							{{structname}}API => {
								let {{structname}} = new {{Structname}}
								this.frontRepo.array_{{Structname}}s.push({{structname}})
								this.frontRepo.map_ID_{{Structname}}.set({{structname}}API.ID, {{structname}})
							}
						)
`,

	NgLibFrontRepoInitMapInstancesFromWebSocket: `
				// init the arrays
				frontRepo.array_{{Structname}}s = []
				frontRepo.map_ID_{{Structname}}.clear()

				backRepoData.{{Structname}}APIs.forEach(
					{{structname}}API => {
						let {{structname}} = new {{Structname}}
						frontRepo.array_{{Structname}}s.push({{structname}})
						frontRepo.map_ID_{{Structname}}.set({{structname}}API.ID, {{structname}})
					}
				)
`,

	NgLibFrontRepoInitFrontObjects: `
						// fill up front objects
						{{structname}}s.forEach(
							{{structname}}API => {
								let {{structname}} = this.frontRepo.map_ID_{{Structname}}.get({{structname}}API.ID)
								Copy{{Structname}}APITo{{Structname}}({{structname}}API, {{structname}}!, this.frontRepo)
							}
						)
`,

	NgLibFrontRepoInitFrontObjectsFromWebSocket: `
				// fill up front objects
				backRepoData.{{Structname}}APIs.forEach(
					{{structname}}API => {
						let {{structname}} = frontRepo.map_ID_{{Structname}}.get({{structname}}API.ID)
						Copy{{Structname}}APITo{{Structname}}({{structname}}API, {{structname}}!, frontRepo)
					}
				)
`,

	NgLibFrontRepoRedeemPointers: `
						{{structname}}s.forEach(
							{{structname}} => {
								// insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming{{` + string(rune(NgFrontRepoPtrToStructRedeeming)) + `}}
							}
						)`,

	NgLibFrontRepoSlicesOfPointersDecode: `
						{{structname}}s.forEach(
							{{structname}} => {
								// insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming{{` + string(rune(NgFrontRepoPtrToStructRedeeming)) + `}}
								// insertion point for pointers decoding{{` + string(rune(NgFrontRepoSliceOfPointerSorting)) + `}}
							}
						)`,

	NgLibFrontRepoPerStructGetUniqueID: `
export function get{{Structname}}UniqueID(id: number): number {
	return {{Prime}} * id
}`,
}

type NgLibFrontRepoServiceSubSubTemplate int

const (
	NgFrontRepoPtrToStructRedeeming NgLibFrontRepoServiceSubSubTemplate = iota + 1789
	NgFrontRepoSliceOfPointerSorting
)

// for each sub sub template, what sub template it relates to
var NgLibFrontRepoPointerToStructSubSubToSubMap map[NgLibFrontRepoServiceSubSubTemplate]string = //
map[NgLibFrontRepoServiceSubSubTemplate]string{
	NgFrontRepoPtrToStructRedeeming: string(rune(NgLibFrontRepoRedeemPointers)),
}

var NgFrontRepoPtrToStructTmplCodes map[NgLibFrontRepoServiceSubSubTemplate]string = // new line
map[NgLibFrontRepoServiceSubSubTemplate]string{
	NgFrontRepoPtrToStructRedeeming: `
								// insertion point for pointer field {{FieldName}} redeeming
								{
									let _{{assocStructName}} = this.frontRepo.{{AssocStructName}}s.get({{structname}}.{{Structname}}PointersEncoding.{{FieldName}}ID.Int64)
									if (_{{assocStructName}}) {
										{{structname}}.{{FieldName}} = _{{assocStructName}}
									}
								}`,
	NgFrontRepoSliceOfPointerSorting: `
								{{structname}}.{{FieldName}} = new Array<{{AssocStructName}}API>()
								for (let _id of {{structname}}.{{Structname}}PointersEncoding.{{FieldName}}) {
									let _{{assocStructName}} = this.frontRepo.{{AssocStructName}}s.get(_id)
									if (_{{assocStructName}} != undefined) {
										{{structname}}.{{FieldName}}.push(_{{assocStructName}}!)
									}
								}`,
}

func CodeGeneratorNgFrontRepo(modelPkg *models.ModelPkg) {

	pkgName := modelPkg.Name
	matTargetPath := modelPkg.NgDataLibrarySourceCodeDirectory
	pkgGoPath := modelPkg.PkgPath

	file, err := os.Create(filepath.Join(matTargetPath, "front-repo.service.ts"))
	if err != nil {
		log.Panic(err)
	}

	code := NgLibFrontRepoServiceTemplate

	perStructCodes := make(map[FrontRepoInsertionPointId]string)
	for subTemplate := range NgFrontRepoPerStructTmplCodes {
		perStructCodes[subTemplate] = ""
	}

	// have alphabetical order generation
	structList := []*models.GongStruct{}
	for _, _struct := range modelPkg.GongStructs {
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
		fieldSliceOfPtrSortingToStructCodes := ""

		// compute code per field
		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *models.PointerToGongStructField:

				fieldName := strings.ToLower(field.Name)
				assocStructName := strings.ToLower(field.GongStruct.Name)

				fieldPointerToStructCodes += models.Replace6(
					NgFrontRepoPtrToStructTmplCodes[NgFrontRepoPtrToStructRedeeming],
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
					NgFrontRepoPtrToStructTmplCodes[NgFrontRepoSliceOfPointerSorting],
					"{{fieldName}}", fieldName,
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", assocStructName,
					"{{Structname}}", _struct.Name,
					"{{structname}}", structName)
			}
		}

		// compute code from sub template
		for subTemplate := range NgFrontRepoPerStructTmplCodes {
			perStructCodes[subTemplate] += models.Replace3(NgFrontRepoPerStructTmplCodes[subTemplate],
				"{{Structname}}", _struct.Name,
				"{{structname}}", structName,
				"{{Prime}}", strconv.Itoa(models.Primes[structIndex+10]))

			toReplace := "{{" + string(rune(NgFrontRepoPtrToStructRedeeming)) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldPointerToStructCodes)

			toReplace = "{{" + string(rune(NgFrontRepoSliceOfPointerSorting)) + "}}"
			perStructCodes[subTemplate] = strings.ReplaceAll(
				perStructCodes[subTemplate],
				toReplace, fieldSliceOfPtrSortingToStructCodes)
		}
	}

	//
	// final per struct replacement
	//
	for perStructTmplCode := range NgFrontRepoPerStructTmplCodes {
		toReplace := "{{" + string(rune(perStructTmplCode)) + "}}"
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
