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

    // this.setEditorRouterOutlet(1)
  }

  // set editor outlet
  setEditorRouterOutlet(classdiagramID: number) {

    console.log("pkgElt setEditorRouterOutlet " + classdiagramID)

    this.router.navigate([{
      outlets: {
        diagrameditor: ["classdiagram-detail", classdiagramID]
      }
    }]).catch(
      reason => {
        console.log(reason)
      }
    );
  }

}
