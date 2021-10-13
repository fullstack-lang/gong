// insertion point sub template for components imports 
  import { AstructsTableComponent } from './astructs-table/astructs-table.component'
  import { AstructSortingComponent } from './astruct-sorting/astruct-sorting.component'
  import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
  import { AstructBstructUseSortingComponent } from './astructbstructuse-sorting/astructbstructuse-sorting.component'
  import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
  import { BstructSortingComponent } from './bstruct-sorting/bstruct-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAstructsComponents: Map<string, any> = new Map([["AstructsTableComponent", AstructsTableComponent],])
  export const MapOfAstructSortingComponents: Map<string, any> = new Map([["AstructSortingComponent", AstructSortingComponent],])
  export const MapOfAstructBstructUsesComponents: Map<string, any> = new Map([["AstructBstructUsesTableComponent", AstructBstructUsesTableComponent],])
  export const MapOfAstructBstructUseSortingComponents: Map<string, any> = new Map([["AstructBstructUseSortingComponent", AstructBstructUseSortingComponent],])
  export const MapOfBstructsComponents: Map<string, any> = new Map([["BstructsTableComponent", BstructsTableComponent],])
  export const MapOfBstructSortingComponents: Map<string, any> = new Map([["BstructSortingComponent", BstructSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Astruct", MapOfAstructsComponents],
      ["AstructBstructUse", MapOfAstructBstructUsesComponents],
      ["Bstruct", MapOfBstructsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Astruct", MapOfAstructSortingComponents],
      ["AstructBstructUse", MapOfAstructBstructUseSortingComponents],
      ["Bstruct", MapOfBstructSortingComponents],
    ]
  )
