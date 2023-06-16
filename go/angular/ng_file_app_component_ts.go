package angular

const NgFileAppComponentTs = `import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as {{pkgname}} from '{{pkgname}}'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = '{{TitlePkgName}} Data/Model'
  view = this.default

  views: string[] = [this.default];

  DataStack = "{{pkgname}}"
  ModelStacks = "{{PkgPathRoot}}/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
`
