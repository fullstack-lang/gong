import { Component, Input, OnInit, EventEmitter, Output, signal  } from '@angular/core';

import * as load from '../../../../load/src/public-api'


@Component({
  selector: 'lib-load-specific',
  imports: [],
  templateUrl: './load-specific.component.html',
  styleUrl: './load-specific.component.css'
})
export class LoadSpecificComponent implements OnInit {
  @Output() fileDropped = new EventEmitter<File>();
  isDragging = false;

  @Input() Name: string = ""

  public frontRepo?: load.FrontRepo;
  public fileToDownload?: load.FileToDownload

  constructor(
    private frontRepoService: load.FrontRepoService,
    private fileToUploadService: load.FileToUploadService,
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

 isDragOver = signal(false);
  isUploading = signal(false);
  uploadStatus = signal<string>('');

  onDragOver(event: DragEvent): void {
    event.preventDefault();
    event.stopPropagation();
    this.isDragOver.set(true);
  }

  onDragLeave(event: DragEvent): void {
    event.preventDefault();
    event.stopPropagation();
    this.isDragOver.set(false);
  }

  onDrop(event: DragEvent): void {
    event.preventDefault();
    event.stopPropagation();
    this.isDragOver.set(false);

    const files = event.dataTransfer?.files;
    if (files && files.length > 0) {
      this.handleFile(files[0]);
    }
  }

  onFileSelected(event: Event): void {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files.length > 0) {
      this.handleFile(input.files[0]);
    }
  }

  private handleFile(file: File): void {
    this.isUploading.set(true);
    this.uploadStatus.set('Reading file...');

    const reader = new FileReader();
    reader.onload = (e) => {
      const fileContent = e.target?.result as string;
      this.uploadFile(file.name, fileContent);
    };

    reader.onerror = () => {
      this.uploadStatus.set('Error reading file');
      this.isUploading.set(false);
    };

    reader.readAsText(file);
  }

  private uploadFile(fileName: string, fileContent: string): void {
    // Keep the content of function empty as requested

    let fileToUpload = new(load.FileToUpload)
    fileToUpload.Name = fileName
    fileToUpload.Content = fileContent

    this.fileToUploadService.postFront(fileToUpload, this.Name).subscribe(
      ()=> {
          this.isUploading.set(false);
          this.uploadStatus.set(`File "${fileName}" processed successfully`);
      }
    )
  }
}
