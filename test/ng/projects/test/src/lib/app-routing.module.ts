import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { AstructsTableComponent } from './astructs-table/astructs-table.component'
import { AstructDetailComponent } from './astruct-detail/astruct-detail.component'
import { AstructPresentationComponent } from './astruct-presentation/astruct-presentation.component'

import { AstructBstruct2UsesTableComponent } from './astructbstruct2uses-table/astructbstruct2uses-table.component'
import { AstructBstruct2UseDetailComponent } from './astructbstruct2use-detail/astructbstruct2use-detail.component'
import { AstructBstruct2UsePresentationComponent } from './astructbstruct2use-presentation/astructbstruct2use-presentation.component'

import { AstructBstructUsesTableComponent } from './astructbstructuses-table/astructbstructuses-table.component'
import { AstructBstructUseDetailComponent } from './astructbstructuse-detail/astructbstructuse-detail.component'
import { AstructBstructUsePresentationComponent } from './astructbstructuse-presentation/astructbstructuse-presentation.component'

import { BstructsTableComponent } from './bstructs-table/bstructs-table.component'
import { BstructDetailComponent } from './bstruct-detail/bstruct-detail.component'
import { BstructPresentationComponent } from './bstruct-presentation/bstruct-presentation.component'

import { DstructsTableComponent } from './dstructs-table/dstructs-table.component'
import { DstructDetailComponent } from './dstruct-detail/dstruct-detail.component'
import { DstructPresentationComponent } from './dstruct-presentation/dstruct-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_test_go-astructs', component: AstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-adder', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-adder/:id/:originStruct/:originStructFieldName', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-detail/:id', component: AstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-presentation/:id', component: AstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astruct-presentation-special/:id', component: AstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_goastructpres' },

	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2uses', component: AstructBstruct2UsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-adder', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-adder/:id/:originStruct/:originStructFieldName', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-detail/:id', component: AstructBstruct2UseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-presentation/:id', component: AstructBstruct2UsePresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstruct2use-presentation-special/:id', component: AstructBstruct2UsePresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_goastructbstruct2usepres' },

	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuses', component: AstructBstructUsesTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-adder', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-adder/:id/:originStruct/:originStructFieldName', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-detail/:id', component: AstructBstructUseDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-presentation/:id', component: AstructBstructUsePresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-astructbstructuse-presentation-special/:id', component: AstructBstructUsePresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_goastructbstructusepres' },

	{ path: 'github_com_fullstack_lang_gong_test_go-bstructs', component: BstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-adder', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-adder/:id/:originStruct/:originStructFieldName', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-detail/:id', component: BstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-presentation/:id', component: BstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-bstruct-presentation-special/:id', component: BstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_gobstructpres' },

	{ path: 'github_com_fullstack_lang_gong_test_go-dstructs', component: DstructsTableComponent, outlet: 'github_com_fullstack_lang_gong_test_go_table' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-adder', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-adder/:id/:originStruct/:originStructFieldName', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-detail/:id', component: DstructDetailComponent, outlet: 'github_com_fullstack_lang_gong_test_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-presentation/:id', component: DstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_test_go-dstruct-presentation-special/:id', component: DstructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_test_godstructpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
