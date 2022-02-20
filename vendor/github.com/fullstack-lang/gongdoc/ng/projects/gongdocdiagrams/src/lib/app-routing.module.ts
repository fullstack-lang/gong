import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ClassDiagramComponent } from './class-diagram/class-diagram.component'
import { UmlscDiagramComponent } from './umlsc-diagram/umlsc-diagram.component'

import { ClassdiagramDetailComponent } from './classdiagram-detail/classdiagram-detail.component'
import { ClassdiagramPresentationComponent } from 'gongdoc'

const routes: Routes = [

  { path: 'classdiagram-detail/:id', component: ClassDiagramComponent, outlet: 'diagrameditor' },
  { path: 'umlsc-detail/:id', component: UmlscDiagramComponent, outlet: 'diagrameditor' },
  { path: 'github_com_fullstack_lang_gongdoc_go-classdiagram-adder', component: ClassdiagramDetailComponent, outlet: 'elementeditor' },
  { path: 'classdiagram-detail/:id', component: ClassdiagramDetailComponent, outlet: 'elementeditor' },
  { path: 'classdiagram-presentation/:id', component: ClassdiagramPresentationComponent, outlet: 'presentation' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
