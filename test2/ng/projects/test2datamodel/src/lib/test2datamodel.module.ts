import { NgModule } from '@angular/core';
import { Test2datamodelComponent } from './test2datamodel.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';



@NgModule({
  declarations: [
    Test2datamodelComponent,
    DataModelPanelComponent
  ],
  imports: [
  ],
  exports: [
    Test2datamodelComponent
  ]
})
export class Test2datamodelModule { }
