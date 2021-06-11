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
	{ path: 'github_com_fullstack_lang_gong_test_go-aclasss', component: AclasssTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-aclass-adder', component: AclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-aclass-detail/:id', component: AclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_goaclasspres' },

	{ path: 'github_com_fullstack_lang_gong_test_go-bclasss', component: BclasssTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bclass-adder', component: BclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bclass-adder/:id/:association', component: BclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bclass-detail/:id', component: BclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bclass-presentation/:id', component: BclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bclass-presentation-special/:id', component: BclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_gobclasspres' },

	{ path: 'github_com_fullstack_lang_gong_test_go-dclasss', component: DclasssTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dclass-adder', component: DclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dclass-adder/:id/:association', component: DclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dclass-detail/:id', component: DclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dclass-presentation/:id', component: DclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dclass-presentation-special/:id', component: DclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_godclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
