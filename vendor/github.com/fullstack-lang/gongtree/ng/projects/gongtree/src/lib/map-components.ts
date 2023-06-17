// insertion point sub template for components imports 
  import { ButtonsTableComponent } from './buttons-table/buttons-table.component'
  import { ButtonSortingComponent } from './button-sorting/button-sorting.component'
  import { NodesTableComponent } from './nodes-table/nodes-table.component'
  import { NodeSortingComponent } from './node-sorting/node-sorting.component'
  import { TreesTableComponent } from './trees-table/trees-table.component'
  import { TreeSortingComponent } from './tree-sorting/tree-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfButtonsComponents: Map<string, any> = new Map([["ButtonsTableComponent", ButtonsTableComponent],])
  export const MapOfButtonSortingComponents: Map<string, any> = new Map([["ButtonSortingComponent", ButtonSortingComponent],])
  export const MapOfNodesComponents: Map<string, any> = new Map([["NodesTableComponent", NodesTableComponent],])
  export const MapOfNodeSortingComponents: Map<string, any> = new Map([["NodeSortingComponent", NodeSortingComponent],])
  export const MapOfTreesComponents: Map<string, any> = new Map([["TreesTableComponent", TreesTableComponent],])
  export const MapOfTreeSortingComponents: Map<string, any> = new Map([["TreeSortingComponent", TreeSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Button", MapOfButtonsComponents],
      ["Node", MapOfNodesComponents],
      ["Tree", MapOfTreesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Button", MapOfButtonSortingComponents],
      ["Node", MapOfNodeSortingComponents],
      ["Tree", MapOfTreeSortingComponents],
    ]
  )
