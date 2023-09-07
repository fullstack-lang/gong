import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gong from 'gong'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gong Data/Model'
  view = this.default

  views: string[] = [this.default];

  DataStack = "gong"
  ModelStacks = "github.com/fullstack-lang/gong/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
