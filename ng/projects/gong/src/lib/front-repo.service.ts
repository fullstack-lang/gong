import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { GongBasicFieldDB } from './gongbasicfield-db'
import { GongBasicField, CopyGongBasicFieldDBToGongBasicField } from './gongbasicfield'
import { GongBasicFieldService } from './gongbasicfield.service'

import { GongEnumDB } from './gongenum-db'
import { GongEnum, CopyGongEnumDBToGongEnum } from './gongenum'
import { GongEnumService } from './gongenum.service'

import { GongEnumValueDB } from './gongenumvalue-db'
import { GongEnumValue, CopyGongEnumValueDBToGongEnumValue } from './gongenumvalue'
import { GongEnumValueService } from './gongenumvalue.service'

import { GongLinkDB } from './gonglink-db'
import { GongLink, CopyGongLinkDBToGongLink } from './gonglink'
import { GongLinkService } from './gonglink.service'

import { GongNoteDB } from './gongnote-db'
import { GongNote, CopyGongNoteDBToGongNote } from './gongnote'
import { GongNoteService } from './gongnote.service'

import { GongStructDB } from './gongstruct-db'
import { GongStruct, CopyGongStructDBToGongStruct } from './gongstruct'
import { GongStructService } from './gongstruct.service'

import { GongTimeFieldDB } from './gongtimefield-db'
import { GongTimeField, CopyGongTimeFieldDBToGongTimeField } from './gongtimefield'
import { GongTimeFieldService } from './gongtimefield.service'

import { MetaDB } from './meta-db'
import { Meta, CopyMetaDBToMeta } from './meta'
import { MetaService } from './meta.service'

import { MetaReferenceDB } from './metareference-db'
import { MetaReference, CopyMetaReferenceDBToMetaReference } from './metareference'
import { MetaReferenceService } from './metareference.service'

import { ModelPkgDB } from './modelpkg-db'
import { ModelPkg, CopyModelPkgDBToModelPkg } from './modelpkg'
import { ModelPkgService } from './modelpkg.service'

import { PointerToGongStructFieldDB } from './pointertogongstructfield-db'
import { PointerToGongStructField, CopyPointerToGongStructFieldDBToPointerToGongStructField } from './pointertogongstructfield'
import { PointerToGongStructFieldService } from './pointertogongstructfield.service'

