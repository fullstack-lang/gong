// insertion point sub template for components imports 
  import { AstructsTableComponent } from './astructs-table/astructs-table.component'
  import { AstructSortingComponent } from './astruct-sorting/astruct-sorting.component'
  import { AstructBstruct2UsesTableComponent } from './astructbstruct2uses-table/astructbstruct2uses-table.component'
  import { AstructBstruct2UseSortingComponent } from './astructbstruct2use-sorting/astructbstruct2use-sorting.component'
  import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
  import { AstructBstructUseSortingComponent } from './astructbstructuse-sorting/astructbstructuse-sorting.component'
  import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
  import { BstructSortingComponent } from './bstruct-sorting/bstruct-sorting.component'
  import { DstructsTableComponent } from './dstructs-table/dstructs-table.component'
  import { DstructSortingComponent } from './dstruct-sorting/dstruct-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAstructsComponents: Map<string, any> = new Map([["AstructsTableComponent", AstructsTableComponent],])
  export const MapOfAstructSortingComponents: Map<string, any> = new Map([["AstructSortingComponent", AstructSortingComponent],])
  export const MapOfAstructBstruct2UsesComponents: Map<string, any> = new Map([["AstructBstruct2UsesTableComponent", AstructBstruct2UsesTableComponent],])
  export const MapOfAstructBstruct2UseSortingComponents: Map<string, any> = new Map([["AstructBstruct2UseSortingComponent", AstructBstruct2UseSortingComponent],])
  export const MapOfAstructBstructUsesComponents: Map<string, any> = new Map([["AstructBstructUsesTableComponent", AstructBstructUsesTableComponent],])
  export const MapOfAstructBstructUseSortingComponents: Map<string, any> = new Map([["AstructBstructUseSortingComponent", AstructBstructUseSortingComponent],])
  export const MapOfBstructsComponents: Map<string, any> = new Map([["BstructsTableComponent", BstructsTableComponent],])
  export const MapOfBstructSortingComponents: Map<string, any> = new Map([["BstructSortingComponent", BstructSortingComponent],])
  export const MapOfDstructsComponents: Map<string, any> = new Map([["DstructsTableComponent", DstructsTableComponent],])
  export const MapOfDstructSortingComponents: Map<string, any> = new Map([["DstructSortingComponent", DstructSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Astruct", MapOfAstructsComponents],
      ["AstructBstruct2Use", MapOfAstructBstruct2UsesComponents],
      ["AstructBstructUse", MapOfAstructBstructUsesComponents],
      ["Bstruct", MapOfBstructsComponents],
      ["Dstruct", MapOfDstructsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Astruct", MapOfAstructSortingComponents],
      ["AstructBstruct2Use", MapOfAstructBstruct2UseSortingComponents],
      ["AstructBstructUse", MapOfAstructBstructUseSortingComponents],
      ["Bstruct", MapOfBstructSortingComponents],
      ["Dstruct", MapOfDstructSortingComponents],
    ]
  )
