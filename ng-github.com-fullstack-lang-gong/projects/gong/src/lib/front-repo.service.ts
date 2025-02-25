import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { GongBasicFieldAPI } from './gongbasicfield-api'
import { GongBasicField, CopyGongBasicFieldAPIToGongBasicField } from './gongbasicfield'
import { GongBasicFieldService } from './gongbasicfield.service'

import { GongEnumAPI } from './gongenum-api'
import { GongEnum, CopyGongEnumAPIToGongEnum } from './gongenum'
import { GongEnumService } from './gongenum.service'

import { GongEnumValueAPI } from './gongenumvalue-api'
import { GongEnumValue, CopyGongEnumValueAPIToGongEnumValue } from './gongenumvalue'
import { GongEnumValueService } from './gongenumvalue.service'

import { GongLinkAPI } from './gonglink-api'
import { GongLink, CopyGongLinkAPIToGongLink } from './gonglink'
import { GongLinkService } from './gonglink.service'

import { GongNoteAPI } from './gongnote-api'
import { GongNote, CopyGongNoteAPIToGongNote } from './gongnote'
import { GongNoteService } from './gongnote.service'

import { GongStructAPI } from './gongstruct-api'
import { GongStruct, CopyGongStructAPIToGongStruct } from './gongstruct'
import { GongStructService } from './gongstruct.service'

import { GongTimeFieldAPI } from './gongtimefield-api'
import { GongTimeField, CopyGongTimeFieldAPIToGongTimeField } from './gongtimefield'
import { GongTimeFieldService } from './gongtimefield.service'

import { MetaAPI } from './meta-api'
import { Meta, CopyMetaAPIToMeta } from './meta'
import { MetaService } from './meta.service'

import { MetaReferenceAPI } from './metareference-api'
import { MetaReference, CopyMetaReferenceAPIToMetaReference } from './metareference'
import { MetaReferenceService } from './metareference.service'

import { ModelPkgAPI } from './modelpkg-api'
import { ModelPkg, CopyModelPkgAPIToModelPkg } from './modelpkg'
import { ModelPkgService } from './modelpkg.service'

import { PointerToGongStructFieldAPI } from './pointertogongstructfield-api'
import { PointerToGongStructField, CopyPointerToGongStructFieldAPIToPointerToGongStructField } from './pointertogongstructfield'
import { PointerToGongStructFieldService } from './pointertogongstructfield.service'

