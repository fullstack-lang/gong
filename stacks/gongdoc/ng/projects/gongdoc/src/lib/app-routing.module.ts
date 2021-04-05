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

import { GongdocCommandsTableComponent } from './gongdoccommands-table/gongdoccommands-table.component'
import { GongdocCommandDetailComponent } from './gongdoccommand-detail/gongdoccommand-detail.component'
import { GongdocCommandPresentationComponent } from './gongdoccommand-presentation/gongdoccommand-presentation.component'

import { GongdocStatussTableComponent } from './gongdocstatuss-table/gongdocstatuss-table.component'
import { GongdocStatusDetailComponent } from './gongdocstatus-detail/gongdocstatus-detail.component'
import { GongdocStatusPresentationComponent } from './gongdocstatus-presentation/gongdocstatus-presentation.component'

import { LinksTableComponent } from './links-table/links-table.component'
import { LinkDetailComponent } from './link-detail/link-detail.component'
import { LinkPresentationComponent } from './link-presentation/link-presentation.component'

import { PkgeltsTableComponent } from './pkgelts-table/pkgelts-table.component'
import { PkgeltDetailComponent } from './pkgelt-detail/pkgelt-detail.component'
import { PkgeltPresentationComponent } from './pkgelt-presentation/pkgelt-presentation.component'

import { PositionsTableComponent } from './positions-table/positions-table.component'
import { PositionDetailComponent } from './position-detail/position-detail.component'
import { PositionPresentationComponent } from './position-presentation/position-presentation.component'

import { StatesTableComponent } from './states-table/states-table.component'
import { StateDetailComponent } from './state-detail/state-detail.component'
import { StatePresentationComponent } from './state-presentation/state-presentation.component'

import { UmlscsTableComponent } from './umlscs-table/umlscs-table.component'
import { UmlscDetailComponent } from './umlsc-detail/umlsc-detail.component'
import { UmlscPresentationComponent } from './umlsc-presentation/umlsc-presentation.component'