import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructField, CopySliceOfPointerToGongStructFieldDBToSliceOfPointerToGongStructField } from './sliceofpointertogongstructfield'
import { SliceOfPointerToGongStructFieldService } from './sliceofpointertogongstructfield.service'

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
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<GongBasicFieldDB[]>,
		Observable<GongEnumDB[]>,
		Observable<GongEnumValueDB[]>,
		Observable<GongLinkDB[]>,
		Observable<GongNoteDB[]>,
		Observable<GongStructDB[]>,
		Observable<GongTimeFieldDB[]>,
		Observable<MetaDB[]>,
		Observable<MetaReferenceDB[]>,
		Observable<ModelPkgDB[]>,
		Observable<PointerToGongStructFieldDB[]>,
		Observable<SliceOfPointerToGongStructFieldDB[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
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
						var gongbasicfields: GongBasicFieldDB[]
						gongbasicfields = gongbasicfields_ as GongBasicFieldDB[]
						var gongenums: GongEnumDB[]
						gongenums = gongenums_ as GongEnumDB[]
						var gongenumvalues: GongEnumValueDB[]
						gongenumvalues = gongenumvalues_ as GongEnumValueDB[]
						var gonglinks: GongLinkDB[]
						gonglinks = gonglinks_ as GongLinkDB[]
						var gongnotes: GongNoteDB[]
						gongnotes = gongnotes_ as GongNoteDB[]
						var gongstructs: GongStructDB[]
						gongstructs = gongstructs_ as GongStructDB[]
						var gongtimefields: GongTimeFieldDB[]
						gongtimefields = gongtimefields_ as GongTimeFieldDB[]
						var metas: MetaDB[]
						metas = metas_ as MetaDB[]
						var metareferences: MetaReferenceDB[]
						metareferences = metareferences_ as MetaReferenceDB[]
						var modelpkgs: ModelPkgDB[]
						modelpkgs = modelpkgs_ as ModelPkgDB[]
						var pointertogongstructfields: PointerToGongStructFieldDB[]
						pointertogongstructfields = pointertogongstructfields_ as PointerToGongStructFieldDB[]
						var sliceofpointertogongstructfields: SliceOfPointerToGongStructFieldDB[]
						sliceofpointertogongstructfields = sliceofpointertogongstructfields_ as SliceOfPointerToGongStructFieldDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_GongBasicFields = []
						this.frontRepo.map_ID_GongBasicField.clear()

						gongbasicfields.forEach(
							gongbasicfieldDB => {
								let gongbasicfield = new GongBasicField
								this.frontRepo.array_GongBasicFields.push(gongbasicfield)
								this.frontRepo.map_ID_GongBasicField.set(gongbasicfieldDB.ID, gongbasicfield)
							}
						)

						// init the arrays
						this.frontRepo.array_GongEnums = []
						this.frontRepo.map_ID_GongEnum.clear()

						gongenums.forEach(
							gongenumDB => {
								let gongenum = new GongEnum
								this.frontRepo.array_GongEnums.push(gongenum)
								this.frontRepo.map_ID_GongEnum.set(gongenumDB.ID, gongenum)
							}
						)

						// init the arrays
						this.frontRepo.array_GongEnumValues = []
						this.frontRepo.map_ID_GongEnumValue.clear()

						gongenumvalues.forEach(
							gongenumvalueDB => {
								let gongenumvalue = new GongEnumValue
								this.frontRepo.array_GongEnumValues.push(gongenumvalue)
								this.frontRepo.map_ID_GongEnumValue.set(gongenumvalueDB.ID, gongenumvalue)
							}
						)

						// init the arrays
						this.frontRepo.array_GongLinks = []
						this.frontRepo.map_ID_GongLink.clear()

						gonglinks.forEach(
							gonglinkDB => {
								let gonglink = new GongLink
								this.frontRepo.array_GongLinks.push(gonglink)
								this.frontRepo.map_ID_GongLink.set(gonglinkDB.ID, gonglink)
							}
						)

						// init the arrays
						this.frontRepo.array_GongNotes = []
						this.frontRepo.map_ID_GongNote.clear()

						gongnotes.forEach(
							gongnoteDB => {
								let gongnote = new GongNote
								this.frontRepo.array_GongNotes.push(gongnote)
								this.frontRepo.map_ID_GongNote.set(gongnoteDB.ID, gongnote)
							}
						)

						// init the arrays
						this.frontRepo.array_GongStructs = []
						this.frontRepo.map_ID_GongStruct.clear()

						gongstructs.forEach(
							gongstructDB => {
								let gongstruct = new GongStruct
								this.frontRepo.array_GongStructs.push(gongstruct)
								this.frontRepo.map_ID_GongStruct.set(gongstructDB.ID, gongstruct)
							}
						)

						// init the arrays
						this.frontRepo.array_GongTimeFields = []
						this.frontRepo.map_ID_GongTimeField.clear()

						gongtimefields.forEach(
							gongtimefieldDB => {
								let gongtimefield = new GongTimeField
								this.frontRepo.array_GongTimeFields.push(gongtimefield)
								this.frontRepo.map_ID_GongTimeField.set(gongtimefieldDB.ID, gongtimefield)
							}
						)

						// init the arrays
						this.frontRepo.array_Metas = []
						this.frontRepo.map_ID_Meta.clear()

						metas.forEach(
							metaDB => {
								let meta = new Meta
								this.frontRepo.array_Metas.push(meta)
								this.frontRepo.map_ID_Meta.set(metaDB.ID, meta)
							}
						)

						// init the arrays
						this.frontRepo.array_MetaReferences = []
						this.frontRepo.map_ID_MetaReference.clear()

						metareferences.forEach(
							metareferenceDB => {
								let metareference = new MetaReference
								this.frontRepo.array_MetaReferences.push(metareference)
								this.frontRepo.map_ID_MetaReference.set(metareferenceDB.ID, metareference)
							}
						)

						// init the arrays
						this.frontRepo.array_ModelPkgs = []
						this.frontRepo.map_ID_ModelPkg.clear()

						modelpkgs.forEach(
							modelpkgDB => {
								let modelpkg = new ModelPkg
								this.frontRepo.array_ModelPkgs.push(modelpkg)
								this.frontRepo.map_ID_ModelPkg.set(modelpkgDB.ID, modelpkg)
							}
						)

						// init the arrays
						this.frontRepo.array_PointerToGongStructFields = []
						this.frontRepo.map_ID_PointerToGongStructField.clear()

						pointertogongstructfields.forEach(
							pointertogongstructfieldDB => {
								let pointertogongstructfield = new PointerToGongStructField
								this.frontRepo.array_PointerToGongStructFields.push(pointertogongstructfield)
								this.frontRepo.map_ID_PointerToGongStructField.set(pointertogongstructfieldDB.ID, pointertogongstructfield)
							}
						)

						// init the arrays
						this.frontRepo.array_SliceOfPointerToGongStructFields = []
						this.frontRepo.map_ID_SliceOfPointerToGongStructField.clear()

						sliceofpointertogongstructfields.forEach(
							sliceofpointertogongstructfieldDB => {
								let sliceofpointertogongstructfield = new SliceOfPointerToGongStructField
								this.frontRepo.array_SliceOfPointerToGongStructFields.push(sliceofpointertogongstructfield)
								this.frontRepo.map_ID_SliceOfPointerToGongStructField.set(sliceofpointertogongstructfieldDB.ID, sliceofpointertogongstructfield)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						gongbasicfields.forEach(
							gongbasicfieldDB => {
								let gongbasicfield = this.frontRepo.map_ID_GongBasicField.get(gongbasicfieldDB.ID)
								CopyGongBasicFieldDBToGongBasicField(gongbasicfieldDB, gongbasicfield!, this.frontRepo)
							}
						)

						// fill up front objects
						gongenums.forEach(
							gongenumDB => {
								let gongenum = this.frontRepo.map_ID_GongEnum.get(gongenumDB.ID)
								CopyGongEnumDBToGongEnum(gongenumDB, gongenum!, this.frontRepo)
							}
						)

						// fill up front objects
						gongenumvalues.forEach(
							gongenumvalueDB => {
								let gongenumvalue = this.frontRepo.map_ID_GongEnumValue.get(gongenumvalueDB.ID)
								CopyGongEnumValueDBToGongEnumValue(gongenumvalueDB, gongenumvalue!, this.frontRepo)
							}
						)

						// fill up front objects
						gonglinks.forEach(
							gonglinkDB => {
								let gonglink = this.frontRepo.map_ID_GongLink.get(gonglinkDB.ID)
								CopyGongLinkDBToGongLink(gonglinkDB, gonglink!, this.frontRepo)
							}
						)

						// fill up front objects
						gongnotes.forEach(
							gongnoteDB => {
								let gongnote = this.frontRepo.map_ID_GongNote.get(gongnoteDB.ID)
								CopyGongNoteDBToGongNote(gongnoteDB, gongnote!, this.frontRepo)
							}
						)

						// fill up front objects
						gongstructs.forEach(
							gongstructDB => {
								let gongstruct = this.frontRepo.map_ID_GongStruct.get(gongstructDB.ID)
								CopyGongStructDBToGongStruct(gongstructDB, gongstruct!, this.frontRepo)
							}
						)

						// fill up front objects
						gongtimefields.forEach(
							gongtimefieldDB => {
								let gongtimefield = this.frontRepo.map_ID_GongTimeField.get(gongtimefieldDB.ID)
								CopyGongTimeFieldDBToGongTimeField(gongtimefieldDB, gongtimefield!, this.frontRepo)
							}
						)

						// fill up front objects
						metas.forEach(
							metaDB => {
								let meta = this.frontRepo.map_ID_Meta.get(metaDB.ID)
								CopyMetaDBToMeta(metaDB, meta!, this.frontRepo)
							}
						)

						// fill up front objects
						metareferences.forEach(
							metareferenceDB => {
								let metareference = this.frontRepo.map_ID_MetaReference.get(metareferenceDB.ID)
								CopyMetaReferenceDBToMetaReference(metareferenceDB, metareference!, this.frontRepo)
							}
						)

						// fill up front objects
						modelpkgs.forEach(
							modelpkgDB => {
								let modelpkg = this.frontRepo.map_ID_ModelPkg.get(modelpkgDB.ID)
								CopyModelPkgDBToModelPkg(modelpkgDB, modelpkg!, this.frontRepo)
							}
						)

						// fill up front objects
						pointertogongstructfields.forEach(
							pointertogongstructfieldDB => {
								let pointertogongstructfield = this.frontRepo.map_ID_PointerToGongStructField.get(pointertogongstructfieldDB.ID)
								CopyPointerToGongStructFieldDBToPointerToGongStructField(pointertogongstructfieldDB, pointertogongstructfield!, this.frontRepo)
							}
						)

						// fill up front objects
						sliceofpointertogongstructfields.forEach(
							sliceofpointertogongstructfieldDB => {
								let sliceofpointertogongstructfield = this.frontRepo.map_ID_SliceOfPointerToGongStructField.get(sliceofpointertogongstructfieldDB.ID)
								CopySliceOfPointerToGongStructFieldDBToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldDB, sliceofpointertogongstructfield!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
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
