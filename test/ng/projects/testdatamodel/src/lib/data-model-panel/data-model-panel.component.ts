import { Component, Input, OnInit } from '@angular/core';
import * as test from 'test';

@Component({
  selector: 'lib-github_com_fullstack_lang_gong_test_go-data-model-panel',
  templateUrl: './data-model-panel.component.html',
})
export class DataModelPanelComponent implements OnInit {

  data = 'Data'
  dataNewNew = 'Data New New'
  tableData = "Table Data"
  view = this.data
  model = 'Model'

  views: string[] = [this.data, this.dataNewNew, this.tableData, this.model]

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

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  constructor(
    private routeService: test.RouteService,
  ) { }

  // this component relies on the the gongtree stack to enable
  // edit of the data
  @Input() GONG__DATA__StackPath: string = ""

  // this component relies on several stacks (gong, gongdoc, gongsvg, ...) for the modeling of the stack
  // a gongdoc stack and a gong stack. All have the same path defined by
  // GONG__MODEL_StackPath
  @Input() GONG__MODEL__StacksPath: string = ""

  tableOutletName = ""
  editorOutletName = ""

  ngOnInit(): void {

    // add the routes that will used by router compnents
    this.routeService.addDataPanelRoutes(this.GONG__DATA__StackPath)
    this.tableOutletName = this.routeService.getTableOutlet(this.GONG__DATA__StackPath)
    this.editorOutletName = this.routeService.getEditorOutlet(this.GONG__DATA__StackPath)
  }
}
