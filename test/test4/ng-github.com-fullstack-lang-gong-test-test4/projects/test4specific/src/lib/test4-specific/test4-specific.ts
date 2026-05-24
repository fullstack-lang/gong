import { Component, Input, OnInit, ChangeDetectorRef } from '@angular/core';

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
    private cdr: ChangeDetectorRef // 1. Inject ChangeDetectorRef
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    // // 2. Explicitly pull the initial state from the WASM backend
    // this.frontRepoService.pull(this.Name).subscribe((frontRepo) => {
    //   this.frontRepo = frontRepo;
    //   console.log("initial frontRepo request via WASM WS:", this.frontRepo);
    //   this.cdr.detectChanges(); // Force UI update
    // });

    // 3. Listen for subsequent updates from the Go backend
    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;
        console.log("frontRepo updated via WASM WS:", this.frontRepo);

        for (let astruct of this.frontRepo.array_Astructs) {
          console.log("astruct:", astruct.Name)
        }

        // Force Angular to update the UI since this executes outside the Angular Zone
        this.cdr.detectChanges();
      },
      error: (err) => {
        console.error("Error connecting to WebSocket:", err);
      }
    })
  }

  onClick(event: MouseEvent, astruct: test4.Astruct) {
    this.astructService.updateFront(astruct, this.Name).subscribe(
      () => {
        console.log("astruct", astruct.Name, "updated");
        // We pull again just to ensure the UI is in sync if the WS is slow
        this.frontRepoService.pull(this.Name).subscribe(repo => {
          this.frontRepo = repo;
          this.cdr.detectChanges();
        });
      }
    )
  }
}