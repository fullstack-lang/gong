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
import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EngineDetailComponent } from './engine-detail/engine-detail.component'
import { EnginePresentationComponent } from './engine-presentation/engine-presentation.component'

import { EventsTableComponent } from './events-table/events-table.component'
import { EventDetailComponent } from './event-detail/event-detail.component'
import { EventPresentationComponent } from './event-presentation/event-presentation.component'

import { GongsimCommandsTableComponent } from './gongsimcommands-table/gongsimcommands-table.component'
import { GongsimCommandDetailComponent } from './gongsimcommand-detail/gongsimcommand-detail.component'
import { GongsimCommandPresentationComponent } from './gongsimcommand-presentation/gongsimcommand-presentation.component'

import { GongsimStatussTableComponent } from './gongsimstatuss-table/gongsimstatuss-table.component'
import { GongsimStatusDetailComponent } from './gongsimstatus-detail/gongsimstatus-detail.component'
import { GongsimStatusPresentationComponent } from './gongsimstatus-presentation/gongsimstatus-presentation.component'

import { UpdateStatesTableComponent } from './updatestates-table/updatestates-table.component'
import { UpdateStateDetailComponent } from './updatestate-detail/updatestate-detail.component'
import { UpdateStatePresentationComponent } from './updatestate-presentation/updatestate-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		EnginesTableComponent,
		EngineDetailComponent,
		EnginePresentationComponent,

		EventsTableComponent,
		EventDetailComponent,
		EventPresentationComponent,

		GongsimCommandsTableComponent,
		GongsimCommandDetailComponent,
		GongsimCommandPresentationComponent,

		GongsimStatussTableComponent,
		GongsimStatusDetailComponent,
		GongsimStatusPresentationComponent,

		UpdateStatesTableComponent,
		UpdateStateDetailComponent,
		UpdateStatePresentationComponent,


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
		EnginesTableComponent,
		EngineDetailComponent,
		EnginePresentationComponent,

		EventsTableComponent,
		EventDetailComponent,
		EventPresentationComponent,

		GongsimCommandsTableComponent,
		GongsimCommandDetailComponent,
		GongsimCommandPresentationComponent,

		GongsimStatussTableComponent,
		GongsimStatusDetailComponent,
		GongsimStatusPresentationComponent,

		UpdateStatesTableComponent,
		UpdateStateDetailComponent,
		UpdateStatePresentationComponent,


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
export class GongsimModule { }
