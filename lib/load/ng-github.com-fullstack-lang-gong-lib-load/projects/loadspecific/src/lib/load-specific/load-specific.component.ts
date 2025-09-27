import { Component, Input, OnInit, OnDestroy, EventEmitter, Output, signal } from '@angular/core';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

import * as load from '../../../../load/src/public-api';

@Component({
  selector: 'lib-load-specific',
  imports: [],
  templateUrl: './load-specific.component.html',
  styleUrl: './load-specific.component.css'
})
export class LoadSpecificComponent implements OnInit, OnDestroy {
  @Output() fileDropped = new EventEmitter<File>();
  isDragging = false;

  @Input() Name: string = "";

  public frontRepo?: load.FrontRepo;
  public fileToDownload?: load.FileToDownload;
  public fileToUpload?: load.FileToUpload;
  public message?: load.Message;

  // 1. Create a subject to notify when the component is destroyed.
  private readonly destroy$ = new Subject<void>();

  constructor(
    private frontRepoService: load.FrontRepoService,
    private fileToUploadService: load.FileToUploadService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.Name)
      .pipe(
        // 2. Use takeUntil to automatically complete the subscription on destroy.
        takeUntil(this.destroy$)
      )
      .subscribe({
        next: (frontRepo) => {
          // This block will now only be executed by this component instance's subscription.
          console.log("WebSocket message received.");

          this.frontRepo = frontRepo;

          for (let message_ of this.frontRepo.getFrontArray<load.Message>(load.Message.GONGSTRUCT_NAME)) {
            this.message = message_;
          }

          for (let file_ of this.frontRepo.getFrontArray<load.FileToDownload>(load.FileToDownload.GONGSTRUCT_NAME)) {
            this.fileToDownload = file_;
          }

          for (let file_ of this.frontRepo.getFrontArray<load.FileToUpload>(load.FileToUpload.GONGSTRUCT_NAME)) {
            this.fileToUpload = file_;
          }

          if (this.fileToDownload == undefined && this.fileToUpload == undefined) {
            return;
          }

          if (this.frontRepo.getFrontArray<load.FileToDownload>(load.FileToDownload.GONGSTRUCT_NAME).length > 1) {
            return;
          }

          if (this.frontRepo.getFrontArray<load.FileToUpload>(load.FileToUpload.GONGSTRUCT_NAME).length > 1) {
            return;
          }

          if (this.fileToDownload) {
            const blob = new Blob([this.fileToDownload.Content], { type: 'text/plain' });
            const url = URL.createObjectURL(blob);
            const link = document.createElement('a');
            link.href = url;
            link.download = this.fileToDownload.Name;
            link.click();
            URL.revokeObjectURL(url);
          }
        }
      });
  }

  ngOnDestroy(): void {
    // 3. Emit a value and complete the subject to trigger takeUntil.
    this.destroy$.next();
    this.destroy$.complete();
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
    const element = event.currentTarget as HTMLInputElement;
    const fileList: FileList | null = element.files;
    if (fileList && fileList.length > 0) {
      this.handleFile(fileList[0]);
    }
  }

  private handleFile(file: File): void {
    if (!file) {
      this.uploadStatus.set("No file selected.");
      return;
    }

    this.isUploading.set(true);
    this.uploadStatus.set(`Preparing to upload ${file.name}...`);

    const reader = new FileReader();
    reader.readAsArrayBuffer(file);

    reader.onload = (e: ProgressEvent<FileReader>) => {
      if (e.target?.result) {
        const fileContent = e.target.result as ArrayBuffer;
        this.uploadFileInternal(file.name, fileContent);
      } else {
        this.isUploading.set(false);
        this.uploadStatus.set(`Error reading file ${file.name}.`);
        console.error("FileReader error: event target result is null or undefined.");
      }
    };

    reader.onerror = (error: ProgressEvent<FileReader>) => {
      this.isUploading.set(false);
      this.uploadStatus.set(`Error reading file ${file.name}: ${reader.error?.message}`);
      console.error("FileReader error:", reader.error);
    };
  }

  private uploadFileInternal(fileName: string, fileContent: string | ArrayBuffer): void {
    if (this.fileToUpload == undefined) {
      this.isUploading.set(false);
      this.uploadStatus.set("Upload cancelled: component not ready.");
      console.warn("FileToUpload instance is not ready for internal upload call.");
      return;
    }

    this.uploadStatus.set(`Encoding ${fileName}...`);
    let base64EncodedContent: string;

    try {
      if (typeof fileContent === 'string') {
        const encoder = new TextEncoder();
        const utf8Bytes = encoder.encode(fileContent);
        let binaryString = '';
        utf8Bytes.forEach((byte) => {
          binaryString += String.fromCharCode(byte);
        });
        base64EncodedContent = btoa(binaryString);
      } else { // fileContent is ArrayBuffer
        const uint8Array = new Uint8Array(fileContent);
        let binaryString = '';
        uint8Array.forEach((byte) => {
          binaryString += String.fromCharCode(byte);
        });
        base64EncodedContent = btoa(binaryString);
      }

      this.fileToUpload.Name = fileName;
      this.fileToUpload.Base64EncodedContent = base64EncodedContent;

      this.uploadStatus.set(`Uploading ${fileName}...`);
      this.fileToUploadService.updateFront(this.fileToUpload, this.Name).subscribe(
        (fileToUploadResponse: load.FileToUpload) => {
          this.isUploading.set(false);
          this.uploadStatus.set(`File "${fileName}" processed successfully`);
          console.log("Upload successful", fileToUploadResponse);
        },
        (serviceError: any) => {
          this.isUploading.set(false);
          this.uploadStatus.set(`Upload failed for "${fileName}": ${serviceError.message || 'Server error'}`);
          console.error("Service upload error:", serviceError);
        }
      );
    } catch (error: any) {
      this.isUploading.set(false);
      this.uploadStatus.set(`Error processing file "${fileName}": ${error.message || 'Encoding failed'}`);
      console.error("Error during file processing or encoding:", error);
    }
  }
}