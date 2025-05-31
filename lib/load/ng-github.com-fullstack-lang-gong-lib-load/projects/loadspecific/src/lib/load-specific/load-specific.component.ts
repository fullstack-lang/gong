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

  // to upload a file, the back must instance a fileToUpload
  // the front will update its name and content
  public fileToUpload?: load.FileToUpload

  public message?: load.Message

  constructor(
    private frontRepoService: load.FrontRepoService,
    private fileToUploadService: load.FileToUploadService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        for (let message_ of this.frontRepo.getFrontArray<load.Message>(load.Message.GONGSTRUCT_NAME)) 
        {
          this.message = message_
        }

        for (let file_ of this.frontRepo.getFrontArray<load.FileToDownload>(load.FileToDownload.GONGSTRUCT_NAME)) 
        {
          this.fileToDownload = file_
        }

        for (let file_ of this.frontRepo.getFrontArray<load.FileToUpload>(load.FileToUpload.GONGSTRUCT_NAME)) 
        {
          this.fileToUpload = file_
        }

        if (this.fileToDownload == undefined && this.fileToUpload == undefined) {
          return
        }

        if (this.frontRepo.getFrontArray<load.FileToDownload>(load.FileToDownload.GONGSTRUCT_NAME).length > 1) {
          return
        }

          if (this.frontRepo.getFrontArray<load.FileToUpload>(load.FileToUpload.GONGSTRUCT_NAME).length > 1) {
          return
        }

        // Create a Blob from the file string
        if (this.fileToDownload) {
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

private uploadFile(fileName: string, fileContent: string | ArrayBuffer): void {
    // Minimal guard for component state, though "no error" implies this should be met.
    if (this.fileToUpload == undefined) {
      // If truly "no error is possible", this path would not be hit.
      // You might choose to remove it if the calling context absolutely guarantees fileToUpload exists.
      console.warn("FileToUpload instance is not ready.");
      this.isUploading.set(false); // Reset UI state if proceed further is impossible
      this.uploadStatus.set("Upload cancelled: component not ready.");
      return;
    }

    let base64EncodedContent: string;

    if (typeof fileContent === 'string') {
      // fileContent is a raw string, needs UTF-8 conversion then Base64 encoding
      const encoder = new TextEncoder(); // Handles UTF-8 correctly
      const utf8Bytes = encoder.encode(fileContent);

      // Convert Uint8Array to a binary string
      let binaryString = '';
      utf8Bytes.forEach((byte) => {
          binaryString += String.fromCharCode(byte);
      });
      base64EncodedContent = btoa(binaryString);
    } else {
      // fileContent is an ArrayBuffer (binary data)
      // The type assertion `string | ArrayBuffer` and "no error" means if it's not string, it's ArrayBuffer.
      const uint8Array = new Uint8Array(fileContent);

      // Convert Uint8Array to a binary string
      let binaryString = '';
      uint8Array.forEach((byte) => {
          binaryString += String.fromCharCode(byte);
      });
      base64EncodedContent = btoa(binaryString);
    }

    this.fileToUpload.Name = fileName;
    this.fileToUpload.Base64EncodedContent = base64EncodedContent;

    // The service call is asynchronous and can still fail (network issues, server errors).
    // The template shows `uploadStatus()`, so handling service call outcomes is still relevant for the UI.
    this.fileToUploadService.updateFront(this.fileToUpload, this.Name /* Assuming this.Name is context for updateFront */).subscribe(
      (fileToUploadResponse: load.FileToUpload) => { // Use the actual type returned by the service
          this.isUploading.set(false);
          this.uploadStatus.set(`File "${fileName}" processed successfully`);
          // console.log("Upload successful", fileToUploadResponse); // Optional: for debugging
      },
      (serviceError: any) => {
          this.isUploading.set(false);
          // Even if "no front-end errors" are assumed in logic, service calls can fail.
          // The UI (uploadStatus) should reflect this.
          this.uploadStatus.set(`Upload failed for "${fileName}": ${serviceError.message || 'Server error'}`);
          // console.error("Service upload error:", serviceError); // Optional: for debugging
      }
    );
}
}
