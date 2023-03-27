import { Component, Input, OnInit } from '@angular/core';

import * as gongdoc from 'gongdoc';

@Component({
  selector: 'lib-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css']
})
export class PanelComponent implements OnInit {

  // choices for the top radio button
  view = 'Default view'
  default = 'Default view'
  meta = 'Meta view'
  gong = 'Gong view'
  views: string[] = [this.default, this.meta, this.gong];

  stacks: string[] = []

  loading = true

  @Input() GONG__StackPath: string = ""

  diagramPackage: gongdoc.DiagramPackageDB = new (gongdoc.DiagramPackageDB);

  constructor(
    private diagramPackageService: gongdoc.DiagramPackageService,
  ) {

  }

  ngOnInit(): void {
    // create a new GongDoc instance
    this.diagramPackageService.getDiagramPackages(this.GONG__StackPath).subscribe(

      diagramPackages => {

        this.diagramPackage = diagramPackages[0];
        console.log("PanelComponent", this.diagramPackage.Name)

        this.loading = false
      })
  }

  refresh() {
    // refresh the view
    this.diagramPackage.IsReloaded = true
    this.diagramPackageService.updateDiagramPackage(this.diagramPackage, "").subscribe(
      diagramPackage => {
        console.log('diagram package refreshed')
      })
  }
}
