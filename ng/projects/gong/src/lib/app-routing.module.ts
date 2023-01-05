import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { GongBasicFieldsTableComponent } from './gongbasicfields-table/gongbasicfields-table.component'
import { GongBasicFieldDetailComponent } from './gongbasicfield-detail/gongbasicfield-detail.component'
import { GongBasicFieldPresentationComponent } from './gongbasicfield-presentation/gongbasicfield-presentation.component'

import { GongEnumsTableComponent } from './gongenums-table/gongenums-table.component'
import { GongEnumDetailComponent } from './gongenum-detail/gongenum-detail.component'
import { GongEnumPresentationComponent } from './gongenum-presentation/gongenum-presentation.component'

import { GongEnumValuesTableComponent } from './gongenumvalues-table/gongenumvalues-table.component'
import { GongEnumValueDetailComponent } from './gongenumvalue-detail/gongenumvalue-detail.component'
import { GongEnumValuePresentationComponent } from './gongenumvalue-presentation/gongenumvalue-presentation.component'

import { GongLinksTableComponent } from './gonglinks-table/gonglinks-table.component'
import { GongLinkDetailComponent } from './gonglink-detail/gonglink-detail.component'
import { GongLinkPresentationComponent } from './gonglink-presentation/gonglink-presentation.component'

import { GongNotesTableComponent } from './gongnotes-table/gongnotes-table.component'
import { GongNoteDetailComponent } from './gongnote-detail/gongnote-detail.component'
import { GongNotePresentationComponent } from './gongnote-presentation/gongnote-presentation.component'

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'
import { GongStructPresentationComponent } from './gongstruct-presentation/gongstruct-presentation.component'

import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
import { GongTimeFieldDetailComponent } from './gongtimefield-detail/gongtimefield-detail.component'
import { GongTimeFieldPresentationComponent } from './gongtimefield-presentation/gongtimefield-presentation.component'

import { MetasTableComponent } from './metas-table/metas-table.component'
import { MetaDetailComponent } from './meta-detail/meta-detail.component'
import { MetaPresentationComponent } from './meta-presentation/meta-presentation.component'

import { MetaReferencesTableComponent } from './metareferences-table/metareferences-table.component'
import { MetaReferenceDetailComponent } from './metareference-detail/metareference-detail.component'
import { MetaReferencePresentationComponent } from './metareference-presentation/metareference-presentation.component'

import { ModelPkgsTableComponent } from './modelpkgs-table/modelpkgs-table.component'
import { ModelPkgDetailComponent } from './modelpkg-detail/modelpkg-detail.component'
import { ModelPkgPresentationComponent } from './modelpkg-presentation/modelpkg-presentation.component'

import { PointerToGongStructFieldsTableComponent } from './pointertogongstructfields-table/pointertogongstructfields-table.component'
import { PointerToGongStructFieldDetailComponent } from './pointertogongstructfield-detail/pointertogongstructfield-detail.component'
import { PointerToGongStructFieldPresentationComponent } from './pointertogongstructfield-presentation/pointertogongstructfield-presentation.component'

