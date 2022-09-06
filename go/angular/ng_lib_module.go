package angular

const NgLibModuleTemplate = `import { NgModule } from '@angular/core';

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
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports {{` + string(NgLibModuleImports) + `}}

@NgModule({
	declarations: [
		// insertion point for declarations {{` + string(NgLibModuleDeclarations) + `}}

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
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations {{` + string(NgLibModuleDeclarations) + `}}

		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class {{PkgName}}Module { }
`

type NgLibModuleServiceSubTemplate string

const (
	NgLibModuleDeclarations NgLibModuleServiceSubTemplate = "NgLibModuleIndivDecls"
	NgLibModuleImports      NgLibModuleServiceSubTemplate = "NgLibModuleImports"
)

var NgLibModuleSubTemplateCode map[string]string = // new line
map[string]string{

	string(NgLibModuleImports): `
import { {{Structname}}sTableComponent } from './{{structname}}s-table/{{structname}}s-table.component'
import { {{Structname}}SortingComponent } from './{{structname}}-sorting/{{structname}}-sorting.component'
import { {{Structname}}DetailComponent } from './{{structname}}-detail/{{structname}}-detail.component'
import { {{Structname}}PresentationComponent } from './{{structname}}-presentation/{{structname}}-presentation.component'
`,

	string(NgLibModuleDeclarations): `
		{{Structname}}sTableComponent,
		{{Structname}}SortingComponent,
		{{Structname}}DetailComponent,
		{{Structname}}PresentationComponent,
`,
}
