import { Component, Input } from '@angular/core';

@Component({
  selector: 'lib-github_com_fullstack_lang_gongtree_go-data-model-panel',
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

  // this component relies on the the gongtree stack to enable
  // edit of the data
  @Input() GONG__DATA__StackPath: string = ""

  // this component relies on several stacks (gong, gongdoc, gongsvg, ...) for the modeling of the stack
  // a gongdoc stack and a gong stack. All have the same path defined by
  // GONG__MODEL_StackPath
  @Input() GONG__MODEL__StacksPath: string = ""

}
