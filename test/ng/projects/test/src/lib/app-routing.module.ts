import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { AclassDetailComponent } from './aclass-detail/aclass-detail.component'
import { AclassPresentationComponent } from './aclass-presentation/aclass-presentation.component'

import { BclasssTableComponent } from './bclasss-table/bclasss-table.component'
import { BclassDetailComponent } from './bclass-detail/bclass-detail.component'
import { BclassPresentationComponent } from './bclass-presentation/bclass-presentation.component'

import { DclasssTableComponent } from './dclasss-table/dclasss-table.component'
import { DclassDetailComponent } from './dclass-detail/dclass-detail.component'
import { DclassPresentationComponent } from './dclass-presentation/dclass-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'aclasss', component: AclasssTableComponent, outlet: 'table' },
	{ path: 'aclass-adder', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'aclass-detail/:id', component: AclassDetailComponent, outlet: 'editor' },
	{ path: 'aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'presentation' },
	{ path: 'aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'aclasspres' },

	{ path: 'bclasss', component: BclasssTableComponent, outlet: 'table' },
	{ path: 'bclass-adder', component: BclassDetailComponent, outlet: 'editor' },
	{ path: 'bclass-adder/:id/:association', component: BclassDetailComponent, outlet: 'editor' },
	{ path: 'bclass-detail/:id', component: BclassDetailComponent, outlet: 'editor' },
	{ path: 'bclass-presentation/:id', component: BclassPresentationComponent, outlet: 'presentation' },
	{ path: 'bclass-presentation-special/:id', component: BclassPresentationComponent, outlet: 'bclasspres' },

	{ path: 'dclasss', component: DclasssTableComponent, outlet: 'table' },
	{ path: 'dclass-adder', component: DclassDetailComponent, outlet: 'editor' },
	{ path: 'dclass-adder/:id/:association', component: DclassDetailComponent, outlet: 'editor' },
	{ path: 'dclass-detail/:id', component: DclassDetailComponent, outlet: 'editor' },
	{ path: 'dclass-presentation/:id', component: DclassPresentationComponent, outlet: 'presentation' },
	{ path: 'dclass-presentation-special/:id', component: DclassPresentationComponent, outlet: 'dclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
