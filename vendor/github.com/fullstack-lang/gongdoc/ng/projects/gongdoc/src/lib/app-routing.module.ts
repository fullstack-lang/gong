import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ClassdiagramsTableComponent } from './classdiagrams-table/classdiagrams-table.component'
import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'
import { ClassdiagramPresentationComponent } from './classdiagram-presentation/classdiagram-presentation.component'

import { ClassshapesTableComponent } from './classshapes-table/classshapes-table.component'
import { ClassshapeDetailComponent } from './classshape-detail/classshape-detail.component'
import { ClassshapePresentationComponent } from './classshape-presentation/classshape-presentation.component'

import { FieldsTableComponent } from './fields-table/fields-table.component'
import { FieldDetailComponent } from './field-detail/field-detail.component'
import { FieldPresentationComponent } from './field-presentation/field-presentation.component'

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'
import { GongStructPresentationComponent } from './gongstruct-presentation/gongstruct-presentation.component'

import { GongdocCommandsTableComponent } from './gongdoccommands-table/gongdoccommands-table.component'
import { GongdocCommandDetailComponent } from './gongdoccommand-detail/gongdoccommand-detail.component'
import { GongdocCommandPresentationComponent } from './gongdoccommand-presentation/gongdoccommand-presentation.component'

import { GongdocStatussTableComponent } from './gongdocstatuss-table/gongdocstatuss-table.component'
import { GongdocStatusDetailComponent } from './gongdocstatus-detail/gongdocstatus-detail.component'
import { GongdocStatusPresentationComponent } from './gongdocstatus-presentation/gongdocstatus-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { NotesTableComponent } from './notes-table/notes-table.component'
import { NoteDetailComponent } from './note-detail/note-detail.component'
import { NotePresentationComponent } from './note-presentation/note-presentation.component'

import { PkgeltsTableComponent } from './pkgelts-table/pkgelts-table.component'
import { PkgeltDetailComponent } from './pkgelt-detail/pkgelt-detail.component'
import { PkgeltPresentationComponent } from './pkgelt-presentation/pkgelt-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

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

	{ path: 'github_com_fullstack_lang_gongdoc_go-fields', component: FieldsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-adder', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-adder/:id/:originStruct/:originStructFieldName', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-detail/:id', component: FieldDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-presentation/:id', component: FieldPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-field-presentation-special/:id', component: FieldPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gofieldpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstructs', component: GongStructsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstruct-adder', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstruct-adder/:id/:originStruct/:originStructFieldName', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstruct-detail/:id', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstruct-presentation/:id', component: GongStructPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongstruct-presentation-special/:id', component: GongStructPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gogongstructpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdoccommands', component: GongdocCommandsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdoccommand-adder', component: GongdocCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdoccommand-adder/:id/:originStruct/:originStructFieldName', component: GongdocCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdoccommand-detail/:id', component: GongdocCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdoccommand-presentation/:id', component: GongdocCommandPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdoccommand-presentation-special/:id', component: GongdocCommandPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gogongdoccommandpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdocstatuss', component: GongdocStatussTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdocstatus-adder', component: GongdocStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdocstatus-adder/:id/:originStruct/:originStructFieldName', component: GongdocStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdocstatus-detail/:id', component: GongdocStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdocstatus-presentation/:id', component: GongdocStatusPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-gongdocstatus-presentation-special/:id', component: GongdocStatusPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gogongdocstatuspres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-links', component: LinksTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-adder', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-adder/:id/:originStruct/:originStructFieldName', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-detail/:id', component: LinkDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-presentation/:id', component: LinkPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-link-presentation-special/:id', component: LinkPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_golinkpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-notes', component: NotesTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-adder', component: NoteDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-adder/:id/:originStruct/:originStructFieldName', component: NoteDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-detail/:id', component: NoteDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-presentation/:id', component: NotePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-note-presentation-special/:id', component: NotePresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gonotepres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-pkgelts', component: PkgeltsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-pkgelt-adder', component: PkgeltDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-pkgelt-adder/:id/:originStruct/:originStructFieldName', component: PkgeltDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-pkgelt-detail/:id', component: PkgeltDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-pkgelt-presentation/:id', component: PkgeltPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-pkgelt-presentation-special/:id', component: PkgeltPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gopkgeltpres' },

	{ path: 'github_com_fullstack_lang_gongdoc_go-positions', component: PositionsTableComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_table' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-adder', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-adder/:id/:originStruct/:originStructFieldName', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-detail/:id', component: PositionDetailComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_editor' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-presentation/:id', component: PositionPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongdoc_go-position-presentation-special/:id', component: PositionPresentationComponent, outlet: 'github_com_fullstack_lang_gongdoc_gopositionpres' },

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
