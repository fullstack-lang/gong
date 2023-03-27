import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ClassDiagramComponent } from './class-diagram/class-diagram.component'

import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'

const routes: Routes = [
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
