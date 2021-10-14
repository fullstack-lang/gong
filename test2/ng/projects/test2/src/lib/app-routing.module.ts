import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AstructsTableComponent } from './astructs-table/astructs-table.component'
import { AstructDetailComponent } from './astruct-detail/astruct-detail.component'
import { AstructPresentationComponent } from './astruct-presentation/astruct-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructs', component: AstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-adder', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-adder/:id/:originStruct/:originStructFieldName', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-detail/:id', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-presentation/:id', component: AstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-presentation-special/:id', component: AstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_goastructpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
