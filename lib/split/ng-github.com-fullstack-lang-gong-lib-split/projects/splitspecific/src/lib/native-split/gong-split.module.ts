import { NgModule } from '@angular/core';
import { GongSplitComponent } from './gong-split.component';
import { GongSplitAreaComponent } from './gong-split-area.component';

export const GONG_SPLIT_DIRECTIVES = [
  GongSplitComponent,
  GongSplitAreaComponent
] as const;

@NgModule({
  imports: [GongSplitComponent, GongSplitAreaComponent],
  exports: [GongSplitComponent, GongSplitAreaComponent]
})
export class GongSplitModule {}
