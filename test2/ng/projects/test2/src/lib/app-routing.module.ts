import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AstructsTableComponent } from './astructs-table/astructs-table.component'
import { AstructDetailComponent } from './astruct-detail/astruct-detail.component'
import { AstructPresentationComponent } from './astruct-presentation/astruct-presentation.component'

import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
import { AstructBstructUseDetailComponent } from './astructbstructuse-detail/astructbstructuse-detail.component'
import { AstructBstructUsePresentationComponent } from './astructbstructuse-presentation/astructbstructuse-presentation.component'

import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
import { BstructDetailComponent } from './bstruct-detail/bstruct-detail.component'
import { BstructPresentationComponent } from './bstruct-presentation/bstruct-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructs', component: AstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-adder', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-adder/:id/:originStruct/:originStructFieldName', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-detail/:id', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-presentation/:id', component: AstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astruct-presentation-special/:id', component: AstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_goastructpres' },

	{ path: 'github_com_fullstack_lang_gong_test2_go-astructbstructuses', component: AstructBstructUsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructbstructuse-adder', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructbstructuse-adder/:id/:originStruct/:originStructFieldName', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructbstructuse-detail/:id', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructbstructuse-presentation/:id', component: AstructBstructUsePresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-astructbstructuse-presentation-special/:id', component: AstructBstructUsePresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_goastructbstructusepres' },

	{ path: 'github_com_fullstack_lang_gong_test2_go-bstructs', component: BstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-bstruct-adder', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-bstruct-adder/:id/:originStruct/:originStructFieldName', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-bstruct-detail/:id', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-bstruct-presentation/:id', component: BstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test2_go-bstruct-presentation-special/:id', component: BstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test2_gobstructpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
