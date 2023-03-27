import { Component, Input, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

@Component({
  selector: 'lib-pkgelt-docs',
  templateUrl: './pkgelt-docs.component.html',
  styleUrls: ['./pkgelt-docs.component.css']
})
export class PkgeltDocsComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {
    console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }
}
