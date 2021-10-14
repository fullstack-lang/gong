import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { AstructDB } from './astruct-db'
import { AstructService } from './astruct.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Astructs_array = new Array<AstructDB>(); // array of repo instances
  Astructs = new Map<number, AstructDB>(); // map of repo instances
  Astructs_batch = new Map<number, AstructDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private astructService: AstructService,
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
    Observable<AstructDB[]>,
  ] = [ // insertion point sub template 
      this.astructService.getAstructs(),
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
            astructs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var astructs: AstructDB[]
            astructs = astructs_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Astructs_array = astructs

            // clear the map that counts Astruct in the GET
            FrontRepoSingloton.Astructs_batch.clear()

            astructs.forEach(
              astruct => {
                FrontRepoSingloton.Astructs.set(astruct.ID, astruct)
                FrontRepoSingloton.Astructs_batch.set(astruct.ID, astruct)
              }
            )

            // clear astructs that are absent from the batch
            FrontRepoSingloton.Astructs.forEach(
              astruct => {
                if (FrontRepoSingloton.Astructs_batch.get(astruct.ID) == undefined) {
                  FrontRepoSingloton.Astructs.delete(astruct.ID)
                }
              }
            )

            // sort Astructs_array array
            FrontRepoSingloton.Astructs_array.sort((t1, t2) => {
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
            astructs.forEach(
              astruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofa redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(astruct.Astruct_AnarrayofaDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anarrayofa == undefined) {
                      _astruct.Anarrayofa = new Array<AstructDB>()
                    }
                    _astruct.Anarrayofa.push(astruct)
                    if (astruct.Astruct_Anarrayofa_reverse == undefined) {
                      astruct.Astruct_Anarrayofa_reverse = _astruct
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

  // AstructPull performs a GET on Astruct of the stack and redeem association pointers 
  AstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructService.getAstructs()
        ]).subscribe(
          ([ // insertion point sub template 
            astructs,
          ]) => {
            // init the array
            FrontRepoSingloton.Astructs_array = astructs

            // clear the map that counts Astruct in the GET
            FrontRepoSingloton.Astructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            astructs.forEach(
              astruct => {
                FrontRepoSingloton.Astructs.set(astruct.ID, astruct)
                FrontRepoSingloton.Astructs_batch.set(astruct.ID, astruct)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofa redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(astruct.Astruct_AnarrayofaDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anarrayofa == undefined) {
                      _astruct.Anarrayofa = new Array<AstructDB>()
                    }
                    _astruct.Anarrayofa.push(astruct)
                    if (astruct.Astruct_Anarrayofa_reverse == undefined) {
                      astruct.Astruct_Anarrayofa_reverse = _astruct
                    }
                  }
                }
              }
            )

            // clear astructs that are absent from the GET
            FrontRepoSingloton.Astructs.forEach(
              astruct => {
                if (FrontRepoSingloton.Astructs_batch.get(astruct.ID) == undefined) {
                  FrontRepoSingloton.Astructs.delete(astruct.ID)
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
export function getAstructUniqueID(id: number): number {
  return 31 * id
}
