import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { AclassDB } from './aclass-db'
import { AclassService } from './aclass.service'

import { BclassDB } from './bclass-db'
import { BclassService } from './bclass.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Aclasss_array = new Array<AclassDB>(); // array of repo instances
  Aclasss = new Map<number, AclassDB>(); // map of repo instances
  Aclasss_batch = new Map<number, AclassDB>(); // same but only in last GET (for finding repo instances to delete)
  Bclasss_array = new Array<BclassDB>(); // array of repo instances
  Bclasss = new Map<number, BclassDB>(); // map of repo instances
  Bclasss_batch = new Map<number, BclassDB>(); // same but only in last GET (for finding repo instances to delete)
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

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
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
    private bclassService: BclassService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<AclassDB[]>,
    Observable<BclassDB[]>,
  ] = [ // insertion point sub template 
      this.aclassService.getAclasss(),
      this.bclassService.getBclasss(),
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
            bclasss_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var aclasss: AclassDB[]
            aclasss = aclasss_
            var bclasss: BclassDB[]
            bclasss = bclasss_

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
            // init the array
            FrontRepoSingloton.Bclasss_array = bclasss

            // clear the map that counts Bclass in the GET
            FrontRepoSingloton.Bclasss_batch.clear()
            
            bclasss.forEach(
              bclass => {
                FrontRepoSingloton.Bclasss.set(bclass.ID, bclass)
                FrontRepoSingloton.Bclasss_batch.set(bclass.ID, bclass)
              }
            )
            
            // clear bclasss that are absent from the batch
            FrontRepoSingloton.Bclasss.forEach(
              bclass => {
                if (FrontRepoSingloton.Bclasss_batch.get(bclass.ID) == undefined) {
                  FrontRepoSingloton.Bclasss.delete(bclass.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            aclasss.forEach(
              aclass => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Associationtob redeeming
                {
                  let _bclass = FrontRepoSingloton.Bclasss.get(aclass.AssociationtobID.Int64)
                  if (_bclass) {
                    aclass.Associationtob = _bclass
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bclass = FrontRepoSingloton.Bclasss.get(aclass.Anotherassociationtob_2ID.Int64)
                  if (_bclass) {
                    aclass.Anotherassociationtob_2 = _bclass
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Aclass.Anarrayofa redeeming
                {
                  let _aclass = FrontRepoSingloton.Aclasss.get(aclass.Aclass_AnarrayofaDBID.Int64)
                  if (_aclass) {
                    if (_aclass.Anarrayofa == undefined) {
                      _aclass.Anarrayofa = new Array<AclassDB>()
                    }
                    _aclass.Anarrayofa.push(aclass)
                    if (aclass.Aclass_Anarrayofa_reverse == undefined) {
                      aclass.Aclass_Anarrayofa_reverse = _aclass
                    }
                  }
                }
              }
            )
            bclasss.forEach(
              bclass => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Aclass.Anarrayofb redeeming
                {
                  let _aclass = FrontRepoSingloton.Aclasss.get(bclass.Aclass_AnarrayofbDBID.Int64)
                  if (_aclass) {
                    if (_aclass.Anarrayofb == undefined) {
                      _aclass.Anarrayofb = new Array<AclassDB>()
                    }
                    _aclass.Anarrayofb.push(bclass)
                    if (bclass.Aclass_Anarrayofb_reverse == undefined) {
                      bclass.Aclass_Anarrayofb_reverse = _aclass
                    }
                  }
                }
                // insertion point for slice of pointer field Aclass.Anotherarrayofb redeeming
                {
                  let _aclass = FrontRepoSingloton.Aclasss.get(bclass.Aclass_AnotherarrayofbDBID.Int64)
                  if (_aclass) {
                    if (_aclass.Anotherarrayofb == undefined) {
                      _aclass.Anotherarrayofb = new Array<AclassDB>()
                    }
                    _aclass.Anotherarrayofb.push(bclass)
                    if (bclass.Aclass_Anotherarrayofb_reverse == undefined) {
                      bclass.Aclass_Anotherarrayofb_reverse = _aclass
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
                // insertion point for pointer field Associationtob redeeming
                {
                  let _bclass = FrontRepoSingloton.Bclasss.get(aclass.AssociationtobID.Int64)
                  if (_bclass) {
                    aclass.Associationtob = _bclass
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bclass = FrontRepoSingloton.Bclasss.get(aclass.Anotherassociationtob_2ID.Int64)
                  if (_bclass) {
                    aclass.Anotherassociationtob_2 = _bclass
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Aclass.Anarrayofa redeeming
                {
                  let _aclass = FrontRepoSingloton.Aclasss.get(aclass.Aclass_AnarrayofaDBID.Int64)
                  if (_aclass) {
                    if (_aclass.Anarrayofa == undefined) {
                      _aclass.Anarrayofa = new Array<AclassDB>()
                    }
                    _aclass.Anarrayofa.push(aclass)
                    if (aclass.Aclass_Anarrayofa_reverse == undefined) {
                      aclass.Aclass_Anarrayofa_reverse = _aclass
                    }
                  }
                }
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

  // BclassPull performs a GET on Bclass of the stack and redeem association pointers 
  BclassPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.bclassService.getBclasss()
        ]).subscribe(
          ([ // insertion point sub template 
            bclasss,
          ]) => {
            // init the array
            FrontRepoSingloton.Bclasss_array = bclasss

            // clear the map that counts Bclass in the GET
            FrontRepoSingloton.Bclasss_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bclasss.forEach(
              bclass => {
                FrontRepoSingloton.Bclasss.set(bclass.ID, bclass)
                FrontRepoSingloton.Bclasss_batch.set(bclass.ID, bclass)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Aclass.Anarrayofb redeeming
                {
                  let _aclass = FrontRepoSingloton.Aclasss.get(bclass.Aclass_AnarrayofbDBID.Int64)
                  if (_aclass) {
                    if (_aclass.Anarrayofb == undefined) {
                      _aclass.Anarrayofb = new Array<AclassDB>()
                    }
                    _aclass.Anarrayofb.push(bclass)
                    if (bclass.Aclass_Anarrayofb_reverse == undefined) {
                      bclass.Aclass_Anarrayofb_reverse = _aclass
                    }
                  }
                }
                // insertion point for slice of pointer field Aclass.Anotherarrayofb redeeming
                {
                  let _aclass = FrontRepoSingloton.Aclasss.get(bclass.Aclass_AnotherarrayofbDBID.Int64)
                  if (_aclass) {
                    if (_aclass.Anotherarrayofb == undefined) {
                      _aclass.Anotherarrayofb = new Array<AclassDB>()
                    }
                    _aclass.Anotherarrayofb.push(bclass)
                    if (bclass.Aclass_Anotherarrayofb_reverse == undefined) {
                      bclass.Aclass_Anotherarrayofb_reverse = _aclass
                    }
                  }
                }
              }
            )

            // clear bclasss that are absent from the GET
            FrontRepoSingloton.Bclasss.forEach(
              bclass => {
                if (FrontRepoSingloton.Bclasss_batch.get(bclass.ID) == undefined) {
                  FrontRepoSingloton.Bclasss.delete(bclass.ID)
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
export function getBclassUniqueID(id: number): number {
  return 37 * id
}
