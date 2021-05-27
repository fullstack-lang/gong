// insertion point sub template for components imports 
  import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
  import { AclassSortingComponent } from './aclass-sorting/aclass-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAclasssComponents: Map<string, any> = new Map([["AclasssTableComponent", AclasssTableComponent],])
  export const MapOfAclassSortingComponents: Map<string, any> = new Map([["AclassSortingComponent", AclassSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Aclass", MapOfAclasssComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Aclass", MapOfAclassSortingComponents],
    ]
  )
