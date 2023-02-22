import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { GongBasicFieldsTableComponent } from './gongbasicfields-table/gongbasicfields-table.component'
import { GongBasicFieldDetailComponent } from './gongbasicfield-detail/gongbasicfield-detail.component'

import { GongEnumsTableComponent } from './gongenums-table/gongenums-table.component'
import { GongEnumDetailComponent } from './gongenum-detail/gongenum-detail.component'

import { GongEnumValuesTableComponent } from './gongenumvalues-table/gongenumvalues-table.component'
import { GongEnumValueDetailComponent } from './gongenumvalue-detail/gongenumvalue-detail.component'

import { GongLinksTableComponent } from './gonglinks-table/gonglinks-table.component'
import { GongLinkDetailComponent } from './gonglink-detail/gonglink-detail.component'

import { GongNotesTableComponent } from './gongnotes-table/gongnotes-table.component'
import { GongNoteDetailComponent } from './gongnote-detail/gongnote-detail.component'

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'

import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
import { GongTimeFieldDetailComponent } from './gongtimefield-detail/gongtimefield-detail.component'

import { MetasTableComponent } from './metas-table/metas-table.component'
import { MetaDetailComponent } from './meta-detail/meta-detail.component'

import { MetaReferencesTableComponent } from './metareferences-table/metareferences-table.component'
import { MetaReferenceDetailComponent } from './metareference-detail/metareference-detail.component'

import { ModelPkgsTableComponent } from './modelpkgs-table/modelpkgs-table.component'
import { ModelPkgDetailComponent } from './modelpkg-detail/modelpkg-detail.component'

import { PointerToGongStructFieldsTableComponent } from './pointertogongstructfields-table/pointertogongstructfields-table.component'
import { PointerToGongStructFieldDetailComponent } from './pointertogongstructfield-detail/pointertogongstructfield-detail.component'

import { SliceOfPointerToGongStructFieldsTableComponent } from './sliceofpointertogongstructfields-table/sliceofpointertogongstructfields-table.component'
import { SliceOfPointerToGongStructFieldDetailComponent } from './sliceofpointertogongstructfield-detail/sliceofpointertogongstructfield-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfields/:GONG__StackPath', component: GongBasicFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-adder/:GONG__StackPath', component: GongBasicFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongBasicFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-detail/:id/:GONG__StackPath', component: GongBasicFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-gongenums/:GONG__StackPath', component: GongEnumsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-adder/:GONG__StackPath', component: GongEnumDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongEnumDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-detail/:id/:GONG__StackPath', component: GongEnumDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalues/:GONG__StackPath', component: GongEnumValuesTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-adder/:GONG__StackPath', component: GongEnumValueDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongEnumValueDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-detail/:id/:GONG__StackPath', component: GongEnumValueDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-gonglinks/:GONG__StackPath', component: GongLinksTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-adder/:GONG__StackPath', component: GongLinkDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongLinkDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-detail/:id/:GONG__StackPath', component: GongLinkDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-gongnotes/:GONG__StackPath', component: GongNotesTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-adder/:GONG__StackPath', component: GongNoteDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongNoteDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-detail/:id/:GONG__StackPath', component: GongNoteDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-gongstructs/:GONG__StackPath', component: GongStructsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-adder/:GONG__StackPath', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-detail/:id/:GONG__StackPath', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-gongtimefields/:GONG__StackPath', component: GongTimeFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-adder/:GONG__StackPath', component: GongTimeFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongTimeFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-detail/:id/:GONG__StackPath', component: GongTimeFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-metas/:GONG__StackPath', component: MetasTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-adder/:GONG__StackPath', component: MetaDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: MetaDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-detail/:id/:GONG__StackPath', component: MetaDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-metareferences/:GONG__StackPath', component: MetaReferencesTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-adder/:GONG__StackPath', component: MetaReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: MetaReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-detail/:id/:GONG__StackPath', component: MetaReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-modelpkgs/:GONG__StackPath', component: ModelPkgsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-adder/:GONG__StackPath', component: ModelPkgDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: ModelPkgDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-detail/:id/:GONG__StackPath', component: ModelPkgDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfields/:GONG__StackPath', component: PointerToGongStructFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-adder/:GONG__StackPath', component: PointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: PointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-detail/:id/:GONG__StackPath', component: PointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfields/:GONG__StackPath', component: SliceOfPointerToGongStructFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-adder/:GONG__StackPath', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-detail/:id/:GONG__StackPath', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
