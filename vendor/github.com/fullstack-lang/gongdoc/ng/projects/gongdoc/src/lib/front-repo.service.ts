import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { ClassdiagramDB } from './classdiagram-db'
import { ClassdiagramService } from './classdiagram.service'

import { ClassshapeDB } from './classshape-db'
import { ClassshapeService } from './classshape.service'

import { DiagramPackageDB } from './diagrampackage-db'
import { DiagramPackageService } from './diagrampackage.service'

import { FieldDB } from './field-db'
import { FieldService } from './field.service'

import { LinkDB } from './link-db'
import { LinkService } from './link.service'

import { NodeDB } from './node-db'
import { NodeService } from './node.service'

import { NoteDB } from './note-db'
import { NoteService } from './note.service'

import { PositionDB } from './position-db'
import { PositionService } from './position.service'

import { ReferenceDB } from './reference-db'
import { ReferenceService } from './reference.service'

import { TreeDB } from './tree-db'
import { TreeService } from './tree.service'

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
  Classshapes_array = new Array<ClassshapeDB>(); // array of repo instances
  Classshapes = new Map<number, ClassshapeDB>(); // map of repo instances
  Classshapes_batch = new Map<number, ClassshapeDB>(); // same but only in last GET (for finding repo instances to delete)
  DiagramPackages_array = new Array<DiagramPackageDB>(); // array of repo instances
  DiagramPackages = new Map<number, DiagramPackageDB>(); // map of repo instances
  DiagramPackages_batch = new Map<number, DiagramPackageDB>(); // same but only in last GET (for finding repo instances to delete)
  Fields_array = new Array<FieldDB>(); // array of repo instances
  Fields = new Map<number, FieldDB>(); // map of repo instances
  Fields_batch = new Map<number, FieldDB>(); // same but only in last GET (for finding repo instances to delete)
  Links_array = new Array<LinkDB>(); // array of repo instances
  Links = new Map<number, LinkDB>(); // map of repo instances
  Links_batch = new Map<number, LinkDB>(); // same but only in last GET (for finding repo instances to delete)
  Nodes_array = new Array<NodeDB>(); // array of repo instances
  Nodes = new Map<number, NodeDB>(); // map of repo instances
  Nodes_batch = new Map<number, NodeDB>(); // same but only in last GET (for finding repo instances to delete)
  Notes_array = new Array<NoteDB>(); // array of repo instances
  Notes = new Map<number, NoteDB>(); // map of repo instances
  Notes_batch = new Map<number, NoteDB>(); // same but only in last GET (for finding repo instances to delete)
  Positions_array = new Array<PositionDB>(); // array of repo instances
  Positions = new Map<number, PositionDB>(); // map of repo instances
  Positions_batch = new Map<number, PositionDB>(); // same but only in last GET (for finding repo instances to delete)
  References_array = new Array<ReferenceDB>(); // array of repo instances
  References = new Map<number, ReferenceDB>(); // map of repo instances
  References_batch = new Map<number, ReferenceDB>(); // same but only in last GET (for finding repo instances to delete)
  Trees_array = new Array<TreeDB>(); // array of repo instances
  Trees = new Map<number, TreeDB>(); // map of repo instances
  Trees_batch = new Map<number, TreeDB>(); // same but only in last GET (for finding repo instances to delete)
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
    private classdiagramService: ClassdiagramService,
    private classshapeService: ClassshapeService,
    private diagrampackageService: DiagramPackageService,
    private fieldService: FieldService,
    private linkService: LinkService,
    private nodeService: NodeService,
    private noteService: NoteService,
    private positionService: PositionService,
    private referenceService: ReferenceService,
    private treeService: TreeService,
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
    Observable<ClassshapeDB[]>,
    Observable<DiagramPackageDB[]>,
    Observable<FieldDB[]>,
    Observable<LinkDB[]>,
    Observable<NodeDB[]>,
    Observable<NoteDB[]>,
    Observable<PositionDB[]>,
    Observable<ReferenceDB[]>,
    Observable<TreeDB[]>,
    Observable<UmlStateDB[]>,
    Observable<UmlscDB[]>,
    Observable<VerticeDB[]>,
  ] = [ // insertion point sub template 
      this.classdiagramService.getClassdiagrams(),
      this.classshapeService.getClassshapes(),
      this.diagrampackageService.getDiagramPackages(),
      this.fieldService.getFields(),
      this.linkService.getLinks(),
      this.nodeService.getNodes(),
      this.noteService.getNotes(),
      this.positionService.getPositions(),
      this.referenceService.getReferences(),
      this.treeService.getTrees(),
      this.umlstateService.getUmlStates(),
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
            diagrampackages_,
            fields_,
            links_,
            nodes_,
            notes_,
            positions_,
            references_,
            trees_,
            umlstates_,
            umlscs_,
            vertices_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var classdiagrams: ClassdiagramDB[]
            classdiagrams = classdiagrams_ as ClassdiagramDB[]
            var classshapes: ClassshapeDB[]
            classshapes = classshapes_ as ClassshapeDB[]
            var diagrampackages: DiagramPackageDB[]
            diagrampackages = diagrampackages_ as DiagramPackageDB[]
            var fields: FieldDB[]
            fields = fields_ as FieldDB[]
            var links: LinkDB[]
            links = links_ as LinkDB[]
            var nodes: NodeDB[]
            nodes = nodes_ as NodeDB[]
            var notes: NoteDB[]
            notes = notes_ as NoteDB[]
            var positions: PositionDB[]
            positions = positions_ as PositionDB[]
            var references: ReferenceDB[]
            references = references_ as ReferenceDB[]
            var trees: TreeDB[]
            trees = trees_ as TreeDB[]
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

            // sort Classdiagrams_array array
            FrontRepoSingloton.Classdiagrams_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

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

            // sort Classshapes_array array
            FrontRepoSingloton.Classshapes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.DiagramPackages_array = diagrampackages

            // clear the map that counts DiagramPackage in the GET
            FrontRepoSingloton.DiagramPackages_batch.clear()

            diagrampackages.forEach(
              diagrampackage => {
                FrontRepoSingloton.DiagramPackages.set(diagrampackage.ID, diagrampackage)
                FrontRepoSingloton.DiagramPackages_batch.set(diagrampackage.ID, diagrampackage)
              }
            )

            // clear diagrampackages that are absent from the batch
            FrontRepoSingloton.DiagramPackages.forEach(
              diagrampackage => {
                if (FrontRepoSingloton.DiagramPackages_batch.get(diagrampackage.ID) == undefined) {
                  FrontRepoSingloton.DiagramPackages.delete(diagrampackage.ID)
                }
              }
            )

            // sort DiagramPackages_array array
            FrontRepoSingloton.DiagramPackages_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

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

            // sort Fields_array array
            FrontRepoSingloton.Fields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

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

            // sort Links_array array
            FrontRepoSingloton.Links_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Nodes_array = nodes

            // clear the map that counts Node in the GET
            FrontRepoSingloton.Nodes_batch.clear()

            nodes.forEach(
              node => {
                FrontRepoSingloton.Nodes.set(node.ID, node)
                FrontRepoSingloton.Nodes_batch.set(node.ID, node)
              }
            )

            // clear nodes that are absent from the batch
            FrontRepoSingloton.Nodes.forEach(
              node => {
                if (FrontRepoSingloton.Nodes_batch.get(node.ID) == undefined) {
                  FrontRepoSingloton.Nodes.delete(node.ID)
                }
              }
            )

            // sort Nodes_array array
            FrontRepoSingloton.Nodes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Notes_array = notes

            // clear the map that counts Note in the GET
            FrontRepoSingloton.Notes_batch.clear()

            notes.forEach(
              note => {
                FrontRepoSingloton.Notes.set(note.ID, note)
                FrontRepoSingloton.Notes_batch.set(note.ID, note)
              }
            )

            // clear notes that are absent from the batch
            FrontRepoSingloton.Notes.forEach(
              note => {
                if (FrontRepoSingloton.Notes_batch.get(note.ID) == undefined) {
                  FrontRepoSingloton.Notes.delete(note.ID)
                }
              }
            )

            // sort Notes_array array
            FrontRepoSingloton.Notes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

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

            // sort Positions_array array
            FrontRepoSingloton.Positions_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.References_array = references

            // clear the map that counts Reference in the GET
            FrontRepoSingloton.References_batch.clear()

            references.forEach(
              reference => {
                FrontRepoSingloton.References.set(reference.ID, reference)
                FrontRepoSingloton.References_batch.set(reference.ID, reference)
              }
            )

            // clear references that are absent from the batch
            FrontRepoSingloton.References.forEach(
              reference => {
                if (FrontRepoSingloton.References_batch.get(reference.ID) == undefined) {
                  FrontRepoSingloton.References.delete(reference.ID)
                }
              }
            )

            // sort References_array array
            FrontRepoSingloton.References_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Trees_array = trees

            // clear the map that counts Tree in the GET
            FrontRepoSingloton.Trees_batch.clear()

            trees.forEach(
              tree => {
                FrontRepoSingloton.Trees.set(tree.ID, tree)
                FrontRepoSingloton.Trees_batch.set(tree.ID, tree)
              }
            )

            // clear trees that are absent from the batch
            FrontRepoSingloton.Trees.forEach(
              tree => {
                if (FrontRepoSingloton.Trees_batch.get(tree.ID) == undefined) {
                  FrontRepoSingloton.Trees.delete(tree.ID)
                }
              }
            )

            // sort Trees_array array
            FrontRepoSingloton.Trees_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.UmlStates_array = umlstates

            // clear the map that counts UmlState in the GET
            FrontRepoSingloton.UmlStates_batch.clear()

            umlstates.forEach(
              umlstate => {
                FrontRepoSingloton.UmlStates.set(umlstate.ID, umlstate)
                FrontRepoSingloton.UmlStates_batch.set(umlstate.ID, umlstate)
              }
            )

            // clear umlstates that are absent from the batch
            FrontRepoSingloton.UmlStates.forEach(
              umlstate => {
                if (FrontRepoSingloton.UmlStates_batch.get(umlstate.ID) == undefined) {
                  FrontRepoSingloton.UmlStates.delete(umlstate.ID)
                }
              }
            )

            // sort UmlStates_array array
            FrontRepoSingloton.UmlStates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

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

            // sort Umlscs_array array
            FrontRepoSingloton.Umlscs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

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

            // sort Vertices_array array
            FrontRepoSingloton.Vertices_array.sort((t1, t2) => {
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
                  let _diagrampackage = FrontRepoSingloton.DiagramPackages.get(classdiagram.DiagramPackage_ClassdiagramsDBID.Int64)
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
                // insertion point for pointer field Reference redeeming
                {
                  let _reference = FrontRepoSingloton.References.get(classshape.ReferenceID.Int64)
                  if (_reference) {
                    classshape.Reference = _reference
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.Classshapes redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(classshape.Classdiagram_ClassshapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.Classshapes == undefined) {
                      _classdiagram.Classshapes = new Array<ClassshapeDB>()
                    }
                    _classdiagram.Classshapes.push(classshape)
                    if (classshape.Classdiagram_Classshapes_reverse == undefined) {
                      classshape.Classdiagram_Classshapes_reverse = _classdiagram
                    }
                  }
                }
              }
            )
            diagrampackages.forEach(
              diagrampackage => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
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
                      _classshape.Fields = new Array<FieldDB>()
                    }
                    _classshape.Fields.push(field)
                    if (field.Classshape_Fields_reverse == undefined) {
                      field.Classshape_Fields_reverse = _classshape
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
                      _classshape.Links = new Array<LinkDB>()
                    }
                    _classshape.Links.push(link)
                    if (link.Classshape_Links_reverse == undefined) {
                      link.Classshape_Links_reverse = _classshape
                    }
                  }
                }
              }
            )
            nodes.forEach(
              node => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Classdiagram redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(node.ClassdiagramID.Int64)
                  if (_classdiagram) {
                    node.Classdiagram = _classdiagram
                  }
                }
                // insertion point for pointer field Umlsc redeeming
                {
                  let _umlsc = FrontRepoSingloton.Umlscs.get(node.UmlscID.Int64)
                  if (_umlsc) {
                    node.Umlsc = _umlsc
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Node.Children redeeming
                {
                  let _node = FrontRepoSingloton.Nodes.get(node.Node_ChildrenDBID.Int64)
                  if (_node) {
                    if (_node.Children == undefined) {
                      _node.Children = new Array<NodeDB>()
                    }
                    _node.Children.push(node)
                    if (node.Node_Children_reverse == undefined) {
                      node.Node_Children_reverse = _node
                    }
                  }
                }
                // insertion point for slice of pointer field Tree.RootNodes redeeming
                {
                  let _tree = FrontRepoSingloton.Trees.get(node.Tree_RootNodesDBID.Int64)
                  if (_tree) {
                    if (_tree.RootNodes == undefined) {
                      _tree.RootNodes = new Array<NodeDB>()
                    }
                    _tree.RootNodes.push(node)
                    if (node.Tree_RootNodes_reverse == undefined) {
                      node.Tree_RootNodes_reverse = _tree
                    }
                  }
                }
              }
            )
            notes.forEach(
              note => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.Notes redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(note.Classdiagram_NotesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.Notes == undefined) {
                      _classdiagram.Notes = new Array<NoteDB>()
                    }
                    _classdiagram.Notes.push(note)
                    if (note.Classdiagram_Notes_reverse == undefined) {
                      note.Classdiagram_Notes_reverse = _classdiagram
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
            references.forEach(
              reference => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            trees.forEach(
              tree => {
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
                  let _umlsc = FrontRepoSingloton.Umlscs.get(umlstate.Umlsc_StatesDBID.Int64)
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
                  let _diagrampackage = FrontRepoSingloton.DiagramPackages.get(umlsc.DiagramPackage_UmlscsDBID.Int64)
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
                // insertion point for slice of pointer field DiagramPackage.Classdiagrams redeeming
                {
                  let _diagrampackage = FrontRepoSingloton.DiagramPackages.get(classdiagram.DiagramPackage_ClassdiagramsDBID.Int64)
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
                // insertion point for pointer field Reference redeeming
                {
                  let _reference = FrontRepoSingloton.References.get(classshape.ReferenceID.Int64)
                  if (_reference) {
                    classshape.Reference = _reference
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.Classshapes redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(classshape.Classdiagram_ClassshapesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.Classshapes == undefined) {
                      _classdiagram.Classshapes = new Array<ClassshapeDB>()
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

  // DiagramPackagePull performs a GET on DiagramPackage of the stack and redeem association pointers 
  DiagramPackagePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.diagrampackageService.getDiagramPackages()
        ]).subscribe(
          ([ // insertion point sub template 
            diagrampackages,
          ]) => {
            // init the array
            FrontRepoSingloton.DiagramPackages_array = diagrampackages

            // clear the map that counts DiagramPackage in the GET
            FrontRepoSingloton.DiagramPackages_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            diagrampackages.forEach(
              diagrampackage => {
                FrontRepoSingloton.DiagramPackages.set(diagrampackage.ID, diagrampackage)
                FrontRepoSingloton.DiagramPackages_batch.set(diagrampackage.ID, diagrampackage)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear diagrampackages that are absent from the GET
            FrontRepoSingloton.DiagramPackages.forEach(
              diagrampackage => {
                if (FrontRepoSingloton.DiagramPackages_batch.get(diagrampackage.ID) == undefined) {
                  FrontRepoSingloton.DiagramPackages.delete(diagrampackage.ID)
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
                      _classshape.Fields = new Array<FieldDB>()
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
                      _classshape.Links = new Array<LinkDB>()
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

  // NodePull performs a GET on Node of the stack and redeem association pointers 
  NodePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.nodeService.getNodes()
        ]).subscribe(
          ([ // insertion point sub template 
            nodes,
          ]) => {
            // init the array
            FrontRepoSingloton.Nodes_array = nodes

            // clear the map that counts Node in the GET
            FrontRepoSingloton.Nodes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            nodes.forEach(
              node => {
                FrontRepoSingloton.Nodes.set(node.ID, node)
                FrontRepoSingloton.Nodes_batch.set(node.ID, node)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Classdiagram redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(node.ClassdiagramID.Int64)
                  if (_classdiagram) {
                    node.Classdiagram = _classdiagram
                  }
                }
                // insertion point for pointer field Umlsc redeeming
                {
                  let _umlsc = FrontRepoSingloton.Umlscs.get(node.UmlscID.Int64)
                  if (_umlsc) {
                    node.Umlsc = _umlsc
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Node.Children redeeming
                {
                  let _node = FrontRepoSingloton.Nodes.get(node.Node_ChildrenDBID.Int64)
                  if (_node) {
                    if (_node.Children == undefined) {
                      _node.Children = new Array<NodeDB>()
                    }
                    _node.Children.push(node)
                    if (node.Node_Children_reverse == undefined) {
                      node.Node_Children_reverse = _node
                    }
                  }
                }
                // insertion point for slice of pointer field Tree.RootNodes redeeming
                {
                  let _tree = FrontRepoSingloton.Trees.get(node.Tree_RootNodesDBID.Int64)
                  if (_tree) {
                    if (_tree.RootNodes == undefined) {
                      _tree.RootNodes = new Array<NodeDB>()
                    }
                    _tree.RootNodes.push(node)
                    if (node.Tree_RootNodes_reverse == undefined) {
                      node.Tree_RootNodes_reverse = _tree
                    }
                  }
                }
              }
            )

            // clear nodes that are absent from the GET
            FrontRepoSingloton.Nodes.forEach(
              node => {
                if (FrontRepoSingloton.Nodes_batch.get(node.ID) == undefined) {
                  FrontRepoSingloton.Nodes.delete(node.ID)
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

  // NotePull performs a GET on Note of the stack and redeem association pointers 
  NotePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.noteService.getNotes()
        ]).subscribe(
          ([ // insertion point sub template 
            notes,
          ]) => {
            // init the array
            FrontRepoSingloton.Notes_array = notes

            // clear the map that counts Note in the GET
            FrontRepoSingloton.Notes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            notes.forEach(
              note => {
                FrontRepoSingloton.Notes.set(note.ID, note)
                FrontRepoSingloton.Notes_batch.set(note.ID, note)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Classdiagram.Notes redeeming
                {
                  let _classdiagram = FrontRepoSingloton.Classdiagrams.get(note.Classdiagram_NotesDBID.Int64)
                  if (_classdiagram) {
                    if (_classdiagram.Notes == undefined) {
                      _classdiagram.Notes = new Array<NoteDB>()
                    }
                    _classdiagram.Notes.push(note)
                    if (note.Classdiagram_Notes_reverse == undefined) {
                      note.Classdiagram_Notes_reverse = _classdiagram
                    }
                  }
                }
              }
            )

            // clear notes that are absent from the GET
            FrontRepoSingloton.Notes.forEach(
              note => {
                if (FrontRepoSingloton.Notes_batch.get(note.ID) == undefined) {
                  FrontRepoSingloton.Notes.delete(note.ID)
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

  // ReferencePull performs a GET on Reference of the stack and redeem association pointers 
  ReferencePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.referenceService.getReferences()
        ]).subscribe(
          ([ // insertion point sub template 
            references,
          ]) => {
            // init the array
            FrontRepoSingloton.References_array = references

            // clear the map that counts Reference in the GET
            FrontRepoSingloton.References_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            references.forEach(
              reference => {
                FrontRepoSingloton.References.set(reference.ID, reference)
                FrontRepoSingloton.References_batch.set(reference.ID, reference)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear references that are absent from the GET
            FrontRepoSingloton.References.forEach(
              reference => {
                if (FrontRepoSingloton.References_batch.get(reference.ID) == undefined) {
                  FrontRepoSingloton.References.delete(reference.ID)
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

  // TreePull performs a GET on Tree of the stack and redeem association pointers 
  TreePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.treeService.getTrees()
        ]).subscribe(
          ([ // insertion point sub template 
            trees,
          ]) => {
            // init the array
            FrontRepoSingloton.Trees_array = trees

            // clear the map that counts Tree in the GET
            FrontRepoSingloton.Trees_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            trees.forEach(
              tree => {
                FrontRepoSingloton.Trees.set(tree.ID, tree)
                FrontRepoSingloton.Trees_batch.set(tree.ID, tree)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear trees that are absent from the GET
            FrontRepoSingloton.Trees.forEach(
              tree => {
                if (FrontRepoSingloton.Trees_batch.get(tree.ID) == undefined) {
                  FrontRepoSingloton.Trees.delete(tree.ID)
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

  // UmlStatePull performs a GET on UmlState of the stack and redeem association pointers 
  UmlStatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.umlstateService.getUmlStates()
        ]).subscribe(
          ([ // insertion point sub template 
            umlstates,
          ]) => {
            // init the array
            FrontRepoSingloton.UmlStates_array = umlstates

            // clear the map that counts UmlState in the GET
            FrontRepoSingloton.UmlStates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            umlstates.forEach(
              umlstate => {
                FrontRepoSingloton.UmlStates.set(umlstate.ID, umlstate)
                FrontRepoSingloton.UmlStates_batch.set(umlstate.ID, umlstate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Umlsc.States redeeming
                {
                  let _umlsc = FrontRepoSingloton.Umlscs.get(umlstate.Umlsc_StatesDBID.Int64)
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
            FrontRepoSingloton.UmlStates.forEach(
              umlstate => {
                if (FrontRepoSingloton.UmlStates_batch.get(umlstate.ID) == undefined) {
                  FrontRepoSingloton.UmlStates.delete(umlstate.ID)
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
                // insertion point for slice of pointer field DiagramPackage.Umlscs redeeming
                {
                  let _diagrampackage = FrontRepoSingloton.DiagramPackages.get(umlsc.DiagramPackage_UmlscsDBID.Int64)
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
export function getDiagramPackageUniqueID(id: number): number {
  return 41 * id
}
export function getFieldUniqueID(id: number): number {
  return 43 * id
}
export function getLinkUniqueID(id: number): number {
  return 47 * id
}
export function getNodeUniqueID(id: number): number {
  return 53 * id
}
export function getNoteUniqueID(id: number): number {
  return 59 * id
}
export function getPositionUniqueID(id: number): number {
  return 61 * id
}
export function getReferenceUniqueID(id: number): number {
  return 67 * id
}
export function getTreeUniqueID(id: number): number {
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
