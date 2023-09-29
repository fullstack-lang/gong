import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

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

import { FstructDB } from './fstruct-db'
import { FstructService } from './fstruct.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Astructs_array = new Array<AstructDB>() // array of repo instances
  Astructs = new Map<number, AstructDB>() // map of repo instances
  Astructs_batch = new Map<number, AstructDB>() // same but only in last GET (for finding repo instances to delete)

  AstructBstruct2Uses_array = new Array<AstructBstruct2UseDB>() // array of repo instances
  AstructBstruct2Uses = new Map<number, AstructBstruct2UseDB>() // map of repo instances
  AstructBstruct2Uses_batch = new Map<number, AstructBstruct2UseDB>() // same but only in last GET (for finding repo instances to delete)

  AstructBstructUses_array = new Array<AstructBstructUseDB>() // array of repo instances
  AstructBstructUses = new Map<number, AstructBstructUseDB>() // map of repo instances
  AstructBstructUses_batch = new Map<number, AstructBstructUseDB>() // same but only in last GET (for finding repo instances to delete)

  Bstructs_array = new Array<BstructDB>() // array of repo instances
  Bstructs = new Map<number, BstructDB>() // map of repo instances
  Bstructs_batch = new Map<number, BstructDB>() // same but only in last GET (for finding repo instances to delete)

  Dstructs_array = new Array<DstructDB>() // array of repo instances
  Dstructs = new Map<number, DstructDB>() // map of repo instances
  Dstructs_batch = new Map<number, DstructDB>() // same but only in last GET (for finding repo instances to delete)

  Fstructs_array = new Array<FstructDB>() // array of repo instances
  Fstructs = new Map<number, FstructDB>() // map of repo instances
  Fstructs_batch = new Map<number, FstructDB>() // same but only in last GET (for finding repo instances to delete)


  getArray<Type>(): Array<Type> {
    const token = this.getToken<Type>();

    switch (token) {
    // insertion point
    case 'AstructDB':
      return this.Astructs_array as unknown as Array<Type>
    case 'AstructBstruct2UseDB':
      return this.AstructBstruct2Uses_array as unknown as Array<Type>
    case 'AstructBstructUseDB':
      return this.AstructBstructUses_array as unknown as Array<Type>
    case 'BstructDB':
      return this.Bstructs_array as unknown as Array<Type>
    case 'DstructDB':
      return this.Dstructs_array as unknown as Array<Type>
    case 'FstructDB':
      return this.Fstructs_array as unknown as Array<Type>
    default:
      throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(): Map<number, Type> {
    const token = this.getToken<Type>();

    switch (token) {
    // insertion point
    case 'AstructDB':
      return this.Astructs_array as unknown as Map<number, Type>
    case 'AstructBstruct2UseDB':
      return this.AstructBstruct2Uses_array as unknown as Map<number, Type>
    case 'AstructBstructUseDB':
      return this.AstructBstructUses_array as unknown as Map<number, Type>
    case 'BstructDB':
      return this.Bstructs_array as unknown as Map<number, Type>
    case 'DstructDB':
      return this.Dstructs_array as unknown as Map<number, Type>
    case 'FstructDB':
      return this.Fstructs_array as unknown as Map<number, Type>
    default:
      throw new Error("Type not recognized");
    }
  }

  // getToken allows for a get function that is robust to refactoring of the named struct name
  private getToken<Type>(): string {
    // insertion point
  if (({} as Type) instanceof AstructDB) return 'AstructDB'
  if (({} as Type) instanceof AstructBstruct2UseDB) return 'AstructBstruct2UseDB'
  if (({} as Type) instanceof AstructBstructUseDB) return 'AstructBstructUseDB'
  if (({} as Type) instanceof BstructDB) return 'BstructDB'
  if (({} as Type) instanceof DstructDB) return 'DstructDB'
  if (({} as Type) instanceof FstructDB) return 'FstructDB'
    return '';
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
    private fstructService: FstructService,
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
    Observable<FstructDB[]>,
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
      this.astructService.getAstructs(this.GONG__StackPath),
      this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath),
      this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath),
      this.bstructService.getBstructs(this.GONG__StackPath),
      this.dstructService.getDstructs(this.GONG__StackPath),
      this.fstructService.getFstructs(this.GONG__StackPath),
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
      this.astructService.getAstructs(this.GONG__StackPath),
      this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath),
      this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath),
      this.bstructService.getBstructs(this.GONG__StackPath),
      this.dstructService.getDstructs(this.GONG__StackPath),
      this.fstructService.getFstructs(this.GONG__StackPath),
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
            fstructs_,
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
            var fstructs: FstructDB[]
            fstructs = fstructs_ as FstructDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Astructs_array = astructs

            // clear the map that counts Astruct in the GET
            this.frontRepo.Astructs_batch.clear()

            astructs.forEach(
              astruct => {
                this.frontRepo.Astructs.set(astruct.ID, astruct)
                this.frontRepo.Astructs_batch.set(astruct.ID, astruct)
              }
            )

            // clear astructs that are absent from the batch
            this.frontRepo.Astructs.forEach(
              astruct => {
                if (this.frontRepo.Astructs_batch.get(astruct.ID) == undefined) {
                  this.frontRepo.Astructs.delete(astruct.ID)
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

            // init the array
            this.frontRepo.AstructBstruct2Uses_array = astructbstruct2uses

            // clear the map that counts AstructBstruct2Use in the GET
            this.frontRepo.AstructBstruct2Uses_batch.clear()

            astructbstruct2uses.forEach(
              astructbstruct2use => {
                this.frontRepo.AstructBstruct2Uses.set(astructbstruct2use.ID, astructbstruct2use)
                this.frontRepo.AstructBstruct2Uses_batch.set(astructbstruct2use.ID, astructbstruct2use)
              }
            )

            // clear astructbstruct2uses that are absent from the batch
            this.frontRepo.AstructBstruct2Uses.forEach(
              astructbstruct2use => {
                if (this.frontRepo.AstructBstruct2Uses_batch.get(astructbstruct2use.ID) == undefined) {
                  this.frontRepo.AstructBstruct2Uses.delete(astructbstruct2use.ID)
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

            // init the array
            this.frontRepo.AstructBstructUses_array = astructbstructuses

            // clear the map that counts AstructBstructUse in the GET
            this.frontRepo.AstructBstructUses_batch.clear()

            astructbstructuses.forEach(
              astructbstructuse => {
                this.frontRepo.AstructBstructUses.set(astructbstructuse.ID, astructbstructuse)
                this.frontRepo.AstructBstructUses_batch.set(astructbstructuse.ID, astructbstructuse)
              }
            )

            // clear astructbstructuses that are absent from the batch
            this.frontRepo.AstructBstructUses.forEach(
              astructbstructuse => {
                if (this.frontRepo.AstructBstructUses_batch.get(astructbstructuse.ID) == undefined) {
                  this.frontRepo.AstructBstructUses.delete(astructbstructuse.ID)
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

            // init the array
            this.frontRepo.Bstructs_array = bstructs

            // clear the map that counts Bstruct in the GET
            this.frontRepo.Bstructs_batch.clear()

            bstructs.forEach(
              bstruct => {
                this.frontRepo.Bstructs.set(bstruct.ID, bstruct)
                this.frontRepo.Bstructs_batch.set(bstruct.ID, bstruct)
              }
            )

            // clear bstructs that are absent from the batch
            this.frontRepo.Bstructs.forEach(
              bstruct => {
                if (this.frontRepo.Bstructs_batch.get(bstruct.ID) == undefined) {
                  this.frontRepo.Bstructs.delete(bstruct.ID)
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

            // init the array
            this.frontRepo.Dstructs_array = dstructs

            // clear the map that counts Dstruct in the GET
            this.frontRepo.Dstructs_batch.clear()

            dstructs.forEach(
              dstruct => {
                this.frontRepo.Dstructs.set(dstruct.ID, dstruct)
                this.frontRepo.Dstructs_batch.set(dstruct.ID, dstruct)
              }
            )

            // clear dstructs that are absent from the batch
            this.frontRepo.Dstructs.forEach(
              dstruct => {
                if (this.frontRepo.Dstructs_batch.get(dstruct.ID) == undefined) {
                  this.frontRepo.Dstructs.delete(dstruct.ID)
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

            // init the array
            this.frontRepo.Fstructs_array = fstructs

            // clear the map that counts Fstruct in the GET
            this.frontRepo.Fstructs_batch.clear()

            fstructs.forEach(
              fstruct => {
                this.frontRepo.Fstructs.set(fstruct.ID, fstruct)
                this.frontRepo.Fstructs_batch.set(fstruct.ID, fstruct)
              }
            )

            // clear fstructs that are absent from the batch
            this.frontRepo.Fstructs.forEach(
              fstruct => {
                if (this.frontRepo.Fstructs_batch.get(fstruct.ID) == undefined) {
                  this.frontRepo.Fstructs.delete(fstruct.ID)
                }
              }
            )

            // sort Fstructs_array array
            this.frontRepo.Fstructs_array.sort((t1, t2) => {
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
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AssociationtobID.Int64)
                  if (_bstruct) {
                    astruct.Associationtob = _bstruct
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.Anotherassociationtob_2ID.Int64)
                  if (_bstruct) {
                    astruct.Anotherassociationtob_2 = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.BstructID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct2 = _bstruct
                  }
                }
                // insertion point for pointer field Dstruct redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.DstructID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct2 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.Dstruct2ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct2 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct3 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.Dstruct3ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct3 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct4 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.Dstruct4ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct4 = _dstruct
                  }
                }
                // insertion point for pointer field AnAstruct redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astruct.AnAstructID.Int64)
                  if (_astruct) {
                    astruct.AnAstruct = _astruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofa redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astruct.Astruct_AnarrayofaDBID.Int64)
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
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstruct2use.Bstrcut2ID.Int64)
                  if (_bstruct) {
                    astructbstruct2use.Bstrcut2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb2Use redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astructbstruct2use.Astruct_Anarrayofb2UseDBID.Int64)
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
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstructuse.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astructbstructuse.Bstruct2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.AnarrayofbUse redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astructbstructuse.Astruct_AnarrayofbUseDBID.Int64)
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
                  let _astruct = this.frontRepo.Astructs.get(bstruct.Astruct_AnarrayofbDBID.Int64)
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
                  let _astruct = this.frontRepo.Astructs.get(bstruct.Astruct_AnotherarrayofbDBID.Int64)
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
                // insertion point for slice of pointer field Dstruct.Anarrayofb redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(bstruct.Dstruct_AnarrayofbDBID.Int64)
                  if (_dstruct) {
                    if (_dstruct.Anarrayofb == undefined) {
                      _dstruct.Anarrayofb = new Array<BstructDB>()
                    }
                    _dstruct.Anarrayofb.push(bstruct)
                    if (bstruct.Dstruct_Anarrayofb_reverse == undefined) {
                      bstruct.Dstruct_Anarrayofb_reverse = _dstruct
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
            fstructs.forEach(
              fstruct => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem 
            astructs.forEach(
              astruct => {
                // insertion point for sorting
                astruct.Anarrayofb?.sort((t1, t2) => {
                  if (t1.Astruct_AnarrayofbDBID_Index.Int64 > t2.Astruct_AnarrayofbDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Astruct_AnarrayofbDBID_Index.Int64 < t2.Astruct_AnarrayofbDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                astruct.Anarrayofa?.sort((t1, t2) => {
                  if (t1.Astruct_AnarrayofaDBID_Index.Int64 > t2.Astruct_AnarrayofaDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Astruct_AnarrayofaDBID_Index.Int64 < t2.Astruct_AnarrayofaDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                astruct.Anotherarrayofb?.sort((t1, t2) => {
                  if (t1.Astruct_AnotherarrayofbDBID_Index.Int64 > t2.Astruct_AnotherarrayofbDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Astruct_AnotherarrayofbDBID_Index.Int64 < t2.Astruct_AnotherarrayofbDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                astruct.AnarrayofbUse?.sort((t1, t2) => {
                  if (t1.Astruct_AnarrayofbUseDBID_Index.Int64 > t2.Astruct_AnarrayofbUseDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Astruct_AnarrayofbUseDBID_Index.Int64 < t2.Astruct_AnarrayofbUseDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                astruct.Anarrayofb2Use?.sort((t1, t2) => {
                  if (t1.Astruct_Anarrayofb2UseDBID_Index.Int64 > t2.Astruct_Anarrayofb2UseDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Astruct_Anarrayofb2UseDBID_Index.Int64 < t2.Astruct_Anarrayofb2UseDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            astructbstruct2uses.forEach(
              astructbstruct2use => {
                // insertion point for sorting
              }
            )
            astructbstructuses.forEach(
              astructbstructuse => {
                // insertion point for sorting
              }
            )
            bstructs.forEach(
              bstruct => {
                // insertion point for sorting
              }
            )
            dstructs.forEach(
              dstruct => {
                // insertion point for sorting
                dstruct.Anarrayofb?.sort((t1, t2) => {
                  if (t1.Dstruct_AnarrayofbDBID_Index.Int64 > t2.Dstruct_AnarrayofbDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Dstruct_AnarrayofbDBID_Index.Int64 < t2.Dstruct_AnarrayofbDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            fstructs.forEach(
              fstruct => {
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

  // AstructPull performs a GET on Astruct of the stack and redeem association pointers 
  AstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.astructService.getAstructs(this.GONG__StackPath)
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
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.AssociationtobID.Int64)
                  if (_bstruct) {
                    astruct.Associationtob = _bstruct
                  }
                }
                // insertion point for pointer field Anotherassociationtob_2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.Anotherassociationtob_2ID.Int64)
                  if (_bstruct) {
                    astruct.Anotherassociationtob_2 = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.BstructID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct = _bstruct
                  }
                }
                // insertion point for pointer field Bstruct2 redeeming
                {
                  let _bstruct = this.frontRepo.Bstructs.get(astruct.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astruct.Bstruct2 = _bstruct
                  }
                }
                // insertion point for pointer field Dstruct redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.DstructID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct2 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.Dstruct2ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct2 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct3 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.Dstruct3ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct3 = _dstruct
                  }
                }
                // insertion point for pointer field Dstruct4 redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(astruct.Dstruct4ID.Int64)
                  if (_dstruct) {
                    astruct.Dstruct4 = _dstruct
                  }
                }
                // insertion point for pointer field AnAstruct redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astruct.AnAstructID.Int64)
                  if (_astruct) {
                    astruct.AnAstruct = _astruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofa redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astruct.Astruct_AnarrayofaDBID.Int64)
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
          this.astructbstruct2useService.getAstructBstruct2Uses(this.GONG__StackPath)
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
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstruct2use.Bstrcut2ID.Int64)
                  if (_bstruct) {
                    astructbstruct2use.Bstrcut2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb2Use redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astructbstruct2use.Astruct_Anarrayofb2UseDBID.Int64)
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
          this.astructbstructuseService.getAstructBstructUses(this.GONG__StackPath)
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
                  let _bstruct = this.frontRepo.Bstructs.get(astructbstructuse.Bstruct2ID.Int64)
                  if (_bstruct) {
                    astructbstructuse.Bstruct2 = _bstruct
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.AnarrayofbUse redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(astructbstructuse.Astruct_AnarrayofbUseDBID.Int64)
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
          this.bstructService.getBstructs(this.GONG__StackPath)
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

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Astruct.Anarrayofb redeeming
                {
                  let _astruct = this.frontRepo.Astructs.get(bstruct.Astruct_AnarrayofbDBID.Int64)
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
                  let _astruct = this.frontRepo.Astructs.get(bstruct.Astruct_AnotherarrayofbDBID.Int64)
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
                // insertion point for slice of pointer field Dstruct.Anarrayofb redeeming
                {
                  let _dstruct = this.frontRepo.Dstructs.get(bstruct.Dstruct_AnarrayofbDBID.Int64)
                  if (_dstruct) {
                    if (_dstruct.Anarrayofb == undefined) {
                      _dstruct.Anarrayofb = new Array<BstructDB>()
                    }
                    _dstruct.Anarrayofb.push(bstruct)
                    if (bstruct.Dstruct_Anarrayofb_reverse == undefined) {
                      bstruct.Dstruct_Anarrayofb_reverse = _dstruct
                    }
                  }
                }
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
          this.dstructService.getDstructs(this.GONG__StackPath)
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

                // insertion point for redeeming ONE-MANY associations
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

  // FstructPull performs a GET on Fstruct of the stack and redeem association pointers 
  FstructPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.fstructService.getFstructs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            fstructs,
          ]) => {
            // init the array
            this.frontRepo.Fstructs_array = fstructs

            // clear the map that counts Fstruct in the GET
            this.frontRepo.Fstructs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            fstructs.forEach(
              fstruct => {
                this.frontRepo.Fstructs.set(fstruct.ID, fstruct)
                this.frontRepo.Fstructs_batch.set(fstruct.ID, fstruct)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear fstructs that are absent from the GET
            this.frontRepo.Fstructs.forEach(
              fstruct => {
                if (this.frontRepo.Fstructs_batch.get(fstruct.ID) == undefined) {
                  this.frontRepo.Fstructs.delete(fstruct.ID)
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
export function getFstructUniqueID(id: number): number {
  return 53 * id
}
