// insertion point sub template for components imports 
  import { GongBasicFieldsTableComponent } from './gongbasicfields-table/gongbasicfields-table.component'
  import { GongBasicFieldSortingComponent } from './gongbasicfield-sorting/gongbasicfield-sorting.component'
  import { GongEnumsTableComponent } from './gongenums-table/gongenums-table.component'
  import { GongEnumSortingComponent } from './gongenum-sorting/gongenum-sorting.component'
  import { GongEnumValuesTableComponent } from './gongenumvalues-table/gongenumvalues-table.component'
  import { GongEnumValueSortingComponent } from './gongenumvalue-sorting/gongenumvalue-sorting.component'
  import { GongNotesTableComponent } from './gongnotes-table/gongnotes-table.component'
  import { GongNoteSortingComponent } from './gongnote-sorting/gongnote-sorting.component'
  import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
  import { GongStructSortingComponent } from './gongstruct-sorting/gongstruct-sorting.component'
  import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
  import { GongTimeFieldSortingComponent } from './gongtimefield-sorting/gongtimefield-sorting.component'
  import { ModelPkgsTableComponent } from './modelpkgs-table/modelpkgs-table.component'
  import { ModelPkgSortingComponent } from './modelpkg-sorting/modelpkg-sorting.component'
  import { PointerToGongStructFieldsTableComponent } from './pointertogongstructfields-table/pointertogongstructfields-table.component'
  import { PointerToGongStructFieldSortingComponent } from './pointertogongstructfield-sorting/pointertogongstructfield-sorting.component'
  import { SliceOfPointerToGongStructFieldsTableComponent } from './sliceofpointertogongstructfields-table/sliceofpointertogongstructfields-table.component'
  import { SliceOfPointerToGongStructFieldSortingComponent } from './sliceofpointertogongstructfield-sorting/sliceofpointertogongstructfield-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfGongBasicFieldsComponents: Map<string, any> = new Map([["GongBasicFieldsTableComponent", GongBasicFieldsTableComponent],])
  export const MapOfGongBasicFieldSortingComponents: Map<string, any> = new Map([["GongBasicFieldSortingComponent", GongBasicFieldSortingComponent],])
  export const MapOfGongEnumsComponents: Map<string, any> = new Map([["GongEnumsTableComponent", GongEnumsTableComponent],])
  export const MapOfGongEnumSortingComponents: Map<string, any> = new Map([["GongEnumSortingComponent", GongEnumSortingComponent],])
  export const MapOfGongEnumValuesComponents: Map<string, any> = new Map([["GongEnumValuesTableComponent", GongEnumValuesTableComponent],])
  export const MapOfGongEnumValueSortingComponents: Map<string, any> = new Map([["GongEnumValueSortingComponent", GongEnumValueSortingComponent],])
  export const MapOfGongNotesComponents: Map<string, any> = new Map([["GongNotesTableComponent", GongNotesTableComponent],])
  export const MapOfGongNoteSortingComponents: Map<string, any> = new Map([["GongNoteSortingComponent", GongNoteSortingComponent],])
  export const MapOfGongStructsComponents: Map<string, any> = new Map([["GongStructsTableComponent", GongStructsTableComponent],])
  export const MapOfGongStructSortingComponents: Map<string, any> = new Map([["GongStructSortingComponent", GongStructSortingComponent],])
  export const MapOfGongTimeFieldsComponents: Map<string, any> = new Map([["GongTimeFieldsTableComponent", GongTimeFieldsTableComponent],])
  export const MapOfGongTimeFieldSortingComponents: Map<string, any> = new Map([["GongTimeFieldSortingComponent", GongTimeFieldSortingComponent],])
  export const MapOfModelPkgsComponents: Map<string, any> = new Map([["ModelPkgsTableComponent", ModelPkgsTableComponent],])
  export const MapOfModelPkgSortingComponents: Map<string, any> = new Map([["ModelPkgSortingComponent", ModelPkgSortingComponent],])
  export const MapOfPointerToGongStructFieldsComponents: Map<string, any> = new Map([["PointerToGongStructFieldsTableComponent", PointerToGongStructFieldsTableComponent],])
  export const MapOfPointerToGongStructFieldSortingComponents: Map<string, any> = new Map([["PointerToGongStructFieldSortingComponent", PointerToGongStructFieldSortingComponent],])
  export const MapOfSliceOfPointerToGongStructFieldsComponents: Map<string, any> = new Map([["SliceOfPointerToGongStructFieldsTableComponent", SliceOfPointerToGongStructFieldsTableComponent],])
  export const MapOfSliceOfPointerToGongStructFieldSortingComponents: Map<string, any> = new Map([["SliceOfPointerToGongStructFieldSortingComponent", SliceOfPointerToGongStructFieldSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["GongBasicField", MapOfGongBasicFieldsComponents],
      ["GongEnum", MapOfGongEnumsComponents],
      ["GongEnumValue", MapOfGongEnumValuesComponents],
      ["GongNote", MapOfGongNotesComponents],
      ["GongStruct", MapOfGongStructsComponents],
      ["GongTimeField", MapOfGongTimeFieldsComponents],
      ["ModelPkg", MapOfModelPkgsComponents],
      ["PointerToGongStructField", MapOfPointerToGongStructFieldsComponents],
      ["SliceOfPointerToGongStructField", MapOfSliceOfPointerToGongStructFieldsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["GongBasicField", MapOfGongBasicFieldSortingComponents],
      ["GongEnum", MapOfGongEnumSortingComponents],
      ["GongEnumValue", MapOfGongEnumValueSortingComponents],
      ["GongNote", MapOfGongNoteSortingComponents],
      ["GongStruct", MapOfGongStructSortingComponents],
      ["GongTimeField", MapOfGongTimeFieldSortingComponents],
      ["ModelPkg", MapOfModelPkgSortingComponents],
      ["PointerToGongStructField", MapOfPointerToGongStructFieldSortingComponents],
      ["SliceOfPointerToGongStructField", MapOfSliceOfPointerToGongStructFieldSortingComponents],
    ]
  )
