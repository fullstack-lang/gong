import { NgModule } from '@angular/core';
import { TestspecificComponent } from './testspecific.component';
import { TestModule } from 'test';



@NgModule({
  declarations: [
    TestspecificComponent
  ],
  imports: [
    TestModule,
  ],
  exports: [
    TestspecificComponent
  ]
})
export class TestspecificModule { }
