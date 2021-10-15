// insertion point sub template for components imports 
  import { AstructsTableComponent } from './astructs-table/astructs-table.component'
  import { AstructSortingComponent } from './astruct-sorting/astruct-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAstructsComponents: Map<string, any> = new Map([["AstructsTableComponent", AstructsTableComponent],])
  export const MapOfAstructSortingComponents: Map<string, any> = new Map([["AstructSortingComponent", AstructSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Astruct", MapOfAstructsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Astruct", MapOfAstructSortingComponents],
    ]
  )
