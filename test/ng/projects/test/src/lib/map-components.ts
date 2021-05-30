// insertion point sub template for components imports 
  import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
  import { AclassSortingComponent } from './aclass-sorting/aclass-sorting.component'
  import { BclasssTableComponent } from './bclasss-table/bclasss-table.component'
  import { BclassSortingComponent } from './bclass-sorting/bclass-sorting.component'
  import { DclasssTableComponent } from './dclasss-table/dclasss-table.component'
  import { DclassSortingComponent } from './dclass-sorting/dclass-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAclasssComponents: Map<string, any> = new Map([["AclasssTableComponent", AclasssTableComponent],])
  export const MapOfAclassSortingComponents: Map<string, any> = new Map([["AclassSortingComponent", AclassSortingComponent],])
  export const MapOfBclasssComponents: Map<string, any> = new Map([["BclasssTableComponent", BclasssTableComponent],])
  export const MapOfBclassSortingComponents: Map<string, any> = new Map([["BclassSortingComponent", BclassSortingComponent],])
  export const MapOfDclasssComponents: Map<string, any> = new Map([["DclasssTableComponent", DclasssTableComponent],])
  export const MapOfDclassSortingComponents: Map<string, any> = new Map([["DclassSortingComponent", DclassSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Aclass", MapOfAclasssComponents],
      ["Bclass", MapOfBclasssComponents],
      ["Dclass", MapOfDclasssComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Aclass", MapOfAclassSortingComponents],
      ["Bclass", MapOfBclassSortingComponents],
      ["Dclass", MapOfDclassSortingComponents],
    ]
  )
