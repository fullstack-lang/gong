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
  GongBasicFields_array = new Array<GongBasicFieldDB>() // array of repo instances
  GongBasicFields = new Map<number, GongBasicFieldDB>() // map of repo instances
  GongBasicFields_batch = new Map<number, GongBasicFieldDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongBasicFields = new Array<GongBasicField>() // array of front instances
  map_ID_GongBasicField = new Map<number, GongBasicField>() // map of front instances

  GongEnums_array = new Array<GongEnumDB>() // array of repo instances
  GongEnums = new Map<number, GongEnumDB>() // map of repo instances
  GongEnums_batch = new Map<number, GongEnumDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongEnums = new Array<GongEnum>() // array of front instances
  map_ID_GongEnum = new Map<number, GongEnum>() // map of front instances

  GongEnumValues_array = new Array<GongEnumValueDB>() // array of repo instances
  GongEnumValues = new Map<number, GongEnumValueDB>() // map of repo instances
  GongEnumValues_batch = new Map<number, GongEnumValueDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongEnumValues = new Array<GongEnumValue>() // array of front instances
  map_ID_GongEnumValue = new Map<number, GongEnumValue>() // map of front instances

  GongLinks_array = new Array<GongLinkDB>() // array of repo instances
  GongLinks = new Map<number, GongLinkDB>() // map of repo instances
  GongLinks_batch = new Map<number, GongLinkDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongLinks = new Array<GongLink>() // array of front instances
  map_ID_GongLink = new Map<number, GongLink>() // map of front instances

  GongNotes_array = new Array<GongNoteDB>() // array of repo instances
  GongNotes = new Map<number, GongNoteDB>() // map of repo instances
  GongNotes_batch = new Map<number, GongNoteDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongNotes = new Array<GongNote>() // array of front instances
  map_ID_GongNote = new Map<number, GongNote>() // map of front instances

  GongStructs_array = new Array<GongStructDB>() // array of repo instances
  GongStructs = new Map<number, GongStructDB>() // map of repo instances
  GongStructs_batch = new Map<number, GongStructDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongStructs = new Array<GongStruct>() // array of front instances
  map_ID_GongStruct = new Map<number, GongStruct>() // map of front instances

  GongTimeFields_array = new Array<GongTimeFieldDB>() // array of repo instances
  GongTimeFields = new Map<number, GongTimeFieldDB>() // map of repo instances
  GongTimeFields_batch = new Map<number, GongTimeFieldDB>() // same but only in last GET (for finding repo instances to delete)

  array_GongTimeFields = new Array<GongTimeField>() // array of front instances
  map_ID_GongTimeField = new Map<number, GongTimeField>() // map of front instances

  Metas_array = new Array<MetaDB>() // array of repo instances
  Metas = new Map<number, MetaDB>() // map of repo instances
  Metas_batch = new Map<number, MetaDB>() // same but only in last GET (for finding repo instances to delete)

  array_Metas = new Array<Meta>() // array of front instances
  map_ID_Meta = new Map<number, Meta>() // map of front instances

  MetaReferences_array = new Array<MetaReferenceDB>() // array of repo instances
  MetaReferences = new Map<number, MetaReferenceDB>() // map of repo instances
  MetaReferences_batch = new Map<number, MetaReferenceDB>() // same but only in last GET (for finding repo instances to delete)

  array_MetaReferences = new Array<MetaReference>() // array of front instances
  map_ID_MetaReference = new Map<number, MetaReference>() // map of front instances

  ModelPkgs_array = new Array<ModelPkgDB>() // array of repo instances
  ModelPkgs = new Map<number, ModelPkgDB>() // map of repo instances
  ModelPkgs_batch = new Map<number, ModelPkgDB>() // same but only in last GET (for finding repo instances to delete)

  array_ModelPkgs = new Array<ModelPkg>() // array of front instances
  map_ID_ModelPkg = new Map<number, ModelPkg>() // map of front instances

  PointerToGongStructFields_array = new Array<PointerToGongStructFieldDB>() // array of repo instances
  PointerToGongStructFields = new Map<number, PointerToGongStructFieldDB>() // map of repo instances
  PointerToGongStructFields_batch = new Map<number, PointerToGongStructFieldDB>() // same but only in last GET (for finding repo instances to delete)

  array_PointerToGongStructFields = new Array<PointerToGongStructField>() // array of front instances
  map_ID_PointerToGongStructField = new Map<number, PointerToGongStructField>() // map of front instances

  SliceOfPointerToGongStructFields_array = new Array<SliceOfPointerToGongStructFieldDB>() // array of repo instances
  SliceOfPointerToGongStructFields = new Map<number, SliceOfPointerToGongStructFieldDB>() // map of repo instances
  SliceOfPointerToGongStructFields_batch = new Map<number, SliceOfPointerToGongStructFieldDB>() // same but only in last GET (for finding repo instances to delete)

  array_SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructField>() // array of front instances
  map_ID_SliceOfPointerToGongStructField = new Map<number, SliceOfPointerToGongStructField>() // map of front instances


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) { // deprecated
      // insertion point
      case 'GongBasicField':
        return this.GongBasicFields_array as unknown as Array<Type>
      case 'GongEnum':
        return this.GongEnums_array as unknown as Array<Type>
      case 'GongEnumValue':
        return this.GongEnumValues_array as unknown as Array<Type>
      case 'GongLink':
        return this.GongLinks_array as unknown as Array<Type>
      case 'GongNote':
        return this.GongNotes_array as unknown as Array<Type>
      case 'GongStruct':
        return this.GongStructs_array as unknown as Array<Type>
      case 'GongTimeField':
        return this.GongTimeFields_array as unknown as Array<Type>
      case 'Meta':
        return this.Metas_array as unknown as Array<Type>
      case 'MetaReference':
        return this.MetaReferences_array as unknown as Array<Type>
      case 'ModelPkg':
        return this.ModelPkgs_array as unknown as Array<Type>
      case 'PointerToGongStructField':
        return this.PointerToGongStructFields_array as unknown as Array<Type>
      case 'SliceOfPointerToGongStructField':
        return this.SliceOfPointerToGongStructFields_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

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

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> { // deprecated
    switch (gongStructName) {
      // insertion point
      case 'GongBasicField':
        return this.GongBasicFields as unknown as Map<number, Type>
      case 'GongEnum':
        return this.GongEnums as unknown as Map<number, Type>
      case 'GongEnumValue':
        return this.GongEnumValues as unknown as Map<number, Type>
      case 'GongLink':
        return this.GongLinks as unknown as Map<number, Type>
      case 'GongNote':
        return this.GongNotes as unknown as Map<number, Type>
      case 'GongStruct':
        return this.GongStructs as unknown as Map<number, Type>
      case 'GongTimeField':
        return this.GongTimeFields as unknown as Map<number, Type>
      case 'Meta':
        return this.Metas as unknown as Map<number, Type>
      case 'MetaReference':
        return this.MetaReferences as unknown as Map<number, Type>
      case 'ModelPkg':
        return this.ModelPkgs as unknown as Map<number, Type>
      case 'PointerToGongStructField':
        return this.PointerToGongStructFields as unknown as Map<number, Type>
      case 'SliceOfPointerToGongStructField':
        return this.SliceOfPointerToGongStructFields as unknown as Map<number, Type>
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
            // init the array
            this.frontRepo.GongBasicFields_array = gongbasicfields

            // clear the map that counts GongBasicField in the GET
            this.frontRepo.GongBasicFields_batch.clear()

            gongbasicfields.forEach(
              gongbasicfieldDB => {
                this.frontRepo.GongBasicFields.set(gongbasicfieldDB.ID, gongbasicfieldDB)
                this.frontRepo.GongBasicFields_batch.set(gongbasicfieldDB.ID, gongbasicfieldDB)
              }
            )

            // clear gongbasicfields that are absent from the batch
            this.frontRepo.GongBasicFields.forEach(
              gongbasicfieldDB => {
                if (this.frontRepo.GongBasicFields_batch.get(gongbasicfieldDB.ID) == undefined) {
                  this.frontRepo.GongBasicFields.delete(gongbasicfieldDB.ID)
                }
              }
            )

            // sort GongBasicFields_array array
            this.frontRepo.GongBasicFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongEnums_array = gongenums

            // clear the map that counts GongEnum in the GET
            this.frontRepo.GongEnums_batch.clear()

            gongenums.forEach(
              gongenumDB => {
                this.frontRepo.GongEnums.set(gongenumDB.ID, gongenumDB)
                this.frontRepo.GongEnums_batch.set(gongenumDB.ID, gongenumDB)
              }
            )

            // clear gongenums that are absent from the batch
            this.frontRepo.GongEnums.forEach(
              gongenumDB => {
                if (this.frontRepo.GongEnums_batch.get(gongenumDB.ID) == undefined) {
                  this.frontRepo.GongEnums.delete(gongenumDB.ID)
                }
              }
            )

            // sort GongEnums_array array
            this.frontRepo.GongEnums_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongEnumValues_array = gongenumvalues

            // clear the map that counts GongEnumValue in the GET
            this.frontRepo.GongEnumValues_batch.clear()

            gongenumvalues.forEach(
              gongenumvalueDB => {
                this.frontRepo.GongEnumValues.set(gongenumvalueDB.ID, gongenumvalueDB)
                this.frontRepo.GongEnumValues_batch.set(gongenumvalueDB.ID, gongenumvalueDB)
              }
            )

            // clear gongenumvalues that are absent from the batch
            this.frontRepo.GongEnumValues.forEach(
              gongenumvalueDB => {
                if (this.frontRepo.GongEnumValues_batch.get(gongenumvalueDB.ID) == undefined) {
                  this.frontRepo.GongEnumValues.delete(gongenumvalueDB.ID)
                }
              }
            )

            // sort GongEnumValues_array array
            this.frontRepo.GongEnumValues_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongLinks_array = gonglinks

            // clear the map that counts GongLink in the GET
            this.frontRepo.GongLinks_batch.clear()

            gonglinks.forEach(
              gonglinkDB => {
                this.frontRepo.GongLinks.set(gonglinkDB.ID, gonglinkDB)
                this.frontRepo.GongLinks_batch.set(gonglinkDB.ID, gonglinkDB)
              }
            )

            // clear gonglinks that are absent from the batch
            this.frontRepo.GongLinks.forEach(
              gonglinkDB => {
                if (this.frontRepo.GongLinks_batch.get(gonglinkDB.ID) == undefined) {
                  this.frontRepo.GongLinks.delete(gonglinkDB.ID)
                }
              }
            )

            // sort GongLinks_array array
            this.frontRepo.GongLinks_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongNotes_array = gongnotes

            // clear the map that counts GongNote in the GET
            this.frontRepo.GongNotes_batch.clear()

            gongnotes.forEach(
              gongnoteDB => {
                this.frontRepo.GongNotes.set(gongnoteDB.ID, gongnoteDB)
                this.frontRepo.GongNotes_batch.set(gongnoteDB.ID, gongnoteDB)
              }
            )

            // clear gongnotes that are absent from the batch
            this.frontRepo.GongNotes.forEach(
              gongnoteDB => {
                if (this.frontRepo.GongNotes_batch.get(gongnoteDB.ID) == undefined) {
                  this.frontRepo.GongNotes.delete(gongnoteDB.ID)
                }
              }
            )

            // sort GongNotes_array array
            this.frontRepo.GongNotes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongStructs_array = gongstructs

            // clear the map that counts GongStruct in the GET
            this.frontRepo.GongStructs_batch.clear()

            gongstructs.forEach(
              gongstructDB => {
                this.frontRepo.GongStructs.set(gongstructDB.ID, gongstructDB)
                this.frontRepo.GongStructs_batch.set(gongstructDB.ID, gongstructDB)
              }
            )

            // clear gongstructs that are absent from the batch
            this.frontRepo.GongStructs.forEach(
              gongstructDB => {
                if (this.frontRepo.GongStructs_batch.get(gongstructDB.ID) == undefined) {
                  this.frontRepo.GongStructs.delete(gongstructDB.ID)
                }
              }
            )

            // sort GongStructs_array array
            this.frontRepo.GongStructs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongTimeFields_array = gongtimefields

            // clear the map that counts GongTimeField in the GET
            this.frontRepo.GongTimeFields_batch.clear()

            gongtimefields.forEach(
              gongtimefieldDB => {
                this.frontRepo.GongTimeFields.set(gongtimefieldDB.ID, gongtimefieldDB)
                this.frontRepo.GongTimeFields_batch.set(gongtimefieldDB.ID, gongtimefieldDB)
              }
            )

            // clear gongtimefields that are absent from the batch
            this.frontRepo.GongTimeFields.forEach(
              gongtimefieldDB => {
                if (this.frontRepo.GongTimeFields_batch.get(gongtimefieldDB.ID) == undefined) {
                  this.frontRepo.GongTimeFields.delete(gongtimefieldDB.ID)
                }
              }
            )

            // sort GongTimeFields_array array
            this.frontRepo.GongTimeFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Metas_array = metas

            // clear the map that counts Meta in the GET
            this.frontRepo.Metas_batch.clear()

            metas.forEach(
              metaDB => {
                this.frontRepo.Metas.set(metaDB.ID, metaDB)
                this.frontRepo.Metas_batch.set(metaDB.ID, metaDB)
              }
            )

            // clear metas that are absent from the batch
            this.frontRepo.Metas.forEach(
              metaDB => {
                if (this.frontRepo.Metas_batch.get(metaDB.ID) == undefined) {
                  this.frontRepo.Metas.delete(metaDB.ID)
                }
              }
            )

            // sort Metas_array array
            this.frontRepo.Metas_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.MetaReferences_array = metareferences

            // clear the map that counts MetaReference in the GET
            this.frontRepo.MetaReferences_batch.clear()

            metareferences.forEach(
              metareferenceDB => {
                this.frontRepo.MetaReferences.set(metareferenceDB.ID, metareferenceDB)
                this.frontRepo.MetaReferences_batch.set(metareferenceDB.ID, metareferenceDB)
              }
            )

            // clear metareferences that are absent from the batch
            this.frontRepo.MetaReferences.forEach(
              metareferenceDB => {
                if (this.frontRepo.MetaReferences_batch.get(metareferenceDB.ID) == undefined) {
                  this.frontRepo.MetaReferences.delete(metareferenceDB.ID)
                }
              }
            )

            // sort MetaReferences_array array
            this.frontRepo.MetaReferences_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.ModelPkgs_array = modelpkgs

            // clear the map that counts ModelPkg in the GET
            this.frontRepo.ModelPkgs_batch.clear()

            modelpkgs.forEach(
              modelpkgDB => {
                this.frontRepo.ModelPkgs.set(modelpkgDB.ID, modelpkgDB)
                this.frontRepo.ModelPkgs_batch.set(modelpkgDB.ID, modelpkgDB)
              }
            )

            // clear modelpkgs that are absent from the batch
            this.frontRepo.ModelPkgs.forEach(
              modelpkgDB => {
                if (this.frontRepo.ModelPkgs_batch.get(modelpkgDB.ID) == undefined) {
                  this.frontRepo.ModelPkgs.delete(modelpkgDB.ID)
                }
              }
            )

            // sort ModelPkgs_array array
            this.frontRepo.ModelPkgs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.PointerToGongStructFields_array = pointertogongstructfields

            // clear the map that counts PointerToGongStructField in the GET
            this.frontRepo.PointerToGongStructFields_batch.clear()

            pointertogongstructfields.forEach(
              pointertogongstructfieldDB => {
                this.frontRepo.PointerToGongStructFields.set(pointertogongstructfieldDB.ID, pointertogongstructfieldDB)
                this.frontRepo.PointerToGongStructFields_batch.set(pointertogongstructfieldDB.ID, pointertogongstructfieldDB)
              }
            )

            // clear pointertogongstructfields that are absent from the batch
            this.frontRepo.PointerToGongStructFields.forEach(
              pointertogongstructfieldDB => {
                if (this.frontRepo.PointerToGongStructFields_batch.get(pointertogongstructfieldDB.ID) == undefined) {
                  this.frontRepo.PointerToGongStructFields.delete(pointertogongstructfieldDB.ID)
                }
              }
            )

            // sort PointerToGongStructFields_array array
            this.frontRepo.PointerToGongStructFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.SliceOfPointerToGongStructFields_array = sliceofpointertogongstructfields

            // clear the map that counts SliceOfPointerToGongStructField in the GET
            this.frontRepo.SliceOfPointerToGongStructFields_batch.clear()

            sliceofpointertogongstructfields.forEach(
              sliceofpointertogongstructfieldDB => {
                this.frontRepo.SliceOfPointerToGongStructFields.set(sliceofpointertogongstructfieldDB.ID, sliceofpointertogongstructfieldDB)
                this.frontRepo.SliceOfPointerToGongStructFields_batch.set(sliceofpointertogongstructfieldDB.ID, sliceofpointertogongstructfieldDB)
              }
            )

            // clear sliceofpointertogongstructfields that are absent from the batch
            this.frontRepo.SliceOfPointerToGongStructFields.forEach(
              sliceofpointertogongstructfieldDB => {
                if (this.frontRepo.SliceOfPointerToGongStructFields_batch.get(sliceofpointertogongstructfieldDB.ID) == undefined) {
                  this.frontRepo.SliceOfPointerToGongStructFields.delete(sliceofpointertogongstructfieldDB.ID)
                }
              }
            )

            // sort SliceOfPointerToGongStructFields_array array
            this.frontRepo.SliceOfPointerToGongStructFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            gongbasicfields.forEach(
              gongbasicfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field GongEnum redeeming
                {
                  let _gongenum = this.frontRepo.GongEnums.get(gongbasicfield.GongBasicFieldPointersEncoding.GongEnumID.Int64)
                  if (_gongenum) {
                    gongbasicfield.GongEnum = _gongenum
                  }
                }
                // insertion point for pointers decoding
              }
            )
            gongenums.forEach(
              gongenum => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                gongenum.GongEnumValues = new Array<GongEnumValueDB>()
                for (let _id of gongenum.GongEnumPointersEncoding.GongEnumValues) {
                  let _gongenumvalue = this.frontRepo.GongEnumValues.get(_id)
                  if (_gongenumvalue != undefined) {
                    gongenum.GongEnumValues.push(_gongenumvalue!)
                  }
                }
              }
            )
            gongenumvalues.forEach(
              gongenumvalue => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            gonglinks.forEach(
              gonglink => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            gongnotes.forEach(
              gongnote => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                gongnote.Links = new Array<GongLinkDB>()
                for (let _id of gongnote.GongNotePointersEncoding.Links) {
                  let _gonglink = this.frontRepo.GongLinks.get(_id)
                  if (_gonglink != undefined) {
                    gongnote.Links.push(_gonglink!)
                  }
                }
              }
            )
            gongstructs.forEach(
              gongstruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                gongstruct.GongBasicFields = new Array<GongBasicFieldDB>()
                for (let _id of gongstruct.GongStructPointersEncoding.GongBasicFields) {
                  let _gongbasicfield = this.frontRepo.GongBasicFields.get(_id)
                  if (_gongbasicfield != undefined) {
                    gongstruct.GongBasicFields.push(_gongbasicfield!)
                  }
                }
                gongstruct.GongTimeFields = new Array<GongTimeFieldDB>()
                for (let _id of gongstruct.GongStructPointersEncoding.GongTimeFields) {
                  let _gongtimefield = this.frontRepo.GongTimeFields.get(_id)
                  if (_gongtimefield != undefined) {
                    gongstruct.GongTimeFields.push(_gongtimefield!)
                  }
                }
                gongstruct.PointerToGongStructFields = new Array<PointerToGongStructFieldDB>()
                for (let _id of gongstruct.GongStructPointersEncoding.PointerToGongStructFields) {
                  let _pointertogongstructfield = this.frontRepo.PointerToGongStructFields.get(_id)
                  if (_pointertogongstructfield != undefined) {
                    gongstruct.PointerToGongStructFields.push(_pointertogongstructfield!)
                  }
                }
                gongstruct.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructFieldDB>()
                for (let _id of gongstruct.GongStructPointersEncoding.SliceOfPointerToGongStructFields) {
                  let _sliceofpointertogongstructfield = this.frontRepo.SliceOfPointerToGongStructFields.get(_id)
                  if (_sliceofpointertogongstructfield != undefined) {
                    gongstruct.SliceOfPointerToGongStructFields.push(_sliceofpointertogongstructfield!)
                  }
                }
              }
            )
            gongtimefields.forEach(
              gongtimefield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            metas.forEach(
              meta => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                meta.MetaReferences = new Array<MetaReferenceDB>()
                for (let _id of meta.MetaPointersEncoding.MetaReferences) {
                  let _metareference = this.frontRepo.MetaReferences.get(_id)
                  if (_metareference != undefined) {
                    meta.MetaReferences.push(_metareference!)
                  }
                }
              }
            )
            metareferences.forEach(
              metareference => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            modelpkgs.forEach(
              modelpkg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            pointertogongstructfields.forEach(
              pointertogongstructfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = this.frontRepo.GongStructs.get(pointertogongstructfield.PointerToGongStructFieldPointersEncoding.GongStructID.Int64)
                  if (_gongstruct) {
                    pointertogongstructfield.GongStruct = _gongstruct
                  }
                }
                // insertion point for pointers decoding
              }
            )
            sliceofpointertogongstructfields.forEach(
              sliceofpointertogongstructfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = this.frontRepo.GongStructs.get(sliceofpointertogongstructfield.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64)
                  if (_gongstruct) {
                    sliceofpointertogongstructfield.GongStruct = _gongstruct
                  }
                }
                // insertion point for pointers decoding
              }
            )

            // 
            // Third Step: reddeem front objects
            // insertion point sub template for redeem 
            
            // init front objects
            this.frontRepo.array_GongBasicFields = []
            this.frontRepo.map_ID_GongBasicField.clear()
            this.frontRepo.GongBasicFields_array.forEach(
              gongbasicfieldDB => {
                let gongbasicfield = new GongBasicField
                CopyGongBasicFieldDBToGongBasicField(gongbasicfieldDB, gongbasicfield, this.frontRepo)
                this.frontRepo.array_GongBasicFields.push(gongbasicfield)
                this.frontRepo.map_ID_GongBasicField.set(gongbasicfield.ID, gongbasicfield)
              }
            )

            
            // init front objects
            this.frontRepo.array_GongEnums = []
            this.frontRepo.map_ID_GongEnum.clear()
            this.frontRepo.GongEnums_array.forEach(
              gongenumDB => {
                let gongenum = new GongEnum
                CopyGongEnumDBToGongEnum(gongenumDB, gongenum, this.frontRepo)
                this.frontRepo.array_GongEnums.push(gongenum)
                this.frontRepo.map_ID_GongEnum.set(gongenum.ID, gongenum)
              }
            )

            
            // init front objects
            this.frontRepo.array_GongEnumValues = []
            this.frontRepo.map_ID_GongEnumValue.clear()
            this.frontRepo.GongEnumValues_array.forEach(
              gongenumvalueDB => {
                let gongenumvalue = new GongEnumValue
                CopyGongEnumValueDBToGongEnumValue(gongenumvalueDB, gongenumvalue, this.frontRepo)
                this.frontRepo.array_GongEnumValues.push(gongenumvalue)
                this.frontRepo.map_ID_GongEnumValue.set(gongenumvalue.ID, gongenumvalue)
              }
            )

            
            // init front objects
            this.frontRepo.array_GongLinks = []
            this.frontRepo.map_ID_GongLink.clear()
            this.frontRepo.GongLinks_array.forEach(
              gonglinkDB => {
                let gonglink = new GongLink
                CopyGongLinkDBToGongLink(gonglinkDB, gonglink, this.frontRepo)
                this.frontRepo.array_GongLinks.push(gonglink)
                this.frontRepo.map_ID_GongLink.set(gonglink.ID, gonglink)
              }
            )

            
            // init front objects
            this.frontRepo.array_GongNotes = []
            this.frontRepo.map_ID_GongNote.clear()
            this.frontRepo.GongNotes_array.forEach(
              gongnoteDB => {
                let gongnote = new GongNote
                CopyGongNoteDBToGongNote(gongnoteDB, gongnote, this.frontRepo)
                this.frontRepo.array_GongNotes.push(gongnote)
                this.frontRepo.map_ID_GongNote.set(gongnote.ID, gongnote)
              }
            )

            
            // init front objects
            this.frontRepo.array_GongStructs = []
            this.frontRepo.map_ID_GongStruct.clear()
            this.frontRepo.GongStructs_array.forEach(
              gongstructDB => {
                let gongstruct = new GongStruct
                CopyGongStructDBToGongStruct(gongstructDB, gongstruct, this.frontRepo)
                this.frontRepo.array_GongStructs.push(gongstruct)
                this.frontRepo.map_ID_GongStruct.set(gongstruct.ID, gongstruct)
              }
            )

            
            // init front objects
            this.frontRepo.array_GongTimeFields = []
            this.frontRepo.map_ID_GongTimeField.clear()
            this.frontRepo.GongTimeFields_array.forEach(
              gongtimefieldDB => {
                let gongtimefield = new GongTimeField
                CopyGongTimeFieldDBToGongTimeField(gongtimefieldDB, gongtimefield, this.frontRepo)
                this.frontRepo.array_GongTimeFields.push(gongtimefield)
                this.frontRepo.map_ID_GongTimeField.set(gongtimefield.ID, gongtimefield)
              }
            )

            
            // init front objects
            this.frontRepo.array_Metas = []
            this.frontRepo.map_ID_Meta.clear()
            this.frontRepo.Metas_array.forEach(
              metaDB => {
                let meta = new Meta
                CopyMetaDBToMeta(metaDB, meta, this.frontRepo)
                this.frontRepo.array_Metas.push(meta)
                this.frontRepo.map_ID_Meta.set(meta.ID, meta)
              }
            )

            
            // init front objects
            this.frontRepo.array_MetaReferences = []
            this.frontRepo.map_ID_MetaReference.clear()
            this.frontRepo.MetaReferences_array.forEach(
              metareferenceDB => {
                let metareference = new MetaReference
                CopyMetaReferenceDBToMetaReference(metareferenceDB, metareference, this.frontRepo)
                this.frontRepo.array_MetaReferences.push(metareference)
                this.frontRepo.map_ID_MetaReference.set(metareference.ID, metareference)
              }
            )

            
            // init front objects
            this.frontRepo.array_ModelPkgs = []
            this.frontRepo.map_ID_ModelPkg.clear()
            this.frontRepo.ModelPkgs_array.forEach(
              modelpkgDB => {
                let modelpkg = new ModelPkg
                CopyModelPkgDBToModelPkg(modelpkgDB, modelpkg, this.frontRepo)
                this.frontRepo.array_ModelPkgs.push(modelpkg)
                this.frontRepo.map_ID_ModelPkg.set(modelpkg.ID, modelpkg)
              }
            )

            
            // init front objects
            this.frontRepo.array_PointerToGongStructFields = []
            this.frontRepo.map_ID_PointerToGongStructField.clear()
            this.frontRepo.PointerToGongStructFields_array.forEach(
              pointertogongstructfieldDB => {
                let pointertogongstructfield = new PointerToGongStructField
                CopyPointerToGongStructFieldDBToPointerToGongStructField(pointertogongstructfieldDB, pointertogongstructfield, this.frontRepo)
                this.frontRepo.array_PointerToGongStructFields.push(pointertogongstructfield)
                this.frontRepo.map_ID_PointerToGongStructField.set(pointertogongstructfield.ID, pointertogongstructfield)
              }
            )

            
            // init front objects
            this.frontRepo.array_SliceOfPointerToGongStructFields = []
            this.frontRepo.map_ID_SliceOfPointerToGongStructField.clear()
            this.frontRepo.SliceOfPointerToGongStructFields_array.forEach(
              sliceofpointertogongstructfieldDB => {
                let sliceofpointertogongstructfield = new SliceOfPointerToGongStructField
                CopySliceOfPointerToGongStructFieldDBToSliceOfPointerToGongStructField(sliceofpointertogongstructfieldDB, sliceofpointertogongstructfield, this.frontRepo)
                this.frontRepo.array_SliceOfPointerToGongStructFields.push(sliceofpointertogongstructfield)
                this.frontRepo.map_ID_SliceOfPointerToGongStructField.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)
              }
            )



            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // GongBasicFieldPull performs a GET on GongBasicField of the stack and redeem association pointers 
  GongBasicFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongbasicfieldService.getGongBasicFields(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gongbasicfields,
          ]) => {
            // init the array
            this.frontRepo.GongBasicFields_array = gongbasicfields

            // clear the map that counts GongBasicField in the GET
            this.frontRepo.GongBasicFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongbasicfields.forEach(
              gongbasicfield => {
                this.frontRepo.GongBasicFields.set(gongbasicfield.ID, gongbasicfield)
                this.frontRepo.GongBasicFields_batch.set(gongbasicfield.ID, gongbasicfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field GongEnum redeeming
                {
                  let _gongenum = this.frontRepo.GongEnums.get(gongbasicfield.GongBasicFieldPointersEncoding.GongEnumID.Int64)
                  if (_gongenum) {
                    gongbasicfield.GongEnum = _gongenum
                  }
                }
              }
            )

            // clear gongbasicfields that are absent from the GET
            this.frontRepo.GongBasicFields.forEach(
              gongbasicfield => {
                if (this.frontRepo.GongBasicFields_batch.get(gongbasicfield.ID) == undefined) {
                  this.frontRepo.GongBasicFields.delete(gongbasicfield.ID)
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
  }

  // GongEnumPull performs a GET on GongEnum of the stack and redeem association pointers 
  GongEnumPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongenumService.getGongEnums(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gongenums,
          ]) => {
            // init the array
            this.frontRepo.GongEnums_array = gongenums

            // clear the map that counts GongEnum in the GET
            this.frontRepo.GongEnums_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongenums.forEach(
              gongenum => {
                this.frontRepo.GongEnums.set(gongenum.ID, gongenum)
                this.frontRepo.GongEnums_batch.set(gongenum.ID, gongenum)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear gongenums that are absent from the GET
            this.frontRepo.GongEnums.forEach(
              gongenum => {
                if (this.frontRepo.GongEnums_batch.get(gongenum.ID) == undefined) {
                  this.frontRepo.GongEnums.delete(gongenum.ID)
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
  }

  // GongEnumValuePull performs a GET on GongEnumValue of the stack and redeem association pointers 
  GongEnumValuePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongenumvalueService.getGongEnumValues(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gongenumvalues,
          ]) => {
            // init the array
            this.frontRepo.GongEnumValues_array = gongenumvalues

            // clear the map that counts GongEnumValue in the GET
            this.frontRepo.GongEnumValues_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongenumvalues.forEach(
              gongenumvalue => {
                this.frontRepo.GongEnumValues.set(gongenumvalue.ID, gongenumvalue)
                this.frontRepo.GongEnumValues_batch.set(gongenumvalue.ID, gongenumvalue)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear gongenumvalues that are absent from the GET
            this.frontRepo.GongEnumValues.forEach(
              gongenumvalue => {
                if (this.frontRepo.GongEnumValues_batch.get(gongenumvalue.ID) == undefined) {
                  this.frontRepo.GongEnumValues.delete(gongenumvalue.ID)
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
  }

  // GongLinkPull performs a GET on GongLink of the stack and redeem association pointers 
  GongLinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gonglinkService.getGongLinks(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gonglinks,
          ]) => {
            // init the array
            this.frontRepo.GongLinks_array = gonglinks

            // clear the map that counts GongLink in the GET
            this.frontRepo.GongLinks_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gonglinks.forEach(
              gonglink => {
                this.frontRepo.GongLinks.set(gonglink.ID, gonglink)
                this.frontRepo.GongLinks_batch.set(gonglink.ID, gonglink)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear gonglinks that are absent from the GET
            this.frontRepo.GongLinks.forEach(
              gonglink => {
                if (this.frontRepo.GongLinks_batch.get(gonglink.ID) == undefined) {
                  this.frontRepo.GongLinks.delete(gonglink.ID)
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
  }

  // GongNotePull performs a GET on GongNote of the stack and redeem association pointers 
  GongNotePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongnoteService.getGongNotes(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gongnotes,
          ]) => {
            // init the array
            this.frontRepo.GongNotes_array = gongnotes

            // clear the map that counts GongNote in the GET
            this.frontRepo.GongNotes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongnotes.forEach(
              gongnote => {
                this.frontRepo.GongNotes.set(gongnote.ID, gongnote)
                this.frontRepo.GongNotes_batch.set(gongnote.ID, gongnote)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear gongnotes that are absent from the GET
            this.frontRepo.GongNotes.forEach(
              gongnote => {
                if (this.frontRepo.GongNotes_batch.get(gongnote.ID) == undefined) {
                  this.frontRepo.GongNotes.delete(gongnote.ID)
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
  }

  // GongStructPull performs a GET on GongStruct of the stack and redeem association pointers 
  GongStructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongstructService.getGongStructs(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gongstructs,
          ]) => {
            // init the array
            this.frontRepo.GongStructs_array = gongstructs

            // clear the map that counts GongStruct in the GET
            this.frontRepo.GongStructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongstructs.forEach(
              gongstruct => {
                this.frontRepo.GongStructs.set(gongstruct.ID, gongstruct)
                this.frontRepo.GongStructs_batch.set(gongstruct.ID, gongstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear gongstructs that are absent from the GET
            this.frontRepo.GongStructs.forEach(
              gongstruct => {
                if (this.frontRepo.GongStructs_batch.get(gongstruct.ID) == undefined) {
                  this.frontRepo.GongStructs.delete(gongstruct.ID)
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
  }

  // GongTimeFieldPull performs a GET on GongTimeField of the stack and redeem association pointers 
  GongTimeFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongtimefieldService.getGongTimeFields(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            gongtimefields,
          ]) => {
            // init the array
            this.frontRepo.GongTimeFields_array = gongtimefields

            // clear the map that counts GongTimeField in the GET
            this.frontRepo.GongTimeFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongtimefields.forEach(
              gongtimefield => {
                this.frontRepo.GongTimeFields.set(gongtimefield.ID, gongtimefield)
                this.frontRepo.GongTimeFields_batch.set(gongtimefield.ID, gongtimefield)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear gongtimefields that are absent from the GET
            this.frontRepo.GongTimeFields.forEach(
              gongtimefield => {
                if (this.frontRepo.GongTimeFields_batch.get(gongtimefield.ID) == undefined) {
                  this.frontRepo.GongTimeFields.delete(gongtimefield.ID)
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
  }

  // MetaPull performs a GET on Meta of the stack and redeem association pointers 
  MetaPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.metaService.getMetas(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            metas,
          ]) => {
            // init the array
            this.frontRepo.Metas_array = metas

            // clear the map that counts Meta in the GET
            this.frontRepo.Metas_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            metas.forEach(
              meta => {
                this.frontRepo.Metas.set(meta.ID, meta)
                this.frontRepo.Metas_batch.set(meta.ID, meta)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear metas that are absent from the GET
            this.frontRepo.Metas.forEach(
              meta => {
                if (this.frontRepo.Metas_batch.get(meta.ID) == undefined) {
                  this.frontRepo.Metas.delete(meta.ID)
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
  }

  // MetaReferencePull performs a GET on MetaReference of the stack and redeem association pointers 
  MetaReferencePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.metareferenceService.getMetaReferences(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            metareferences,
          ]) => {
            // init the array
            this.frontRepo.MetaReferences_array = metareferences

            // clear the map that counts MetaReference in the GET
            this.frontRepo.MetaReferences_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            metareferences.forEach(
              metareference => {
                this.frontRepo.MetaReferences.set(metareference.ID, metareference)
                this.frontRepo.MetaReferences_batch.set(metareference.ID, metareference)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear metareferences that are absent from the GET
            this.frontRepo.MetaReferences.forEach(
              metareference => {
                if (this.frontRepo.MetaReferences_batch.get(metareference.ID) == undefined) {
                  this.frontRepo.MetaReferences.delete(metareference.ID)
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
  }

  // ModelPkgPull performs a GET on ModelPkg of the stack and redeem association pointers 
  ModelPkgPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.modelpkgService.getModelPkgs(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            modelpkgs,
          ]) => {
            // init the array
            this.frontRepo.ModelPkgs_array = modelpkgs

            // clear the map that counts ModelPkg in the GET
            this.frontRepo.ModelPkgs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            modelpkgs.forEach(
              modelpkg => {
                this.frontRepo.ModelPkgs.set(modelpkg.ID, modelpkg)
                this.frontRepo.ModelPkgs_batch.set(modelpkg.ID, modelpkg)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear modelpkgs that are absent from the GET
            this.frontRepo.ModelPkgs.forEach(
              modelpkg => {
                if (this.frontRepo.ModelPkgs_batch.get(modelpkg.ID) == undefined) {
                  this.frontRepo.ModelPkgs.delete(modelpkg.ID)
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
  }

  // PointerToGongStructFieldPull performs a GET on PointerToGongStructField of the stack and redeem association pointers 
  PointerToGongStructFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pointertogongstructfieldService.getPointerToGongStructFields(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            pointertogongstructfields,
          ]) => {
            // init the array
            this.frontRepo.PointerToGongStructFields_array = pointertogongstructfields

            // clear the map that counts PointerToGongStructField in the GET
            this.frontRepo.PointerToGongStructFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            pointertogongstructfields.forEach(
              pointertogongstructfield => {
                this.frontRepo.PointerToGongStructFields.set(pointertogongstructfield.ID, pointertogongstructfield)
                this.frontRepo.PointerToGongStructFields_batch.set(pointertogongstructfield.ID, pointertogongstructfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = this.frontRepo.GongStructs.get(pointertogongstructfield.PointerToGongStructFieldPointersEncoding.GongStructID.Int64)
                  if (_gongstruct) {
                    pointertogongstructfield.GongStruct = _gongstruct
                  }
                }
              }
            )

            // clear pointertogongstructfields that are absent from the GET
            this.frontRepo.PointerToGongStructFields.forEach(
              pointertogongstructfield => {
                if (this.frontRepo.PointerToGongStructFields_batch.get(pointertogongstructfield.ID) == undefined) {
                  this.frontRepo.PointerToGongStructFields.delete(pointertogongstructfield.ID)
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
  }

  // SliceOfPointerToGongStructFieldPull performs a GET on SliceOfPointerToGongStructField of the stack and redeem association pointers 
  SliceOfPointerToGongStructFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.sliceofpointertogongstructfieldService.getSliceOfPointerToGongStructFields(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            sliceofpointertogongstructfields,
          ]) => {
            // init the array
            this.frontRepo.SliceOfPointerToGongStructFields_array = sliceofpointertogongstructfields

            // clear the map that counts SliceOfPointerToGongStructField in the GET
            this.frontRepo.SliceOfPointerToGongStructFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            sliceofpointertogongstructfields.forEach(
              sliceofpointertogongstructfield => {
                this.frontRepo.SliceOfPointerToGongStructFields.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)
                this.frontRepo.SliceOfPointerToGongStructFields_batch.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = this.frontRepo.GongStructs.get(sliceofpointertogongstructfield.SliceOfPointerToGongStructFieldPointersEncoding.GongStructID.Int64)
                  if (_gongstruct) {
                    sliceofpointertogongstructfield.GongStruct = _gongstruct
                  }
                }
              }
            )

            // clear sliceofpointertogongstructfields that are absent from the GET
            this.frontRepo.SliceOfPointerToGongStructFields.forEach(
              sliceofpointertogongstructfield => {
                if (this.frontRepo.SliceOfPointerToGongStructFields_batch.get(sliceofpointertogongstructfield.ID) == undefined) {
                  this.frontRepo.SliceOfPointerToGongStructFields.delete(sliceofpointertogongstructfield.ID)
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
