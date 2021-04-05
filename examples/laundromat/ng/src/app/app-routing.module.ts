import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import * as laundromat from 'laundromat'

import { UmlscDiagramComponent } from 'gongdocdiagrams';

export enum Paths {
  MACHINE_PRESENTATION_SPECIAL_PATH = "machine-presentation-special",
  WASHER_PRESENTATION_SPECIAL_PATH = "washer-presentation-special",
  UMLSC_WASHER_STATES_PATH = "umlsc-washer-diagram",
  UMLSC_MACHINE_STATES_PATH = "umlsc-machine-diagram"
}

const routes: Routes = [
	{ path: Paths.MACHINE_PRESENTATION_SPECIAL_PATH+'/:id', component: laundromat.MachinePresentationComponent, outlet: 'machinepres' },
	{ path: Paths.WASHER_PRESENTATION_SPECIAL_PATH+'/:id', component: laundromat.WasherPresentationComponent, outlet: 'washerpres' },

  { path: Paths.UMLSC_WASHER_STATES_PATH+'/:id', component: UmlscDiagramComponent, outlet: 'washerstates' },
  { path: Paths.UMLSC_MACHINE_STATES_PATH+'/:id', component: UmlscDiagramComponent, outlet: 'machinestates' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
