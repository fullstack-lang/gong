import { Component, Input, OnInit } from '@angular/core';

import * as load from '../../../../load/src/public-api'


@Component({
  selector: 'lib-load-specific',
  imports: [],
  templateUrl: './load-specific.component.html',
  styleUrl: './load-specific.component.css'
})
export class LoadSpecificComponent implements OnInit {

  @Input() Name: string = ""

  public frontRepo?: load.FrontRepo;
  public fileToDownload?: load.FileToDownload

  constructor(
    private frontRepoService: load.FrontRepoService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        for (let file_ of this.frontRepo.getFrontArray<load.FileToDownload>(load.FileToDownload.GONGSTRUCT_NAME)) 
        {
          this.fileToDownload = file_
        }

        if (this.fileToDownload == undefined) {
          return
        }

        if (this.frontRepo.getFrontArray<load.FileToDownload>(load.FileToDownload.GONGSTRUCT_NAME).length > 1) {
          return
        }

        // Create a Blob from the file string
        const blob = new Blob([this.fileToDownload.Content], { type: 'text/plain' });

        // Generate a temporary URL for the Blob
        const url = URL.createObjectURL(blob);

        // Create a link element, set its href to the Blob URL, and initiate download
        const link = document.createElement('a');
        link.href = url;
        link.download = this.fileToDownload.Name;  // The name of the file to be downloaded
        link.click();

        // Clean up the URL object once finished
        URL.revokeObjectURL(url);
      }
    }
    )
  }
}
