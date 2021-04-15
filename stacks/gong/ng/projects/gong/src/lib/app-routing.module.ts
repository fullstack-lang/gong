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

import { GongStructsTableComponent } from './gongstructs-table/gongstructs-table.component'
import { GongStructDetailComponent } from './gongstruct-detail/gongstruct-detail.component'
import { GongStructPresentationComponent } from './gongstruct-presentation/gongstruct-presentation.component'

import { GongTimeFieldsTableComponent } from './gongtimefields-table/gongtimefields-table.component'
import { GongTimeFieldDetailComponent } from './gongtimefield-detail/gongtimefield-detail.component'
import { GongTimeFieldPresentationComponent } from './gongtimefield-presentation/gongtimefield-presentation.component'

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
	{ path: 'gongbasicfields', component: GongBasicFieldsTableComponent, outlet: 'table' },
	{ path: 'gongbasicfield-adder', component: GongBasicFieldDetailComponent, outlet: 'editor' },
	{ path: 'gongbasicfield-adder/:id/:association', component: GongBasicFieldDetailComponent, outlet: 'editor' },
	{ path: 'gongbasicfield-detail/:id', component: GongBasicFieldDetailComponent, outlet: 'editor' },
	{ path: 'gongbasicfield-presentation/:id', component: GongBasicFieldPresentationComponent, outlet: 'presentation' },
	{ path: 'gongbasicfield-presentation-special/:id', component: GongBasicFieldPresentationComponent, outlet: 'gongbasicfieldpres' },

	{ path: 'gongenums', component: GongEnumsTableComponent, outlet: 'table' },
	{ path: 'gongenum-adder', component: GongEnumDetailComponent, outlet: 'editor' },
	{ path: 'gongenum-adder/:id/:association', component: GongEnumDetailComponent, outlet: 'editor' },
	{ path: 'gongenum-detail/:id', component: GongEnumDetailComponent, outlet: 'editor' },
	{ path: 'gongenum-presentation/:id', component: GongEnumPresentationComponent, outlet: 'presentation' },
	{ path: 'gongenum-presentation-special/:id', component: GongEnumPresentationComponent, outlet: 'gongenumpres' },

	{ path: 'gongenumvalues', component: GongEnumValuesTableComponent, outlet: 'table' },
	{ path: 'gongenumvalue-adder', component: GongEnumValueDetailComponent, outlet: 'editor' },
	{ path: 'gongenumvalue-adder/:id/:association', component: GongEnumValueDetailComponent, outlet: 'editor' },
	{ path: 'gongenumvalue-detail/:id', component: GongEnumValueDetailComponent, outlet: 'editor' },
	{ path: 'gongenumvalue-presentation/:id', component: GongEnumValuePresentationComponent, outlet: 'presentation' },
	{ path: 'gongenumvalue-presentation-special/:id', component: GongEnumValuePresentationComponent, outlet: 'gongenumvaluepres' },

	{ path: 'gongstructs', component: GongStructsTableComponent, outlet: 'table' },
	{ path: 'gongstruct-adder', component: GongStructDetailComponent, outlet: 'editor' },
	{ path: 'gongstruct-adder/:id/:association', component: GongStructDetailComponent, outlet: 'editor' },
	{ path: 'gongstruct-detail/:id', component: GongStructDetailComponent, outlet: 'editor' },
	{ path: 'gongstruct-presentation/:id', component: GongStructPresentationComponent, outlet: 'presentation' },
	{ path: 'gongstruct-presentation-special/:id', component: GongStructPresentationComponent, outlet: 'gongstructpres' },

	{ path: 'gongtimefields', component: GongTimeFieldsTableComponent, outlet: 'table' },
	{ path: 'gongtimefield-adder', component: GongTimeFieldDetailComponent, outlet: 'editor' },
	{ path: 'gongtimefield-adder/:id/:association', component: GongTimeFieldDetailComponent, outlet: 'editor' },
	{ path: 'gongtimefield-detail/:id', component: GongTimeFieldDetailComponent, outlet: 'editor' },
	{ path: 'gongtimefield-presentation/:id', component: GongTimeFieldPresentationComponent, outlet: 'presentation' },
	{ path: 'gongtimefield-presentation-special/:id', component: GongTimeFieldPresentationComponent, outlet: 'gongtimefieldpres' },

	{ path: 'modelpkgs', component: ModelPkgsTableComponent, outlet: 'table' },
	{ path: 'modelpkg-adder', component: ModelPkgDetailComponent, outlet: 'editor' },
	{ path: 'modelpkg-adder/:id/:association', component: ModelPkgDetailComponent, outlet: 'editor' },
	{ path: 'modelpkg-detail/:id', component: ModelPkgDetailComponent, outlet: 'editor' },
	{ path: 'modelpkg-presentation/:id', component: ModelPkgPresentationComponent, outlet: 'presentation' },
	{ path: 'modelpkg-presentation-special/:id', component: ModelPkgPresentationComponent, outlet: 'modelpkgpres' },

	{ path: 'pointertogongstructfields', component: PointerToGongStructFieldsTableComponent, outlet: 'table' },
	{ path: 'pointertogongstructfield-adder', component: PointerToGongStructFieldDetailComponent, outlet: 'editor' },
	{ path: 'pointertogongstructfield-adder/:id/:association', component: PointerToGongStructFieldDetailComponent, outlet: 'editor' },
	{ path: 'pointertogongstructfield-detail/:id', component: PointerToGongStructFieldDetailComponent, outlet: 'editor' },
	{ path: 'pointertogongstructfield-presentation/:id', component: PointerToGongStructFieldPresentationComponent, outlet: 'presentation' },
	{ path: 'pointertogongstructfield-presentation-special/:id', component: PointerToGongStructFieldPresentationComponent, outlet: 'pointertogongstructfieldpres' },

	{ path: 'sliceofpointertogongstructfields', component: SliceOfPointerToGongStructFieldsTableComponent, outlet: 'table' },
	{ path: 'sliceofpointertogongstructfield-adder', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'editor' },
	{ path: 'sliceofpointertogongstructfield-adder/:id/:association', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'editor' },
	{ path: 'sliceofpointertogongstructfield-detail/:id', component: SliceOfPointerToGongStructFieldDetailComponent, outlet: 'editor' },
	{ path: 'sliceofpointertogongstructfield-presentation/:id', component: SliceOfPointerToGongStructFieldPresentationComponent, outlet: 'presentation' },
	{ path: 'sliceofpointertogongstructfield-presentation-special/:id', component: SliceOfPointerToGongStructFieldPresentationComponent, outlet: 'sliceofpointertogongstructfieldpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
