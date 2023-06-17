import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { ClassdiagramDB } from './classdiagram-db'
import { ClassdiagramService } from './classdiagram.service'

import { DiagramPackageDB } from './diagrampackage-db'
import { DiagramPackageService } from './diagrampackage.service'

import { FieldDB } from './field-db'
import { FieldService } from './field.service'

import { GongEnumShapeDB } from './gongenumshape-db'
import { GongEnumShapeService } from './gongenumshape.service'

import { GongEnumValueEntryDB } from './gongenumvalueentry-db'
import { GongEnumValueEntryService } from './gongenumvalueentry.service'

import { GongStructShapeDB } from './gongstructshape-db'
import { GongStructShapeService } from './gongstructshape.service'

import { LinkDB } from './link-db'
import { LinkService } from './link.service'

import { NoteShapeDB } from './noteshape-db'
import { NoteShapeService } from './noteshape.service'

import { NoteShapeLinkDB } from './noteshapelink-db'
import { NoteShapeLinkService } from './noteshapelink.service'

import { PositionDB } from './position-db'
import { PositionService } from './position.service'

import { UmlStateDB } from './umlstate-db'
import { UmlStateService } from './umlstate.service'

import { UmlscDB } from './umlsc-db'
import { UmlscService } from './umlsc.service'

