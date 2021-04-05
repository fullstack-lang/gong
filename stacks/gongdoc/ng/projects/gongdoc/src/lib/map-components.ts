// insertion point sub template for components imports 
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
import { FieldsTableComponent } from './fields-table/fields-table.component'
import { GongdocCommandsTableComponent } from './gongdoccommands-table/gongdoccommands-table.component'
import { GongdocStatussTableComponent } from './gongdocstatuss-table/gongdocstatuss-table.component'
import { LinksTableComponent } from './links-table/links-table.component'
import { PkgeltsTableComponent } from './pkgelts-table/pkgelts-table.component'
import { PositionsTableComponent } from './positions-table/positions-table.component'
import { StatesTableComponent } from './states-table/states-table.component'
import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { VerticesTableComponent } from './vertices-table/vertices-table.component'

// insertion point sub template for map of components per struct 
export const MapOfClassdiagramsComponents: Map<string, any> = new Map([["ClassdiagramsTableComponent", ClassdiagramsTableComponent],])
export const MapOfClassshapesComponents: Map<string, any> = new Map([["ClassshapesTableComponent", ClassshapesTableComponent],])
export const MapOfFieldsComponents: Map<string, any> = new Map([["FieldsTableComponent", FieldsTableComponent],])
export const MapOfGongdocCommandsComponents: Map<string, any> = new Map([["GongdocCommandsTableComponent", GongdocCommandsTableComponent],])
export const MapOfGongdocStatussComponents: Map<string, any> = new Map([["GongdocStatussTableComponent", GongdocStatussTableComponent],])
export const MapOfLinksComponents: Map<string, any> = new Map([["LinksTableComponent", LinksTableComponent],])
export const MapOfPkgeltsComponents: Map<string, any> = new Map([["PkgeltsTableComponent", PkgeltsTableComponent],])
export const MapOfPositionsComponents: Map<string, any> = new Map([["PositionsTableComponent", PositionsTableComponent],])
export const MapOfStatesComponents: Map<string, any> = new Map([["StatesTableComponent", StatesTableComponent],])
export const MapOfUmlscsComponents: Map<string, any> = new Map([["UmlscsTableComponent", UmlscsTableComponent],])
export const MapOfVerticesComponents: Map<string, any> = new Map([["VerticesTableComponent", VerticesTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Classdiagram", MapOfClassdiagramsComponents],
      ["Classshape", MapOfClassshapesComponents],
      ["Field", MapOfFieldsComponents],
      ["GongdocCommand", MapOfGongdocCommandsComponents],
      ["GongdocStatus", MapOfGongdocStatussComponents],
      ["Link", MapOfLinksComponents],
      ["Pkgelt", MapOfPkgeltsComponents],
      ["Position", MapOfPositionsComponents],
      ["State", MapOfStatesComponents],
      ["Umlsc", MapOfUmlscsComponents],
      ["Vertice", MapOfVerticesComponents],
    ]
  )