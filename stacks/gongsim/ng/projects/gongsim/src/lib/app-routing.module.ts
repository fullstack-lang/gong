import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

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


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'engines', component: EnginesTableComponent, outlet: 'table' },
	{ path: 'engine-adder', component: EngineDetailComponent, outlet: 'editor' },
	{ path: 'engine-adder/:id/:association', component: EngineDetailComponent, outlet: 'editor' },
	{ path: 'engine-detail/:id', component: EngineDetailComponent, outlet: 'editor' },
	{ path: 'engine-presentation/:id', component: EnginePresentationComponent, outlet: 'presentation' },
	{ path: 'engine-presentation-special/:id', component: EnginePresentationComponent, outlet: 'enginepres' },

	{ path: 'events', component: EventsTableComponent, outlet: 'table' },
	{ path: 'event-adder', component: EventDetailComponent, outlet: 'editor' },
	{ path: 'event-adder/:id/:association', component: EventDetailComponent, outlet: 'editor' },
	{ path: 'event-detail/:id', component: EventDetailComponent, outlet: 'editor' },
	{ path: 'event-presentation/:id', component: EventPresentationComponent, outlet: 'presentation' },
	{ path: 'event-presentation-special/:id', component: EventPresentationComponent, outlet: 'eventpres' },

	{ path: 'gongsimcommands', component: GongsimCommandsTableComponent, outlet: 'table' },
	{ path: 'gongsimcommand-adder', component: GongsimCommandDetailComponent, outlet: 'editor' },
	{ path: 'gongsimcommand-adder/:id/:association', component: GongsimCommandDetailComponent, outlet: 'editor' },
	{ path: 'gongsimcommand-detail/:id', component: GongsimCommandDetailComponent, outlet: 'editor' },
	{ path: 'gongsimcommand-presentation/:id', component: GongsimCommandPresentationComponent, outlet: 'presentation' },
	{ path: 'gongsimcommand-presentation-special/:id', component: GongsimCommandPresentationComponent, outlet: 'gongsimcommandpres' },

	{ path: 'gongsimstatuss', component: GongsimStatussTableComponent, outlet: 'table' },
	{ path: 'gongsimstatus-adder', component: GongsimStatusDetailComponent, outlet: 'editor' },
	{ path: 'gongsimstatus-adder/:id/:association', component: GongsimStatusDetailComponent, outlet: 'editor' },
	{ path: 'gongsimstatus-detail/:id', component: GongsimStatusDetailComponent, outlet: 'editor' },
	{ path: 'gongsimstatus-presentation/:id', component: GongsimStatusPresentationComponent, outlet: 'presentation' },
	{ path: 'gongsimstatus-presentation-special/:id', component: GongsimStatusPresentationComponent, outlet: 'gongsimstatuspres' },

	{ path: 'updatestates', component: UpdateStatesTableComponent, outlet: 'table' },
	{ path: 'updatestate-adder', component: UpdateStateDetailComponent, outlet: 'editor' },
	{ path: 'updatestate-adder/:id/:association', component: UpdateStateDetailComponent, outlet: 'editor' },
	{ path: 'updatestate-detail/:id', component: UpdateStateDetailComponent, outlet: 'editor' },
	{ path: 'updatestate-presentation/:id', component: UpdateStatePresentationComponent, outlet: 'presentation' },
	{ path: 'updatestate-presentation-special/:id', component: UpdateStatePresentationComponent, outlet: 'updatestatepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
