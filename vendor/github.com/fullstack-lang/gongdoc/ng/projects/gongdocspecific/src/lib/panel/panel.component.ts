import { Component, Input, OnInit } from '@angular/core';

import * as gongdoc from 'gongdoc';

@Component({
  selector: 'lib-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css']
})
export class PanelComponent implements OnInit {

  // choices for the top radio button
  view = 'Diagrams'
  diagrams = 'Diagrams'
  diagrams_data = 'Diagrams data'
  model_data = 'Model data'
  svg_data = 'SVG data'
  tree_data = 'Tree data'
  views: string[] = [this.diagrams, 
    this.diagrams_data, 
    this.model_data, 
    this.svg_data, 
    this.tree_data];

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
    this.diagramPackageService.updateDiagramPackage(this.diagramPackage, this.GONG__StackPath).subscribe(
      diagramPackage => {
        console.log('diagram package refreshed')
      })
  }
}
