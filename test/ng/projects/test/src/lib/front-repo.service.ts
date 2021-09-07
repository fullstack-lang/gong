import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { AstructDB } from './astruct-db'
import { AstructService } from './astruct.service'

import { AstructBstruct2UseDB } from './astructbstruct2use-db'
import { AstructBstruct2UseService } from './astructbstruct2use.service'

import { AstructBstructUseDB } from './astructbstructuse-db'
import { AstructBstructUseService } from './astructbstructuse.service'

import { BstructDB } from './bstruct-db'
import { BstructService } from './bstruct.service'

import { DstructDB } from './dstruct-db'
import { DstructService } from './dstruct.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Astructs_array = new Array<AstructDB>(); // array of repo instances
  Astructs = new Map<number, AstructDB>(); // map of repo instances
  Astructs_batch = new Map<number, AstructDB>(); // same but only in last GET (for finding repo instances to delete)
  AstructBstruct2Uses_array = new Array<AstructBstruct2UseDB>(); // array of repo instances
  AstructBstruct2Uses = new Map<number, AstructBstruct2UseDB>(); // map of repo instances
  AstructBstruct2Uses_batch = new Map<number, AstructBstruct2UseDB>(); // same but only in last GET (for finding repo instances to delete)
  AstructBstructUses_array = new Array<AstructBstructUseDB>(); // array of repo instances
  AstructBstructUses = new Map<number, AstructBstructUseDB>(); // map of repo instances
  AstructBstructUses_batch = new Map<number, AstructBstructUseDB>(); // same but only in last GET (for finding repo instances to delete)
  Bstructs_array = new Array<BstructDB>(); // array of repo instances
  Bstructs = new Map<number, BstructDB>(); // map of repo instances
  Bstructs_batch = new Map<number, BstructDB>(); // same but only in last GET (for finding repo instances to delete)
  Dstructs_array = new Array<DstructDB>(); // array of repo instances
  Dstructs = new Map<number, DstructDB>(); // map of repo instances
  Dstructs_batch = new Map<number, DstructDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private astructService: AstructService,
    private astructbstruct2useService: AstructBstruct2UseService,
    private astructbstructuseService: AstructBstructUseService,
    private bstructService: BstructService,
    private dstructService: DstructService,
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
    Observable<AstructDB[]>,
    Observable<AstructBstruct2UseDB[]>,
    Observable<AstructBstructUseDB[]>,
    Observable<BstructDB[]>,
    Observable<DstructDB[]>,
  ] = [ // insertion point sub template 
      this.astructService.getAstructs(),
      this.astructbstruct2useService.getAstructBstruct2Uses(),
      this.astructbstructuseService.getAstructBstructUses(),
      this.bstructService.getBstructs(),
      this.dstructService.getDstructs(),
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
            astructbstruct2uses_,
            astructbstructuses_,
            bstructs_,
            dstructs_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var astructs: AstructDB[]
            astructs = astructs_
            var astructbstruct2uses: AstructBstruct2UseDB[]
            astructbstruct2uses = astructbstruct2uses_
            var astructbstructuses: AstructBstructUseDB[]
            astructbstructuses = astructbstructuses_
            var bstructs: BstructDB[]
            bstructs = bstructs_
            var dstructs: DstructDB[]
            dstructs = dstructs_

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

            // init the array
            FrontRepoSingloton.AstructBstruct2Uses_array = astructbstruct2uses

            // clear the map that counts AstructBstruct2Use in the GET
            FrontRepoSingloton.AstructBstruct2Uses_batch.clear()

            astructbstruct2uses.forEach(
              astructbstruct2use => {
                FrontRepoSingloton.AstructBstruct2Uses.set(astructbstruct2use.ID, astructbstruct2use)
                FrontRepoSingloton.AstructBstruct2Uses_batch.set(astructbstruct2use.ID, astructbstruct2use)
              }
            )

            // clear astructbstruct2uses that are absent from the batch
            FrontRepoSingloton.AstructBstruct2Uses.forEach(
              astructbstruct2use => {
                if (FrontRepoSingloton.AstructBstruct2Uses_batch.get(astructbstruct2use.ID) == undefined) {
                  FrontRepoSingloton.AstructBstruct2Uses.delete(astructbstruct2use.ID)
                }
              }
            )

            // sort AstructBstruct2Uses_array array
            FrontRepoSingloton.AstructBstruct2Uses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.AstructBstructUses_array = astructbstructuses

            // clear the map that counts AstructBstructUse in the GET
            FrontRepoSingloton.AstructBstructUses_batch.clear()

            astructbstructuses.forEach(
              astructbstructuse => {
                FrontRepoSingloton.AstructBstructUses.set(astructbstructuse.ID, astructbstructuse)
                FrontRepoSingloton.AstructBstructUses_batch.set(astructbstructuse.ID, astructbstructuse)
              }
            )

            // clear astructbstructuses that are absent from the batch
            FrontRepoSingloton.AstructBstructUses.forEach(
              astructbstructuse => {
                if (FrontRepoSingloton.AstructBstructUses_batch.get(astructbstructuse.ID) == undefined) {
                  FrontRepoSingloton.AstructBstructUses.delete(astructbstructuse.ID)
                }
              }
            )

            // sort AstructBstructUses_array array
            FrontRepoSingloton.AstructBstructUses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Bstructs_array = bstructs

            // clear the map that counts Bstruct in the GET
            FrontRepoSingloton.Bstructs_batch.clear()

            bstructs.forEach(
              bstruct => {
                FrontRepoSingloton.Bstructs.set(bstruct.ID, bstruct)
                FrontRepoSingloton.Bstructs_batch.set(bstruct.ID, bstruct)
              }
            )

            // clear bstructs that are absent from the batch
            FrontRepoSingloton.Bstructs.forEach(
              bstruct => {
                if (FrontRepoSingloton.Bstructs_batch.get(bstruct.ID) == undefined) {
                  FrontRepoSingloton.Bstructs.delete(bstruct.ID)
                }
              }
            )

            // sort Bstructs_array array
            FrontRepoSingloton.Bstructs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Dstructs_array = dstructs

            // clear the map that counts Dstruct in the GET
            FrontRepoSingloton.Dstructs_batch.clear()

            dstructs.forEach(
              dstruct => {
                FrontRepoSingloton.Dstructs.set(dstruct.ID, dstruct)
                FrontRepoSingloton.Dstructs_batch.set(dstruct.ID, dstruct)
              }
            )

            // clear dstructs that are absent from the batch
            FrontRepoSingloton.Dstructs.forEach(
              dstruct => {
                if (FrontRepoSingloton.Dstructs_batch.get(dstruct.ID) == undefined) {
                  FrontRepoSingloton.Dstructs.delete(dstruct.ID)
                }
              }
            )

            // sort Dstructs_array array
            FrontRepoSingloton.Dstructs_array.sort((t1, t2) => {
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
                // insertion point for pointer field Associationtob redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astruct.AssociationtobID.Int64)
                  if (_bstruct) {
                    astruct.Associationtob = _bstruct
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astruct.Anotherassociationtob_2ID.Int64)
                  if (_bstruct) {
                    astruct.Anotherassociationtob_2 = _bstruct
                  }
                }

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
            astructbstruct2uses.forEach(
              astructbstruct2use => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Bstrcut2 redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astructbstruct2use.Bstrcut2ID.Int64)
                  if (_bstruct) {
                    astructbstruct2use.Bstrcut2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb2Use redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(astructbstruct2use.Astruct_Anarrayofb2UseDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anarrayofb2Use == undefined) {
                      _astruct.Anarrayofb2Use = new Array<AstructBstruct2UseDB>()
                    }
                    _astruct.Anarrayofb2Use.push(astructbstruct2use)
                    if (astructbstruct2use.Astruct_Anarrayofb2Use_reverse == undefined) {
                      astructbstruct2use.Astruct_Anarrayofb2Use_reverse = _astruct
                    }
                  }
                }
              }
            )
            astructbstructuses.forEach(
              astructbstructuse => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astructbstructuse.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astructbstructuse.Bstruct2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.AnarrayofbUse redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(astructbstructuse.Astruct_AnarrayofbUseDBID.Int64)
                  if (_astruct) {
                    if (_astruct.AnarrayofbUse == undefined) {
                      _astruct.AnarrayofbUse = new Array<AstructBstructUseDB>()
                    }
                    _astruct.AnarrayofbUse.push(astructbstructuse)
                    if (astructbstructuse.Astruct_AnarrayofbUse_reverse == undefined) {
                      astructbstructuse.Astruct_AnarrayofbUse_reverse = _astruct
                    }
                  }
                }
              }
            )
            bstructs.forEach(
              bstruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(bstruct.Astruct_AnarrayofbDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anarrayofb == undefined) {
                      _astruct.Anarrayofb = new Array<BstructDB>()
                    }
                    _astruct.Anarrayofb.push(bstruct)
                    if (bstruct.Astruct_Anarrayofb_reverse == undefined) {
                      bstruct.Astruct_Anarrayofb_reverse = _astruct
                    }
                  }
                }
                // insertion point for slice of pointer field Astruct.Anotherarrayofb redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(bstruct.Astruct_AnotherarrayofbDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anotherarrayofb == undefined) {
                      _astruct.Anotherarrayofb = new Array<BstructDB>()
                    }
                    _astruct.Anotherarrayofb.push(bstruct)
                    if (bstruct.Astruct_Anotherarrayofb_reverse == undefined) {
                      bstruct.Astruct_Anotherarrayofb_reverse = _astruct
                    }
                  }
                }
              }
            )
            dstructs.forEach(
              dstruct => {
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
                // insertion point for pointer field Associationtob redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astruct.AssociationtobID.Int64)
                  if (_bstruct) {
                    astruct.Associationtob = _bstruct
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astruct.Anotherassociationtob_2ID.Int64)
                  if (_bstruct) {
                    astruct.Anotherassociationtob_2 = _bstruct
                  }
                }

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

  // AstructBstruct2UsePull performs a GET on AstructBstruct2Use of the stack and redeem association pointers 
  AstructBstruct2UsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructbstruct2useService.getAstructBstruct2Uses()
        ]).subscribe(
          ([ // insertion point sub template 
            astructbstruct2uses,
          ]) => {
            // init the array
            FrontRepoSingloton.AstructBstruct2Uses_array = astructbstruct2uses

            // clear the map that counts AstructBstruct2Use in the GET
            FrontRepoSingloton.AstructBstruct2Uses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            astructbstruct2uses.forEach(
              astructbstruct2use => {
                FrontRepoSingloton.AstructBstruct2Uses.set(astructbstruct2use.ID, astructbstruct2use)
                FrontRepoSingloton.AstructBstruct2Uses_batch.set(astructbstruct2use.ID, astructbstruct2use)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Bstrcut2 redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astructbstruct2use.Bstrcut2ID.Int64)
                  if (_bstruct) {
                    astructbstruct2use.Bstrcut2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb2Use redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(astructbstruct2use.Astruct_Anarrayofb2UseDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anarrayofb2Use == undefined) {
                      _astruct.Anarrayofb2Use = new Array<AstructBstruct2UseDB>()
                    }
                    _astruct.Anarrayofb2Use.push(astructbstruct2use)
                    if (astructbstruct2use.Astruct_Anarrayofb2Use_reverse == undefined) {
                      astructbstruct2use.Astruct_Anarrayofb2Use_reverse = _astruct
                    }
                  }
                }
              }
            )

            // clear astructbstruct2uses that are absent from the GET
            FrontRepoSingloton.AstructBstruct2Uses.forEach(
              astructbstruct2use => {
                if (FrontRepoSingloton.AstructBstruct2Uses_batch.get(astructbstruct2use.ID) == undefined) {
                  FrontRepoSingloton.AstructBstruct2Uses.delete(astructbstruct2use.ID)
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

  // AstructBstructUsePull performs a GET on AstructBstructUse of the stack and redeem association pointers 
  AstructBstructUsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructbstructuseService.getAstructBstructUses()
        ]).subscribe(
          ([ // insertion point sub template 
            astructbstructuses,
          ]) => {
            // init the array
            FrontRepoSingloton.AstructBstructUses_array = astructbstructuses

            // clear the map that counts AstructBstructUse in the GET
            FrontRepoSingloton.AstructBstructUses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            astructbstructuses.forEach(
              astructbstructuse => {
                FrontRepoSingloton.AstructBstructUses.set(astructbstructuse.ID, astructbstructuse)
                FrontRepoSingloton.AstructBstructUses_batch.set(astructbstructuse.ID, astructbstructuse)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = FrontRepoSingloton.Bstructs.get(astructbstructuse.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astructbstructuse.Bstruct2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.AnarrayofbUse redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(astructbstructuse.Astruct_AnarrayofbUseDBID.Int64)
                  if (_astruct) {
                    if (_astruct.AnarrayofbUse == undefined) {
                      _astruct.AnarrayofbUse = new Array<AstructBstructUseDB>()
                    }
                    _astruct.AnarrayofbUse.push(astructbstructuse)
                    if (astructbstructuse.Astruct_AnarrayofbUse_reverse == undefined) {
                      astructbstructuse.Astruct_AnarrayofbUse_reverse = _astruct
                    }
                  }
                }
              }
            )

            // clear astructbstructuses that are absent from the GET
            FrontRepoSingloton.AstructBstructUses.forEach(
              astructbstructuse => {
                if (FrontRepoSingloton.AstructBstructUses_batch.get(astructbstructuse.ID) == undefined) {
                  FrontRepoSingloton.AstructBstructUses.delete(astructbstructuse.ID)
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

  // BstructPull performs a GET on Bstruct of the stack and redeem association pointers 
  BstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.bstructService.getBstructs()
        ]).subscribe(
          ([ // insertion point sub template 
            bstructs,
          ]) => {
            // init the array
            FrontRepoSingloton.Bstructs_array = bstructs

            // clear the map that counts Bstruct in the GET
            FrontRepoSingloton.Bstructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bstructs.forEach(
              bstruct => {
                FrontRepoSingloton.Bstructs.set(bstruct.ID, bstruct)
                FrontRepoSingloton.Bstructs_batch.set(bstruct.ID, bstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(bstruct.Astruct_AnarrayofbDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anarrayofb == undefined) {
                      _astruct.Anarrayofb = new Array<BstructDB>()
                    }
                    _astruct.Anarrayofb.push(bstruct)
                    if (bstruct.Astruct_Anarrayofb_reverse == undefined) {
                      bstruct.Astruct_Anarrayofb_reverse = _astruct
                    }
                  }
                }
                // insertion point for slice of pointer field Astruct.Anotherarrayofb redeeming
                {
                  let _astruct = FrontRepoSingloton.Astructs.get(bstruct.Astruct_AnotherarrayofbDBID.Int64)
                  if (_astruct) {
                    if (_astruct.Anotherarrayofb == undefined) {
                      _astruct.Anotherarrayofb = new Array<BstructDB>()
                    }
                    _astruct.Anotherarrayofb.push(bstruct)
                    if (bstruct.Astruct_Anotherarrayofb_reverse == undefined) {
                      bstruct.Astruct_Anotherarrayofb_reverse = _astruct
                    }
                  }
                }
              }
            )

            // clear bstructs that are absent from the GET
            FrontRepoSingloton.Bstructs.forEach(
              bstruct => {
                if (FrontRepoSingloton.Bstructs_batch.get(bstruct.ID) == undefined) {
                  FrontRepoSingloton.Bstructs.delete(bstruct.ID)
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

  // DstructPull performs a GET on Dstruct of the stack and redeem association pointers 
  DstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.dstructService.getDstructs()
        ]).subscribe(
          ([ // insertion point sub template 
            dstructs,
          ]) => {
            // init the array
            FrontRepoSingloton.Dstructs_array = dstructs

            // clear the map that counts Dstruct in the GET
            FrontRepoSingloton.Dstructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            dstructs.forEach(
              dstruct => {
                FrontRepoSingloton.Dstructs.set(dstruct.ID, dstruct)
                FrontRepoSingloton.Dstructs_batch.set(dstruct.ID, dstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear dstructs that are absent from the GET
            FrontRepoSingloton.Dstructs.forEach(
              dstruct => {
                if (FrontRepoSingloton.Dstructs_batch.get(dstruct.ID) == undefined) {
                  FrontRepoSingloton.Dstructs.delete(dstruct.ID)
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
