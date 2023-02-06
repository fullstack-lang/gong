import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AstructsTableComponent } from './astructs-table/astructs-table.component'
import { AstructDetailComponent } from './astruct-detail/astruct-detail.component'

import { AstructBstruct2UsesTableComponent } from './astructbstruct2uses-table/astructbstruct2uses-table.component'
import { AstructBstruct2UseDetailComponent } from './astructbstruct2use-detail/astructbstruct2use-detail.component'

import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
import { AstructBstructUseDetailComponent } from './astructbstructuse-detail/astructbstructuse-detail.component'

import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
import { BstructDetailComponent } from './bstruct-detail/bstruct-detail.component'

import { DstructsTableComponent } from './dstructs-table/dstructs-table.component'
import { DstructDetailComponent } from './dstruct-detail/dstruct-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_test_go-astructs', component: AstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-adder', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-adder/:id/:originStruct/:originStructFieldName', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-detail/:id', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2uses', component: AstructBstruct2UsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-adder', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-adder/:id/:originStruct/:originStructFieldName', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-detail/:id', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuses', component: AstructBstructUsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-adder', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-adder/:id/:originStruct/:originStructFieldName', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-detail/:id', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_test_go-bstructs', component: BstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-adder', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-adder/:id/:originStruct/:originStructFieldName', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-detail/:id', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_test_go-dstructs', component: DstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-adder', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-adder/:id/:originStruct/:originStructFieldName', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-detail/:id', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
