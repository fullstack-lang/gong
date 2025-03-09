package angular

const NgFileAppComponentTs = `import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as {{pkgname}} from '../../projects/{{pkgname}}/src/public-api'

import { {{TitlePkgName}}SpecificComponent } from '../../projects/{{pkgname}}specific/src/lib/{{pkgname}}-specific/{{pkgname}}-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    {{TitlePkgName}}SpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = "{{pkgname}}"

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
`
