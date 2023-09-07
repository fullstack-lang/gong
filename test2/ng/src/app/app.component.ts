import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as test2 from 'test2'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Test2 Data/Model'
  view = this.default

  views: string[] = [this.default];

  DataStack = "test2"
  ModelStacks = "github.com/fullstack-lang/gong/test2/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
