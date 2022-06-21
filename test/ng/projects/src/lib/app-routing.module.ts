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
	{ path: '-astructs', component: AstructsTableComponent, outlet: '_table' },
	{ path: '-astruct-adder', component: AstructDetailComponent, outlet: '_editor' },
	{ path: '-astruct-adder/:id/:originStruct/:originStructFieldName', component: AstructDetailComponent, outlet: '_editor' },
	{ path: '-astruct-detail/:id', component: AstructDetailComponent, outlet: '_editor' },
	{ path: '-astruct-presentation/:id', component: AstructPresentationComponent, outlet: '_presentation' },
	{ path: '-astruct-presentation-special/:id', component: AstructPresentationComponent, outlet: 'astructpres' },

	{ path: '-astructbstruct2uses', component: AstructBstruct2UsesTableComponent, outlet: '_table' },
	{ path: '-astructbstruct2use-adder', component: AstructBstruct2UseDetailComponent, outlet: '_editor' },
	{ path: '-astructbstruct2use-adder/:id/:originStruct/:originStructFieldName', component: AstructBstruct2UseDetailComponent, outlet: '_editor' },
	{ path: '-astructbstruct2use-detail/:id', component: AstructBstruct2UseDetailComponent, outlet: '_editor' },
	{ path: '-astructbstruct2use-presentation/:id', component: AstructBstruct2UsePresentationComponent, outlet: '_presentation' },
	{ path: '-astructbstruct2use-presentation-special/:id', component: AstructBstruct2UsePresentationComponent, outlet: 'astructbstruct2usepres' },

	{ path: '-astructbstructuses', component: AstructBstructUsesTableComponent, outlet: '_table' },
	{ path: '-astructbstructuse-adder', component: AstructBstructUseDetailComponent, outlet: '_editor' },
	{ path: '-astructbstructuse-adder/:id/:originStruct/:originStructFieldName', component: AstructBstructUseDetailComponent, outlet: '_editor' },
	{ path: '-astructbstructuse-detail/:id', component: AstructBstructUseDetailComponent, outlet: '_editor' },
	{ path: '-astructbstructuse-presentation/:id', component: AstructBstructUsePresentationComponent, outlet: '_presentation' },
	{ path: '-astructbstructuse-presentation-special/:id', component: AstructBstructUsePresentationComponent, outlet: 'astructbstructusepres' },

	{ path: '-bstructs', component: BstructsTableComponent, outlet: '_table' },
	{ path: '-bstruct-adder', component: BstructDetailComponent, outlet: '_editor' },
	{ path: '-bstruct-adder/:id/:originStruct/:originStructFieldName', component: BstructDetailComponent, outlet: '_editor' },
	{ path: '-bstruct-detail/:id', component: BstructDetailComponent, outlet: '_editor' },
	{ path: '-bstruct-presentation/:id', component: BstructPresentationComponent, outlet: '_presentation' },
	{ path: '-bstruct-presentation-special/:id', component: BstructPresentationComponent, outlet: 'bstructpres' },

	{ path: '-dstructs', component: DstructsTableComponent, outlet: '_table' },
	{ path: '-dstruct-adder', component: DstructDetailComponent, outlet: '_editor' },
	{ path: '-dstruct-adder/:id/:originStruct/:originStructFieldName', component: DstructDetailComponent, outlet: '_editor' },
	{ path: '-dstruct-detail/:id', component: DstructDetailComponent, outlet: '_editor' },
	{ path: '-dstruct-presentation/:id', component: DstructPresentationComponent, outlet: '_presentation' },
	{ path: '-dstruct-presentation-special/:id', component: DstructPresentationComponent, outlet: 'dstructpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
