import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

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


@NgModule({
	declarations: [
		// insertion point for declarations 
		GongBasicFieldsTableComponent,
		GongBasicFieldDetailComponent,
		GongBasicFieldPresentationComponent,

		GongEnumsTableComponent,
		GongEnumDetailComponent,
		GongEnumPresentationComponent,

		GongEnumValuesTableComponent,
		GongEnumValueDetailComponent,
		GongEnumValuePresentationComponent,

		GongStructsTableComponent,
		GongStructDetailComponent,
		GongStructPresentationComponent,

		GongTimeFieldsTableComponent,
		GongTimeFieldDetailComponent,
		GongTimeFieldPresentationComponent,

		ModelPkgsTableComponent,
		ModelPkgDetailComponent,
		ModelPkgPresentationComponent,

		PointerToGongStructFieldsTableComponent,
		PointerToGongStructFieldDetailComponent,
		PointerToGongStructFieldPresentationComponent,

		SliceOfPointerToGongStructFieldsTableComponent,
		SliceOfPointerToGongStructFieldDetailComponent,
		SliceOfPointerToGongStructFieldPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		GongBasicFieldsTableComponent,
		GongBasicFieldDetailComponent,
		GongBasicFieldPresentationComponent,

		GongEnumsTableComponent,
		GongEnumDetailComponent,
		GongEnumPresentationComponent,

		GongEnumValuesTableComponent,
		GongEnumValueDetailComponent,
		GongEnumValuePresentationComponent,

		GongStructsTableComponent,
		GongStructDetailComponent,
		GongStructPresentationComponent,

		GongTimeFieldsTableComponent,
		GongTimeFieldDetailComponent,
		GongTimeFieldPresentationComponent,

		ModelPkgsTableComponent,
		ModelPkgDetailComponent,
		ModelPkgPresentationComponent,

		PointerToGongStructFieldsTableComponent,
		PointerToGongStructFieldDetailComponent,
		PointerToGongStructFieldPresentationComponent,

		SliceOfPointerToGongStructFieldsTableComponent,
		SliceOfPointerToGongStructFieldDetailComponent,
		SliceOfPointerToGongStructFieldPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongModule { }
