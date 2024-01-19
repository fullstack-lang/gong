import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AstructDB } from './astruct-db'
import { Astruct, CopyAstructDBToAstruct } from './astruct'
import { AstructService } from './astruct.service'

import { AstructBstruct2UseDB } from './astructbstruct2use-db'
import { AstructBstruct2Use, CopyAstructBstruct2UseDBToAstructBstruct2Use } from './astructbstruct2use'
import { AstructBstruct2UseService } from './astructbstruct2use.service'

import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstructUse, CopyAstructBstructUseDBToAstructBstructUse } from './astructbstructuse'
import { AstructBstructUseService } from './astructbstructuse.service'

import { BstructDB } from './bstruct-db'
import { Bstruct, CopyBstructDBToBstruct } from './bstruct'
import { BstructService } from './bstruct.service'

import { DstructDB } from './dstruct-db'
import { Dstruct, CopyDstructDBToDstruct } from './dstruct'
import { DstructService } from './dstruct.service'

export const StackType = "github.com/fullstack-lang/gong/test/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  Astructs_array = new Array<AstructDB>() // array of repo instances
  Astructs = new Map<number, AstructDB>() // map of repo instances
  Astructs_batch = new Map<number, AstructDB>() // same but only in last GET (for finding repo instances to delete)

  array_Astructs = new Array<Astruct>() // array of front instances
  map_ID_Astruct = new Map<number, Astruct>() // map of front instances

  AstructBstruct2Uses_array = new Array<AstructBstruct2UseDB>() // array of repo instances
  AstructBstruct2Uses = new Map<number, AstructBstruct2UseDB>() // map of repo instances
  AstructBstruct2Uses_batch = new Map<number, AstructBstruct2UseDB>() // same but only in last GET (for finding repo instances to delete)

  array_AstructBstruct2Uses = new Array<AstructBstruct2Use>() // array of front instances
  map_ID_AstructBstruct2Use = new Map<number, AstructBstruct2Use>() // map of front instances

  AstructBstructUses_array = new Array<AstructBstructUseDB>() // array of repo instances
  AstructBstructUses = new Map<number, AstructBstructUseDB>() // map of repo instances
  AstructBstructUses_batch = new Map<number, AstructBstructUseDB>() // same but only in last GET (for finding repo instances to delete)

  array_AstructBstructUses = new Array<AstructBstructUse>() // array of front instances
  map_ID_AstructBstructUse = new Map<number, AstructBstructUse>() // map of front instances

  Bstructs_array = new Array<BstructDB>() // array of repo instances
  Bstructs = new Map<number, BstructDB>() // map of repo instances
  Bstructs_batch = new Map<number, BstructDB>() // same but only in last GET (for finding repo instances to delete)

  array_Bstructs = new Array<Bstruct>() // array of front instances
  map_ID_Bstruct = new Map<number, Bstruct>() // map of front instances

  Dstructs_array = new Array<DstructDB>() // array of repo instances
  Dstructs = new Map<number, DstructDB>() // map of repo instances
  Dstructs_batch = new Map<number, DstructDB>() // same but only in last GET (for finding repo instances to delete)

  array_Dstructs = new Array<Dstruct>() // array of front instances
  map_ID_Dstruct = new Map<number, Dstruct>() // map of front instances


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) { // deprecated
      // insertion point
      case 'Astruct':
        return this.Astructs_array as unknown as Array<Type>
      case 'AstructBstruct2Use':
        return this.AstructBstruct2Uses_array as unknown as Array<Type>
      case 'AstructBstructUse':
        return this.AstructBstructUses_array as unknown as Array<Type>
      case 'Bstruct':
        return this.Bstructs_array as unknown as Array<Type>
      case 'Dstruct':
        return this.Dstructs_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  getFrontArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'Astruct':
        return this.array_Astructs as unknown as Array<Type>
      case 'AstructBstruct2Use':
        return this.array_AstructBstruct2Uses as unknown as Array<Type>
      case 'AstructBstructUse':
        return this.array_AstructBstructUses as unknown as Array<Type>
      case 'Bstruct':
        return this.array_Bstructs as unknown as Array<Type>
      case 'Dstruct':
        return this.array_Dstructs as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }


  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> { // deprecated
    switch (gongStructName) {
      // insertion point
      case 'Astruct':
        return this.Astructs as unknown as Map<number, Type>
      case 'AstructBstruct2Use':
        return this.AstructBstruct2Uses as unknown as Map<number, Type>
      case 'AstructBstructUse':
        return this.AstructBstructUses as unknown as Map<number, Type>
      case 'Bstruct':
        return this.Bstructs as unknown as Map<number, Type>
      case 'Dstruct':
        return this.Dstructs as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
  getFrontMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'Astruct':
        return this.map_ID_Astruct as unknown as Map<number, Type>
      case 'AstructBstruct2Use':
        return this.map_ID_AstructBstruct2Use as unknown as Map<number, Type>
      case 'AstructBstructUse':
        return this.map_ID_AstructBstructUse as unknown as Map<number, Type>
      case 'Bstruct':
        return this.map_ID_Bstruct as unknown as Map<number, Type>
      case 'Dstruct':
        return this.map_ID_Dstruct as unknown as Map<number, Type>
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
    private astructService: AstructService,
    private astructbstruct2useService: AstructBstruct2UseService,
    private astructbstructuseService: AstructBstructUseService,
    private bstructService: BstructService,
    private dstructService: DstructService,
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
    Observable<AstructDB[]>,
    Observable<AstructBstruct2UseDB[]>,
    Observable<AstructBstructUseDB[]>,
    Observable<BstructDB[]>,
    Observable<DstructDB[]>,
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
      this.astructService.getAstructs(this.GONG__StackPath, this.frontRepo),
      this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath, this.frontRepo),
      this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath, this.frontRepo),
      this.bstructService.getBstructs(this.GONG__StackPath, this.frontRepo),
      this.dstructService.getDstructs(this.GONG__StackPath, this.frontRepo),
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
      this.astructService.getAstructs(this.GONG__StackPath, this.frontRepo),
      this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath, this.frontRepo),
      this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath, this.frontRepo),
      this.bstructService.getBstructs(this.GONG__StackPath, this.frontRepo),
      this.dstructService.getDstructs(this.GONG__StackPath, this.frontRepo),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            astructs_,
            astructbstruct2uses_,
            astructbstructuses_,
            bstructs_,
            dstructs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var astructs: AstructDB[]
            astructs = astructs_ as AstructDB[]
            var astructbstruct2uses: AstructBstruct2UseDB[]
            astructbstruct2uses = astructbstruct2uses_ as AstructBstruct2UseDB[]
            var astructbstructuses: AstructBstructUseDB[]
            astructbstructuses = astructbstructuses_ as AstructBstructUseDB[]
            var bstructs: BstructDB[]
            bstructs = bstructs_ as BstructDB[]
            var dstructs: DstructDB[]
            dstructs = dstructs_ as DstructDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Astructs_array = astructs

            // clear the map that counts Astruct in the GET
            this.frontRepo.Astructs_batch.clear()

            astructs.forEach(
              astructDB => {
                this.frontRepo.Astructs.set(astructDB.ID, astructDB)
                this.frontRepo.Astructs_batch.set(astructDB.ID, astructDB)
              }
            )

            // clear astructs that are absent from the batch
            this.frontRepo.Astructs.forEach(
              astructDB => {
                if (this.frontRepo.Astructs_batch.get(astructDB.ID) == undefined) {
                  this.frontRepo.Astructs.delete(astructDB.ID)
                }
              }
            )

            // sort Astructs_array array
            this.frontRepo.Astructs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init front objects
            this.frontRepo.array_Astructs = []
            this.frontRepo.map_ID_Astruct.clear()
            this.frontRepo.Astructs_array.forEach(
              astructDB => {
                let astruct = new Astruct
                CopyAstructDBToAstruct(astructDB, astruct, this.frontRepo)
                this.frontRepo.array_Astructs.push(astruct)
                this.frontRepo.map_ID_Astruct.set(astruct.ID, astruct)
              }
            )

            // init the array
            this.frontRepo.AstructBstruct2Uses_array = astructbstruct2uses

            // clear the map that counts AstructBstruct2Use in the GET
            this.frontRepo.AstructBstruct2Uses_batch.clear()

            astructbstruct2uses.forEach(
              astructbstruct2useDB => {
                this.frontRepo.AstructBstruct2Uses.set(astructbstruct2useDB.ID, astructbstruct2useDB)
                this.frontRepo.AstructBstruct2Uses_batch.set(astructbstruct2useDB.ID, astructbstruct2useDB)
              }
            )

            // clear astructbstruct2uses that are absent from the batch
            this.frontRepo.AstructBstruct2Uses.forEach(
              astructbstruct2useDB => {
                if (this.frontRepo.AstructBstruct2Uses_batch.get(astructbstruct2useDB.ID) == undefined) {
                  this.frontRepo.AstructBstruct2Uses.delete(astructbstruct2useDB.ID)
                }
              }
            )

            // sort AstructBstruct2Uses_array array
            this.frontRepo.AstructBstruct2Uses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init front objects
            this.frontRepo.array_AstructBstruct2Uses = []
            this.frontRepo.map_ID_AstructBstruct2Use.clear()
            this.frontRepo.AstructBstruct2Uses_array.forEach(
              astructbstruct2useDB => {
                let astructbstruct2use = new AstructBstruct2Use
                CopyAstructBstruct2UseDBToAstructBstruct2Use(astructbstruct2useDB, astructbstruct2use, this.frontRepo)
                this.frontRepo.array_AstructBstruct2Uses.push(astructbstruct2use)
                this.frontRepo.map_ID_AstructBstruct2Use.set(astructbstruct2use.ID, astructbstruct2use)
              }
            )

            // init the array
            this.frontRepo.AstructBstructUses_array = astructbstructuses

            // clear the map that counts AstructBstructUse in the GET
            this.frontRepo.AstructBstructUses_batch.clear()

            astructbstructuses.forEach(
              astructbstructuseDB => {
                this.frontRepo.AstructBstructUses.set(astructbstructuseDB.ID, astructbstructuseDB)
                this.frontRepo.AstructBstructUses_batch.set(astructbstructuseDB.ID, astructbstructuseDB)
              }
            )

            // clear astructbstructuses that are absent from the batch
            this.frontRepo.AstructBstructUses.forEach(
              astructbstructuseDB => {
                if (this.frontRepo.AstructBstructUses_batch.get(astructbstructuseDB.ID) == undefined) {
                  this.frontRepo.AstructBstructUses.delete(astructbstructuseDB.ID)
                }
              }
            )

            // sort AstructBstructUses_array array
            this.frontRepo.AstructBstructUses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init front objects
            this.frontRepo.array_AstructBstructUses = []
            this.frontRepo.map_ID_AstructBstructUse.clear()
            this.frontRepo.AstructBstructUses_array.forEach(
              astructbstructuseDB => {
                let astructbstructuse = new AstructBstructUse
                CopyAstructBstructUseDBToAstructBstructUse(astructbstructuseDB, astructbstructuse, this.frontRepo)
                this.frontRepo.array_AstructBstructUses.push(astructbstructuse)
                this.frontRepo.map_ID_AstructBstructUse.set(astructbstructuse.ID, astructbstructuse)
              }
            )

            // init the array
            this.frontRepo.Bstructs_array = bstructs

            // clear the map that counts Bstruct in the GET
            this.frontRepo.Bstructs_batch.clear()

            bstructs.forEach(
              bstructDB => {
                this.frontRepo.Bstructs.set(bstructDB.ID, bstructDB)
                this.frontRepo.Bstructs_batch.set(bstructDB.ID, bstructDB)
              }
            )

            // clear bstructs that are absent from the batch
            this.frontRepo.Bstructs.forEach(
              bstructDB => {
                if (this.frontRepo.Bstructs_batch.get(bstructDB.ID) == undefined) {
                  this.frontRepo.Bstructs.delete(bstructDB.ID)
                }
              }
            )

            // sort Bstructs_array array
            this.frontRepo.Bstructs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init front objects
            this.frontRepo.array_Bstructs = []
            this.frontRepo.map_ID_Bstruct.clear()
            this.frontRepo.Bstructs_array.forEach(
              bstructDB => {
                let bstruct = new Bstruct
                CopyBstructDBToBstruct(bstructDB, bstruct, this.frontRepo)
                this.frontRepo.array_Bstructs.push(bstruct)
                this.frontRepo.map_ID_Bstruct.set(bstruct.ID, bstruct)
              }
            )

            // init the array
            this.frontRepo.Dstructs_array = dstructs

            // clear the map that counts Dstruct in the GET
            this.frontRepo.Dstructs_batch.clear()

            dstructs.forEach(
              dstructDB => {
                this.frontRepo.Dstructs.set(dstructDB.ID, dstructDB)
                this.frontRepo.Dstructs_batch.set(dstructDB.ID, dstructDB)
              }
            )

            // clear dstructs that are absent from the batch
            this.frontRepo.Dstructs.forEach(
              dstructDB => {
                if (this.frontRepo.Dstructs_batch.get(dstructDB.ID) == undefined) {
                  this.frontRepo.Dstructs.delete(dstructDB.ID)
                }
              }
            )

            // sort Dstructs_array array
            this.frontRepo.Dstructs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init front objects
            this.frontRepo.array_Dstructs = []
            this.frontRepo.map_ID_Dstruct.clear()
            this.frontRepo.Dstructs_array.forEach(
              dstructDB => {
                let dstruct = new Dstruct
                CopyDstructDBToDstruct(dstructDB, dstruct, this.frontRepo)
                this.frontRepo.array_Dstructs.push(dstruct)
                this.frontRepo.map_ID_Dstruct.set(dstruct.ID, dstruct)
              }
            )


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            astructs.forEach(
              astruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Associationtob redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.AssociationtobID.Int64)
                  if (_bstruct) {
                    astruct.Associationtob = _bstruct
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.Anotherassociationtob_2ID.Int64)
                  if (_bstruct) {
                    astruct.Anotherassociationtob_2 = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.BstructID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct2 = _bstruct
                  }
                }
                // insertion point for pointer field Dstruct redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.DstructID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct2 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.Dstruct2ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct2 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct3 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.Dstruct3ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct3 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct4 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.Dstruct4ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct4 = _dstruct
                  }
                }
                // insertion point for pointer field AnAstruct redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astruct.AstructPointersEncoding.AnAstructID.Int64)
                  if (_astruct) {
                    astruct.AnAstruct = _astruct
                  }
                }
                // insertion point for pointers decoding
                astruct.Anarrayofb = new Array<BstructDB>()
                for (let _id of astruct.AstructPointersEncoding.Anarrayofb) {
                  let _bstruct = this.frontRepo.Bstructs.get(_id)
                  if (_bstruct != undefined) {
                    astruct.Anarrayofb.push(_bstruct!)
                  }
                }
                astruct.Anarrayofa = new Array<AstructDB>()
                for (let _id of astruct.AstructPointersEncoding.Anarrayofa) {
                  let _astruct = this.frontRepo.Astructs.get(_id)
                  if (_astruct != undefined) {
                    astruct.Anarrayofa.push(_astruct!)
                  }
                }
                astruct.Anotherarrayofb = new Array<BstructDB>()
                for (let _id of astruct.AstructPointersEncoding.Anotherarrayofb) {
                  let _bstruct = this.frontRepo.Bstructs.get(_id)
                  if (_bstruct != undefined) {
                    astruct.Anotherarrayofb.push(_bstruct!)
                  }
                }
                astruct.AnarrayofbUse = new Array<AstructBstructUseDB>()
                for (let _id of astruct.AstructPointersEncoding.AnarrayofbUse) {
                  let _astructbstructuse = this.frontRepo.AstructBstructUses.get(_id)
                  if (_astructbstructuse != undefined) {
                    astruct.AnarrayofbUse.push(_astructbstructuse!)
                  }
                }
                astruct.Anarrayofb2Use = new Array<AstructBstruct2UseDB>()
                for (let _id of astruct.AstructPointersEncoding.Anarrayofb2Use) {
                  let _astructbstruct2use = this.frontRepo.AstructBstruct2Uses.get(_id)
                  if (_astructbstruct2use != undefined) {
                    astruct.Anarrayofb2Use.push(_astructbstruct2use!)
                  }
                }
              }
            )
            astructbstruct2uses.forEach(
              astructbstruct2use => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Bstrcut2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstruct2use.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64)
                  if (_bstruct) {
                    astructbstruct2use.Bstrcut2 = _bstruct
                  }
                }
                // insertion point for pointers decoding
              }
            )
            astructbstructuses.forEach(
              astructbstructuse => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstructuse.AstructBstructUsePointersEncoding.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astructbstructuse.Bstruct2 = _bstruct
                  }
                }
                // insertion point for pointers decoding
              }
            )
            bstructs.forEach(
              bstruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            dstructs.forEach(
              dstruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                dstruct.Anarrayofb = new Array<BstructDB>()
                for (let _id of dstruct.DstructPointersEncoding.Anarrayofb) {
                  let _bstruct = this.frontRepo.Bstructs.get(_id)
                  if (_bstruct != undefined) {
                    dstruct.Anarrayofb.push(_bstruct!)
                  }
                }
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

  // AstructPull performs a GET on Astruct of the stack and redeem association pointers 
  AstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructService.getAstructs(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            astructs,
          ]) => {
            // init the array
            this.frontRepo.Astructs_array = astructs

            // clear the map that counts Astruct in the GET
            this.frontRepo.Astructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            astructs.forEach(
              astruct => {
                this.frontRepo.Astructs.set(astruct.ID, astruct)
                this.frontRepo.Astructs_batch.set(astruct.ID, astruct)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Associationtob redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.AssociationtobID.Int64)
                  if (_bstruct) {
                    astruct.Associationtob = _bstruct
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.Anotherassociationtob_2ID.Int64)
                  if (_bstruct) {
                    astruct.Anotherassociationtob_2 = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.BstructID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AstructPointersEncoding.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct2 = _bstruct
                  }
                }
                // insertion point for pointer field Dstruct redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.DstructID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct2 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.Dstruct2ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct2 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct3 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.Dstruct3ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct3 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct4 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.AstructPointersEncoding.Dstruct4ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct4 = _dstruct
                  }
                }
                // insertion point for pointer field AnAstruct redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astruct.AstructPointersEncoding.AnAstructID.Int64)
                  if (_astruct) {
                    astruct.AnAstruct = _astruct
                  }
                }
              }
            )

            // clear astructs that are absent from the GET
            this.frontRepo.Astructs.forEach(
              astruct => {
                if (this.frontRepo.Astructs_batch.get(astruct.ID) == undefined) {
                  this.frontRepo.Astructs.delete(astruct.ID)
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

  // AstructBstruct2UsePull performs a GET on AstructBstruct2Use of the stack and redeem association pointers 
  AstructBstruct2UsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            astructbstruct2uses,
          ]) => {
            // init the array
            this.frontRepo.AstructBstruct2Uses_array = astructbstruct2uses

            // clear the map that counts AstructBstruct2Use in the GET
            this.frontRepo.AstructBstruct2Uses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            astructbstruct2uses.forEach(
              astructbstruct2use => {
                this.frontRepo.AstructBstruct2Uses.set(astructbstruct2use.ID, astructbstruct2use)
                this.frontRepo.AstructBstruct2Uses_batch.set(astructbstruct2use.ID, astructbstruct2use)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Bstrcut2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstruct2use.AstructBstruct2UsePointersEncoding.Bstrcut2ID.Int64)
                  if (_bstruct) {
                    astructbstruct2use.Bstrcut2 = _bstruct
                  }
                }
              }
            )

            // clear astructbstruct2uses that are absent from the GET
            this.frontRepo.AstructBstruct2Uses.forEach(
              astructbstruct2use => {
                if (this.frontRepo.AstructBstruct2Uses_batch.get(astructbstruct2use.ID) == undefined) {
                  this.frontRepo.AstructBstruct2Uses.delete(astructbstruct2use.ID)
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

  // AstructBstructUsePull performs a GET on AstructBstructUse of the stack and redeem association pointers 
  AstructBstructUsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            astructbstructuses,
          ]) => {
            // init the array
            this.frontRepo.AstructBstructUses_array = astructbstructuses

            // clear the map that counts AstructBstructUse in the GET
            this.frontRepo.AstructBstructUses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            astructbstructuses.forEach(
              astructbstructuse => {
                this.frontRepo.AstructBstructUses.set(astructbstructuse.ID, astructbstructuse)
                this.frontRepo.AstructBstructUses_batch.set(astructbstructuse.ID, astructbstructuse)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstructuse.AstructBstructUsePointersEncoding.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astructbstructuse.Bstruct2 = _bstruct
                  }
                }
              }
            )

            // clear astructbstructuses that are absent from the GET
            this.frontRepo.AstructBstructUses.forEach(
              astructbstructuse => {
                if (this.frontRepo.AstructBstructUses_batch.get(astructbstructuse.ID) == undefined) {
                  this.frontRepo.AstructBstructUses.delete(astructbstructuse.ID)
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

  // BstructPull performs a GET on Bstruct of the stack and redeem association pointers 
  BstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.bstructService.getBstructs(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            bstructs,
          ]) => {
            // init the array
            this.frontRepo.Bstructs_array = bstructs

            // clear the map that counts Bstruct in the GET
            this.frontRepo.Bstructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bstructs.forEach(
              bstruct => {
                this.frontRepo.Bstructs.set(bstruct.ID, bstruct)
                this.frontRepo.Bstructs_batch.set(bstruct.ID, bstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear bstructs that are absent from the GET
            this.frontRepo.Bstructs.forEach(
              bstruct => {
                if (this.frontRepo.Bstructs_batch.get(bstruct.ID) == undefined) {
                  this.frontRepo.Bstructs.delete(bstruct.ID)
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

  // DstructPull performs a GET on Dstruct of the stack and redeem association pointers 
  DstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.dstructService.getDstructs(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            dstructs,
          ]) => {
            // init the array
            this.frontRepo.Dstructs_array = dstructs

            // clear the map that counts Dstruct in the GET
            this.frontRepo.Dstructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            dstructs.forEach(
              dstruct => {
                this.frontRepo.Dstructs.set(dstruct.ID, dstruct)
                this.frontRepo.Dstructs_batch.set(dstruct.ID, dstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear dstructs that are absent from the GET
            this.frontRepo.Dstructs.forEach(
              dstruct => {
                if (this.frontRepo.Dstructs_batch.get(dstruct.ID) == undefined) {
                  this.frontRepo.Dstructs.delete(dstruct.ID)
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
export function getAstructUniqueID(id: number): number {
  return 31 * id
}
export function getAstructBstruct2UseUniqueID(id: number): number {
  return 37 * id
}
export function getAstructBstructUseUniqueID(id: number): number {
  return 41 * id
}
export function getBstructUniqueID(id: number): number {
  return 43 * id
}
export function getDstructUniqueID(id: number): number {
  return 47 * id
}
