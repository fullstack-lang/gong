// insertion point sub template for components imports 
  import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
  import { ClassdiagramSortingComponent } from './classdiagram-sorting/classdiagram-sorting.component'
  import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
  import { ClassshapeSortingComponent } from './classshape-sorting/classshape-sorting.component'
  import { DiagramPackagesTableComponent } from './diagrampackages-table/diagrampackages-table.component'
  import { DiagramPackageSortingComponent } from './diagrampackage-sorting/diagrampackage-sorting.component'
  import { FieldsTableComponent } from './fields-table/fields-table.component'
  import { FieldSortingComponent } from './field-sorting/field-sorting.component'
  import { LinksTableComponent } from './links-table/links-table.component'
  import { LinkSortingComponent } from './link-sorting/link-sorting.component'
  import { NodesTableComponent } from './nodes-table/nodes-table.component'
  import { NodeSortingComponent } from './node-sorting/node-sorting.component'
  import { NotesTableComponent } from './notes-table/notes-table.component'
  import { NoteSortingComponent } from './note-sorting/note-sorting.component'
  import { PositionsTableComponent } from './positions-table/positions-table.component'
  import { PositionSortingComponent } from './position-sorting/position-sorting.component'
  import { ReferencesTableComponent } from './references-table/references-table.component'
  import { ReferenceSortingComponent } from './reference-sorting/reference-sorting.component'
  import { TreesTableComponent } from './trees-table/trees-table.component'
  import { TreeSortingComponent } from './tree-sorting/tree-sorting.component'
  import { UmlStatesTableComponent } from './umlstates-table/umlstates-table.component'
  import { UmlStateSortingComponent } from './umlstate-sorting/umlstate-sorting.component'
  import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
  import { UmlscSortingComponent } from './umlsc-sorting/umlsc-sorting.component'
  import { VerticesTableComponent } from './vertices-table/vertices-table.component'
  import { VerticeSortingComponent } from './vertice-sorting/vertice-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfClassdiagramsComponents: Map<string, any> = new Map([["ClassdiagramsTableComponent", ClassdiagramsTableComponent],])
  export const MapOfClassdiagramSortingComponents: Map<string, any> = new Map([["ClassdiagramSortingComponent", ClassdiagramSortingComponent],])
  export const MapOfClassshapesComponents: Map<string, any> = new Map([["ClassshapesTableComponent", ClassshapesTableComponent],])
  export const MapOfClassshapeSortingComponents: Map<string, any> = new Map([["ClassshapeSortingComponent", ClassshapeSortingComponent],])
  export const MapOfDiagramPackagesComponents: Map<string, any> = new Map([["DiagramPackagesTableComponent", DiagramPackagesTableComponent],])
  export const MapOfDiagramPackageSortingComponents: Map<string, any> = new Map([["DiagramPackageSortingComponent", DiagramPackageSortingComponent],])
  export const MapOfFieldsComponents: Map<string, any> = new Map([["FieldsTableComponent", FieldsTableComponent],])
  export const MapOfFieldSortingComponents: Map<string, any> = new Map([["FieldSortingComponent", FieldSortingComponent],])
  export const MapOfLinksComponents: Map<string, any> = new Map([["LinksTableComponent", LinksTableComponent],])
  export const MapOfLinkSortingComponents: Map<string, any> = new Map([["LinkSortingComponent", LinkSortingComponent],])
  export const MapOfNodesComponents: Map<string, any> = new Map([["NodesTableComponent", NodesTableComponent],])
  export const MapOfNodeSortingComponents: Map<string, any> = new Map([["NodeSortingComponent", NodeSortingComponent],])
  export const MapOfNotesComponents: Map<string, any> = new Map([["NotesTableComponent", NotesTableComponent],])
  export const MapOfNoteSortingComponents: Map<string, any> = new Map([["NoteSortingComponent", NoteSortingComponent],])
  export const MapOfPositionsComponents: Map<string, any> = new Map([["PositionsTableComponent", PositionsTableComponent],])
  export const MapOfPositionSortingComponents: Map<string, any> = new Map([["PositionSortingComponent", PositionSortingComponent],])
  export const MapOfReferencesComponents: Map<string, any> = new Map([["ReferencesTableComponent", ReferencesTableComponent],])
  export const MapOfReferenceSortingComponents: Map<string, any> = new Map([["ReferenceSortingComponent", ReferenceSortingComponent],])
  export const MapOfTreesComponents: Map<string, any> = new Map([["TreesTableComponent", TreesTableComponent],])
  export const MapOfTreeSortingComponents: Map<string, any> = new Map([["TreeSortingComponent", TreeSortingComponent],])
  export const MapOfUmlStatesComponents: Map<string, any> = new Map([["UmlStatesTableComponent", UmlStatesTableComponent],])
  export const MapOfUmlStateSortingComponents: Map<string, any> = new Map([["UmlStateSortingComponent", UmlStateSortingComponent],])
  export const MapOfUmlscsComponents: Map<string, any> = new Map([["UmlscsTableComponent", UmlscsTableComponent],])
  export const MapOfUmlscSortingComponents: Map<string, any> = new Map([["UmlscSortingComponent", UmlscSortingComponent],])
  export const MapOfVerticesComponents: Map<string, any> = new Map([["VerticesTableComponent", VerticesTableComponent],])
  export const MapOfVerticeSortingComponents: Map<string, any> = new Map([["VerticeSortingComponent", VerticeSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Classdiagram", MapOfClassdiagramsComponents],
      ["Classshape", MapOfClassshapesComponents],
      ["DiagramPackage", MapOfDiagramPackagesComponents],
      ["Field", MapOfFieldsComponents],
      ["Link", MapOfLinksComponents],
      ["Node", MapOfNodesComponents],
      ["Note", MapOfNotesComponents],
      ["Position", MapOfPositionsComponents],
      ["Reference", MapOfReferencesComponents],
      ["Tree", MapOfTreesComponents],
      ["UmlState", MapOfUmlStatesComponents],
      ["Umlsc", MapOfUmlscsComponents],
      ["Vertice", MapOfVerticesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Classdiagram", MapOfClassdiagramSortingComponents],
      ["Classshape", MapOfClassshapeSortingComponents],
      ["DiagramPackage", MapOfDiagramPackageSortingComponents],
      ["Field", MapOfFieldSortingComponents],
      ["Link", MapOfLinkSortingComponents],
      ["Node", MapOfNodeSortingComponents],
      ["Note", MapOfNoteSortingComponents],
      ["Position", MapOfPositionSortingComponents],
      ["Reference", MapOfReferenceSortingComponents],
      ["Tree", MapOfTreeSortingComponents],
      ["UmlState", MapOfUmlStateSortingComponents],
      ["Umlsc", MapOfUmlscSortingComponents],
      ["Vertice", MapOfVerticeSortingComponents],
    ]
  )
