// insertion point sub template for components imports 
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { BclasssTableComponent } from './bclasss-table/bclasss-table.component'

// insertion point sub template for map of components per struct 
export const MapOfAclasssComponents: Map<string, any> = new Map([["AclasssTableComponent", AclasssTableComponent],])
export const MapOfBclasssComponents: Map<string, any> = new Map([["BclasssTableComponent", BclasssTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Aclass", MapOfAclasssComponents],
      ["Bclass", MapOfBclasssComponents],
    ]
  )