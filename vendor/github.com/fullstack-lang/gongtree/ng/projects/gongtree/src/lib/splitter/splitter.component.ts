import { Component, Input, OnInit } from '@angular/core';

import { RouteService } from '../route-service';

@Component({
  selector: 'app-gongtree-splitter',
  templateUrl: './splitter.component.html',
  styleUrls: ['./splitter.component.css'],
})
export class SplitterComponent implements OnInit {

  @Input() GONG__StackPath: string = ""
  tableOutletName: string = ""
  editorOutletName: string = ""

  constructor(private routeService: RouteService) { }

  ngOnInit(): void {
    console.log("Splitter: " + this.GONG__StackPath)

    this.tableOutletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    this.editorOutletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
  }
}
