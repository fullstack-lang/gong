import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { GongBasicFieldDB } from './gongbasicfield-db'
import { GongBasicFieldService } from './gongbasicfield.service'

import { GongEnumDB } from './gongenum-db'
import { GongEnumService } from './gongenum.service'

import { GongEnumValueDB } from './gongenumvalue-db'
import { GongEnumValueService } from './gongenumvalue.service'

import { GongLinkDB } from './gonglink-db'
import { GongLinkService } from './gonglink.service'

import { GongNoteDB } from './gongnote-db'
import { GongNoteService } from './gongnote.service'

import { GongStructDB } from './gongstruct-db'
import { GongStructService } from './gongstruct.service'

import { GongTimeFieldDB } from './gongtimefield-db'
import { GongTimeFieldService } from './gongtimefield.service'

import { MetaDB } from './meta-db'
import { MetaService } from './meta.service'

import { MetaReferenceDB } from './metareference-db'
import { MetaReferenceService } from './metareference.service'

import { ModelPkgDB } from './modelpkg-db'
import { ModelPkgService } from './modelpkg.service'

import { PointerToGongStructFieldDB } from './pointertogongstructfield-db'
import { PointerToGongStructFieldService } from './pointertogongstructfield.service'

import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructFieldService } from './sliceofpointertogongstructfield.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  GongBasicFields_array = new Array<GongBasicFieldDB>(); // array of repo instances
  GongBasicFields = new Map<number, GongBasicFieldDB>(); // map of repo instances
  GongBasicFields_batch = new Map<number, GongBasicFieldDB>(); // same but only in last GET (for finding repo instances to delete)
  GongEnums_array = new Array<GongEnumDB>(); // array of repo instances
  GongEnums = new Map<number, GongEnumDB>(); // map of repo instances
  GongEnums_batch = new Map<number, GongEnumDB>(); // same but only in last GET (for finding repo instances to delete)
  GongEnumValues_array = new Array<GongEnumValueDB>(); // array of repo instances
  GongEnumValues = new Map<number, GongEnumValueDB>(); // map of repo instances
  GongEnumValues_batch = new Map<number, GongEnumValueDB>(); // same but only in last GET (for finding repo instances to delete)
  GongLinks_array = new Array<GongLinkDB>(); // array of repo instances
  GongLinks = new Map<number, GongLinkDB>(); // map of repo instances
  GongLinks_batch = new Map<number, GongLinkDB>(); // same but only in last GET (for finding repo instances to delete)
  GongNotes_array = new Array<GongNoteDB>(); // array of repo instances
  GongNotes = new Map<number, GongNoteDB>(); // map of repo instances
  GongNotes_batch = new Map<number, GongNoteDB>(); // same but only in last GET (for finding repo instances to delete)
  GongStructs_array = new Array<GongStructDB>(); // array of repo instances
  GongStructs = new Map<number, GongStructDB>(); // map of repo instances
  GongStructs_batch = new Map<number, GongStructDB>(); // same but only in last GET (for finding repo instances to delete)
  GongTimeFields_array = new Array<GongTimeFieldDB>(); // array of repo instances
  GongTimeFields = new Map<number, GongTimeFieldDB>(); // map of repo instances
  GongTimeFields_batch = new Map<number, GongTimeFieldDB>(); // same but only in last GET (for finding repo instances to delete)
  Metas_array = new Array<MetaDB>(); // array of repo instances
  Metas = new Map<number, MetaDB>(); // map of repo instances
  Metas_batch = new Map<number, MetaDB>(); // same but only in last GET (for finding repo instances to delete)
  MetaReferences_array = new Array<MetaReferenceDB>(); // array of repo instances
  MetaReferences = new Map<number, MetaReferenceDB>(); // map of repo instances
  MetaReferences_batch = new Map<number, MetaReferenceDB>(); // same but only in last GET (for finding repo instances to delete)
  ModelPkgs_array = new Array<ModelPkgDB>(); // array of repo instances
  ModelPkgs = new Map<number, ModelPkgDB>(); // map of repo instances
  ModelPkgs_batch = new Map<number, ModelPkgDB>(); // same but only in last GET (for finding repo instances to delete)
  PointerToGongStructFields_array = new Array<PointerToGongStructFieldDB>(); // array of repo instances
  PointerToGongStructFields = new Map<number, PointerToGongStructFieldDB>(); // map of repo instances
  PointerToGongStructFields_batch = new Map<number, PointerToGongStructFieldDB>(); // same but only in last GET (for finding repo instances to delete)
  SliceOfPointerToGongStructFields_array = new Array<SliceOfPointerToGongStructFieldDB>(); // array of repo instances
  SliceOfPointerToGongStructFields = new Map<number, SliceOfPointerToGongStructFieldDB>(); // map of repo instances
  SliceOfPointerToGongStructFields_batch = new Map<number, SliceOfPointerToGongStructFieldDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

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

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

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
  observableFrontRepo: [ // insertion point sub template 
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
  ] = [ // insertion point sub template 
      this.gongbasicfieldService.getGongBasicFields(),
      this.gongenumService.getGongEnums(),
      this.gongenumvalueService.getGongEnumValues(),
      this.gonglinkService.getGongLinks(),
      this.gongnoteService.getGongNotes(),
      this.gongstructService.getGongStructs(),
      this.gongtimefieldService.getGongTimeFields(),
      this.metaService.getMetas(),
      this.metareferenceService.getMetaReferences(),
      this.modelpkgService.getModelPkgs(),
      this.pointertogongstructfieldService.getPointerToGongStructFields(),
      this.sliceofpointertogongstructfieldService.getSliceOfPointerToGongStructFields(),
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
          ([ // insertion point sub template for declarations 
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
            FrontRepoSingloton.GongBasicFields_array = gongbasicfields

            // clear the map that counts GongBasicField in the GET
            FrontRepoSingloton.GongBasicFields_batch.clear()

            gongbasicfields.forEach(
              gongbasicfield => {
                FrontRepoSingloton.GongBasicFields.set(gongbasicfield.ID, gongbasicfield)
                FrontRepoSingloton.GongBasicFields_batch.set(gongbasicfield.ID, gongbasicfield)
              }
            )

            // clear gongbasicfields that are absent from the batch
            FrontRepoSingloton.GongBasicFields.forEach(
              gongbasicfield => {
                if (FrontRepoSingloton.GongBasicFields_batch.get(gongbasicfield.ID) == undefined) {
                  FrontRepoSingloton.GongBasicFields.delete(gongbasicfield.ID)
                }
              }
            )

            // sort GongBasicFields_array array
            FrontRepoSingloton.GongBasicFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongEnums_array = gongenums

            // clear the map that counts GongEnum in the GET
            FrontRepoSingloton.GongEnums_batch.clear()

            gongenums.forEach(
              gongenum => {
                FrontRepoSingloton.GongEnums.set(gongenum.ID, gongenum)
                FrontRepoSingloton.GongEnums_batch.set(gongenum.ID, gongenum)
              }
            )

            // clear gongenums that are absent from the batch
            FrontRepoSingloton.GongEnums.forEach(
              gongenum => {
                if (FrontRepoSingloton.GongEnums_batch.get(gongenum.ID) == undefined) {
                  FrontRepoSingloton.GongEnums.delete(gongenum.ID)
                }
              }
            )

            // sort GongEnums_array array
            FrontRepoSingloton.GongEnums_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongEnumValues_array = gongenumvalues

            // clear the map that counts GongEnumValue in the GET
            FrontRepoSingloton.GongEnumValues_batch.clear()

            gongenumvalues.forEach(
              gongenumvalue => {
                FrontRepoSingloton.GongEnumValues.set(gongenumvalue.ID, gongenumvalue)
                FrontRepoSingloton.GongEnumValues_batch.set(gongenumvalue.ID, gongenumvalue)
              }
            )

            // clear gongenumvalues that are absent from the batch
            FrontRepoSingloton.GongEnumValues.forEach(
              gongenumvalue => {
                if (FrontRepoSingloton.GongEnumValues_batch.get(gongenumvalue.ID) == undefined) {
                  FrontRepoSingloton.GongEnumValues.delete(gongenumvalue.ID)
                }
              }
            )

            // sort GongEnumValues_array array
            FrontRepoSingloton.GongEnumValues_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongLinks_array = gonglinks

            // clear the map that counts GongLink in the GET
            FrontRepoSingloton.GongLinks_batch.clear()

            gonglinks.forEach(
              gonglink => {
                FrontRepoSingloton.GongLinks.set(gonglink.ID, gonglink)
                FrontRepoSingloton.GongLinks_batch.set(gonglink.ID, gonglink)
              }
            )

            // clear gonglinks that are absent from the batch
            FrontRepoSingloton.GongLinks.forEach(
              gonglink => {
                if (FrontRepoSingloton.GongLinks_batch.get(gonglink.ID) == undefined) {
                  FrontRepoSingloton.GongLinks.delete(gonglink.ID)
                }
              }
            )

            // sort GongLinks_array array
            FrontRepoSingloton.GongLinks_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongNotes_array = gongnotes

            // clear the map that counts GongNote in the GET
            FrontRepoSingloton.GongNotes_batch.clear()

            gongnotes.forEach(
              gongnote => {
                FrontRepoSingloton.GongNotes.set(gongnote.ID, gongnote)
                FrontRepoSingloton.GongNotes_batch.set(gongnote.ID, gongnote)
              }
            )

            // clear gongnotes that are absent from the batch
            FrontRepoSingloton.GongNotes.forEach(
              gongnote => {
                if (FrontRepoSingloton.GongNotes_batch.get(gongnote.ID) == undefined) {
                  FrontRepoSingloton.GongNotes.delete(gongnote.ID)
                }
              }
            )

            // sort GongNotes_array array
            FrontRepoSingloton.GongNotes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongStructs_array = gongstructs

            // clear the map that counts GongStruct in the GET
            FrontRepoSingloton.GongStructs_batch.clear()

            gongstructs.forEach(
              gongstruct => {
                FrontRepoSingloton.GongStructs.set(gongstruct.ID, gongstruct)
                FrontRepoSingloton.GongStructs_batch.set(gongstruct.ID, gongstruct)
              }
            )

            // clear gongstructs that are absent from the batch
            FrontRepoSingloton.GongStructs.forEach(
              gongstruct => {
                if (FrontRepoSingloton.GongStructs_batch.get(gongstruct.ID) == undefined) {
                  FrontRepoSingloton.GongStructs.delete(gongstruct.ID)
                }
              }
            )

            // sort GongStructs_array array
            FrontRepoSingloton.GongStructs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.GongTimeFields_array = gongtimefields

            // clear the map that counts GongTimeField in the GET
            FrontRepoSingloton.GongTimeFields_batch.clear()

            gongtimefields.forEach(
              gongtimefield => {
                FrontRepoSingloton.GongTimeFields.set(gongtimefield.ID, gongtimefield)
                FrontRepoSingloton.GongTimeFields_batch.set(gongtimefield.ID, gongtimefield)
              }
            )

            // clear gongtimefields that are absent from the batch
            FrontRepoSingloton.GongTimeFields.forEach(
              gongtimefield => {
                if (FrontRepoSingloton.GongTimeFields_batch.get(gongtimefield.ID) == undefined) {
                  FrontRepoSingloton.GongTimeFields.delete(gongtimefield.ID)
                }
              }
            )

            // sort GongTimeFields_array array
            FrontRepoSingloton.GongTimeFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Metas_array = metas

            // clear the map that counts Meta in the GET
            FrontRepoSingloton.Metas_batch.clear()

            metas.forEach(
              meta => {
                FrontRepoSingloton.Metas.set(meta.ID, meta)
                FrontRepoSingloton.Metas_batch.set(meta.ID, meta)
              }
            )

            // clear metas that are absent from the batch
            FrontRepoSingloton.Metas.forEach(
              meta => {
                if (FrontRepoSingloton.Metas_batch.get(meta.ID) == undefined) {
                  FrontRepoSingloton.Metas.delete(meta.ID)
                }
              }
            )

            // sort Metas_array array
            FrontRepoSingloton.Metas_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.MetaReferences_array = metareferences

            // clear the map that counts MetaReference in the GET
            FrontRepoSingloton.MetaReferences_batch.clear()

            metareferences.forEach(
              metareference => {
                FrontRepoSingloton.MetaReferences.set(metareference.ID, metareference)
                FrontRepoSingloton.MetaReferences_batch.set(metareference.ID, metareference)
              }
            )

            // clear metareferences that are absent from the batch
            FrontRepoSingloton.MetaReferences.forEach(
              metareference => {
                if (FrontRepoSingloton.MetaReferences_batch.get(metareference.ID) == undefined) {
                  FrontRepoSingloton.MetaReferences.delete(metareference.ID)
                }
              }
            )

            // sort MetaReferences_array array
            FrontRepoSingloton.MetaReferences_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.ModelPkgs_array = modelpkgs

            // clear the map that counts ModelPkg in the GET
            FrontRepoSingloton.ModelPkgs_batch.clear()

            modelpkgs.forEach(
              modelpkg => {
                FrontRepoSingloton.ModelPkgs.set(modelpkg.ID, modelpkg)
                FrontRepoSingloton.ModelPkgs_batch.set(modelpkg.ID, modelpkg)
              }
            )

            // clear modelpkgs that are absent from the batch
            FrontRepoSingloton.ModelPkgs.forEach(
              modelpkg => {
                if (FrontRepoSingloton.ModelPkgs_batch.get(modelpkg.ID) == undefined) {
                  FrontRepoSingloton.ModelPkgs.delete(modelpkg.ID)
                }
              }
            )

            // sort ModelPkgs_array array
            FrontRepoSingloton.ModelPkgs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.PointerToGongStructFields_array = pointertogongstructfields

            // clear the map that counts PointerToGongStructField in the GET
            FrontRepoSingloton.PointerToGongStructFields_batch.clear()

            pointertogongstructfields.forEach(
              pointertogongstructfield => {
                FrontRepoSingloton.PointerToGongStructFields.set(pointertogongstructfield.ID, pointertogongstructfield)
                FrontRepoSingloton.PointerToGongStructFields_batch.set(pointertogongstructfield.ID, pointertogongstructfield)
              }
            )

            // clear pointertogongstructfields that are absent from the batch
            FrontRepoSingloton.PointerToGongStructFields.forEach(
              pointertogongstructfield => {
                if (FrontRepoSingloton.PointerToGongStructFields_batch.get(pointertogongstructfield.ID) == undefined) {
                  FrontRepoSingloton.PointerToGongStructFields.delete(pointertogongstructfield.ID)
                }
              }
            )

            // sort PointerToGongStructFields_array array
            FrontRepoSingloton.PointerToGongStructFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.SliceOfPointerToGongStructFields_array = sliceofpointertogongstructfields

            // clear the map that counts SliceOfPointerToGongStructField in the GET
            FrontRepoSingloton.SliceOfPointerToGongStructFields_batch.clear()

            sliceofpointertogongstructfields.forEach(
              sliceofpointertogongstructfield => {
                FrontRepoSingloton.SliceOfPointerToGongStructFields.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)
                FrontRepoSingloton.SliceOfPointerToGongStructFields_batch.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)
              }
            )

            // clear sliceofpointertogongstructfields that are absent from the batch
            FrontRepoSingloton.SliceOfPointerToGongStructFields.forEach(
              sliceofpointertogongstructfield => {
                if (FrontRepoSingloton.SliceOfPointerToGongStructFields_batch.get(sliceofpointertogongstructfield.ID) == undefined) {
                  FrontRepoSingloton.SliceOfPointerToGongStructFields.delete(sliceofpointertogongstructfield.ID)
                }
              }
            )

            // sort SliceOfPointerToGongStructFields_array array
            FrontRepoSingloton.SliceOfPointerToGongStructFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            gongbasicfields.forEach(
              gongbasicfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field GongEnum redeeming
                {
                  let _gongenum = FrontRepoSingloton.GongEnums.get(gongbasicfield.GongEnumID.Int64)
                  if (_gongenum) {
                    gongbasicfield.GongEnum = _gongenum
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.GongBasicFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(gongbasicfield.GongStruct_GongBasicFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.GongBasicFields == undefined) {
                      _gongstruct.GongBasicFields = new Array<GongBasicFieldDB>()
                    }
                    _gongstruct.GongBasicFields.push(gongbasicfield)
                    if (gongbasicfield.GongStruct_GongBasicFields_reverse == undefined) {
                      gongbasicfield.GongStruct_GongBasicFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )
            gongenums.forEach(
              gongenum => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            gongenumvalues.forEach(
              gongenumvalue => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongEnum.GongEnumValues redeeming
                {
                  let _gongenum = FrontRepoSingloton.GongEnums.get(gongenumvalue.GongEnum_GongEnumValuesDBID.Int64)
                  if (_gongenum) {
                    if (_gongenum.GongEnumValues == undefined) {
                      _gongenum.GongEnumValues = new Array<GongEnumValueDB>()
                    }
                    _gongenum.GongEnumValues.push(gongenumvalue)
                    if (gongenumvalue.GongEnum_GongEnumValues_reverse == undefined) {
                      gongenumvalue.GongEnum_GongEnumValues_reverse = _gongenum
                    }
                  }
                }
              }
            )
            gonglinks.forEach(
              gonglink => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongNote.Links redeeming
                {
                  let _gongnote = FrontRepoSingloton.GongNotes.get(gonglink.GongNote_LinksDBID.Int64)
                  if (_gongnote) {
                    if (_gongnote.Links == undefined) {
                      _gongnote.Links = new Array<GongLinkDB>()
                    }
                    _gongnote.Links.push(gonglink)
                    if (gonglink.GongNote_Links_reverse == undefined) {
                      gonglink.GongNote_Links_reverse = _gongnote
                    }
                  }
                }
              }
            )
            gongnotes.forEach(
              gongnote => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            gongstructs.forEach(
              gongstruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            gongtimefields.forEach(
              gongtimefield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.GongTimeFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(gongtimefield.GongStruct_GongTimeFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.GongTimeFields == undefined) {
                      _gongstruct.GongTimeFields = new Array<GongTimeFieldDB>()
                    }
                    _gongstruct.GongTimeFields.push(gongtimefield)
                    if (gongtimefield.GongStruct_GongTimeFields_reverse == undefined) {
                      gongtimefield.GongStruct_GongTimeFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )
            metas.forEach(
              meta => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            metareferences.forEach(
              metareference => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Meta.MetaReferences redeeming
                {
                  let _meta = FrontRepoSingloton.Metas.get(metareference.Meta_MetaReferencesDBID.Int64)
                  if (_meta) {
                    if (_meta.MetaReferences == undefined) {
                      _meta.MetaReferences = new Array<MetaReferenceDB>()
                    }
                    _meta.MetaReferences.push(metareference)
                    if (metareference.Meta_MetaReferences_reverse == undefined) {
                      metareference.Meta_MetaReferences_reverse = _meta
                    }
                  }
                }
              }
            )
            modelpkgs.forEach(
              modelpkg => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            pointertogongstructfields.forEach(
              pointertogongstructfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(pointertogongstructfield.GongStructID.Int64)
                  if (_gongstruct) {
                    pointertogongstructfield.GongStruct = _gongstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.PointerToGongStructFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.PointerToGongStructFields == undefined) {
                      _gongstruct.PointerToGongStructFields = new Array<PointerToGongStructFieldDB>()
                    }
                    _gongstruct.PointerToGongStructFields.push(pointertogongstructfield)
                    if (pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse == undefined) {
                      pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )
            sliceofpointertogongstructfields.forEach(
              sliceofpointertogongstructfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(sliceofpointertogongstructfield.GongStructID.Int64)
                  if (_gongstruct) {
                    sliceofpointertogongstructfield.GongStruct = _gongstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.SliceOfPointerToGongStructFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.SliceOfPointerToGongStructFields == undefined) {
                      _gongstruct.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructFieldDB>()
                    }
                    _gongstruct.SliceOfPointerToGongStructFields.push(sliceofpointertogongstructfield)
                    if (sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse == undefined) {
                      sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.gongbasicfieldService.getGongBasicFields()
        ]).subscribe(
          ([ // insertion point sub template 
            gongbasicfields,
          ]) => {
            // init the array
            FrontRepoSingloton.GongBasicFields_array = gongbasicfields

            // clear the map that counts GongBasicField in the GET
            FrontRepoSingloton.GongBasicFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongbasicfields.forEach(
              gongbasicfield => {
                FrontRepoSingloton.GongBasicFields.set(gongbasicfield.ID, gongbasicfield)
                FrontRepoSingloton.GongBasicFields_batch.set(gongbasicfield.ID, gongbasicfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field GongEnum redeeming
                {
                  let _gongenum = FrontRepoSingloton.GongEnums.get(gongbasicfield.GongEnumID.Int64)
                  if (_gongenum) {
                    gongbasicfield.GongEnum = _gongenum
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.GongBasicFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(gongbasicfield.GongStruct_GongBasicFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.GongBasicFields == undefined) {
                      _gongstruct.GongBasicFields = new Array<GongBasicFieldDB>()
                    }
                    _gongstruct.GongBasicFields.push(gongbasicfield)
                    if (gongbasicfield.GongStruct_GongBasicFields_reverse == undefined) {
                      gongbasicfield.GongStruct_GongBasicFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )

            // clear gongbasicfields that are absent from the GET
            FrontRepoSingloton.GongBasicFields.forEach(
              gongbasicfield => {
                if (FrontRepoSingloton.GongBasicFields_batch.get(gongbasicfield.ID) == undefined) {
                  FrontRepoSingloton.GongBasicFields.delete(gongbasicfield.ID)
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
  }

  // GongEnumPull performs a GET on GongEnum of the stack and redeem association pointers 
  GongEnumPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongenumService.getGongEnums()
        ]).subscribe(
          ([ // insertion point sub template 
            gongenums,
          ]) => {
            // init the array
            FrontRepoSingloton.GongEnums_array = gongenums

            // clear the map that counts GongEnum in the GET
            FrontRepoSingloton.GongEnums_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongenums.forEach(
              gongenum => {
                FrontRepoSingloton.GongEnums.set(gongenum.ID, gongenum)
                FrontRepoSingloton.GongEnums_batch.set(gongenum.ID, gongenum)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear gongenums that are absent from the GET
            FrontRepoSingloton.GongEnums.forEach(
              gongenum => {
                if (FrontRepoSingloton.GongEnums_batch.get(gongenum.ID) == undefined) {
                  FrontRepoSingloton.GongEnums.delete(gongenum.ID)
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
  }

  // GongEnumValuePull performs a GET on GongEnumValue of the stack and redeem association pointers 
  GongEnumValuePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongenumvalueService.getGongEnumValues()
        ]).subscribe(
          ([ // insertion point sub template 
            gongenumvalues,
          ]) => {
            // init the array
            FrontRepoSingloton.GongEnumValues_array = gongenumvalues

            // clear the map that counts GongEnumValue in the GET
            FrontRepoSingloton.GongEnumValues_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongenumvalues.forEach(
              gongenumvalue => {
                FrontRepoSingloton.GongEnumValues.set(gongenumvalue.ID, gongenumvalue)
                FrontRepoSingloton.GongEnumValues_batch.set(gongenumvalue.ID, gongenumvalue)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongEnum.GongEnumValues redeeming
                {
                  let _gongenum = FrontRepoSingloton.GongEnums.get(gongenumvalue.GongEnum_GongEnumValuesDBID.Int64)
                  if (_gongenum) {
                    if (_gongenum.GongEnumValues == undefined) {
                      _gongenum.GongEnumValues = new Array<GongEnumValueDB>()
                    }
                    _gongenum.GongEnumValues.push(gongenumvalue)
                    if (gongenumvalue.GongEnum_GongEnumValues_reverse == undefined) {
                      gongenumvalue.GongEnum_GongEnumValues_reverse = _gongenum
                    }
                  }
                }
              }
            )

            // clear gongenumvalues that are absent from the GET
            FrontRepoSingloton.GongEnumValues.forEach(
              gongenumvalue => {
                if (FrontRepoSingloton.GongEnumValues_batch.get(gongenumvalue.ID) == undefined) {
                  FrontRepoSingloton.GongEnumValues.delete(gongenumvalue.ID)
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
  }

  // GongLinkPull performs a GET on GongLink of the stack and redeem association pointers 
  GongLinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gonglinkService.getGongLinks()
        ]).subscribe(
          ([ // insertion point sub template 
            gonglinks,
          ]) => {
            // init the array
            FrontRepoSingloton.GongLinks_array = gonglinks

            // clear the map that counts GongLink in the GET
            FrontRepoSingloton.GongLinks_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gonglinks.forEach(
              gonglink => {
                FrontRepoSingloton.GongLinks.set(gonglink.ID, gonglink)
                FrontRepoSingloton.GongLinks_batch.set(gonglink.ID, gonglink)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongNote.Links redeeming
                {
                  let _gongnote = FrontRepoSingloton.GongNotes.get(gonglink.GongNote_LinksDBID.Int64)
                  if (_gongnote) {
                    if (_gongnote.Links == undefined) {
                      _gongnote.Links = new Array<GongLinkDB>()
                    }
                    _gongnote.Links.push(gonglink)
                    if (gonglink.GongNote_Links_reverse == undefined) {
                      gonglink.GongNote_Links_reverse = _gongnote
                    }
                  }
                }
              }
            )

            // clear gonglinks that are absent from the GET
            FrontRepoSingloton.GongLinks.forEach(
              gonglink => {
                if (FrontRepoSingloton.GongLinks_batch.get(gonglink.ID) == undefined) {
                  FrontRepoSingloton.GongLinks.delete(gonglink.ID)
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
  }

  // GongNotePull performs a GET on GongNote of the stack and redeem association pointers 
  GongNotePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongnoteService.getGongNotes()
        ]).subscribe(
          ([ // insertion point sub template 
            gongnotes,
          ]) => {
            // init the array
            FrontRepoSingloton.GongNotes_array = gongnotes

            // clear the map that counts GongNote in the GET
            FrontRepoSingloton.GongNotes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongnotes.forEach(
              gongnote => {
                FrontRepoSingloton.GongNotes.set(gongnote.ID, gongnote)
                FrontRepoSingloton.GongNotes_batch.set(gongnote.ID, gongnote)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear gongnotes that are absent from the GET
            FrontRepoSingloton.GongNotes.forEach(
              gongnote => {
                if (FrontRepoSingloton.GongNotes_batch.get(gongnote.ID) == undefined) {
                  FrontRepoSingloton.GongNotes.delete(gongnote.ID)
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
  }

  // GongStructPull performs a GET on GongStruct of the stack and redeem association pointers 
  GongStructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongstructService.getGongStructs()
        ]).subscribe(
          ([ // insertion point sub template 
            gongstructs,
          ]) => {
            // init the array
            FrontRepoSingloton.GongStructs_array = gongstructs

            // clear the map that counts GongStruct in the GET
            FrontRepoSingloton.GongStructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongstructs.forEach(
              gongstruct => {
                FrontRepoSingloton.GongStructs.set(gongstruct.ID, gongstruct)
                FrontRepoSingloton.GongStructs_batch.set(gongstruct.ID, gongstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear gongstructs that are absent from the GET
            FrontRepoSingloton.GongStructs.forEach(
              gongstruct => {
                if (FrontRepoSingloton.GongStructs_batch.get(gongstruct.ID) == undefined) {
                  FrontRepoSingloton.GongStructs.delete(gongstruct.ID)
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
  }

  // GongTimeFieldPull performs a GET on GongTimeField of the stack and redeem association pointers 
  GongTimeFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongtimefieldService.getGongTimeFields()
        ]).subscribe(
          ([ // insertion point sub template 
            gongtimefields,
          ]) => {
            // init the array
            FrontRepoSingloton.GongTimeFields_array = gongtimefields

            // clear the map that counts GongTimeField in the GET
            FrontRepoSingloton.GongTimeFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongtimefields.forEach(
              gongtimefield => {
                FrontRepoSingloton.GongTimeFields.set(gongtimefield.ID, gongtimefield)
                FrontRepoSingloton.GongTimeFields_batch.set(gongtimefield.ID, gongtimefield)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.GongTimeFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(gongtimefield.GongStruct_GongTimeFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.GongTimeFields == undefined) {
                      _gongstruct.GongTimeFields = new Array<GongTimeFieldDB>()
                    }
                    _gongstruct.GongTimeFields.push(gongtimefield)
                    if (gongtimefield.GongStruct_GongTimeFields_reverse == undefined) {
                      gongtimefield.GongStruct_GongTimeFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )

            // clear gongtimefields that are absent from the GET
            FrontRepoSingloton.GongTimeFields.forEach(
              gongtimefield => {
                if (FrontRepoSingloton.GongTimeFields_batch.get(gongtimefield.ID) == undefined) {
                  FrontRepoSingloton.GongTimeFields.delete(gongtimefield.ID)
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
  }

  // MetaPull performs a GET on Meta of the stack and redeem association pointers 
  MetaPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.metaService.getMetas()
        ]).subscribe(
          ([ // insertion point sub template 
            metas,
          ]) => {
            // init the array
            FrontRepoSingloton.Metas_array = metas

            // clear the map that counts Meta in the GET
            FrontRepoSingloton.Metas_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            metas.forEach(
              meta => {
                FrontRepoSingloton.Metas.set(meta.ID, meta)
                FrontRepoSingloton.Metas_batch.set(meta.ID, meta)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear metas that are absent from the GET
            FrontRepoSingloton.Metas.forEach(
              meta => {
                if (FrontRepoSingloton.Metas_batch.get(meta.ID) == undefined) {
                  FrontRepoSingloton.Metas.delete(meta.ID)
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
  }

  // MetaReferencePull performs a GET on MetaReference of the stack and redeem association pointers 
  MetaReferencePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.metareferenceService.getMetaReferences()
        ]).subscribe(
          ([ // insertion point sub template 
            metareferences,
          ]) => {
            // init the array
            FrontRepoSingloton.MetaReferences_array = metareferences

            // clear the map that counts MetaReference in the GET
            FrontRepoSingloton.MetaReferences_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            metareferences.forEach(
              metareference => {
                FrontRepoSingloton.MetaReferences.set(metareference.ID, metareference)
                FrontRepoSingloton.MetaReferences_batch.set(metareference.ID, metareference)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Meta.MetaReferences redeeming
                {
                  let _meta = FrontRepoSingloton.Metas.get(metareference.Meta_MetaReferencesDBID.Int64)
                  if (_meta) {
                    if (_meta.MetaReferences == undefined) {
                      _meta.MetaReferences = new Array<MetaReferenceDB>()
                    }
                    _meta.MetaReferences.push(metareference)
                    if (metareference.Meta_MetaReferences_reverse == undefined) {
                      metareference.Meta_MetaReferences_reverse = _meta
                    }
                  }
                }
              }
            )

            // clear metareferences that are absent from the GET
            FrontRepoSingloton.MetaReferences.forEach(
              metareference => {
                if (FrontRepoSingloton.MetaReferences_batch.get(metareference.ID) == undefined) {
                  FrontRepoSingloton.MetaReferences.delete(metareference.ID)
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
  }

  // ModelPkgPull performs a GET on ModelPkg of the stack and redeem association pointers 
  ModelPkgPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.modelpkgService.getModelPkgs()
        ]).subscribe(
          ([ // insertion point sub template 
            modelpkgs,
          ]) => {
            // init the array
            FrontRepoSingloton.ModelPkgs_array = modelpkgs

            // clear the map that counts ModelPkg in the GET
            FrontRepoSingloton.ModelPkgs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            modelpkgs.forEach(
              modelpkg => {
                FrontRepoSingloton.ModelPkgs.set(modelpkg.ID, modelpkg)
                FrontRepoSingloton.ModelPkgs_batch.set(modelpkg.ID, modelpkg)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear modelpkgs that are absent from the GET
            FrontRepoSingloton.ModelPkgs.forEach(
              modelpkg => {
                if (FrontRepoSingloton.ModelPkgs_batch.get(modelpkg.ID) == undefined) {
                  FrontRepoSingloton.ModelPkgs.delete(modelpkg.ID)
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
  }

  // PointerToGongStructFieldPull performs a GET on PointerToGongStructField of the stack and redeem association pointers 
  PointerToGongStructFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pointertogongstructfieldService.getPointerToGongStructFields()
        ]).subscribe(
          ([ // insertion point sub template 
            pointertogongstructfields,
          ]) => {
            // init the array
            FrontRepoSingloton.PointerToGongStructFields_array = pointertogongstructfields

            // clear the map that counts PointerToGongStructField in the GET
            FrontRepoSingloton.PointerToGongStructFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            pointertogongstructfields.forEach(
              pointertogongstructfield => {
                FrontRepoSingloton.PointerToGongStructFields.set(pointertogongstructfield.ID, pointertogongstructfield)
                FrontRepoSingloton.PointerToGongStructFields_batch.set(pointertogongstructfield.ID, pointertogongstructfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(pointertogongstructfield.GongStructID.Int64)
                  if (_gongstruct) {
                    pointertogongstructfield.GongStruct = _gongstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.PointerToGongStructFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(pointertogongstructfield.GongStruct_PointerToGongStructFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.PointerToGongStructFields == undefined) {
                      _gongstruct.PointerToGongStructFields = new Array<PointerToGongStructFieldDB>()
                    }
                    _gongstruct.PointerToGongStructFields.push(pointertogongstructfield)
                    if (pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse == undefined) {
                      pointertogongstructfield.GongStruct_PointerToGongStructFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )

            // clear pointertogongstructfields that are absent from the GET
            FrontRepoSingloton.PointerToGongStructFields.forEach(
              pointertogongstructfield => {
                if (FrontRepoSingloton.PointerToGongStructFields_batch.get(pointertogongstructfield.ID) == undefined) {
                  FrontRepoSingloton.PointerToGongStructFields.delete(pointertogongstructfield.ID)
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
  }

  // SliceOfPointerToGongStructFieldPull performs a GET on SliceOfPointerToGongStructField of the stack and redeem association pointers 
  SliceOfPointerToGongStructFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.sliceofpointertogongstructfieldService.getSliceOfPointerToGongStructFields()
        ]).subscribe(
          ([ // insertion point sub template 
            sliceofpointertogongstructfields,
          ]) => {
            // init the array
            FrontRepoSingloton.SliceOfPointerToGongStructFields_array = sliceofpointertogongstructfields

            // clear the map that counts SliceOfPointerToGongStructField in the GET
            FrontRepoSingloton.SliceOfPointerToGongStructFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            sliceofpointertogongstructfields.forEach(
              sliceofpointertogongstructfield => {
                FrontRepoSingloton.SliceOfPointerToGongStructFields.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)
                FrontRepoSingloton.SliceOfPointerToGongStructFields_batch.set(sliceofpointertogongstructfield.ID, sliceofpointertogongstructfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field GongStruct redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(sliceofpointertogongstructfield.GongStructID.Int64)
                  if (_gongstruct) {
                    sliceofpointertogongstructfield.GongStruct = _gongstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStruct.SliceOfPointerToGongStructFields redeeming
                {
                  let _gongstruct = FrontRepoSingloton.GongStructs.get(sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64)
                  if (_gongstruct) {
                    if (_gongstruct.SliceOfPointerToGongStructFields == undefined) {
                      _gongstruct.SliceOfPointerToGongStructFields = new Array<SliceOfPointerToGongStructFieldDB>()
                    }
                    _gongstruct.SliceOfPointerToGongStructFields.push(sliceofpointertogongstructfield)
                    if (sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse == undefined) {
                      sliceofpointertogongstructfield.GongStruct_SliceOfPointerToGongStructFields_reverse = _gongstruct
                    }
                  }
                }
              }
            )

            // clear sliceofpointertogongstructfields that are absent from the GET
            FrontRepoSingloton.SliceOfPointerToGongStructFields.forEach(
              sliceofpointertogongstructfield => {
                if (FrontRepoSingloton.SliceOfPointerToGongStructFields_batch.get(sliceofpointertogongstructfield.ID) == undefined) {
                  FrontRepoSingloton.SliceOfPointerToGongStructFields.delete(sliceofpointertogongstructfield.ID)
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
