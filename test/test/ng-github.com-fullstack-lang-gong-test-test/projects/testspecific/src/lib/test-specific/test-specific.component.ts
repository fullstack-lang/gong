import { Component, Input, OnInit } from '@angular/core';

import * as test from '../../../../test/src/public-api'

import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'lib-test-specific',
  imports: [
    MatButtonModule,
  ],
  templateUrl: './test-specific.component.html',
  styleUrl: './test-specific.component.css'
})
export class TestSpecificComponent implements OnInit {


  @Input() Name: string = ""

  public frontRepo?: test.FrontRepo;

  constructor(
    private frontRepoService: test.FrontRepoService,
    private astructService: test.AstructService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        for (let astruct of this.frontRepo.array_Astructs) {
          // console.log("astruct:", astruct.Name)
        }
      }
    })

  }

  onClick(event: MouseEvent, astruct: test.Astruct) {
    this.astructService.updateFront(astruct, this.Name).subscribe(
      () => {
        console.log("astruct", astruct.Name, "updated");
      }
    )
  }
}
