import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { ClassdiagramDB } from './classdiagram-db'
import { ClassdiagramService } from './classdiagram.service'

import { ClassshapeDB } from './classshape-db'
import { ClassshapeService } from './classshape.service'

import { FieldDB } from './field-db'
import { FieldService } from './field.service'

import { GongdocCommandDB } from './gongdoccommand-db'
import { GongdocCommandService } from './gongdoccommand.service'

import { GongdocStatusDB } from './gongdocstatus-db'
import { GongdocStatusService } from './gongdocstatus.service'

import { LinkDB } from './link-db'
import { LinkService } from './link.service'

import { PkgeltDB } from './pkgelt-db'
import { PkgeltService } from './pkgelt.service'

import { PositionDB } from './position-db'
import { PositionService } from './position.service'

import { StateDB } from './state-db'
import { StateService } from './state.service'

import { UmlscDB } from './umlsc-db'
import { UmlscService } from './umlsc.service'

import { VerticeDB } from './vertice-db'
import { VerticeService } from './vertice.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Classdiagrams_array = new Array<ClassdiagramDB>(); // array of repo instances
  Classdiagrams = new Map<number, ClassdiagramDB>(); // map of repo instances
  Classdiagrams_batch = new Map<number, ClassdiagramDB>(); // same but only in last GET (for finding repo instances to delete)
  Classshapes_array = new Array<ClassshapeDB>(); // array of repo instances
  Classshapes = new Map<number, ClassshapeDB>(); // map of repo instances
  Classshapes_batch = new Map<number, ClassshapeDB>(); // same but only in last GET (for finding repo instances to delete)
  Fields_array = new Array<FieldDB>(); // array of repo instances
  Fields = new Map<number, FieldDB>(); // map of repo instances
  Fields_batch = new Map<number, FieldDB>(); // same but only in last GET (for finding repo instances to delete)
  GongdocCommands_array = new Array<GongdocCommandDB>(); // array of repo instances
  GongdocCommands = new Map<number, GongdocCommandDB>(); // map of repo instances
  GongdocCommands_batch = new Map<number, GongdocCommandDB>(); // same but only in last GET (for finding repo instances to delete)
  GongdocStatuss_array = new Array<GongdocStatusDB>(); // array of repo instances
  GongdocStatuss = new Map<number, GongdocStatusDB>(); // map of repo instances
  GongdocStatuss_batch = new Map<number, GongdocStatusDB>(); // same but only in last GET (for finding repo instances to delete)
  Links_array = new Array<LinkDB>(); // array of repo instances
  Links = new Map<number, LinkDB>(); // map of repo instances
  Links_batch = new Map<number, LinkDB>(); // same but only in last GET (for finding repo instances to delete)
  Pkgelts_array = new Array<PkgeltDB>(); // array of repo instances
  Pkgelts = new Map<number, PkgeltDB>(); // map of repo instances
  Pkgelts_batch = new Map<number, PkgeltDB>(); // same but only in last GET (for finding repo instances to delete)
  Positions_array = new Array<PositionDB>(); // array of repo instances
  Positions = new Map<number, PositionDB>(); // map of repo instances
  Positions_batch = new Map<number, PositionDB>(); // same but only in last GET (for finding repo instances to delete)
  States_array = new Array<StateDB>(); // array of repo instances
  States = new Map<number, StateDB>(); // map of repo instances
  States_batch = new Map<number, StateDB>(); // same but only in last GET (for finding repo instances to delete)
  Umlscs_array = new Array<UmlscDB>(); // array of repo instances
  Umlscs = new Map<number, UmlscDB>(); // map of repo instances
  Umlscs_batch = new Map<number, UmlscDB>(); // same but only in last GET (for finding repo instances to delete)
  Vertices_array = new Array<VerticeDB>(); // array of repo instances
  Vertices = new Map<number, VerticeDB>(); // map of repo instances
  Vertices_batch = new Map<number, VerticeDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private classdiagramService: ClassdiagramService,
    private classshapeService: ClassshapeService,
    private fieldService: FieldService,
    private gongdoccommandService: GongdocCommandService,
    private gongdocstatusService: GongdocStatusService,
    private linkService: LinkService,
    private pkgeltService: PkgeltService,
    private positionService: PositionService,
    private stateService: StateService,
    private umlscService: UmlscService,
    private verticeService: VerticeService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<ClassdiagramDB[]>,
    Observable<ClassshapeDB[]>,
    Observable<FieldDB[]>,
    Observable<GongdocCommandDB[]>,
    Observable<GongdocStatusDB[]>,
    Observable<LinkDB[]>,
    Observable<PkgeltDB[]>,
    Observable<PositionDB[]>,
    Observable<StateDB[]>,
    Observable<UmlscDB[]>,
    Observable<VerticeDB[]>,
  ] = [ // insertion point sub template 
      this.classdiagramService.getClassdiagrams(),
      this.classshapeService.getClassshapes(),
      this.fieldService.getFields(),
      this.gongdoccommandService.getGongdocCommands(),
      this.gongdocstatusService.getGongdocStatuss(),
      this.linkService.getLinks(),
      this.pkgeltService.getPkgelts(),
      this.positionService.getPositions(),
      this.stateService.getStates(),
      this.umlscService.getUmlscs(),
      this.verticeService.getVertices(),
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
            classdiagrams_,
            classshapes_,
            fields_,
            gongdoccommands_,
            gongdocstatuss_,
            links_,
            pkgelts_,
            positions_,
            states_,
            umlscs_,
            vertices_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var classdiagrams: ClassdiagramDB[]
            classdiagrams = classdiagrams_
            var classshapes: ClassshapeDB[]
            classshapes = classshapes_
            var fields: FieldDB[]
            fields = fields_
            var gongdoccommands: GongdocCommandDB[]
            gongdoccommands = gongdoccommands_
            var gongdocstatuss: GongdocStatusDB[]
            gongdocstatuss = gongdocstatuss_
            var links: LinkDB[]
            links = links_
            var pkgelts: PkgeltDB[]
            pkgelts = pkgelts_
            var positions: PositionDB[]
            positions = positions_
            var states: StateDB[]
            states = states_
            var umlscs: UmlscDB[]
            umlscs = umlscs_
            var vertices: VerticeDB[]
            vertices = vertices_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Classdiagrams_array = classdiagrams

            // clear the map that counts Classdiagram in the GET
            FrontRepoSingloton.Classdiagrams_batch.clear()
            
            classdiagrams.forEach(
              classdiagram => {
                FrontRepoSingloton.Classdiagrams.set(classdiagram.ID, classdiagram)
                FrontRepoSingloton.Classdiagrams_batch.set(classdiagram.ID, classdiagram)
              }
            )
            
            // clear classdiagrams that are absent from the batch
            FrontRepoSingloton.Classdiagrams.forEach(
              classdiagram => {
                if (FrontRepoSingloton.Classdiagrams_batch.get(classdiagram.ID) == undefined) {
                  FrontRepoSingloton.Classdiagrams.delete(classdiagram.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Classshapes_array = classshapes

            // clear the map that counts Classshape in the GET
            FrontRepoSingloton.Classshapes_batch.clear()
            
            classshapes.forEach(
              classshape => {
                FrontRepoSingloton.Classshapes.set(classshape.ID, classshape)
                FrontRepoSingloton.Classshapes_batch.set(classshape.ID, classshape)
              }
            )
            
            // clear classshapes that are absent from the batch
            FrontRepoSingloton.Classshapes.forEach(
              classshape => {
                if (FrontRepoSingloton.Classshapes_batch.get(classshape.ID) == undefined) {
                  FrontRepoSingloton.Classshapes.delete(classshape.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Fields_array = fields

            // clear the map that counts Field in the GET
            FrontRepoSingloton.Fields_batch.clear()
            
            fields.forEach(
              field => {
                FrontRepoSingloton.Fields.set(field.ID, field)
                FrontRepoSingloton.Fields_batch.set(field.ID, field)
              }
            )
            
            // clear fields that are absent from the batch
            FrontRepoSingloton.Fields.forEach(
              field => {
                if (FrontRepoSingloton.Fields_batch.get(field.ID) == undefined) {
                  FrontRepoSingloton.Fields.delete(field.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.GongdocCommands_array = gongdoccommands

            // clear the map that counts GongdocCommand in the GET
            FrontRepoSingloton.GongdocCommands_batch.clear()
            
            gongdoccommands.forEach(
              gongdoccommand => {
                FrontRepoSingloton.GongdocCommands.set(gongdoccommand.ID, gongdoccommand)
                FrontRepoSingloton.GongdocCommands_batch.set(gongdoccommand.ID, gongdoccommand)
              }
            )
            
            // clear gongdoccommands that are absent from the batch
            FrontRepoSingloton.GongdocCommands.forEach(
              gongdoccommand => {
                if (FrontRepoSingloton.GongdocCommands_batch.get(gongdoccommand.ID) == undefined) {
                  FrontRepoSingloton.GongdocCommands.delete(gongdoccommand.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.GongdocStatuss_array = gongdocstatuss

            // clear the map that counts GongdocStatus in the GET
            FrontRepoSingloton.GongdocStatuss_batch.clear()
            
            gongdocstatuss.forEach(
              gongdocstatus => {
                FrontRepoSingloton.GongdocStatuss.set(gongdocstatus.ID, gongdocstatus)
                FrontRepoSingloton.GongdocStatuss_batch.set(gongdocstatus.ID, gongdocstatus)
              }
            )
            
            // clear gongdocstatuss that are absent from the batch
            FrontRepoSingloton.GongdocStatuss.forEach(
              gongdocstatus => {
                if (FrontRepoSingloton.GongdocStatuss_batch.get(gongdocstatus.ID) == undefined) {
                  FrontRepoSingloton.GongdocStatuss.delete(gongdocstatus.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Links_array = links

            // clear the map that counts Link in the GET
            FrontRepoSingloton.Links_batch.clear()
            
            links.forEach(
              link => {
                FrontRepoSingloton.Links.set(link.ID, link)
                FrontRepoSingloton.Links_batch.set(link.ID, link)
              }
            )
            
            // clear links that are absent from the batch
            FrontRepoSingloton.Links.forEach(
              link => {
                if (FrontRepoSingloton.Links_batch.get(link.ID) == undefined) {
                  FrontRepoSingloton.Links.delete(link.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Pkgelts_array = pkgelts

            // clear the map that counts Pkgelt in the GET
            FrontRepoSingloton.Pkgelts_batch.clear()
            
            pkgelts.forEach(
              pkgelt => {
                FrontRepoSingloton.Pkgelts.set(pkgelt.ID, pkgelt)
                FrontRepoSingloton.Pkgelts_batch.set(pkgelt.ID, pkgelt)
              }
            )
            
            // clear pkgelts that are absent from the batch
            FrontRepoSingloton.Pkgelts.forEach(
              pkgelt => {
                if (FrontRepoSingloton.Pkgelts_batch.get(pkgelt.ID) == undefined) {
                  FrontRepoSingloton.Pkgelts.delete(pkgelt.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Positions_array = positions

            // clear the map that counts Position in the GET
            FrontRepoSingloton.Positions_batch.clear()
            
            positions.forEach(
              position => {
                FrontRepoSingloton.Positions.set(position.ID, position)
                FrontRepoSingloton.Positions_batch.set(position.ID, position)
              }
            )
            
            // clear positions that are absent from the batch
            FrontRepoSingloton.Positions.forEach(
              position => {
                if (FrontRepoSingloton.Positions_batch.get(position.ID) == undefined) {
                  FrontRepoSingloton.Positions.delete(position.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.States_array = states

            // clear the map that counts State in the GET
            FrontRepoSingloton.States_batch.clear()
            
            states.forEach(
              state => {
                FrontRepoSingloton.States.set(state.ID, state)
                FrontRepoSingloton.States_batch.set(state.ID, state)
              }
            )
            
            // clear states that are absent from the batch
            FrontRepoSingloton.States.forEach(
              state => {
                if (FrontRepoSingloton.States_batch.get(state.ID) == undefined) {
                  FrontRepoSingloton.States.delete(state.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Umlscs_array = umlscs

            // clear the map that counts Umlsc in the GET
            FrontRepoSingloton.Umlscs_batch.clear()
            
            umlscs.forEach(
              umlsc => {
                FrontRepoSingloton.Umlscs.set(umlsc.ID, umlsc)
                FrontRepoSingloton.Umlscs_batch.set(umlsc.ID, umlsc)
              }
            )
            
            // clear umlscs that are absent from the batch
            FrontRepoSingloton.Umlscs.forEach(
              umlsc => {
                if (FrontRepoSingloton.Umlscs_batch.get(umlsc.ID) == undefined) {
                  FrontRepoSingloton.Umlscs.delete(umlsc.ID)
                }
              }
            )
            // init the array
            FrontRepoSingloton.Vertices_array = vertices

            // clear the map that counts Vertice in the GET
            FrontRepoSingloton.Vertices_batch.clear()
            
            vertices.forEach(
              vertice => {
                FrontRepoSingloton.Vertices.set(vertice.ID, vertice)
                FrontRepoSingloton.Vertices_batch.set(vertice.ID, vertice)
              }
            )
            
            // clear vertices that are absent from the batch
            FrontRepoSingloton.Vertices.forEach(
              vertice => {
                if (FrontRepoSingloton.Vertices_batch.get(vertice.ID) == undefined) {
                  FrontRepoSingloton.Vertices.delete(vertice.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            classdiagrams.forEach(
              classdiagram => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Pkgelt.Classdiagrams redeeming
                {
                  let _pkgelt = FrontRepoSingloton.Pkgelts.get(classdiagram.Pkgelt_ClassdiagramsDBID.Int64)
                  if (_pkgelt) {
                    if (_pkgelt.Classdiagrams == undefined) {
                      _pkgelt.Classdiagrams = new Array<PkgeltDB>()
                    }
                    _pkgelt.Classdiagrams.push(classdiagram)
                    if (classdiagram.Pkgelt_Classdiagrams_reverse == undefined) {
                      classdiagram.Pkgelt_Classdiagrams_reverse = _pkgelt
                    }
                  }
                }
              }
            )
            classshapes.forEach(
              classshape => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Position redeeming
                {
                  let _position = FrontRepoSingloton.Positions.get(classshape.PositionID.Int64)
                  if (_position) {
                    classshape.Position = _position
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.Classshapes redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(classshape.Classdiagram_ClassshapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.Classshapes == undefined) {
                      _classdiagram.Classshapes = new Array<ClassdiagramDB>()
                    }
                    _classdiagram.Classshapes.push(classshape)
                    if (classshape.Classdiagram_Classshapes_reverse == undefined) {
                      classshape.Classdiagram_Classshapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )
            fields.forEach(
              field => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classshape.Fields redeeming
                {
                  let _classshape = FrontRepoSingloton.Classshapes.get(field.Classshape_FieldsDBID.Int64)
                  if (_classshape) {
                    if (_classshape.Fields == undefined) {
                      _classshape.Fields = new Array<ClassshapeDB>()
                    }
                    _classshape.Fields.push(field)
                    if (field.Classshape_Fields_reverse == undefined) {
                      field.Classshape_Fields_reverse = _classshape
                    }
                  }
                }
              }
            )
            gongdoccommands.forEach(
              gongdoccommand => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            gongdocstatuss.forEach(
              gongdocstatus => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            links.forEach(
              link => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Middlevertice redeeming
                {
                  let _vertice = FrontRepoSingloton.Vertices.get(link.MiddleverticeID.Int64)
                  if (_vertice) {
                    link.Middlevertice = _vertice
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classshape.Links redeeming
                {
                  let _classshape = FrontRepoSingloton.Classshapes.get(link.Classshape_LinksDBID.Int64)
                  if (_classshape) {
                    if (_classshape.Links == undefined) {
                      _classshape.Links = new Array<ClassshapeDB>()
                    }
                    _classshape.Links.push(link)
                    if (link.Classshape_Links_reverse == undefined) {
                      link.Classshape_Links_reverse = _classshape
                    }
                  }
                }
              }
            )
            pkgelts.forEach(
              pkgelt => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            positions.forEach(
              position => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            states.forEach(
              state => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Umlsc.States redeeming
                {
                  let _umlsc = FrontRepoSingloton.Umlscs.get(state.Umlsc_StatesDBID.Int64)
                  if (_umlsc) {
                    if (_umlsc.States == undefined) {
                      _umlsc.States = new Array<UmlscDB>()
                    }
                    _umlsc.States.push(state)
                    if (state.Umlsc_States_reverse == undefined) {
                      state.Umlsc_States_reverse = _umlsc
                    }
                  }
                }
              }
            )
            umlscs.forEach(
              umlsc => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Pkgelt.Umlscs redeeming
                {
                  let _pkgelt = FrontRepoSingloton.Pkgelts.get(umlsc.Pkgelt_UmlscsDBID.Int64)
                  if (_pkgelt) {
                    if (_pkgelt.Umlscs == undefined) {
                      _pkgelt.Umlscs = new Array<PkgeltDB>()
                    }
                    _pkgelt.Umlscs.push(umlsc)
                    if (umlsc.Pkgelt_Umlscs_reverse == undefined) {
                      umlsc.Pkgelt_Umlscs_reverse = _pkgelt
                    }
                  }
                }
              }
            )
            vertices.forEach(
              vertice => {
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

  // ClassdiagramPull performs a GET on Classdiagram of the stack and redeem association pointers 
  ClassdiagramPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.classdiagramService.getClassdiagrams()
        ]).subscribe(
          ([ // insertion point sub template 
            classdiagrams,
          ]) => {
            // init the array
            FrontRepoSingloton.Classdiagrams_array = classdiagrams

            // clear the map that counts Classdiagram in the GET
            FrontRepoSingloton.Classdiagrams_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            classdiagrams.forEach(
              classdiagram => {
                FrontRepoSingloton.Classdiagrams.set(classdiagram.ID, classdiagram)
                FrontRepoSingloton.Classdiagrams_batch.set(classdiagram.ID, classdiagram)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Pkgelt.Classdiagrams redeeming
                {
                  let _pkgelt = FrontRepoSingloton.Pkgelts.get(classdiagram.Pkgelt_ClassdiagramsDBID.Int64)
                  if (_pkgelt) {
                    if (_pkgelt.Classdiagrams == undefined) {
                      _pkgelt.Classdiagrams = new Array<PkgeltDB>()
                    }
                    _pkgelt.Classdiagrams.push(classdiagram)
                    if (classdiagram.Pkgelt_Classdiagrams_reverse == undefined) {
                      classdiagram.Pkgelt_Classdiagrams_reverse = _pkgelt
                    }
                  }
                }
              }
            )

            // clear classdiagrams that are absent from the GET
            FrontRepoSingloton.Classdiagrams.forEach(
              classdiagram => {
                if (FrontRepoSingloton.Classdiagrams_batch.get(classdiagram.ID) == undefined) {
                  FrontRepoSingloton.Classdiagrams.delete(classdiagram.ID)
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

  // ClassshapePull performs a GET on Classshape of the stack and redeem association pointers 
  ClassshapePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.classshapeService.getClassshapes()
        ]).subscribe(
          ([ // insertion point sub template 
            classshapes,
          ]) => {
            // init the array
            FrontRepoSingloton.Classshapes_array = classshapes

            // clear the map that counts Classshape in the GET
            FrontRepoSingloton.Classshapes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            classshapes.forEach(
              classshape => {
                FrontRepoSingloton.Classshapes.set(classshape.ID, classshape)
                FrontRepoSingloton.Classshapes_batch.set(classshape.ID, classshape)

                // insertion point for redeeming ONE/ZERO-ONE associations 
                // insertion point for pointer field Position redeeming
                {
                  let _position = FrontRepoSingloton.Positions.get(classshape.PositionID.Int64)
                  if (_position) {
                    classshape.Position = _position
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Classdiagram.Classshapes redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(classshape.Classdiagram_ClassshapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.Classshapes == undefined) {
                      _classdiagram.Classshapes = new Array<ClassdiagramDB>()
                    }
                    _classdiagram.Classshapes.push(classshape)
                    if (classshape.Classdiagram_Classshapes_reverse == undefined) {
                      classshape.Classdiagram_Classshapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )

            // clear classshapes that are absent from the GET
            FrontRepoSingloton.Classshapes.forEach(
              classshape => {
                if (FrontRepoSingloton.Classshapes_batch.get(classshape.ID) == undefined) {
                  FrontRepoSingloton.Classshapes.delete(classshape.ID)
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

  // FieldPull performs a GET on Field of the stack and redeem association pointers 
  FieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.fieldService.getFields()
        ]).subscribe(
          ([ // insertion point sub template 
            fields,
          ]) => {
            // init the array
            FrontRepoSingloton.Fields_array = fields

            // clear the map that counts Field in the GET
            FrontRepoSingloton.Fields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            fields.forEach(
              field => {
                FrontRepoSingloton.Fields.set(field.ID, field)
                FrontRepoSingloton.Fields_batch.set(field.ID, field)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Classshape.Fields redeeming
                {
                  let _classshape = FrontRepoSingloton.Classshapes.get(field.Classshape_FieldsDBID.Int64)
                  if (_classshape) {
                    if (_classshape.Fields == undefined) {
                      _classshape.Fields = new Array<ClassshapeDB>()
                    }
                    _classshape.Fields.push(field)
                    if (field.Classshape_Fields_reverse == undefined) {
                      field.Classshape_Fields_reverse = _classshape
                    }
                  }
                }
              }
            )

            // clear fields that are absent from the GET
            FrontRepoSingloton.Fields.forEach(
              field => {
                if (FrontRepoSingloton.Fields_batch.get(field.ID) == undefined) {
                  FrontRepoSingloton.Fields.delete(field.ID)
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

  // GongdocCommandPull performs a GET on GongdocCommand of the stack and redeem association pointers 
  GongdocCommandPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongdoccommandService.getGongdocCommands()
        ]).subscribe(
          ([ // insertion point sub template 
            gongdoccommands,
          ]) => {
            // init the array
            FrontRepoSingloton.GongdocCommands_array = gongdoccommands

            // clear the map that counts GongdocCommand in the GET
            FrontRepoSingloton.GongdocCommands_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongdoccommands.forEach(
              gongdoccommand => {
                FrontRepoSingloton.GongdocCommands.set(gongdoccommand.ID, gongdoccommand)
                FrontRepoSingloton.GongdocCommands_batch.set(gongdoccommand.ID, gongdoccommand)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear gongdoccommands that are absent from the GET
            FrontRepoSingloton.GongdocCommands.forEach(
              gongdoccommand => {
                if (FrontRepoSingloton.GongdocCommands_batch.get(gongdoccommand.ID) == undefined) {
                  FrontRepoSingloton.GongdocCommands.delete(gongdoccommand.ID)
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

  // GongdocStatusPull performs a GET on GongdocStatus of the stack and redeem association pointers 
  GongdocStatusPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongdocstatusService.getGongdocStatuss()
        ]).subscribe(
          ([ // insertion point sub template 
            gongdocstatuss,
          ]) => {
            // init the array
            FrontRepoSingloton.GongdocStatuss_array = gongdocstatuss

            // clear the map that counts GongdocStatus in the GET
            FrontRepoSingloton.GongdocStatuss_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongdocstatuss.forEach(
              gongdocstatus => {
                FrontRepoSingloton.GongdocStatuss.set(gongdocstatus.ID, gongdocstatus)
                FrontRepoSingloton.GongdocStatuss_batch.set(gongdocstatus.ID, gongdocstatus)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear gongdocstatuss that are absent from the GET
            FrontRepoSingloton.GongdocStatuss.forEach(
              gongdocstatus => {
                if (FrontRepoSingloton.GongdocStatuss_batch.get(gongdocstatus.ID) == undefined) {
                  FrontRepoSingloton.GongdocStatuss.delete(gongdocstatus.ID)
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

  // LinkPull performs a GET on Link of the stack and redeem association pointers 
  LinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linkService.getLinks()
        ]).subscribe(
          ([ // insertion point sub template 
            links,
          ]) => {
            // init the array
            FrontRepoSingloton.Links_array = links

            // clear the map that counts Link in the GET
            FrontRepoSingloton.Links_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            links.forEach(
              link => {
                FrontRepoSingloton.Links.set(link.ID, link)
                FrontRepoSingloton.Links_batch.set(link.ID, link)

                // insertion point for redeeming ONE/ZERO-ONE associations 
                // insertion point for pointer field Middlevertice redeeming
                {
                  let _vertice = FrontRepoSingloton.Vertices.get(link.MiddleverticeID.Int64)
                  if (_vertice) {
                    link.Middlevertice = _vertice
                  }
                }

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Classshape.Links redeeming
                {
                  let _classshape = FrontRepoSingloton.Classshapes.get(link.Classshape_LinksDBID.Int64)
                  if (_classshape) {
                    if (_classshape.Links == undefined) {
                      _classshape.Links = new Array<ClassshapeDB>()
                    }
                    _classshape.Links.push(link)
                    if (link.Classshape_Links_reverse == undefined) {
                      link.Classshape_Links_reverse = _classshape
                    }
                  }
                }
              }
            )

            // clear links that are absent from the GET
            FrontRepoSingloton.Links.forEach(
              link => {
                if (FrontRepoSingloton.Links_batch.get(link.ID) == undefined) {
                  FrontRepoSingloton.Links.delete(link.ID)
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

  // PkgeltPull performs a GET on Pkgelt of the stack and redeem association pointers 
  PkgeltPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.pkgeltService.getPkgelts()
        ]).subscribe(
          ([ // insertion point sub template 
            pkgelts,
          ]) => {
            // init the array
            FrontRepoSingloton.Pkgelts_array = pkgelts

            // clear the map that counts Pkgelt in the GET
            FrontRepoSingloton.Pkgelts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            pkgelts.forEach(
              pkgelt => {
                FrontRepoSingloton.Pkgelts.set(pkgelt.ID, pkgelt)
                FrontRepoSingloton.Pkgelts_batch.set(pkgelt.ID, pkgelt)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear pkgelts that are absent from the GET
            FrontRepoSingloton.Pkgelts.forEach(
              pkgelt => {
                if (FrontRepoSingloton.Pkgelts_batch.get(pkgelt.ID) == undefined) {
                  FrontRepoSingloton.Pkgelts.delete(pkgelt.ID)
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

  // PositionPull performs a GET on Position of the stack and redeem association pointers 
  PositionPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.positionService.getPositions()
        ]).subscribe(
          ([ // insertion point sub template 
            positions,
          ]) => {
            // init the array
            FrontRepoSingloton.Positions_array = positions

            // clear the map that counts Position in the GET
            FrontRepoSingloton.Positions_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            positions.forEach(
              position => {
                FrontRepoSingloton.Positions.set(position.ID, position)
                FrontRepoSingloton.Positions_batch.set(position.ID, position)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear positions that are absent from the GET
            FrontRepoSingloton.Positions.forEach(
              position => {
                if (FrontRepoSingloton.Positions_batch.get(position.ID) == undefined) {
                  FrontRepoSingloton.Positions.delete(position.ID)
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

  // StatePull performs a GET on State of the stack and redeem association pointers 
  StatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.stateService.getStates()
        ]).subscribe(
          ([ // insertion point sub template 
            states,
          ]) => {
            // init the array
            FrontRepoSingloton.States_array = states

            // clear the map that counts State in the GET
            FrontRepoSingloton.States_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            states.forEach(
              state => {
                FrontRepoSingloton.States.set(state.ID, state)
                FrontRepoSingloton.States_batch.set(state.ID, state)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Umlsc.States redeeming
                {
                  let _umlsc = FrontRepoSingloton.Umlscs.get(state.Umlsc_StatesDBID.Int64)
                  if (_umlsc) {
                    if (_umlsc.States == undefined) {
                      _umlsc.States = new Array<UmlscDB>()
                    }
                    _umlsc.States.push(state)
                    if (state.Umlsc_States_reverse == undefined) {
                      state.Umlsc_States_reverse = _umlsc
                    }
                  }
                }
              }
            )

            // clear states that are absent from the GET
            FrontRepoSingloton.States.forEach(
              state => {
                if (FrontRepoSingloton.States_batch.get(state.ID) == undefined) {
                  FrontRepoSingloton.States.delete(state.ID)
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

  // UmlscPull performs a GET on Umlsc of the stack and redeem association pointers 
  UmlscPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.umlscService.getUmlscs()
        ]).subscribe(
          ([ // insertion point sub template 
            umlscs,
          ]) => {
            // init the array
            FrontRepoSingloton.Umlscs_array = umlscs

            // clear the map that counts Umlsc in the GET
            FrontRepoSingloton.Umlscs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            umlscs.forEach(
              umlsc => {
                FrontRepoSingloton.Umlscs.set(umlsc.ID, umlsc)
                FrontRepoSingloton.Umlscs_batch.set(umlsc.ID, umlsc)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Pkgelt.Umlscs redeeming
                {
                  let _pkgelt = FrontRepoSingloton.Pkgelts.get(umlsc.Pkgelt_UmlscsDBID.Int64)
                  if (_pkgelt) {
                    if (_pkgelt.Umlscs == undefined) {
                      _pkgelt.Umlscs = new Array<PkgeltDB>()
                    }
                    _pkgelt.Umlscs.push(umlsc)
                    if (umlsc.Pkgelt_Umlscs_reverse == undefined) {
                      umlsc.Pkgelt_Umlscs_reverse = _pkgelt
                    }
                  }
                }
              }
            )

            // clear umlscs that are absent from the GET
            FrontRepoSingloton.Umlscs.forEach(
              umlsc => {
                if (FrontRepoSingloton.Umlscs_batch.get(umlsc.ID) == undefined) {
                  FrontRepoSingloton.Umlscs.delete(umlsc.ID)
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

  // VerticePull performs a GET on Vertice of the stack and redeem association pointers 
  VerticePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.verticeService.getVertices()
        ]).subscribe(
          ([ // insertion point sub template 
            vertices,
          ]) => {
            // init the array
            FrontRepoSingloton.Vertices_array = vertices

            // clear the map that counts Vertice in the GET
            FrontRepoSingloton.Vertices_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            vertices.forEach(
              vertice => {
                FrontRepoSingloton.Vertices.set(vertice.ID, vertice)
                FrontRepoSingloton.Vertices_batch.set(vertice.ID, vertice)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear vertices that are absent from the GET
            FrontRepoSingloton.Vertices.forEach(
              vertice => {
                if (FrontRepoSingloton.Vertices_batch.get(vertice.ID) == undefined) {
                  FrontRepoSingloton.Vertices.delete(vertice.ID)
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
export function getClassdiagramUniqueID(id: number): number {
  return 31 * id
}
export function getClassshapeUniqueID(id: number): number {
  return 37 * id
}
export function getFieldUniqueID(id: number): number {
  return 41 * id
}
export function getGongdocCommandUniqueID(id: number): number {
  return 43 * id
}
export function getGongdocStatusUniqueID(id: number): number {
  return 47 * id
}
export function getLinkUniqueID(id: number): number {
  return 53 * id
}
export function getPkgeltUniqueID(id: number): number {
  return 59 * id
}
export function getPositionUniqueID(id: number): number {
  return 61 * id
}
export function getStateUniqueID(id: number): number {
  return 67 * id
}
export function getUmlscUniqueID(id: number): number {
  return 71 * id
}
export function getVerticeUniqueID(id: number): number {
  return 73 * id
}
