import { Component, Input, OnInit } from '@angular/core';

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import { PkgeltDocsComponent } from '../pkgelt-docs/pkgelt-docs.component'

@Component({
  selector: 'lib-panel',
  templateUrl: './panel.component.html',
  styleUrls: ['./panel.component.css'],
  standalone: true,
  imports: [
    MatRadioModule,
    FormsModule,
    CommonModule,
    MatButtonModule,
    MatIconModule,

    AngularSplitModule,

    PkgeltDocsComponent
  ],
})
export class PanelComponent implements OnInit {

  // choices for the top radio button
  view = 'Diagrams'
  diagrams = 'Diagrams'
  views: string[] = [this.diagrams];

  stacks: string[] = []

  loading = true

  @Input() GONG__StackPath: string = ""

  constructor(
  ) {

  }

  ngOnInit(): void {

  }

  refresh() {
  }
}
