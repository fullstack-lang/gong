import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-github_com_fullstack_lang_gong_go-data-model-panel',
  templateUrl: './data-model-panel.component.html',
})
export class DataModelPanelComponent {

  view = 'Data'
  default = this.view
  model = 'Model'

  views: string[] = [this.default, this.model]

  textStyle = {
    'color': 'rgba(0, 0, 0, 0.87)', // These are just examples.
    'font-size': '14px' // Adjust them according to your needs.
  }

  containerStyle = {
    'display': 'flex',
    // 'justify-content': 'space-between',
    'align-items': 'center',
    'width': '100%' // Adjust as necessary.
  }

  radioGroupStyle = {
    'display': 'flex',
    'flex-direction': 'row',
    'align-items': 'center',
    'justify-content': 'flex-start'
  }

  @Input() GONG__StackPath: string = ""

}