import { VerticeDB } from './vertice-db'
import { VerticeService } from './vertice.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Classdiagrams_array = new Array<ClassdiagramDB>(); // array of repo instances
  Classdiagrams = new Map<number, ClassdiagramDB>(); // map of repo instances
  Classdiagrams_batch = new Map<number, ClassdiagramDB>(); // same but only in last GET (for finding repo instances to delete)
  DiagramPackages_array = new Array<DiagramPackageDB>(); // array of repo instances
  DiagramPackages = new Map<number, DiagramPackageDB>(); // map of repo instances
  DiagramPackages_batch = new Map<number, DiagramPackageDB>(); // same but only in last GET (for finding repo instances to delete)
  Fields_array = new Array<FieldDB>(); // array of repo instances
  Fields = new Map<number, FieldDB>(); // map of repo instances
  Fields_batch = new Map<number, FieldDB>(); // same but only in last GET (for finding repo instances to delete)
  GongEnumShapes_array = new Array<GongEnumShapeDB>(); // array of repo instances
  GongEnumShapes = new Map<number, GongEnumShapeDB>(); // map of repo instances
  GongEnumShapes_batch = new Map<number, GongEnumShapeDB>(); // same but only in last GET (for finding repo instances to delete)
  GongEnumValueEntrys_array = new Array<GongEnumValueEntryDB>(); // array of repo instances
  GongEnumValueEntrys = new Map<number, GongEnumValueEntryDB>(); // map of repo instances
  GongEnumValueEntrys_batch = new Map<number, GongEnumValueEntryDB>(); // same but only in last GET (for finding repo instances to delete)
  GongStructShapes_array = new Array<GongStructShapeDB>(); // array of repo instances
  GongStructShapes = new Map<number, GongStructShapeDB>(); // map of repo instances
  GongStructShapes_batch = new Map<number, GongStructShapeDB>(); // same but only in last GET (for finding repo instances to delete)
  Links_array = new Array<LinkDB>(); // array of repo instances
  Links = new Map<number, LinkDB>(); // map of repo instances
  Links_batch = new Map<number, LinkDB>(); // same but only in last GET (for finding repo instances to delete)
  NoteShapes_array = new Array<NoteShapeDB>(); // array of repo instances
  NoteShapes = new Map<number, NoteShapeDB>(); // map of repo instances
  NoteShapes_batch = new Map<number, NoteShapeDB>(); // same but only in last GET (for finding repo instances to delete)
  NoteShapeLinks_array = new Array<NoteShapeLinkDB>(); // array of repo instances
  NoteShapeLinks = new Map<number, NoteShapeLinkDB>(); // map of repo instances
  NoteShapeLinks_batch = new Map<number, NoteShapeLinkDB>(); // same but only in last GET (for finding repo instances to delete)
  Positions_array = new Array<PositionDB>(); // array of repo instances
  Positions = new Map<number, PositionDB>(); // map of repo instances
  Positions_batch = new Map<number, PositionDB>(); // same but only in last GET (for finding repo instances to delete)
  UmlStates_array = new Array<UmlStateDB>(); // array of repo instances
  UmlStates = new Map<number, UmlStateDB>(); // map of repo instances
  UmlStates_batch = new Map<number, UmlStateDB>(); // same but only in last GET (for finding repo instances to delete)
  Umlscs_array = new Array<UmlscDB>(); // array of repo instances
  Umlscs = new Map<number, UmlscDB>(); // map of repo instances
  Umlscs_batch = new Map<number, UmlscDB>(); // same but only in last GET (for finding repo instances to delete)
  Vertices_array = new Array<VerticeDB>(); // array of repo instances
  Vertices = new Map<number, VerticeDB>(); // map of repo instances
  Vertices_batch = new Map<number, VerticeDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private classdiagramService: ClassdiagramService,
    private diagrampackageService: DiagramPackageService,
    private fieldService: FieldService,
    private gongenumshapeService: GongEnumShapeService,
    private gongenumvalueentryService: GongEnumValueEntryService,
    private gongstructshapeService: GongStructShapeService,
    private linkService: LinkService,
    private noteshapeService: NoteShapeService,
    private noteshapelinkService: NoteShapeLinkService,
    private positionService: PositionService,
    private umlstateService: UmlStateService,
    private umlscService: UmlscService,
    private verticeService: VerticeService,
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
    Observable<ClassdiagramDB[]>,
    Observable<DiagramPackageDB[]>,
    Observable<FieldDB[]>,
    Observable<GongEnumShapeDB[]>,
    Observable<GongEnumValueEntryDB[]>,
    Observable<GongStructShapeDB[]>,
    Observable<LinkDB[]>,
    Observable<NoteShapeDB[]>,
    Observable<NoteShapeLinkDB[]>,
    Observable<PositionDB[]>,
    Observable<UmlStateDB[]>,
    Observable<UmlscDB[]>,
    Observable<VerticeDB[]>,
  ] = [ // insertion point sub template
      this.classdiagramService.getClassdiagrams(this.GONG__StackPath),
      this.diagrampackageService.getDiagramPackages(this.GONG__StackPath),
      this.fieldService.getFields(this.GONG__StackPath),
      this.gongenumshapeService.getGongEnumShapes(this.GONG__StackPath),
      this.gongenumvalueentryService.getGongEnumValueEntrys(this.GONG__StackPath),
      this.gongstructshapeService.getGongStructShapes(this.GONG__StackPath),
      this.linkService.getLinks(this.GONG__StackPath),
      this.noteshapeService.getNoteShapes(this.GONG__StackPath),
      this.noteshapelinkService.getNoteShapeLinks(this.GONG__StackPath),
      this.positionService.getPositions(this.GONG__StackPath),
      this.umlstateService.getUmlStates(this.GONG__StackPath),
      this.umlscService.getUmlscs(this.GONG__StackPath),
      this.verticeService.getVertices(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.classdiagramService.getClassdiagrams(this.GONG__StackPath),
      this.diagrampackageService.getDiagramPackages(this.GONG__StackPath),
      this.fieldService.getFields(this.GONG__StackPath),
      this.gongenumshapeService.getGongEnumShapes(this.GONG__StackPath),
      this.gongenumvalueentryService.getGongEnumValueEntrys(this.GONG__StackPath),
      this.gongstructshapeService.getGongStructShapes(this.GONG__StackPath),
      this.linkService.getLinks(this.GONG__StackPath),
      this.noteshapeService.getNoteShapes(this.GONG__StackPath),
      this.noteshapelinkService.getNoteShapeLinks(this.GONG__StackPath),
      this.positionService.getPositions(this.GONG__StackPath),
      this.umlstateService.getUmlStates(this.GONG__StackPath),
      this.umlscService.getUmlscs(this.GONG__StackPath),
      this.verticeService.getVertices(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            classdiagrams_,
            diagrampackages_,
            fields_,
            gongenumshapes_,
            gongenumvalueentrys_,
            gongstructshapes_,
            links_,
            noteshapes_,
            noteshapelinks_,
            positions_,
            umlstates_,
            umlscs_,
            vertices_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var classdiagrams: ClassdiagramDB[]
            classdiagrams = classdiagrams_ as ClassdiagramDB[]
            var diagrampackages: DiagramPackageDB[]
            diagrampackages = diagrampackages_ as DiagramPackageDB[]
            var fields: FieldDB[]
            fields = fields_ as FieldDB[]
            var gongenumshapes: GongEnumShapeDB[]
            gongenumshapes = gongenumshapes_ as GongEnumShapeDB[]
            var gongenumvalueentrys: GongEnumValueEntryDB[]
            gongenumvalueentrys = gongenumvalueentrys_ as GongEnumValueEntryDB[]
            var gongstructshapes: GongStructShapeDB[]
            gongstructshapes = gongstructshapes_ as GongStructShapeDB[]
            var links: LinkDB[]
            links = links_ as LinkDB[]
            var noteshapes: NoteShapeDB[]
            noteshapes = noteshapes_ as NoteShapeDB[]
            var noteshapelinks: NoteShapeLinkDB[]
            noteshapelinks = noteshapelinks_ as NoteShapeLinkDB[]
            var positions: PositionDB[]
            positions = positions_ as PositionDB[]
            var umlstates: UmlStateDB[]
            umlstates = umlstates_ as UmlStateDB[]
            var umlscs: UmlscDB[]
            umlscs = umlscs_ as UmlscDB[]
            var vertices: VerticeDB[]
            vertices = vertices_ as VerticeDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Classdiagrams_array = classdiagrams

            // clear the map that counts Classdiagram in the GET
            this.frontRepo.Classdiagrams_batch.clear()

            classdiagrams.forEach(
              classdiagram => {
                this.frontRepo.Classdiagrams.set(classdiagram.ID, classdiagram)
                this.frontRepo.Classdiagrams_batch.set(classdiagram.ID, classdiagram)
              }
            )

            // clear classdiagrams that are absent from the batch
            this.frontRepo.Classdiagrams.forEach(
              classdiagram => {
                if (this.frontRepo.Classdiagrams_batch.get(classdiagram.ID) == undefined) {
                  this.frontRepo.Classdiagrams.delete(classdiagram.ID)
                }
              }
            )

            // sort Classdiagrams_array array
            this.frontRepo.Classdiagrams_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.DiagramPackages_array = diagrampackages

            // clear the map that counts DiagramPackage in the GET
            this.frontRepo.DiagramPackages_batch.clear()

            diagrampackages.forEach(
              diagrampackage => {
                this.frontRepo.DiagramPackages.set(diagrampackage.ID, diagrampackage)
                this.frontRepo.DiagramPackages_batch.set(diagrampackage.ID, diagrampackage)
              }
            )

            // clear diagrampackages that are absent from the batch
            this.frontRepo.DiagramPackages.forEach(
              diagrampackage => {
                if (this.frontRepo.DiagramPackages_batch.get(diagrampackage.ID) == undefined) {
                  this.frontRepo.DiagramPackages.delete(diagrampackage.ID)
                }
              }
            )

            // sort DiagramPackages_array array
            this.frontRepo.DiagramPackages_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Fields_array = fields

            // clear the map that counts Field in the GET
            this.frontRepo.Fields_batch.clear()

            fields.forEach(
              field => {
                this.frontRepo.Fields.set(field.ID, field)
                this.frontRepo.Fields_batch.set(field.ID, field)
              }
            )

            // clear fields that are absent from the batch
            this.frontRepo.Fields.forEach(
              field => {
                if (this.frontRepo.Fields_batch.get(field.ID) == undefined) {
                  this.frontRepo.Fields.delete(field.ID)
                }
              }
            )

            // sort Fields_array array
            this.frontRepo.Fields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongEnumShapes_array = gongenumshapes

            // clear the map that counts GongEnumShape in the GET
            this.frontRepo.GongEnumShapes_batch.clear()

            gongenumshapes.forEach(
              gongenumshape => {
                this.frontRepo.GongEnumShapes.set(gongenumshape.ID, gongenumshape)
                this.frontRepo.GongEnumShapes_batch.set(gongenumshape.ID, gongenumshape)
              }
            )

            // clear gongenumshapes that are absent from the batch
            this.frontRepo.GongEnumShapes.forEach(
              gongenumshape => {
                if (this.frontRepo.GongEnumShapes_batch.get(gongenumshape.ID) == undefined) {
                  this.frontRepo.GongEnumShapes.delete(gongenumshape.ID)
                }
              }
            )

            // sort GongEnumShapes_array array
            this.frontRepo.GongEnumShapes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongEnumValueEntrys_array = gongenumvalueentrys

            // clear the map that counts GongEnumValueEntry in the GET
            this.frontRepo.GongEnumValueEntrys_batch.clear()

            gongenumvalueentrys.forEach(
              gongenumvalueentry => {
                this.frontRepo.GongEnumValueEntrys.set(gongenumvalueentry.ID, gongenumvalueentry)
                this.frontRepo.GongEnumValueEntrys_batch.set(gongenumvalueentry.ID, gongenumvalueentry)
              }
            )

            // clear gongenumvalueentrys that are absent from the batch
            this.frontRepo.GongEnumValueEntrys.forEach(
              gongenumvalueentry => {
                if (this.frontRepo.GongEnumValueEntrys_batch.get(gongenumvalueentry.ID) == undefined) {
                  this.frontRepo.GongEnumValueEntrys.delete(gongenumvalueentry.ID)
                }
              }
            )

            // sort GongEnumValueEntrys_array array
            this.frontRepo.GongEnumValueEntrys_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.GongStructShapes_array = gongstructshapes

            // clear the map that counts GongStructShape in the GET
            this.frontRepo.GongStructShapes_batch.clear()

            gongstructshapes.forEach(
              gongstructshape => {
                this.frontRepo.GongStructShapes.set(gongstructshape.ID, gongstructshape)
                this.frontRepo.GongStructShapes_batch.set(gongstructshape.ID, gongstructshape)
              }
            )

            // clear gongstructshapes that are absent from the batch
            this.frontRepo.GongStructShapes.forEach(
              gongstructshape => {
                if (this.frontRepo.GongStructShapes_batch.get(gongstructshape.ID) == undefined) {
                  this.frontRepo.GongStructShapes.delete(gongstructshape.ID)
                }
              }
            )

            // sort GongStructShapes_array array
            this.frontRepo.GongStructShapes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Links_array = links

            // clear the map that counts Link in the GET
            this.frontRepo.Links_batch.clear()

            links.forEach(
              link => {
                this.frontRepo.Links.set(link.ID, link)
                this.frontRepo.Links_batch.set(link.ID, link)
              }
            )

            // clear links that are absent from the batch
            this.frontRepo.Links.forEach(
              link => {
                if (this.frontRepo.Links_batch.get(link.ID) == undefined) {
                  this.frontRepo.Links.delete(link.ID)
                }
              }
            )

            // sort Links_array array
            this.frontRepo.Links_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.NoteShapes_array = noteshapes

            // clear the map that counts NoteShape in the GET
            this.frontRepo.NoteShapes_batch.clear()

            noteshapes.forEach(
              noteshape => {
                this.frontRepo.NoteShapes.set(noteshape.ID, noteshape)
                this.frontRepo.NoteShapes_batch.set(noteshape.ID, noteshape)
              }
            )

            // clear noteshapes that are absent from the batch
            this.frontRepo.NoteShapes.forEach(
              noteshape => {
                if (this.frontRepo.NoteShapes_batch.get(noteshape.ID) == undefined) {
                  this.frontRepo.NoteShapes.delete(noteshape.ID)
                }
              }
            )

            // sort NoteShapes_array array
            this.frontRepo.NoteShapes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.NoteShapeLinks_array = noteshapelinks

            // clear the map that counts NoteShapeLink in the GET
            this.frontRepo.NoteShapeLinks_batch.clear()

            noteshapelinks.forEach(
              noteshapelink => {
                this.frontRepo.NoteShapeLinks.set(noteshapelink.ID, noteshapelink)
                this.frontRepo.NoteShapeLinks_batch.set(noteshapelink.ID, noteshapelink)
              }
            )

            // clear noteshapelinks that are absent from the batch
            this.frontRepo.NoteShapeLinks.forEach(
              noteshapelink => {
                if (this.frontRepo.NoteShapeLinks_batch.get(noteshapelink.ID) == undefined) {
                  this.frontRepo.NoteShapeLinks.delete(noteshapelink.ID)
                }
              }
            )

            // sort NoteShapeLinks_array array
            this.frontRepo.NoteShapeLinks_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Positions_array = positions

            // clear the map that counts Position in the GET
            this.frontRepo.Positions_batch.clear()

            positions.forEach(
              position => {
                this.frontRepo.Positions.set(position.ID, position)
                this.frontRepo.Positions_batch.set(position.ID, position)
              }
            )

            // clear positions that are absent from the batch
            this.frontRepo.Positions.forEach(
              position => {
                if (this.frontRepo.Positions_batch.get(position.ID) == undefined) {
                  this.frontRepo.Positions.delete(position.ID)
                }
              }
            )

            // sort Positions_array array
            this.frontRepo.Positions_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.UmlStates_array = umlstates

            // clear the map that counts UmlState in the GET
            this.frontRepo.UmlStates_batch.clear()

            umlstates.forEach(
              umlstate => {
                this.frontRepo.UmlStates.set(umlstate.ID, umlstate)
                this.frontRepo.UmlStates_batch.set(umlstate.ID, umlstate)
              }
            )

            // clear umlstates that are absent from the batch
            this.frontRepo.UmlStates.forEach(
              umlstate => {
                if (this.frontRepo.UmlStates_batch.get(umlstate.ID) == undefined) {
                  this.frontRepo.UmlStates.delete(umlstate.ID)
                }
              }
            )

            // sort UmlStates_array array
            this.frontRepo.UmlStates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Umlscs_array = umlscs

            // clear the map that counts Umlsc in the GET
            this.frontRepo.Umlscs_batch.clear()

            umlscs.forEach(
              umlsc => {
                this.frontRepo.Umlscs.set(umlsc.ID, umlsc)
                this.frontRepo.Umlscs_batch.set(umlsc.ID, umlsc)
              }
            )

            // clear umlscs that are absent from the batch
            this.frontRepo.Umlscs.forEach(
              umlsc => {
                if (this.frontRepo.Umlscs_batch.get(umlsc.ID) == undefined) {
                  this.frontRepo.Umlscs.delete(umlsc.ID)
                }
              }
            )

            // sort Umlscs_array array
            this.frontRepo.Umlscs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Vertices_array = vertices

            // clear the map that counts Vertice in the GET
            this.frontRepo.Vertices_batch.clear()

            vertices.forEach(
              vertice => {
                this.frontRepo.Vertices.set(vertice.ID, vertice)
                this.frontRepo.Vertices_batch.set(vertice.ID, vertice)
              }
            )

            // clear vertices that are absent from the batch
            this.frontRepo.Vertices.forEach(
              vertice => {
                if (this.frontRepo.Vertices_batch.get(vertice.ID) == undefined) {
                  this.frontRepo.Vertices.delete(vertice.ID)
                }
              }
            )

            // sort Vertices_array array
            this.frontRepo.Vertices_array.sort((t1, t2) => {
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
            classdiagrams.forEach(
              classdiagram => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field DiagramPackage.Classdiagrams redeeming
                {
                  let _diagrampackage = this.frontRepo.DiagramPackages.get(classdiagram.DiagramPackage_ClassdiagramsDBID.Int64)
                  if (_diagrampackage) {
                    if (_diagrampackage.Classdiagrams == undefined) {
                      _diagrampackage.Classdiagrams = new Array<ClassdiagramDB>()
                    }
                    _diagrampackage.Classdiagrams.push(classdiagram)
                    if (classdiagram.DiagramPackage_Classdiagrams_reverse == undefined) {
                      classdiagram.DiagramPackage_Classdiagrams_reverse = _diagrampackage
                    }
                  }
                }
              }
            )
            diagrampackages.forEach(
              diagrampackage => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field SelectedClassdiagram redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(diagrampackage.SelectedClassdiagramID.Int64)
                  if (_classdiagram) {
                    diagrampackage.SelectedClassdiagram = _classdiagram
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            fields.forEach(
              field => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStructShape.Fields redeeming
                {
                  let _gongstructshape = this.frontRepo.GongStructShapes.get(field.GongStructShape_FieldsDBID.Int64)
                  if (_gongstructshape) {
                    if (_gongstructshape.Fields == undefined) {
                      _gongstructshape.Fields = new Array<FieldDB>()
                    }
                    _gongstructshape.Fields.push(field)
                    if (field.GongStructShape_Fields_reverse == undefined) {
                      field.GongStructShape_Fields_reverse = _gongstructshape
                    }
                  }
                }
              }
            )
            gongenumshapes.forEach(
              gongenumshape => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Position redeeming
                {
                  let _position = this.frontRepo.Positions.get(gongenumshape.PositionID.Int64)
                  if (_position) {
                    gongenumshape.Position = _position
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.GongEnumShapes redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(gongenumshape.Classdiagram_GongEnumShapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.GongEnumShapes == undefined) {
                      _classdiagram.GongEnumShapes = new Array<GongEnumShapeDB>()
                    }
                    _classdiagram.GongEnumShapes.push(gongenumshape)
                    if (gongenumshape.Classdiagram_GongEnumShapes_reverse == undefined) {
                      gongenumshape.Classdiagram_GongEnumShapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )
            gongenumvalueentrys.forEach(
              gongenumvalueentry => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongEnumShape.GongEnumValueEntrys redeeming
                {
                  let _gongenumshape = this.frontRepo.GongEnumShapes.get(gongenumvalueentry.GongEnumShape_GongEnumValueEntrysDBID.Int64)
                  if (_gongenumshape) {
                    if (_gongenumshape.GongEnumValueEntrys == undefined) {
                      _gongenumshape.GongEnumValueEntrys = new Array<GongEnumValueEntryDB>()
                    }
                    _gongenumshape.GongEnumValueEntrys.push(gongenumvalueentry)
                    if (gongenumvalueentry.GongEnumShape_GongEnumValueEntrys_reverse == undefined) {
                      gongenumvalueentry.GongEnumShape_GongEnumValueEntrys_reverse = _gongenumshape
                    }
                  }
                }
              }
            )
            gongstructshapes.forEach(
              gongstructshape => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Position redeeming
                {
                  let _position = this.frontRepo.Positions.get(gongstructshape.PositionID.Int64)
                  if (_position) {
                    gongstructshape.Position = _position
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.GongStructShapes redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(gongstructshape.Classdiagram_GongStructShapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.GongStructShapes == undefined) {
                      _classdiagram.GongStructShapes = new Array<GongStructShapeDB>()
                    }
                    _classdiagram.GongStructShapes.push(gongstructshape)
                    if (gongstructshape.Classdiagram_GongStructShapes_reverse == undefined) {
                      gongstructshape.Classdiagram_GongStructShapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )
            links.forEach(
              link => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Middlevertice redeeming
                {
                  let _vertice = this.frontRepo.Vertices.get(link.MiddleverticeID.Int64)
                  if (_vertice) {
                    link.Middlevertice = _vertice
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStructShape.Links redeeming
                {
                  let _gongstructshape = this.frontRepo.GongStructShapes.get(link.GongStructShape_LinksDBID.Int64)
                  if (_gongstructshape) {
                    if (_gongstructshape.Links == undefined) {
                      _gongstructshape.Links = new Array<LinkDB>()
                    }
                    _gongstructshape.Links.push(link)
                    if (link.GongStructShape_Links_reverse == undefined) {
                      link.GongStructShape_Links_reverse = _gongstructshape
                    }
                  }
                }
              }
            )
            noteshapes.forEach(
              noteshape => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.NoteShapes redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(noteshape.Classdiagram_NoteShapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.NoteShapes == undefined) {
                      _classdiagram.NoteShapes = new Array<NoteShapeDB>()
                    }
                    _classdiagram.NoteShapes.push(noteshape)
                    if (noteshape.Classdiagram_NoteShapes_reverse == undefined) {
                      noteshape.Classdiagram_NoteShapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )
            noteshapelinks.forEach(
              noteshapelink => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field NoteShape.NoteShapeLinks redeeming
                {
                  let _noteshape = this.frontRepo.NoteShapes.get(noteshapelink.NoteShape_NoteShapeLinksDBID.Int64)
                  if (_noteshape) {
                    if (_noteshape.NoteShapeLinks == undefined) {
                      _noteshape.NoteShapeLinks = new Array<NoteShapeLinkDB>()
                    }
                    _noteshape.NoteShapeLinks.push(noteshapelink)
                    if (noteshapelink.NoteShape_NoteShapeLinks_reverse == undefined) {
                      noteshapelink.NoteShape_NoteShapeLinks_reverse = _noteshape
                    }
                  }
                }
              }
            )
            positions.forEach(
              position => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            umlstates.forEach(
              umlstate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Umlsc.States redeeming
                {
                  let _umlsc = this.frontRepo.Umlscs.get(umlstate.Umlsc_StatesDBID.Int64)
                  if (_umlsc) {
                    if (_umlsc.States == undefined) {
                      _umlsc.States = new Array<UmlStateDB>()
                    }
                    _umlsc.States.push(umlstate)
                    if (umlstate.Umlsc_States_reverse == undefined) {
                      umlstate.Umlsc_States_reverse = _umlsc
                    }
                  }
                }
              }
            )
            umlscs.forEach(
              umlsc => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field DiagramPackage.Umlscs redeeming
                {
                  let _diagrampackage = this.frontRepo.DiagramPackages.get(umlsc.DiagramPackage_UmlscsDBID.Int64)
                  if (_diagrampackage) {
                    if (_diagrampackage.Umlscs == undefined) {
                      _diagrampackage.Umlscs = new Array<UmlscDB>()
                    }
                    _diagrampackage.Umlscs.push(umlsc)
                    if (umlsc.DiagramPackage_Umlscs_reverse == undefined) {
                      umlsc.DiagramPackage_Umlscs_reverse = _diagrampackage
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
            observer.next(this.frontRepo)
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
          this.classdiagramService.getClassdiagrams(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            classdiagrams,
          ]) => {
            // init the array
            this.frontRepo.Classdiagrams_array = classdiagrams

            // clear the map that counts Classdiagram in the GET
            this.frontRepo.Classdiagrams_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            classdiagrams.forEach(
              classdiagram => {
                this.frontRepo.Classdiagrams.set(classdiagram.ID, classdiagram)
                this.frontRepo.Classdiagrams_batch.set(classdiagram.ID, classdiagram)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field DiagramPackage.Classdiagrams redeeming
                {
                  let _diagrampackage = this.frontRepo.DiagramPackages.get(classdiagram.DiagramPackage_ClassdiagramsDBID.Int64)
                  if (_diagrampackage) {
                    if (_diagrampackage.Classdiagrams == undefined) {
                      _diagrampackage.Classdiagrams = new Array<ClassdiagramDB>()
                    }
                    _diagrampackage.Classdiagrams.push(classdiagram)
                    if (classdiagram.DiagramPackage_Classdiagrams_reverse == undefined) {
                      classdiagram.DiagramPackage_Classdiagrams_reverse = _diagrampackage
                    }
                  }
                }
              }
            )

            // clear classdiagrams that are absent from the GET
            this.frontRepo.Classdiagrams.forEach(
              classdiagram => {
                if (this.frontRepo.Classdiagrams_batch.get(classdiagram.ID) == undefined) {
                  this.frontRepo.Classdiagrams.delete(classdiagram.ID)
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

  // DiagramPackagePull performs a GET on DiagramPackage of the stack and redeem association pointers 
  DiagramPackagePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.diagrampackageService.getDiagramPackages(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            diagrampackages,
          ]) => {
            // init the array
            this.frontRepo.DiagramPackages_array = diagrampackages

            // clear the map that counts DiagramPackage in the GET
            this.frontRepo.DiagramPackages_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            diagrampackages.forEach(
              diagrampackage => {
                this.frontRepo.DiagramPackages.set(diagrampackage.ID, diagrampackage)
                this.frontRepo.DiagramPackages_batch.set(diagrampackage.ID, diagrampackage)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field SelectedClassdiagram redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(diagrampackage.SelectedClassdiagramID.Int64)
                  if (_classdiagram) {
                    diagrampackage.SelectedClassdiagram = _classdiagram
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear diagrampackages that are absent from the GET
            this.frontRepo.DiagramPackages.forEach(
              diagrampackage => {
                if (this.frontRepo.DiagramPackages_batch.get(diagrampackage.ID) == undefined) {
                  this.frontRepo.DiagramPackages.delete(diagrampackage.ID)
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

  // FieldPull performs a GET on Field of the stack and redeem association pointers 
  FieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.fieldService.getFields(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            fields,
          ]) => {
            // init the array
            this.frontRepo.Fields_array = fields

            // clear the map that counts Field in the GET
            this.frontRepo.Fields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            fields.forEach(
              field => {
                this.frontRepo.Fields.set(field.ID, field)
                this.frontRepo.Fields_batch.set(field.ID, field)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStructShape.Fields redeeming
                {
                  let _gongstructshape = this.frontRepo.GongStructShapes.get(field.GongStructShape_FieldsDBID.Int64)
                  if (_gongstructshape) {
                    if (_gongstructshape.Fields == undefined) {
                      _gongstructshape.Fields = new Array<FieldDB>()
                    }
                    _gongstructshape.Fields.push(field)
                    if (field.GongStructShape_Fields_reverse == undefined) {
                      field.GongStructShape_Fields_reverse = _gongstructshape
                    }
                  }
                }
              }
            )

            // clear fields that are absent from the GET
            this.frontRepo.Fields.forEach(
              field => {
                if (this.frontRepo.Fields_batch.get(field.ID) == undefined) {
                  this.frontRepo.Fields.delete(field.ID)
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

  // GongEnumShapePull performs a GET on GongEnumShape of the stack and redeem association pointers 
  GongEnumShapePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongenumshapeService.getGongEnumShapes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            gongenumshapes,
          ]) => {
            // init the array
            this.frontRepo.GongEnumShapes_array = gongenumshapes

            // clear the map that counts GongEnumShape in the GET
            this.frontRepo.GongEnumShapes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongenumshapes.forEach(
              gongenumshape => {
                this.frontRepo.GongEnumShapes.set(gongenumshape.ID, gongenumshape)
                this.frontRepo.GongEnumShapes_batch.set(gongenumshape.ID, gongenumshape)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Position redeeming
                {
                  let _position = this.frontRepo.Positions.get(gongenumshape.PositionID.Int64)
                  if (_position) {
                    gongenumshape.Position = _position
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.GongEnumShapes redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(gongenumshape.Classdiagram_GongEnumShapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.GongEnumShapes == undefined) {
                      _classdiagram.GongEnumShapes = new Array<GongEnumShapeDB>()
                    }
                    _classdiagram.GongEnumShapes.push(gongenumshape)
                    if (gongenumshape.Classdiagram_GongEnumShapes_reverse == undefined) {
                      gongenumshape.Classdiagram_GongEnumShapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )

            // clear gongenumshapes that are absent from the GET
            this.frontRepo.GongEnumShapes.forEach(
              gongenumshape => {
                if (this.frontRepo.GongEnumShapes_batch.get(gongenumshape.ID) == undefined) {
                  this.frontRepo.GongEnumShapes.delete(gongenumshape.ID)
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

  // GongEnumValueEntryPull performs a GET on GongEnumValueEntry of the stack and redeem association pointers 
  GongEnumValueEntryPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongenumvalueentryService.getGongEnumValueEntrys(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            gongenumvalueentrys,
          ]) => {
            // init the array
            this.frontRepo.GongEnumValueEntrys_array = gongenumvalueentrys

            // clear the map that counts GongEnumValueEntry in the GET
            this.frontRepo.GongEnumValueEntrys_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongenumvalueentrys.forEach(
              gongenumvalueentry => {
                this.frontRepo.GongEnumValueEntrys.set(gongenumvalueentry.ID, gongenumvalueentry)
                this.frontRepo.GongEnumValueEntrys_batch.set(gongenumvalueentry.ID, gongenumvalueentry)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongEnumShape.GongEnumValueEntrys redeeming
                {
                  let _gongenumshape = this.frontRepo.GongEnumShapes.get(gongenumvalueentry.GongEnumShape_GongEnumValueEntrysDBID.Int64)
                  if (_gongenumshape) {
                    if (_gongenumshape.GongEnumValueEntrys == undefined) {
                      _gongenumshape.GongEnumValueEntrys = new Array<GongEnumValueEntryDB>()
                    }
                    _gongenumshape.GongEnumValueEntrys.push(gongenumvalueentry)
                    if (gongenumvalueentry.GongEnumShape_GongEnumValueEntrys_reverse == undefined) {
                      gongenumvalueentry.GongEnumShape_GongEnumValueEntrys_reverse = _gongenumshape
                    }
                  }
                }
              }
            )

            // clear gongenumvalueentrys that are absent from the GET
            this.frontRepo.GongEnumValueEntrys.forEach(
              gongenumvalueentry => {
                if (this.frontRepo.GongEnumValueEntrys_batch.get(gongenumvalueentry.ID) == undefined) {
                  this.frontRepo.GongEnumValueEntrys.delete(gongenumvalueentry.ID)
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

  // GongStructShapePull performs a GET on GongStructShape of the stack and redeem association pointers 
  GongStructShapePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.gongstructshapeService.getGongStructShapes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            gongstructshapes,
          ]) => {
            // init the array
            this.frontRepo.GongStructShapes_array = gongstructshapes

            // clear the map that counts GongStructShape in the GET
            this.frontRepo.GongStructShapes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gongstructshapes.forEach(
              gongstructshape => {
                this.frontRepo.GongStructShapes.set(gongstructshape.ID, gongstructshape)
                this.frontRepo.GongStructShapes_batch.set(gongstructshape.ID, gongstructshape)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Position redeeming
                {
                  let _position = this.frontRepo.Positions.get(gongstructshape.PositionID.Int64)
                  if (_position) {
                    gongstructshape.Position = _position
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.GongStructShapes redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(gongstructshape.Classdiagram_GongStructShapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.GongStructShapes == undefined) {
                      _classdiagram.GongStructShapes = new Array<GongStructShapeDB>()
                    }
                    _classdiagram.GongStructShapes.push(gongstructshape)
                    if (gongstructshape.Classdiagram_GongStructShapes_reverse == undefined) {
                      gongstructshape.Classdiagram_GongStructShapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )

            // clear gongstructshapes that are absent from the GET
            this.frontRepo.GongStructShapes.forEach(
              gongstructshape => {
                if (this.frontRepo.GongStructShapes_batch.get(gongstructshape.ID) == undefined) {
                  this.frontRepo.GongStructShapes.delete(gongstructshape.ID)
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

  // LinkPull performs a GET on Link of the stack and redeem association pointers 
  LinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linkService.getLinks(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            links,
          ]) => {
            // init the array
            this.frontRepo.Links_array = links

            // clear the map that counts Link in the GET
            this.frontRepo.Links_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            links.forEach(
              link => {
                this.frontRepo.Links.set(link.ID, link)
                this.frontRepo.Links_batch.set(link.ID, link)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Middlevertice redeeming
                {
                  let _vertice = this.frontRepo.Vertices.get(link.MiddleverticeID.Int64)
                  if (_vertice) {
                    link.Middlevertice = _vertice
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field GongStructShape.Links redeeming
                {
                  let _gongstructshape = this.frontRepo.GongStructShapes.get(link.GongStructShape_LinksDBID.Int64)
                  if (_gongstructshape) {
                    if (_gongstructshape.Links == undefined) {
                      _gongstructshape.Links = new Array<LinkDB>()
                    }
                    _gongstructshape.Links.push(link)
                    if (link.GongStructShape_Links_reverse == undefined) {
                      link.GongStructShape_Links_reverse = _gongstructshape
                    }
                  }
                }
              }
            )

            // clear links that are absent from the GET
            this.frontRepo.Links.forEach(
              link => {
                if (this.frontRepo.Links_batch.get(link.ID) == undefined) {
                  this.frontRepo.Links.delete(link.ID)
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

  // NoteShapePull performs a GET on NoteShape of the stack and redeem association pointers 
  NoteShapePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.noteshapeService.getNoteShapes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            noteshapes,
          ]) => {
            // init the array
            this.frontRepo.NoteShapes_array = noteshapes

            // clear the map that counts NoteShape in the GET
            this.frontRepo.NoteShapes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            noteshapes.forEach(
              noteshape => {
                this.frontRepo.NoteShapes.set(noteshape.ID, noteshape)
                this.frontRepo.NoteShapes_batch.set(noteshape.ID, noteshape)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.NoteShapes redeeming
                {
                  let _classdiagram = this.frontRepo.Classdiagrams.get(noteshape.Classdiagram_NoteShapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.NoteShapes == undefined) {
                      _classdiagram.NoteShapes = new Array<NoteShapeDB>()
                    }
                    _classdiagram.NoteShapes.push(noteshape)
                    if (noteshape.Classdiagram_NoteShapes_reverse == undefined) {
                      noteshape.Classdiagram_NoteShapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )

            // clear noteshapes that are absent from the GET
            this.frontRepo.NoteShapes.forEach(
              noteshape => {
                if (this.frontRepo.NoteShapes_batch.get(noteshape.ID) == undefined) {
                  this.frontRepo.NoteShapes.delete(noteshape.ID)
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

  // NoteShapeLinkPull performs a GET on NoteShapeLink of the stack and redeem association pointers 
  NoteShapeLinkPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.noteshapelinkService.getNoteShapeLinks(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            noteshapelinks,
          ]) => {
            // init the array
            this.frontRepo.NoteShapeLinks_array = noteshapelinks

            // clear the map that counts NoteShapeLink in the GET
            this.frontRepo.NoteShapeLinks_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            noteshapelinks.forEach(
              noteshapelink => {
                this.frontRepo.NoteShapeLinks.set(noteshapelink.ID, noteshapelink)
                this.frontRepo.NoteShapeLinks_batch.set(noteshapelink.ID, noteshapelink)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field NoteShape.NoteShapeLinks redeeming
                {
                  let _noteshape = this.frontRepo.NoteShapes.get(noteshapelink.NoteShape_NoteShapeLinksDBID.Int64)
                  if (_noteshape) {
                    if (_noteshape.NoteShapeLinks == undefined) {
                      _noteshape.NoteShapeLinks = new Array<NoteShapeLinkDB>()
                    }
                    _noteshape.NoteShapeLinks.push(noteshapelink)
                    if (noteshapelink.NoteShape_NoteShapeLinks_reverse == undefined) {
                      noteshapelink.NoteShape_NoteShapeLinks_reverse = _noteshape
                    }
                  }
                }
              }
            )

            // clear noteshapelinks that are absent from the GET
            this.frontRepo.NoteShapeLinks.forEach(
              noteshapelink => {
                if (this.frontRepo.NoteShapeLinks_batch.get(noteshapelink.ID) == undefined) {
                  this.frontRepo.NoteShapeLinks.delete(noteshapelink.ID)
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

  // PositionPull performs a GET on Position of the stack and redeem association pointers 
  PositionPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.positionService.getPositions(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            positions,
          ]) => {
            // init the array
            this.frontRepo.Positions_array = positions

            // clear the map that counts Position in the GET
            this.frontRepo.Positions_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            positions.forEach(
              position => {
                this.frontRepo.Positions.set(position.ID, position)
                this.frontRepo.Positions_batch.set(position.ID, position)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear positions that are absent from the GET
            this.frontRepo.Positions.forEach(
              position => {
                if (this.frontRepo.Positions_batch.get(position.ID) == undefined) {
                  this.frontRepo.Positions.delete(position.ID)
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

  // UmlStatePull performs a GET on UmlState of the stack and redeem association pointers 
  UmlStatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.umlstateService.getUmlStates(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            umlstates,
          ]) => {
            // init the array
            this.frontRepo.UmlStates_array = umlstates

            // clear the map that counts UmlState in the GET
            this.frontRepo.UmlStates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            umlstates.forEach(
              umlstate => {
                this.frontRepo.UmlStates.set(umlstate.ID, umlstate)
                this.frontRepo.UmlStates_batch.set(umlstate.ID, umlstate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Umlsc.States redeeming
                {
                  let _umlsc = this.frontRepo.Umlscs.get(umlstate.Umlsc_StatesDBID.Int64)
                  if (_umlsc) {
                    if (_umlsc.States == undefined) {
                      _umlsc.States = new Array<UmlStateDB>()
                    }
                    _umlsc.States.push(umlstate)
                    if (umlstate.Umlsc_States_reverse == undefined) {
                      umlstate.Umlsc_States_reverse = _umlsc
                    }
                  }
                }
              }
            )

            // clear umlstates that are absent from the GET
            this.frontRepo.UmlStates.forEach(
              umlstate => {
                if (this.frontRepo.UmlStates_batch.get(umlstate.ID) == undefined) {
                  this.frontRepo.UmlStates.delete(umlstate.ID)
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

  // UmlscPull performs a GET on Umlsc of the stack and redeem association pointers 
  UmlscPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.umlscService.getUmlscs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            umlscs,
          ]) => {
            // init the array
            this.frontRepo.Umlscs_array = umlscs

            // clear the map that counts Umlsc in the GET
            this.frontRepo.Umlscs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            umlscs.forEach(
              umlsc => {
                this.frontRepo.Umlscs.set(umlsc.ID, umlsc)
                this.frontRepo.Umlscs_batch.set(umlsc.ID, umlsc)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field DiagramPackage.Umlscs redeeming
                {
                  let _diagrampackage = this.frontRepo.DiagramPackages.get(umlsc.DiagramPackage_UmlscsDBID.Int64)
                  if (_diagrampackage) {
                    if (_diagrampackage.Umlscs == undefined) {
                      _diagrampackage.Umlscs = new Array<UmlscDB>()
                    }
                    _diagrampackage.Umlscs.push(umlsc)
                    if (umlsc.DiagramPackage_Umlscs_reverse == undefined) {
                      umlsc.DiagramPackage_Umlscs_reverse = _diagrampackage
                    }
                  }
                }
              }
            )

            // clear umlscs that are absent from the GET
            this.frontRepo.Umlscs.forEach(
              umlsc => {
                if (this.frontRepo.Umlscs_batch.get(umlsc.ID) == undefined) {
                  this.frontRepo.Umlscs.delete(umlsc.ID)
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

  // VerticePull performs a GET on Vertice of the stack and redeem association pointers 
  VerticePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.verticeService.getVertices(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            vertices,
          ]) => {
            // init the array
            this.frontRepo.Vertices_array = vertices

            // clear the map that counts Vertice in the GET
            this.frontRepo.Vertices_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            vertices.forEach(
              vertice => {
                this.frontRepo.Vertices.set(vertice.ID, vertice)
                this.frontRepo.Vertices_batch.set(vertice.ID, vertice)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear vertices that are absent from the GET
            this.frontRepo.Vertices.forEach(
              vertice => {
                if (this.frontRepo.Vertices_batch.get(vertice.ID) == undefined) {
                  this.frontRepo.Vertices.delete(vertice.ID)
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
export function getClassdiagramUniqueID(id: number): number {
  return 31 * id
}
export function getDiagramPackageUniqueID(id: number): number {
  return 37 * id
}
export function getFieldUniqueID(id: number): number {
  return 41 * id
}
export function getGongEnumShapeUniqueID(id: number): number {
  return 43 * id
}
export function getGongEnumValueEntryUniqueID(id: number): number {
  return 47 * id
}
export function getGongStructShapeUniqueID(id: number): number {
  return 53 * id
}
export function getLinkUniqueID(id: number): number {
  return 59 * id
}
export function getNoteShapeUniqueID(id: number): number {
  return 61 * id
}
export function getNoteShapeLinkUniqueID(id: number): number {
  return 67 * id
}
export function getPositionUniqueID(id: number): number {
  return 71 * id
}
export function getUmlStateUniqueID(id: number): number {
  return 73 * id
}
export function getUmlscUniqueID(id: number): number {
  return 79 * id
}
export function getVerticeUniqueID(id: number): number {
  return 83 * id
}
