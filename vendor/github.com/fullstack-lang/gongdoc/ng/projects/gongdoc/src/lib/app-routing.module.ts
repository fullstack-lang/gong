import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'
import { ClassdiagramPresentationComponent } from './classdiagram-presentation/classdiagram-presentation.component'

import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
import { ClassshapeDetailComponent } from './classshape-detail/classshape-detail.component'
import { ClassshapePresentationComponent } from './classshape-presentation/classshape-presentation.component'

import { DiagramPackagesTableComponent } from './diagrampackages-table/diagrampackages-table.component'
import { DiagramPackageDetailComponent } from './diagrampackage-detail/diagrampackage-detail.component'
import { DiagramPackagePresentationComponent } from './diagrampackage-presentation/diagrampackage-presentation.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'
import { FieldPresentationComponent } from './field-presentation/field-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'
import { NodePresentationComponent } from './node-presentation/node-presentation.component'

import { NotesTableComponent } from './notes-table/notes-table.component'
import { NoteDetailComponent } from './note-detail/note-detail.component'
import { NotePresentationComponent } from './note-presentation/note-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

import { ReferencesTableComponent } from './references-table/references-table.component'
import { ReferenceDetailComponent } from './reference-detail/reference-detail.component'
import { ReferencePresentationComponent } from './reference-presentation/reference-presentation.component'

import { TreesTableComponent } from './trees-table/trees-table.component'
import { TreeDetailComponent } from './tree-detail/tree-detail.component'
import { TreePresentationComponent } from './tree-presentation/tree-presentation.component'

import { UmlStatesTableComponent } from './umlstates-table/umlstates-table.component'
import { UmlStateDetailComponent } from './umlstate-detail/umlstate-detail.component'
import { UmlStatePresentationComponent } from './umlstate-presentation/umlstate-presentation.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'
import { UmlscPresentationComponent } from './umlsc-presentation/umlsc-presentation.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'
import { VerticePresentationComponent } from './vertice-presentation/vertice-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagrams', component: ClassdiagramsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-adder', component: ClassdiagramDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-adder/:id/:originStruct/:originStructFieldName', component: ClassdiagramDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-detail/:id', component: ClassdiagramDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-presentation/:id', component: ClassdiagramPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-presentation-special/:id', component: ClassdiagramPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_goclassdiagrampres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-classshapes', component: ClassshapesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classshape-adder', component: ClassshapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classshape-adder/:id/:originStruct/:originStructFieldName', component: ClassshapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classshape-detail/:id', component: ClassshapeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classshape-presentation/:id', component: ClassshapePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-classshape-presentation-special/:id', component: ClassshapePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_goclassshapepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackages', component: DiagramPackagesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-adder', component: DiagramPackageDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-adder/:id/:originStruct/:originStructFieldName', component: DiagramPackageDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-detail/:id', component: DiagramPackageDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-presentation/:id', component: DiagramPackagePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-diagrampackage-presentation-special/:id', component: DiagramPackagePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_godiagrampackagepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-fields', component: FieldsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-adder', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-adder/:id/:originStruct/:originStructFieldName', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-detail/:id', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-presentation/:id', component: FieldPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-presentation-special/:id', component: FieldPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gofieldpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-links', component: LinksTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-adder', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-adder/:id/:originStruct/:originStructFieldName', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-detail/:id', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-presentation/:id', component: LinkPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-presentation-special/:id', component: LinkPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_golinkpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-nodes', component: NodesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-adder', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-adder/:id/:originStruct/:originStructFieldName', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-detail/:id', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-presentation/:id', component: NodePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-node-presentation-special/:id', component: NodePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gonodepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-notes', component: NotesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-adder', component: NoteDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-adder/:id/:originStruct/:originStructFieldName', component: NoteDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-detail/:id', component: NoteDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-presentation/:id', component: NotePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-presentation-special/:id', component: NotePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gonotepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-positions', component: PositionsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-adder', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-adder/:id/:originStruct/:originStructFieldName', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-detail/:id', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-presentation/:id', component: PositionPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-presentation-special/:id', component: PositionPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gopositionpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-references', component: ReferencesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-reference-adder', component: ReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-reference-adder/:id/:originStruct/:originStructFieldName', component: ReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-reference-detail/:id', component: ReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-reference-presentation/:id', component: ReferencePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-reference-presentation-special/:id', component: ReferencePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_goreferencepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-trees', component: TreesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-adder', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-adder/:id/:originStruct/:originStructFieldName', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-detail/:id', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-presentation/:id', component: TreePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-tree-presentation-special/:id', component: TreePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gotreepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstates', component: UmlStatesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-adder', component: UmlStateDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-adder/:id/:originStruct/:originStructFieldName', component: UmlStateDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-detail/:id', component: UmlStateDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-presentation/:id', component: UmlStatePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlstate-presentation-special/:id', component: UmlStatePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_goumlstatepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-umlscs', component: UmlscsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-adder', component: UmlscDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-adder/:id/:originStruct/:originStructFieldName', component: UmlscDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-detail/:id', component: UmlscDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-presentation/:id', component: UmlscPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-umlsc-presentation-special/:id', component: UmlscPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_goumlscpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-vertices', component: VerticesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-adder', component: VerticeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-adder/:id/:originStruct/:originStructFieldName', component: VerticeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-detail/:id', component: VerticeDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-presentation/:id', component: VerticePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-vertice-presentation-special/:id', component: VerticePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_goverticepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
