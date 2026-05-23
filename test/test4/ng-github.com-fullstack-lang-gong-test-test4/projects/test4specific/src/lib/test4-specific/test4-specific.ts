import { Component, Input, OnInit } from '@angular/core';

import * as test4 from '../../../../test4/src/public-api'


import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'lib-test4-specific',
  imports: [
    MatButtonModule,
  ],
  templateUrl: './test4-specific.html',
  styleUrl: './test4-specific.css'
})
export class Test4Specific implements OnInit {


  @Input() Name: string = ""

  public frontRepo?: test4.FrontRepo;

  constructor(
    private frontRepoService: test4.FrontRepoService,
    private astructService: test4.AstructService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        console.log("frontRepo:", this.frontRepo)

        for (let astruct of this.frontRepo.array_Astructs) {
          console.log("astruct:", astruct.Name)
        }
      }
    })

  }

  onClick(event: MouseEvent, astruct: test4.Astruct) {
    this.astructService.updateFront(astruct, this.Name).subscribe(
      () => {
        console.log("astruct", astruct.Name, "updated");
      }
    )
  }
}