import { VerticesTableComponent } from './vertices-table/vertices-table.component'
import { VerticeDetailComponent } from './vertice-detail/vertice-detail.component'
import { VerticePresentationComponent } from './vertice-presentation/vertice-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'classdiagrams', component: ClassdiagramsTableComponent, outlet: 'table' },
	{ path: 'classdiagram-adder', component: ClassdiagramDetailComponent, outlet: 'editor' },
	{ path: 'classdiagram-adder/:id/:association', component: ClassdiagramDetailComponent, outlet: 'editor' },
	{ path: 'classdiagram-detail/:id', component: ClassdiagramDetailComponent, outlet: 'editor' },
	{ path: 'classdiagram-presentation/:id', component: ClassdiagramPresentationComponent, outlet: 'presentation' },
	{ path: 'classdiagram-presentation-special/:id', component: ClassdiagramPresentationComponent, outlet: 'classdiagrampres' },

	{ path: 'classshapes', component: ClassshapesTableComponent, outlet: 'table' },
	{ path: 'classshape-adder', component: ClassshapeDetailComponent, outlet: 'editor' },
	{ path: 'classshape-adder/:id/:association', component: ClassshapeDetailComponent, outlet: 'editor' },
	{ path: 'classshape-detail/:id', component: ClassshapeDetailComponent, outlet: 'editor' },
	{ path: 'classshape-presentation/:id', component: ClassshapePresentationComponent, outlet: 'presentation' },
	{ path: 'classshape-presentation-special/:id', component: ClassshapePresentationComponent, outlet: 'classshapepres' },

	{ path: 'fields', component: FieldsTableComponent, outlet: 'table' },
	{ path: 'field-adder', component: FieldDetailComponent, outlet: 'editor' },
	{ path: 'field-adder/:id/:association', component: FieldDetailComponent, outlet: 'editor' },
	{ path: 'field-detail/:id', component: FieldDetailComponent, outlet: 'editor' },
	{ path: 'field-presentation/:id', component: FieldPresentationComponent, outlet: 'presentation' },
	{ path: 'field-presentation-special/:id', component: FieldPresentationComponent, outlet: 'fieldpres' },

	{ path: 'gongdoccommands', component: GongdocCommandsTableComponent, outlet: 'table' },
	{ path: 'gongdoccommand-adder', component: GongdocCommandDetailComponent, outlet: 'editor' },
	{ path: 'gongdoccommand-adder/:id/:association', component: GongdocCommandDetailComponent, outlet: 'editor' },
	{ path: 'gongdoccommand-detail/:id', component: GongdocCommandDetailComponent, outlet: 'editor' },
	{ path: 'gongdoccommand-presentation/:id', component: GongdocCommandPresentationComponent, outlet: 'presentation' },
	{ path: 'gongdoccommand-presentation-special/:id', component: GongdocCommandPresentationComponent, outlet: 'gongdoccommandpres' },

	{ path: 'gongdocstatuss', component: GongdocStatussTableComponent, outlet: 'table' },
	{ path: 'gongdocstatus-adder', component: GongdocStatusDetailComponent, outlet: 'editor' },
	{ path: 'gongdocstatus-adder/:id/:association', component: GongdocStatusDetailComponent, outlet: 'editor' },
	{ path: 'gongdocstatus-detail/:id', component: GongdocStatusDetailComponent, outlet: 'editor' },
	{ path: 'gongdocstatus-presentation/:id', component: GongdocStatusPresentationComponent, outlet: 'presentation' },
	{ path: 'gongdocstatus-presentation-special/:id', component: GongdocStatusPresentationComponent, outlet: 'gongdocstatuspres' },

	{ path: 'links', component: LinksTableComponent, outlet: 'table' },
	{ path: 'link-adder', component: LinkDetailComponent, outlet: 'editor' },
	{ path: 'link-adder/:id/:association', component: LinkDetailComponent, outlet: 'editor' },
	{ path: 'link-detail/:id', component: LinkDetailComponent, outlet: 'editor' },
	{ path: 'link-presentation/:id', component: LinkPresentationComponent, outlet: 'presentation' },
	{ path: 'link-presentation-special/:id', component: LinkPresentationComponent, outlet: 'linkpres' },

	{ path: 'pkgelts', component: PkgeltsTableComponent, outlet: 'table' },
	{ path: 'pkgelt-adder', component: PkgeltDetailComponent, outlet: 'editor' },
	{ path: 'pkgelt-adder/:id/:association', component: PkgeltDetailComponent, outlet: 'editor' },
	{ path: 'pkgelt-detail/:id', component: PkgeltDetailComponent, outlet: 'editor' },
	{ path: 'pkgelt-presentation/:id', component: PkgeltPresentationComponent, outlet: 'presentation' },
	{ path: 'pkgelt-presentation-special/:id', component: PkgeltPresentationComponent, outlet: 'pkgeltpres' },

	{ path: 'positions', component: PositionsTableComponent, outlet: 'table' },
	{ path: 'position-adder', component: PositionDetailComponent, outlet: 'editor' },
	{ path: 'position-adder/:id/:association', component: PositionDetailComponent, outlet: 'editor' },
	{ path: 'position-detail/:id', component: PositionDetailComponent, outlet: 'editor' },
	{ path: 'position-presentation/:id', component: PositionPresentationComponent, outlet: 'presentation' },
	{ path: 'position-presentation-special/:id', component: PositionPresentationComponent, outlet: 'positionpres' },

	{ path: 'states', component: StatesTableComponent, outlet: 'table' },
	{ path: 'state-adder', component: StateDetailComponent, outlet: 'editor' },
	{ path: 'state-adder/:id/:association', component: StateDetailComponent, outlet: 'editor' },
	{ path: 'state-detail/:id', component: StateDetailComponent, outlet: 'editor' },
	{ path: 'state-presentation/:id', component: StatePresentationComponent, outlet: 'presentation' },
	{ path: 'state-presentation-special/:id', component: StatePresentationComponent, outlet: 'statepres' },

	{ path: 'umlscs', component: UmlscsTableComponent, outlet: 'table' },
	{ path: 'umlsc-adder', component: UmlscDetailComponent, outlet: 'editor' },
	{ path: 'umlsc-adder/:id/:association', component: UmlscDetailComponent, outlet: 'editor' },
	{ path: 'umlsc-detail/:id', component: UmlscDetailComponent, outlet: 'editor' },
	{ path: 'umlsc-presentation/:id', component: UmlscPresentationComponent, outlet: 'presentation' },
	{ path: 'umlsc-presentation-special/:id', component: UmlscPresentationComponent, outlet: 'umlscpres' },

	{ path: 'vertices', component: VerticesTableComponent, outlet: 'table' },
	{ path: 'vertice-adder', component: VerticeDetailComponent, outlet: 'editor' },
	{ path: 'vertice-adder/:id/:association', component: VerticeDetailComponent, outlet: 'editor' },
	{ path: 'vertice-detail/:id', component: VerticeDetailComponent, outlet: 'editor' },
	{ path: 'vertice-presentation/:id', component: VerticePresentationComponent, outlet: 'presentation' },
	{ path: 'vertice-presentation-special/:id', component: VerticePresentationComponent, outlet: 'verticepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
