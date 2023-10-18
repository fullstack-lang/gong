import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { ADB } from './a-db'
import { AService } from './a.service'

import { BDB } from './b-db'
import { BService } from './b.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  As_array = new Array<ADB>() // array of repo instances
  As = new Map<number, ADB>() // map of repo instances
  As_batch = new Map<number, ADB>() // same but only in last GET (for finding repo instances to delete)

  Bs_array = new Array<BDB>() // array of repo instances
  Bs = new Map<number, BDB>() // map of repo instances
  Bs_batch = new Map<number, BDB>() // same but only in last GET (for finding repo instances to delete)


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'A':
        return this.As_array as unknown as Array<Type>
      case 'B':
        return this.Bs_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'A':
        return this.As_array as unknown as Map<number, Type>
      case 'B':
        return this.Bs_array as unknown as Map<number, Type>
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
    private aService: AService,
    private bService: BService,
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
    Observable<ADB[]>,
    Observable<BDB[]>,
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
      this.aService.getAs(this.GONG__StackPath),
      this.bService.getBs(this.GONG__StackPath),
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
      this.aService.getAs(this.GONG__StackPath),
      this.bService.getBs(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            as_,
            bs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var as: ADB[]
            as = as_ as ADB[]
            var bs: BDB[]
            bs = bs_ as BDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.As_array = as

            // clear the map that counts A in the GET
            this.frontRepo.As_batch.clear()

            as.forEach(
              a => {
                this.frontRepo.As.set(a.ID, a)
                this.frontRepo.As_batch.set(a.ID, a)
              }
            )

            // clear as that are absent from the batch
            this.frontRepo.As.forEach(
              a => {
                if (this.frontRepo.As_batch.get(a.ID) == undefined) {
                  this.frontRepo.As.delete(a.ID)
                }
              }
            )

            // sort As_array array
            this.frontRepo.As_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Bs_array = bs

            // clear the map that counts B in the GET
            this.frontRepo.Bs_batch.clear()

            bs.forEach(
              b => {
                this.frontRepo.Bs.set(b.ID, b)
                this.frontRepo.Bs_batch.set(b.ID, b)
              }
            )

            // clear bs that are absent from the batch
            this.frontRepo.Bs.forEach(
              b => {
                if (this.frontRepo.Bs_batch.get(b.ID) == undefined) {
                  this.frontRepo.Bs.delete(b.ID)
                }
              }
            )

            // sort Bs_array array
            this.frontRepo.Bs_array.sort((t1, t2) => {
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
            as.forEach(
              a => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field B redeeming
                {
                  let _b = this.frontRepo.Bs.get(a.APointersEncoding.BID.Int64)
                  if (_b) {
                    a.B = _b
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            bs.forEach(
              b => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field A.Bs redeeming
                // to be removed
                {
                  let _id = b.BPointersEncoding.A_BsDBID.Int64
                  let _a = this.frontRepo.As.get(_id)
                  if (_a) {
                    if (_a.Bs == undefined) {
                      _a.Bs = new Array<BDB>()
                    }
                    _a.Bs.push(b)
                    if (b.BPointersEncoding.A_Bs_reverse == undefined) {
                      b.BPointersEncoding.A_Bs_reverse = _a
                    }
                  }
                }
              }
            )

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem 
            as.forEach(
              a => {
                // insertion point for sorting
                // to be removed
                a.Bs?.sort((t1, t2) => {
                  if (t1.BPointersEncoding.A_BsDBID_Index.Int64 > t2.BPointersEncoding.A_BsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.BPointersEncoding.A_BsDBID_Index.Int64 < t2.BPointersEncoding.A_BsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            bs.forEach(
              b => {
                // insertion point for sorting
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

  // APull performs a GET on A of the stack and redeem association pointers 
  APull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.aService.getAs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            as,
          ]) => {
            // init the array
            this.frontRepo.As_array = as

            // clear the map that counts A in the GET
            this.frontRepo.As_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            as.forEach(
              a => {
                this.frontRepo.As.set(a.ID, a)
                this.frontRepo.As_batch.set(a.ID, a)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field B redeeming
                {
                  let _b = this.frontRepo.Bs.get(a.APointersEncoding.BID.Int64)
                  if (_b) {
                    a.B = _b
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear as that are absent from the GET
            this.frontRepo.As.forEach(
              a => {
                if (this.frontRepo.As_batch.get(a.ID) == undefined) {
                  this.frontRepo.As.delete(a.ID)
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

  // BPull performs a GET on B of the stack and redeem association pointers 
  BPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.bService.getBs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            bs,
          ]) => {
            // init the array
            this.frontRepo.Bs_array = bs

            // clear the map that counts B in the GET
            this.frontRepo.Bs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bs.forEach(
              b => {
                this.frontRepo.Bs.set(b.ID, b)
                this.frontRepo.Bs_batch.set(b.ID, b)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field A.Bs redeeming
                // to be removed
                {
                  let _id = b.BPointersEncoding.A_BsDBID.Int64
                  let _a = this.frontRepo.As.get(_id)
                  if (_a) {
                    if (_a.Bs == undefined) {
                      _a.Bs = new Array<BDB>()
                    }
                    _a.Bs.push(b)
                    if (b.BPointersEncoding.A_Bs_reverse == undefined) {
                      b.BPointersEncoding.A_Bs_reverse = _a
                    }
                  }
                }
              }
            )

            // clear bs that are absent from the GET
            this.frontRepo.Bs.forEach(
              b => {
                if (this.frontRepo.Bs_batch.get(b.ID) == undefined) {
                  this.frontRepo.Bs.delete(b.ID)
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
export function getAUniqueID(id: number): number {
  return 31 * id
}
export function getBUniqueID(id: number): number {
  return 37 * id
}
