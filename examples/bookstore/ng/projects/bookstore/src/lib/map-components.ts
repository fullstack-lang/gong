// insertion point sub template for components imports 
import { AreasTableComponent } from './areas-table/areas-table.component'
import { BooksTableComponent } from './books-table/books-table.component'
import { EditorsTableComponent } from './editors-table/editors-table.component'

// insertion point sub template for map of components per struct 
export const MapOfAreasComponents: Map<string, any> = new Map([["AreasTableComponent", AreasTableComponent],])
export const MapOfBooksComponents: Map<string, any> = new Map([["BooksTableComponent", BooksTableComponent],])
export const MapOfEditorsComponents: Map<string, any> = new Map([["EditorsTableComponent", EditorsTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Area", MapOfAreasComponents],
      ["Book", MapOfBooksComponents],
      ["Editor", MapOfEditorsComponents],
    ]
  )