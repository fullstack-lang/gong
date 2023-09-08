// insertion point sub template for components imports 
  import { DummysTableComponent } from './dummys-table/dummys-table.component'
  import { DummySortingComponent } from './dummy-sorting/dummy-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfDummysComponents: Map<string, any> = new Map([["DummysTableComponent", DummysTableComponent],])
  export const MapOfDummySortingComponents: Map<string, any> = new Map([["DummySortingComponent", DummySortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Dummy", MapOfDummysComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Dummy", MapOfDummySortingComponents],
    ]
  )