import { SliceOfPointerToGongStructFieldsTableComponent } from './sliceofpointertogongstructfields-table/sliceofpointertogongstructfields-table.component'
import { SliceOfPointerToGongStructFieldDetailComponent } from './sliceofpointertogongstructfield-detail/sliceofpointertogongstructfield-detail.component'
import { SliceOfPointerToGongStructFieldPresentationComponent } from './sliceofpointertogongstructfield-presentation/sliceofpointertogongstructfield-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfields', component: GongBasicFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-adder', component: GongBasicFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-adder/:id/:originStruct/:originStructFieldName', component: GongBasicFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-detail/:id', component: GongBasicFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-presentation/:id', component: GongBasicFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gongbasicfield-presentation-special/:id', component: GongBasicFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogongbasicfieldpres' },

	{ path: 'github_com_fullstack_lang_gong_go-gongenums', component: GongEnumsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-adder', component: GongEnumDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-adder/:id/:originStruct/:originStructFieldName', component: GongEnumDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-detail/:id', component: GongEnumDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-presentation/:id', component: GongEnumPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenum-presentation-special/:id', component: GongEnumPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogongenumpres' },

	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalues', component: GongEnumValuesTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-adder', component: GongEnumValueDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-adder/:id/:originStruct/:originStructFieldName', component: GongEnumValueDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-detail/:id', component: GongEnumValueDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-presentation/:id', component: GongEnumValuePresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gongenumvalue-presentation-special/:id', component: GongEnumValuePresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogongenumvaluepres' },

	{ path: 'github_com_fullstack_lang_gong_go-gonglinks', component: GongLinksTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-adder', component: GongLinkDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-adder/:id/:originStruct/:originStructFieldName', component: GongLinkDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-detail/:id', component: GongLinkDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-presentation/:id', component: GongLinkPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gonglink-presentation-special/:id', component: GongLinkPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogonglinkpres' },

	{ path: 'github_com_fullstack_lang_gong_go-gongnotes', component: GongNotesTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-adder', component: GongNoteDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-adder/:id/:originStruct/:originStructFieldName', component: GongNoteDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-detail/:id', component: GongNoteDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-presentation/:id', component: GongNotePresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gongnote-presentation-special/:id', component: GongNotePresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogongnotepres' },

	{ path: 'github_com_fullstack_lang_gong_go-gongstructs', component: GongStructsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-adder', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-adder/:id/:originStruct/:originStructFieldName', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-detail/:id', component: GongStructDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-presentation/:id', component: GongStructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gongstruct-presentation-special/:id', component: GongStructPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogongstructpres' },

	{ path: 'github_com_fullstack_lang_gong_go-gongtimefields', component: GongTimeFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-adder', component: GongTimeFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-adder/:id/:originStruct/:originStructFieldName', component: GongTimeFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-detail/:id', component: GongTimeFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-presentation/:id', component: GongTimeFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-gongtimefield-presentation-special/:id', component: GongTimeFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gogongtimefieldpres' },

	{ path: 'github_com_fullstack_lang_gong_go-metas', component: MetasTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-adder', component: MetaDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-adder/:id/:originStruct/:originStructFieldName', component: MetaDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-detail/:id', component: MetaDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-presentation/:id', component: MetaPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-meta-presentation-special/:id', component: MetaPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gometapres' },

	{ path: 'github_com_fullstack_lang_gong_go-metareferences', component: MetaReferencesTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-adder', component: MetaReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-adder/:id/:originStruct/:originStructFieldName', component: MetaReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-detail/:id', component: MetaReferenceDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-presentation/:id', component: MetaReferencePresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-metareference-presentation-special/:id', component: MetaReferencePresentationComponent, outlet: 'github_com_fullstack_lang_gong_gometareferencepres' },

	{ path: 'github_com_fullstack_lang_gong_go-modelpkgs', component: ModelPkgsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-adder', component: ModelPkgDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-adder/:id/:originStruct/:originStructFieldName', component: ModelPkgDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-detail/:id', component: ModelPkgDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-presentation/:id', component: ModelPkgPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-modelpkg-presentation-special/:id', component: ModelPkgPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gomodelpkgpres' },

	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfields', component: PointerToGongStructFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-adder', component: PointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-adder/:id/:originStruct/:originStructFieldName', component: PointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-detail/:id', component: PointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-presentation/:id', component: PointerToGongStructFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-pointertogongstructfield-presentation-special/:id', component: PointerToGongStructFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gopointertogongstructfieldpres' },

	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfields', component: SliceOfPointerToGongStructFieldsTableComponent, outlet: 'github_com_fullstack_lang_gong_go_table' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-adder', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-adder/:id/:originStruct/:originStructFieldName', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-detail/:id', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'github_com_fullstack_lang_gong_go_editor' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-presentation/:id', component: SliceOfPointerToGongStructFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_go_presentation' },
	{ path: 'github_com_fullstack_lang_gong_go-sliceofpointertogongstructfield-presentation-special/:id', component: SliceOfPointerToGongStructFieldPresentationComponent, outlet: 'github_com_fullstack_lang_gong_gosliceofpointertogongstructfieldpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
