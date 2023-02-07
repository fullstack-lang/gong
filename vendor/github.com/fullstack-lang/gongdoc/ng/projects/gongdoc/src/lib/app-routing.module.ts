import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'

import { DiagramPackagesTableComponent } from './diagrampackages-table/diagrampackages-table.component'
import { DiagramPackageDetailComponent } from './diagrampackage-detail/diagrampackage-detail.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'

import { GongEnumShapesTableComponent } from './gongenumshapes-table/gongenumshapes-table.component'
import { GongEnumShapeDetailComponent } from './gongenumshape-detail/gongenumshape-detail.component'

import { GongEnumValueEntrysTableComponent } from './gongenumvalueentrys-table/gongenumvalueentrys-table.component'
import { GongEnumValueEntryDetailComponent } from './gongenumvalueentry-detail/gongenumvalueentry-detail.component'

import { GongStructShapesTableComponent } from './gongstructshapes-table/gongstructshapes-table.component'
import { GongStructShapeDetailComponent } from './gongstructshape-detail/gongstructshape-detail.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'

import { NoteShapesTableComponent } from './noteshapes-table/noteshapes-table.component'
import { NoteShapeDetailComponent } from './noteshape-detail/noteshape-detail.component'

import { NoteShapeLinksTableComponent } from './noteshapelinks-table/noteshapelinks-table.component'
import { NoteShapeLinkDetailComponent } from './noteshapelink-detail/noteshapelink-detail.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'

import { TreesTableComponent } from './trees-table/trees-table.component'
import { TreeDetailComponent } from './tree-detail/tree-detail.component'

import { UmlStatesTableComponent } from './umlstates-table/umlstates-table.component'
import { UmlStateDetailComponent } from './umlstate-detail/umlstate-detail.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagrams', component: ClassdiagramsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-adder', component: ClassdiagramDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-adder/:id/:originStruct/:originStructFieldName', component: ClassdiagramDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-detail/:id', component: ClassdiagramDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackages', component: DiagramPackagesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-adder', component: DiagramPackageDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-adder/:id/:originStruct/:originStructFieldName', component: DiagramPackageDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-detail/:id', component: DiagramPackageDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-fields', component: FieldsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-adder', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-adder/:id/:originStruct/:originStructFieldName', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-detail/:id', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumshapes', component: GongEnumShapesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumshape-adder', component: GongEnumShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumshape-adder/:id/:originStruct/:originStructFieldName', component: GongEnumShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumshape-detail/:id', component: GongEnumShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumvalueentrys', component: GongEnumValueEntrysTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumvalueentry-adder', component: GongEnumValueEntryDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumvalueentry-adder/:id/:originStruct/:originStructFieldName', component: GongEnumValueEntryDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongenumvalueentry-detail/:id', component: GongEnumValueEntryDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstructshapes', component: GongStructShapesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstructshape-adder', component: GongStructShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstructshape-adder/:id/:originStruct/:originStructFieldName', component: GongStructShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstructshape-detail/:id', component: GongStructShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-links', component: LinksTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-adder', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-adder/:id/:originStruct/:originStructFieldName', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-detail/:id', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-nodes', component: NodesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-adder', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-adder/:id/:originStruct/:originStructFieldName', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-detail/:id', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshapes', component: NoteShapesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshape-adder', component: NoteShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshape-adder/:id/:originStruct/:originStructFieldName', component: NoteShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshape-detail/:id', component: NoteShapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshapelinks', component: NoteShapeLinksTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshapelink-adder', component: NoteShapeLinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshapelink-adder/:id/:originStruct/:originStructFieldName', component: NoteShapeLinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-noteshapelink-detail/:id', component: NoteShapeLinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-positions', component: PositionsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-adder', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-adder/:id/:originStruct/:originStructFieldName', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-detail/:id', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-trees', component: TreesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-adder', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-adder/:id/:originStruct/:originStructFieldName', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-detail/:id', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstates', component: UmlStatesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-adder', component: UmlStateDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-adder/:id/:originStruct/:originStructFieldName', component: UmlStateDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-detail/:id', component: UmlStateDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-umlscs', component: UmlscsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-adder', component: UmlscDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-adder/:id/:originStruct/:originStructFieldName', component: UmlscDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-detail/:id', component: UmlscDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-vertices', component: VerticesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-adder', component: VerticeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-adder/:id/:originStruct/:originStructFieldName', component: VerticeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-detail/:id', component: VerticeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
