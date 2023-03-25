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

    // the class diagram editor is setup with a component that has no data
    // this component will fetch the new commits from the back
    this.setEditorRouterOutlet()
    console.log("PkgeltDocsComponent->GONG__StackPath : ", this.GONG__StackPath)
  }

  // set editor outlet
  setEditorRouterOutlet() {

    console.log("pkgElt setEditorRouterOutlet with GONG__StackPath", this.GONG__StackPath)

    this.router.navigate([{
      outlets: {
        diagrameditor: ["classdiagram-editor", this.GONG__StackPath]
      }
    }]).catch(
      reason => {
        console.log(reason)
      }
    );
  }

}
