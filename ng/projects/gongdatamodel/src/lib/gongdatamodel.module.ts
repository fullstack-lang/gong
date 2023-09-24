import { NgModule } from '@angular/core';
import { GongdatamodelComponent } from './gongdatamodel.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';



@NgModule({
  declarations: [
    GongdatamodelComponent,
    DataModelPanelComponent
  ],
  imports: [
  ],
  exports: [
    GongdatamodelComponent
  ]
})
export class GongdatamodelModule { }
