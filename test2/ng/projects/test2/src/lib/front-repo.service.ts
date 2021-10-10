import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { AclassDB } from './aclass-db'
import { AclassService } from './aclass.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Aclasss_array = new Array<AclassDB>(); // array of repo instances
  Aclasss = new Map<number, AclassDB>(); // map of repo instances
  Aclasss_batch = new Map<number, AclassDB>(); // same but only in last GET (for finding repo instances to delete)
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

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number; // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean; // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode;

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string;  // The "Aclass"
  SourceField: string; // the "AnarrayofbUse"
  IntermediateStruct: string; // the "AclassBclassUse" 
  IntermediateStructField: string; // the "Bclass" as field
  NextAssociationStruct: string; // the "Bclass"
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
    private aclassService: AclassService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service"]
    service["post" + structName](instanceToBePosted).subscribe(
      instance => {
        service[structName + "ServiceChanged"].next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service"]
    service["delete" + structName](instanceToBeDeleted).subscribe(
      instance => {
        service[structName + "ServiceChanged"].next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<AclassDB[]>,
  ] = [ // insertion point sub template 
      this.aclassService.getAclasss(),
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
            aclasss_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var aclasss: AclassDB[]
            aclasss = aclasss_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Aclasss_array = aclasss

            // clear the map that counts Aclass in the GET
            FrontRepoSingloton.Aclasss_batch.clear()

            aclasss.forEach(
              aclass => {
                FrontRepoSingloton.Aclasss.set(aclass.ID, aclass)
                FrontRepoSingloton.Aclasss_batch.set(aclass.ID, aclass)
              }
            )

            // clear aclasss that are absent from the batch
            FrontRepoSingloton.Aclasss.forEach(
              aclass => {
                if (FrontRepoSingloton.Aclasss_batch.get(aclass.ID) == undefined) {
                  FrontRepoSingloton.Aclasss.delete(aclass.ID)
                }
              }
            )

            // sort Aclasss_array array
            FrontRepoSingloton.Aclasss_array.sort((t1, t2) => {
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
            aclasss.forEach(
              aclass => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
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

  // AclassPull performs a GET on Aclass of the stack and redeem association pointers 
  AclassPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.aclassService.getAclasss()
        ]).subscribe(
          ([ // insertion point sub template 
            aclasss,
          ]) => {
            // init the array
            FrontRepoSingloton.Aclasss_array = aclasss

            // clear the map that counts Aclass in the GET
            FrontRepoSingloton.Aclasss_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            aclasss.forEach(
              aclass => {
                FrontRepoSingloton.Aclasss.set(aclass.ID, aclass)
                FrontRepoSingloton.Aclasss_batch.set(aclass.ID, aclass)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear aclasss that are absent from the GET
            FrontRepoSingloton.Aclasss.forEach(
              aclass => {
                if (FrontRepoSingloton.Aclasss_batch.get(aclass.ID) == undefined) {
                  FrontRepoSingloton.Aclasss.delete(aclass.ID)
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
export function getAclassUniqueID(id: number): number {
  return 31 * id
}
