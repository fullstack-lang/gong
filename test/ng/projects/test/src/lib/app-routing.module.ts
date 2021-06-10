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
	{ path: 'githubcomfullstacklanggongtestgo-aclasss', component: AclasssTableComponent, outlet: 'githubcomfullstacklanggongtestgotable' },
	{ path: 'githubcomfullstacklanggongtestgo-aclass-adder', component: AclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-aclass-detail/:id', component: AclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'githubcomfullstacklanggongtestgopresentation' },
	{ path: 'githubcomfullstacklanggongtestgo-aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'githubcomfullstacklanggongtestgoaclasspres' },

	{ path: 'githubcomfullstacklanggongtestgo-bclasss', component: BclasssTableComponent, outlet: 'githubcomfullstacklanggongtestgotable' },
	{ path: 'githubcomfullstacklanggongtestgo-bclass-adder', component: BclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-bclass-adder/:id/:association', component: BclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-bclass-detail/:id', component: BclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-bclass-presentation/:id', component: BclassPresentationComponent, outlet: 'githubcomfullstacklanggongtestgopresentation' },
	{ path: 'githubcomfullstacklanggongtestgo-bclass-presentation-special/:id', component: BclassPresentationComponent, outlet: 'githubcomfullstacklanggongtestgobclasspres' },

	{ path: 'githubcomfullstacklanggongtestgo-dclasss', component: DclasssTableComponent, outlet: 'githubcomfullstacklanggongtestgotable' },
	{ path: 'githubcomfullstacklanggongtestgo-dclass-adder', component: DclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-dclass-adder/:id/:association', component: DclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-dclass-detail/:id', component: DclassDetailComponent, outlet: 'githubcomfullstacklanggongtestgoeditor' },
	{ path: 'githubcomfullstacklanggongtestgo-dclass-presentation/:id', component: DclassPresentationComponent, outlet: 'githubcomfullstacklanggongtestgopresentation' },
	{ path: 'githubcomfullstacklanggongtestgo-dclass-presentation-special/:id', component: DclassPresentationComponent, outlet: 'githubcomfullstacklanggongtestgodclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
