// insertion point sub template for components imports 
  import { OutletsTableComponent } from './outlets-table/outlets-table.component'
  import { OutletSortingComponent } from './outlet-sorting/outlet-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfOutletsComponents: Map<string, any> = new Map([["OutletsTableComponent", OutletsTableComponent],])
  export const MapOfOutletSortingComponents: Map<string, any> = new Map([["OutletSortingComponent", OutletSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Outlet", MapOfOutletsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Outlet", MapOfOutletSortingComponents],
    ]
  )
