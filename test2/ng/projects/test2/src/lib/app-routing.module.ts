import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AclasssTableComponent } from './aclasss-table/aclasss-table.component'
import { AclassDetailComponent } from './aclass-detail/aclass-detail.component'
import { AclassPresentationComponent } from './aclass-presentation/aclass-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_test2_go-aclasss', component: AclasssTableComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-aclass-adder', component: AclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-aclass-adder/:id/:association', component: AclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-aclass-detail/:id', component: AclassDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-aclass-presentation/:id', component: AclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-aclass-presentation-special/:id', component: AclassPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_goaclasspres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
