import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-github_com_fullstack_lang_gong_go-data-model-panel',
  templateUrl: './data-model-panel.component.html',
})
export class DataModelPanelComponent {

  view = 'Data'
  default = 'Data'
  model = 'Model'

  views: string[] = [this.default, this.model]

  @Input() GONG__StackPath: string = ""

}
