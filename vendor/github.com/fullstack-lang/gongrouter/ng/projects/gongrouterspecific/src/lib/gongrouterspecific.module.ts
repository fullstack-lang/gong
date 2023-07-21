import { NgModule } from '@angular/core';
import { GongrouterspecificComponent } from './gongrouterspecific.component';

import { Routes, RouterModule } from '@angular/router';
import { GongrouterOutletComponent } from './gongrouter-outlet/gongrouter-outlet.component';


@NgModule({
  declarations: [
    GongrouterspecificComponent,
    GongrouterOutletComponent
  ],
  imports: [
    RouterModule,
  ],
  exports: [
    GongrouterspecificComponent,
    GongrouterOutletComponent
  ]
})
export class GongrouterspecificModule { }
