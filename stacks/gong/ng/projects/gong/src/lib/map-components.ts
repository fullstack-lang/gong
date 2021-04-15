// insertion point sub template for components imports 
import { GongBasicFieldsTableComponent } from './gongbasicfields-table/gongbasicfields-table.component'
import { GongEnumsTableComponent } from './gongenums-table/gongenums-table.component'
import { GongEnumValuesTableComponent } from './gongenumvalues-table/gongenumvalues-table.component'
import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
import { ModelPkgsTableComponent } from './modelpkgs-table/modelpkgs-table.component'
import { PointerToGongStructFieldsTableComponent } from './pointertogongstructfields-table/pointertogongstructfields-table.component'
import { SliceOfPointerToGongStructFieldsTableComponent } from './sliceofpointertogongstructfields-table/sliceofpointertogongstructfields-table.component'

// insertion point sub template for map of components per struct 
export const MapOfGongBasicFieldsComponents: Map<string, any> = new Map([["GongBasicFieldsTableComponent", GongBasicFieldsTableComponent],])
export const MapOfGongEnumsComponents: Map<string, any> = new Map([["GongEnumsTableComponent", GongEnumsTableComponent],])
export const MapOfGongEnumValuesComponents: Map<string, any> = new Map([["GongEnumValuesTableComponent", GongEnumValuesTableComponent],])
export const MapOfGongStructsComponents: Map<string, any> = new Map([["GongStructsTableComponent", GongStructsTableComponent],])
export const MapOfGongTimeFieldsComponents: Map<string, any> = new Map([["GongTimeFieldsTableComponent", GongTimeFieldsTableComponent],])
export const MapOfModelPkgsComponents: Map<string, any> = new Map([["ModelPkgsTableComponent", ModelPkgsTableComponent],])
export const MapOfPointerToGongStructFieldsComponents: Map<string, any> = new Map([["PointerToGongStructFieldsTableComponent", PointerToGongStructFieldsTableComponent],])
export const MapOfSliceOfPointerToGongStructFieldsComponents: Map<string, any> = new Map([["SliceOfPointerToGongStructFieldsTableComponent", SliceOfPointerToGongStructFieldsTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["GongBasicField", MapOfGongBasicFieldsComponents],
      ["GongEnum", MapOfGongEnumsComponents],
      ["GongEnumValue", MapOfGongEnumValuesComponents],
      ["GongStruct", MapOfGongStructsComponents],
      ["GongTimeField", MapOfGongTimeFieldsComponents],
      ["ModelPkg", MapOfModelPkgsComponents],
      ["PointerToGongStructField", MapOfPointerToGongStructFieldsComponents],
      ["SliceOfPointerToGongStructField", MapOfSliceOfPointerToGongStructFieldsComponents],
    ]
  )