import { SliceOfPointerToGongStructFieldAPI } from './sliceofpointertogongstructfield-api'
import { SliceOfPointerToGongStructField, CopySliceOfPointerToGongStructFieldAPIToSliceOfPointerToGongStructField } from './sliceofpointertogongstructfield'
import { SliceOfPointerToGongStructFieldService } from './sliceofpointertogongstructfield.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gong/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_GongBasicFields = new Array<GongBasicField>() // array of front instances
	map_ID_GongBasicField = new Map<number, GongBasicField>() // map of front instances

	array_GongEnums = new Array<GongEnum>() // array of front instances
	map_ID_GongEnum = new Map<number, GongEnum>() // map of front instances

	array_GongEnumValues = new Array<GongEnumValue>() // array of front instances
	map_ID_GongEnumValue = new Map<number, GongEnumValue>() // map of front instances

	array_GongLinks = new Array<GongLink>() // array of front instances
	map_ID_GongLink = new Map<number, GongLink>() // map of front instances

	array_GongNotes = new Array<GongNote>() // array of front instances
	map_ID_GongNote = new Map<number, GongNote>() // map of front instances

	array_GongStructs = new Array<GongStruct>() // array of front instances
	map_ID_GongStruct = new Map<number, GongStruct>() // map of front instances

	array_GongTimeFields = new Array<GongTimeField>() // array of front instances
	map_ID_GongTimeField = new Map<number, GongTimeField>() // map of front instances

	array_Metas = new Array<Meta>() // array of front instances
	map_ID_Meta = new Map<number, Meta>() // map of front instances

	array_MetaReferences = new Array<MetaReference>() // array of front instances
	map_ID_MetaReference = new Map<number, MetaReference>() // map of front instances

	array_ModelPkgs = new Array<ModelPkg>() // array of front instances
	map_ID_ModelPkg = new Map<number, ModelPkg>() // map of front instances

	array_PointerToGongStructFields = new Array<PointerToGongStructField>() // array of front instances
	map_ID_PointerToGongStructField = new Map<number, PointerToGongStructField>() // map of front instances

	array_SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructField>() // array of front instances
	map_ID_SliceOfPointerToGongStructField = new Map<number, SliceOfPointerToGongStructField>() // map of front instances


	public GONG__Index = -1

	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'GongBasicField':
				return this.array_GongBasicFields as unknown as Array<Type>
			case 'GongEnum':
				return this.array_GongEnums as unknown as Array<Type>
			case 'GongEnumValue':
				return this.array_GongEnumValues as unknown as Array<Type>
			case 'GongLink':
				return this.array_GongLinks as unknown as Array<Type>
			case 'GongNote':
				return this.array_GongNotes as unknown as Array<Type>
			case 'GongStruct':
				return this.array_GongStructs as unknown as Array<Type>
			case 'GongTimeField':
				return this.array_GongTimeFields as unknown as Array<Type>
			case 'Meta':
				return this.array_Metas as unknown as Array<Type>
			case 'MetaReference':
				return this.array_MetaReferences as unknown as Array<Type>
			case 'ModelPkg':
				return this.array_ModelPkgs as unknown as Array<Type>
			case 'PointerToGongStructField':
				return this.array_PointerToGongStructFields as unknown as Array<Type>
			case 'SliceOfPointerToGongStructField':
				return this.array_SliceOfPointerToGongStructFields as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'GongBasicField':
				return this.map_ID_GongBasicField as unknown as Map<number, Type>
			case 'GongEnum':
				return this.map_ID_GongEnum as unknown as Map<number, Type>
			case 'GongEnumValue':
				return this.map_ID_GongEnumValue as unknown as Map<number, Type>
			case 'GongLink':
				return this.map_ID_GongLink as unknown as Map<number, Type>
			case 'GongNote':
				return this.map_ID_GongNote as unknown as Map<number, Type>
			case 'GongStruct':
				return this.map_ID_GongStruct as unknown as Map<number, Type>
			case 'GongTimeField':
				return this.map_ID_GongTimeField as unknown as Map<number, Type>
			case 'Meta':
				return this.map_ID_Meta as unknown as Map<number, Type>
			case 'MetaReference':
				return this.map_ID_MetaReference as unknown as Map<number, Type>
			case 'ModelPkg':
				return this.map_ID_ModelPkg as unknown as Map<number, Type>
			case 'PointerToGongStructField':
				return this.map_ID_PointerToGongStructField as unknown as Map<number, Type>
			case 'SliceOfPointerToGongStructField':
				return this.map_ID_SliceOfPointerToGongStructField as unknown as Map<number, Type>
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
		private http: HttpClient, // insertion point sub template 
		private gongbasicfieldService: GongBasicFieldService,
		private gongenumService: GongEnumService,
		private gongenumvalueService: GongEnumValueService,
		private gonglinkService: GongLinkService,
		private gongnoteService: GongNoteService,
		private gongstructService: GongStructService,
		private gongtimefieldService: GongTimeFieldService,
		private metaService: MetaService,
		private metareferenceService: MetaReferenceService,
		private modelpkgService: ModelPkgService,
		private pointertogongstructfieldService: PointerToGongStructFieldService,
		private sliceofpointertogongstructfieldService: SliceOfPointerToGongStructFieldService,
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
	observableFrontRepo!: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<GongBasicFieldAPI[]>,
		Observable<GongEnumAPI[]>,
		Observable<GongEnumValueAPI[]>,
		Observable<GongLinkAPI[]>,
		Observable<GongNoteAPI[]>,
		Observable<GongStructAPI[]>,
		Observable<GongTimeFieldAPI[]>,
		Observable<MetaAPI[]>,
		Observable<MetaReferenceAPI[]>,
		Observable<ModelPkgAPI[]>,
		Observable<PointerToGongStructFieldAPI[]>,
		Observable<SliceOfPointerToGongStructFieldAPI[]>,
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
			// insertion point sub template
			this.gongbasicfieldService.getGongBasicFields(this.GONG__StackPath, this.frontRepo),
			this.gongenumService.getGongEnums(this.GONG__StackPath, this.frontRepo),
			this.gongenumvalueService.getGongEnumValues(this.GONG__StackPath, this.frontRepo),
			this.gonglinkService.getGongLinks(this.GONG__StackPath, this.frontRepo),
			this.gongnoteService.getGongNotes(this.GONG__StackPath, this.frontRepo),
			this.gongstructService.getGongStructs(this.GONG__StackPath, this.frontRepo),
			this.gongtimefieldService.getGongTimeFields(this.GONG__StackPath, this.frontRepo),
			this.metaService.getMetas(this.GONG__StackPath, this.frontRepo),
			this.metareferenceService.getMetaReferences(this.GONG__StackPath, this.frontRepo),
			this.modelpkgService.getModelPkgs(this.GONG__StackPath, this.frontRepo),
			this.pointertogongstructfieldService.getPointerToGongStructFields(this.GONG__StackPath, this.frontRepo),
			this.sliceofpointertogongstructfieldService.getSliceOfPointerToGongStructFields(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						gongbasicfields_,
						gongenums_,
						gongenumvalues_,
						gonglinks_,
						gongnotes_,
						gongstructs_,
						gongtimefields_,
						metas_,
						metareferences_,
						modelpkgs_,
						pointertogongstructfields_,
						sliceofpointertogongstructfields_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var gongbasicfields: GongBasicFieldAPI[]
						gongbasicfields = gongbasicfields_ as GongBasicFieldAPI[]
						var gongenums: GongEnumAPI[]
						gongenums = gongenums_ as GongEnumAPI[]
						var gongenumvalues: GongEnumValueAPI[]
						gongenumvalues = gongenumvalues_ as GongEnumValueAPI[]
						var gonglinks: GongLinkAPI[]
						gonglinks = gonglinks_ as GongLinkAPI[]
						var gongnotes: GongNoteAPI[]
						gongnotes = gongnotes_ as GongNoteAPI[]
						var gongstructs: GongStructAPI[]
						gongstructs = gongstructs_ as GongStructAPI[]
						var gongtimefields: GongTimeFieldAPI[]
						gongtimefields = gongtimefields_ as GongTimeFieldAPI[]
						var metas: MetaAPI[]
						metas = metas_ as MetaAPI[]
						var metareferences: MetaReferenceAPI[]
						metareferences = metareferences_ as MetaReferenceAPI[]
						var modelpkgs: ModelPkgAPI[]
						modelpkgs = modelpkgs_ as ModelPkgAPI[]
						var pointertogongstructfields: PointerToGongStructFieldAPI[]
						pointertogongstructfields = pointertogongstructfields_ as PointerToGongStructFieldAPI[]
						var sliceofpointertogongstructfields: SliceOfPointerToGongStructFieldAPI[]
						sliceofpointertogongstructfields = sliceofpointertogongstructfields_ as SliceOfPointerToGongStructFieldAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_GongBasicFields = []
						this.frontRepo.map_ID_GongBasicField.clear()

						gongbasicfields.forEach(
							gongbasicfieldAPI => {
								let gongbasicfield = new GongBasicField
								this.frontRepo.array_GongBasicFields.push(gongbasicfield)
								this.frontRepo.map_ID_GongBasicField.set(gongbasicfieldAPI.ID, gongbasicfield)
							}
						)

						// init the arrays
						this.frontRepo.array_GongEnums = []
						this.frontRepo.map_ID_GongEnum.clear()

						gongenums.forEach(
							gongenumAPI => {
								let gongenum = new GongEnum
								this.frontRepo.array_GongEnums.push(gongenum)
								this.frontRepo.map_ID_GongEnum.set(gongenumAPI.ID, gongenum)
							}
						)

						// init the arrays
						this.frontRepo.array_GongEnumValues = []
						this.frontRepo.map_ID_GongEnumValue.clear()

						gongenumvalues.forEach(
							gongenumvalueAPI => {
								let gongenumvalue = new GongEnumValue
								this.frontRepo.array_GongEnumValues.push(gongenumvalue)
								this.frontRepo.map_ID_GongEnumValue.set(gongenumvalueAPI.ID, gongenumvalue)
							}
						)

						// init the arrays
						this.frontRepo.array_GongLinks = []
						this.frontRepo.map_ID_GongLink.clear()

						gonglinks.forEach(
							gonglinkAPI => {
								let gonglink = new GongLink
								this.frontRepo.array_GongLinks.push(gonglink)
								this.frontRepo.map_ID_GongLink.set(gonglinkAPI.ID, gonglink)
							}
						)

						// init the arrays
						this.frontRepo.array_GongNotes = []
						this.frontRepo.map_ID_GongNote.clear()

						gongnotes.forEach(
							gongnoteAPI => {
								let gongnote = new GongNote
								this.frontRepo.array_GongNotes.push(gongnote)
								this.frontRepo.map_ID_GongNote.set(gongnoteAPI.ID, gongnote)
							}
						)

						// init the arrays
						this.frontRepo.array_GongStructs = []
						this.frontRepo.map_ID_GongStruct.clear()

						gongstructs.forEach(
							gongstructAPI => {
								let gongstruct = new GongStruct
								this.frontRepo.array_GongStructs.push(gongstruct)
								this.frontRepo.map_ID_GongStruct.set(gongstructAPI.ID, gongstruct)
							}
						)

						// init the arrays
						this.frontRepo.array_GongTimeFields = []
						this.frontRepo.map_ID_GongTimeField.clear()

						gongtimefields.forEach(
							gongtimefieldAPI => {
								let gongtimefield = new GongTimeField
								this.frontRepo.array_GongTimeFields.push(gongtimefield)
								this.frontRepo.map_ID_GongTimeField.set(gongtimefieldAPI.ID, gongtimefield)
							}
						)

						// init the arrays
						this.frontRepo.array_Metas = []
						this.frontRepo.map_ID_Meta.clear()

						metas.forEach(
							metaAPI => {
								let meta = new Meta
								this.frontRepo.array_Metas.push(meta)
								this.frontRepo.map_ID_Meta.set(metaAPI.ID, meta)
							}
						)

						// init the arrays
						this.frontRepo.array_MetaReferences = []
						this.frontRepo.map_ID_MetaReference.clear()

						metareferences.forEach(
							metareferenceAPI => {
								let metareference = new MetaReference
								this.frontRepo.array_MetaReferences.push(metareference)
								this.frontRepo.map_ID_MetaReference.set(metareferenceAPI.ID, metareference)
							}
						)

						// init the arrays
						this.frontRepo.array_ModelPkgs = []
						this.frontRepo.map_ID_ModelPkg.clear()

						modelpkgs.forEach(
							modelpkgAPI => {
								let modelpkg = new ModelPkg
								this.frontRepo.array_ModelPkgs.push(modelpkg)
								this.frontRepo.map_ID_ModelPkg.set(modelpkgAPI.ID, modelpkg)
							}
						)

						// init the arrays
						this.frontRepo.array_PointerToGongStructFields = []
						this.frontRepo.map_ID_PointerToGongStructField.clear()

						pointertogongstructfields.forEach(
							pointertogongstructfieldAPI => {
								let pointertogongstructfield = new PointerToGongStructField
								this.frontRepo.array_PointerToGongStructFields.push(pointertogongstructfield)
								this.frontRepo.map_ID_PointerToGongStructField.set(pointertogongstructfieldAPI.ID, pointertogongstructfield)
							}
						)

						// init the arrays
						this.frontRepo.array_SliceOfPointerToGongStructFields = []
						this.frontRepo.map_ID_SliceOfPointerToGongStructField.clear()

						sliceofpointertogongstructfields.forEach(
							sliceofpointertogongstructfieldAPI => {
								let sliceofpointertogongstructfield = new SliceOfPointerToGongStructField
								this.frontRepo.array_SliceOfPointerToGongStructFields.push(sliceofpointertogongstructfield)
								this.frontRepo.map_ID_SliceOfPointerToGongStructField.set(sliceofpointertogongstructfieldAPI.ID, sliceofpointertogongstructfield)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						gongbasicfields.forEach(
							gongbasicfieldAPI => {
								let gongbasicfield = this.frontRepo.map_ID_GongBasicField.get(gongbasicfieldAPI.ID)
								CopyGongBasicFieldAPIToGongBasicField(gongbasicfieldAPI, gongbasicfield!, this.frontRepo)
							}
						)

						// fill up front objects
						gongenums.forEach(
							gongenumAPI => {
								let gongenum = this.frontRepo.map_ID_GongEnum.get(gongenumAPI.ID)
								CopyGongEnumAPIToGongEnum(gongenumAPI, gongenum!, this.frontRepo)
							}
						)

						// fill up front objects
						gongenumvalues.forEach(
							gongenumvalueAPI => {
								let gongenumvalue = this.frontRepo.map_ID_GongEnumValue.get(gongenumvalueAPI.ID)
								CopyGongEnumValueAPIToGongEnumValue(gongenumvalueAPI, gongenumvalue!, this.frontRepo)
							}
						)

						// fill up front objects
						gonglinks.forEach(
							gonglinkAPI => {
								let gonglink = this.frontRepo.map_ID_GongLink.get(gonglinkAPI.ID)
								CopyGongLinkAPIToGongLink(gonglinkAPI, gonglink!, this.frontRepo)
							}
						)

						// fill up front objects
						gongnotes.forEach(
							gongnoteAPI => {
								let gongnote = this.frontRepo.map_ID_GongNote.get(gongnoteAPI.ID)
								CopyGongNoteAPIToGongNote(gongnoteAPI, gongnote!, this.frontRepo)
							}
						)

						// fill up front objects
						gongstructs.forEach(
							gongstructAPI => {
								let gongstruct = this.frontRepo.map_ID_GongStruct.get(gongstructAPI.ID)
								CopyGongStructAPIToGongStruct(gongstructAPI, gongstruct!, this.frontRepo)
							}
						)

						// fill up front objects
						gongtimefields.forEach(
							gongtimefieldAPI => {
								let gongtimefield = this.frontRepo.map_ID_GongTimeField.get(gongtimefieldAPI.ID)
								CopyGongTimeFieldAPIToGongTimeField(gongtimefieldAPI, gongtimefield!, this.frontRepo)
							}
						)

						// fill up front objects
						metas.forEach(
							metaAPI => {
								let meta = this.frontRepo.map_ID_Meta.get(metaAPI.ID)
								CopyMetaAPIToMeta(metaAPI, meta!, this.frontRepo)
							}
						)

						// fill up front objects
						metareferences.forEach(
							metareferenceAPI => {
								let metareference = this.frontRepo.map_ID_MetaReference.get(metareferenceAPI.ID)
								CopyMetaReferenceAPIToMetaReference(metareferenceAPI, metareference!, this.frontRepo)
							}
						)

						// fill up front objects
						modelpkgs.forEach(
							modelpkgAPI => {
								let modelpkg = this.frontRepo.map_ID_ModelPkg.get(modelpkgAPI.ID)
								CopyModelPkgAPIToModelPkg(modelpkgAPI, modelpkg!, this.frontRepo)
							}
						)

						// fill up front objects
						pointertogongstructfields.forEach(
							pointertogongstructfieldAPI => {
								let pointertogongstructfield = this.frontRepo.map_ID_PointerToGongStructField.get(pointertogongstructfieldAPI.ID)
								CopyPointerToGongStructFieldAPIToPointerToGongStructField(pointertogongstructfieldAPI, pointertogongstructfield!, this.frontRepo)
							}
						)

						// fill up front objects
						sliceofpointertogongstructfields.forEach(
							sliceofpointertogongstructfieldAPI => {
								let sliceofpointertogongstructfield = this.frontRepo.map_ID_SliceOfPointerToGongStructField.get(sliceofpointertogongstructfieldAPI.ID)
								CopySliceOfPointerToGongStructFieldAPIToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldAPI, sliceofpointertogongstructfield!, this.frontRepo)
							}
						)


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
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gong/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {


				const backRepoData = new BackRepoData(JSON.parse(event.data))

				let frontRepo = new (FrontRepo)
				frontRepo.GONG__Index = backRepoData.GONG__Index

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				frontRepo.array_GongBasicFields = []
				frontRepo.map_ID_GongBasicField.clear()

				backRepoData.GongBasicFieldAPIs.forEach(
					gongbasicfieldAPI => {
						let gongbasicfield = new GongBasicField
						frontRepo.array_GongBasicFields.push(gongbasicfield)
						frontRepo.map_ID_GongBasicField.set(gongbasicfieldAPI.ID, gongbasicfield)
					}
				)

				// init the arrays
				frontRepo.array_GongEnums = []
				frontRepo.map_ID_GongEnum.clear()

				backRepoData.GongEnumAPIs.forEach(
					gongenumAPI => {
						let gongenum = new GongEnum
						frontRepo.array_GongEnums.push(gongenum)
						frontRepo.map_ID_GongEnum.set(gongenumAPI.ID, gongenum)
					}
				)

				// init the arrays
				frontRepo.array_GongEnumValues = []
				frontRepo.map_ID_GongEnumValue.clear()

				backRepoData.GongEnumValueAPIs.forEach(
					gongenumvalueAPI => {
						let gongenumvalue = new GongEnumValue
						frontRepo.array_GongEnumValues.push(gongenumvalue)
						frontRepo.map_ID_GongEnumValue.set(gongenumvalueAPI.ID, gongenumvalue)
					}
				)

				// init the arrays
				frontRepo.array_GongLinks = []
				frontRepo.map_ID_GongLink.clear()

				backRepoData.GongLinkAPIs.forEach(
					gonglinkAPI => {
						let gonglink = new GongLink
						frontRepo.array_GongLinks.push(gonglink)
						frontRepo.map_ID_GongLink.set(gonglinkAPI.ID, gonglink)
					}
				)

				// init the arrays
				frontRepo.array_GongNotes = []
				frontRepo.map_ID_GongNote.clear()

				backRepoData.GongNoteAPIs.forEach(
					gongnoteAPI => {
						let gongnote = new GongNote
						frontRepo.array_GongNotes.push(gongnote)
						frontRepo.map_ID_GongNote.set(gongnoteAPI.ID, gongnote)
					}
				)

				// init the arrays
				frontRepo.array_GongStructs = []
				frontRepo.map_ID_GongStruct.clear()

				backRepoData.GongStructAPIs.forEach(
					gongstructAPI => {
						let gongstruct = new GongStruct
						frontRepo.array_GongStructs.push(gongstruct)
						frontRepo.map_ID_GongStruct.set(gongstructAPI.ID, gongstruct)
					}
				)

				// init the arrays
				frontRepo.array_GongTimeFields = []
				frontRepo.map_ID_GongTimeField.clear()

				backRepoData.GongTimeFieldAPIs.forEach(
					gongtimefieldAPI => {
						let gongtimefield = new GongTimeField
						frontRepo.array_GongTimeFields.push(gongtimefield)
						frontRepo.map_ID_GongTimeField.set(gongtimefieldAPI.ID, gongtimefield)
					}
				)

				// init the arrays
				frontRepo.array_Metas = []
				frontRepo.map_ID_Meta.clear()

				backRepoData.MetaAPIs.forEach(
					metaAPI => {
						let meta = new Meta
						frontRepo.array_Metas.push(meta)
						frontRepo.map_ID_Meta.set(metaAPI.ID, meta)
					}
				)

				// init the arrays
				frontRepo.array_MetaReferences = []
				frontRepo.map_ID_MetaReference.clear()

				backRepoData.MetaReferenceAPIs.forEach(
					metareferenceAPI => {
						let metareference = new MetaReference
						frontRepo.array_MetaReferences.push(metareference)
						frontRepo.map_ID_MetaReference.set(metareferenceAPI.ID, metareference)
					}
				)

				// init the arrays
				frontRepo.array_ModelPkgs = []
				frontRepo.map_ID_ModelPkg.clear()

				backRepoData.ModelPkgAPIs.forEach(
					modelpkgAPI => {
						let modelpkg = new ModelPkg
						frontRepo.array_ModelPkgs.push(modelpkg)
						frontRepo.map_ID_ModelPkg.set(modelpkgAPI.ID, modelpkg)
					}
				)

				// init the arrays
				frontRepo.array_PointerToGongStructFields = []
				frontRepo.map_ID_PointerToGongStructField.clear()

				backRepoData.PointerToGongStructFieldAPIs.forEach(
					pointertogongstructfieldAPI => {
						let pointertogongstructfield = new PointerToGongStructField
						frontRepo.array_PointerToGongStructFields.push(pointertogongstructfield)
						frontRepo.map_ID_PointerToGongStructField.set(pointertogongstructfieldAPI.ID, pointertogongstructfield)
					}
				)

				// init the arrays
				frontRepo.array_SliceOfPointerToGongStructFields = []
				frontRepo.map_ID_SliceOfPointerToGongStructField.clear()

				backRepoData.SliceOfPointerToGongStructFieldAPIs.forEach(
					sliceofpointertogongstructfieldAPI => {
						let sliceofpointertogongstructfield = new SliceOfPointerToGongStructField
						frontRepo.array_SliceOfPointerToGongStructFields.push(sliceofpointertogongstructfield)
						frontRepo.map_ID_SliceOfPointerToGongStructField.set(sliceofpointertogongstructfieldAPI.ID, sliceofpointertogongstructfield)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.GongBasicFieldAPIs.forEach(
					gongbasicfieldAPI => {
						let gongbasicfield = frontRepo.map_ID_GongBasicField.get(gongbasicfieldAPI.ID)
						CopyGongBasicFieldAPIToGongBasicField(gongbasicfieldAPI, gongbasicfield!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongEnumAPIs.forEach(
					gongenumAPI => {
						let gongenum = frontRepo.map_ID_GongEnum.get(gongenumAPI.ID)
						CopyGongEnumAPIToGongEnum(gongenumAPI, gongenum!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongEnumValueAPIs.forEach(
					gongenumvalueAPI => {
						let gongenumvalue = frontRepo.map_ID_GongEnumValue.get(gongenumvalueAPI.ID)
						CopyGongEnumValueAPIToGongEnumValue(gongenumvalueAPI, gongenumvalue!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongLinkAPIs.forEach(
					gonglinkAPI => {
						let gonglink = frontRepo.map_ID_GongLink.get(gonglinkAPI.ID)
						CopyGongLinkAPIToGongLink(gonglinkAPI, gonglink!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongNoteAPIs.forEach(
					gongnoteAPI => {
						let gongnote = frontRepo.map_ID_GongNote.get(gongnoteAPI.ID)
						CopyGongNoteAPIToGongNote(gongnoteAPI, gongnote!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongStructAPIs.forEach(
					gongstructAPI => {
						let gongstruct = frontRepo.map_ID_GongStruct.get(gongstructAPI.ID)
						CopyGongStructAPIToGongStruct(gongstructAPI, gongstruct!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongTimeFieldAPIs.forEach(
					gongtimefieldAPI => {
						let gongtimefield = frontRepo.map_ID_GongTimeField.get(gongtimefieldAPI.ID)
						CopyGongTimeFieldAPIToGongTimeField(gongtimefieldAPI, gongtimefield!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MetaAPIs.forEach(
					metaAPI => {
						let meta = frontRepo.map_ID_Meta.get(metaAPI.ID)
						CopyMetaAPIToMeta(metaAPI, meta!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.MetaReferenceAPIs.forEach(
					metareferenceAPI => {
						let metareference = frontRepo.map_ID_MetaReference.get(metareferenceAPI.ID)
						CopyMetaReferenceAPIToMetaReference(metareferenceAPI, metareference!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.ModelPkgAPIs.forEach(
					modelpkgAPI => {
						let modelpkg = frontRepo.map_ID_ModelPkg.get(modelpkgAPI.ID)
						CopyModelPkgAPIToModelPkg(modelpkgAPI, modelpkg!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.PointerToGongStructFieldAPIs.forEach(
					pointertogongstructfieldAPI => {
						let pointertogongstructfield = frontRepo.map_ID_PointerToGongStructField.get(pointertogongstructfieldAPI.ID)
						CopyPointerToGongStructFieldAPIToPointerToGongStructField(pointertogongstructfieldAPI, pointertogongstructfield!, frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SliceOfPointerToGongStructFieldAPIs.forEach(
					sliceofpointertogongstructfieldAPI => {
						let sliceofpointertogongstructfield = frontRepo.map_ID_SliceOfPointerToGongStructField.get(sliceofpointertogongstructfieldAPI.ID)
						CopySliceOfPointerToGongStructFieldAPIToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldAPI, sliceofpointertogongstructfield!, frontRepo)
					}
				)



				observer.next(frontRepo)
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

// insertion point for get unique ID per struct 
export function getGongBasicFieldUniqueID(id: number): number {
	return 31 * id
}
export function getGongEnumUniqueID(id: number): number {
	return 37 * id
}
export function getGongEnumValueUniqueID(id: number): number {
	return 41 * id
}
export function getGongLinkUniqueID(id: number): number {
	return 43 * id
}
export function getGongNoteUniqueID(id: number): number {
	return 47 * id
}
export function getGongStructUniqueID(id: number): number {
	return 53 * id
}
export function getGongTimeFieldUniqueID(id: number): number {
	return 59 * id
}
export function getMetaUniqueID(id: number): number {
	return 61 * id
}
export function getMetaReferenceUniqueID(id: number): number {
	return 67 * id
}
export function getModelPkgUniqueID(id: number): number {
	return 71 * id
}
export function getPointerToGongStructFieldUniqueID(id: number): number {
	return 73 * id
}
export function getSliceOfPointerToGongStructFieldUniqueID(id: number): number {
	return 79 * id
}
