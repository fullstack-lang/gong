import { Component, OnInit } from '@angular/core';

import { Router, RouterState } from '@angular/router';

@Component({
  selector: 'lib-pkgelt-docs',
  templateUrl: './pkgelt-docs.component.html',
  styleUrls: ['./pkgelt-docs.component.css']
})
export class PkgeltDocsComponent implements OnInit {

  constructor(
    private router: Router,
  ) { }

  ngOnInit(): void {

    this.setEditorRouterOutlet()
  }

  // set editor outlet
  setEditorRouterOutlet() {

    console.log("pkgElt setEditorRouterOutlet ")

    this.router.navigate([{
      outlets: {
        diagrameditor: ["classdiagram-detail"]
      }
    }]).catch(
      reason => {
        console.log(reason)
      }
    );
  }

}
