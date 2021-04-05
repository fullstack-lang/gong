import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { MachinesTableComponent } from './machines-table/machines-table.component'
import { MachineDetailComponent } from './machine-detail/machine-detail.component'
import { MachinePresentationComponent } from './machine-presentation/machine-presentation.component'

import { SimulationsTableComponent } from './simulations-table/simulations-table.component'
import { SimulationDetailComponent } from './simulation-detail/simulation-detail.component'
import { SimulationPresentationComponent } from './simulation-presentation/simulation-presentation.component'

import { WashersTableComponent } from './washers-table/washers-table.component'
import { WasherDetailComponent } from './washer-detail/washer-detail.component'
import { WasherPresentationComponent } from './washer-presentation/washer-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'machines', component: MachinesTableComponent, outlet: 'table' },
	{ path: 'machine-adder', component: MachineDetailComponent, outlet: 'editor' },
	{ path: 'machine-adder/:id/:association', component: MachineDetailComponent, outlet: 'editor' },
	{ path: 'machine-detail/:id', component: MachineDetailComponent, outlet: 'editor' },
	{ path: 'machine-presentation/:id', component: MachinePresentationComponent, outlet: 'presentation' },
	{ path: 'machine-presentation-special/:id', component: MachinePresentationComponent, outlet: 'machinepres' },

	{ path: 'simulations', component: SimulationsTableComponent, outlet: 'table' },
	{ path: 'simulation-adder', component: SimulationDetailComponent, outlet: 'editor' },
	{ path: 'simulation-adder/:id/:association', component: SimulationDetailComponent, outlet: 'editor' },
	{ path: 'simulation-detail/:id', component: SimulationDetailComponent, outlet: 'editor' },
	{ path: 'simulation-presentation/:id', component: SimulationPresentationComponent, outlet: 'presentation' },
	{ path: 'simulation-presentation-special/:id', component: SimulationPresentationComponent, outlet: 'simulationpres' },

	{ path: 'washers', component: WashersTableComponent, outlet: 'table' },
	{ path: 'washer-adder', component: WasherDetailComponent, outlet: 'editor' },
	{ path: 'washer-adder/:id/:association', component: WasherDetailComponent, outlet: 'editor' },
	{ path: 'washer-detail/:id', component: WasherDetailComponent, outlet: 'editor' },
	{ path: 'washer-presentation/:id', component: WasherPresentationComponent, outlet: 'presentation' },
	{ path: 'washer-presentation-special/:id', component: WasherPresentationComponent, outlet: 'washerpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
