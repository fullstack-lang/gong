import { Component, Input, OnInit, OnDestroy, EventEmitter, Output, signal } from '@angular/core';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';

import * as load from '../../../../load/src/public-api';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'lib-load-specific',
  imports: [MatIconModule],
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
          this.fileToDownload = undefined
          this.fileToUpload = undefined

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

          // If we are currently waiting for the user to click the "Save As" button, 
          // ignore any subsequent WebSocket messages that might clear the FileToDownload
          // (such as the backend's stager.load() reset after 1 second).
          if (this.saveAsReady() && this.fileToDownload === undefined) {
             return;
          }

          if (this.fileToDownload) {
            // Decode the base64 string to binary data
            const binaryString = window.atob(this.fileToDownload.Base64EncodedContent);
            const len = binaryString.length;
            const bytes = new Uint8Array(len);
            for (let i = 0; i < len; i++) {
              bytes[i] = binaryString.charCodeAt(i);
            }

            // Create Blob from the binary array instead of the raw string
            const blob = new Blob([bytes], { type: 'application/octet-stream' });
            let filename = this.fileToDownload.Name;

            if (filename.startsWith("PROMPT_SAVE_FILE_DIALOG_")) {
              filename = filename.substring("PROMPT_SAVE_FILE_DIALOG_".length);
              this.saveAsBlob = blob;
              this.saveAsName = filename;
              this.saveAsReady.set(true);
            } else {
              this.saveAsReady.set(false);
              const url = URL.createObjectURL(blob);
              const link = document.createElement('a');
              link.href = url;
              link.download = filename;
              link.click();
              URL.revokeObjectURL(url);
            }
          } else {
            this.saveAsReady.set(false);
          }
        }
      });
  }

  public saveAsReady = signal(false);
  public saveAsBlob: Blob | null = null;
  public saveAsName: string = '';

  async saveAs() {
    console.log("saveAs() triggered. saveAsName:", this.saveAsName, "saveAsBlob:", this.saveAsBlob);
    if (!this.saveAsBlob || !this.saveAsName) {
      console.warn("saveAs canceled because blob or name is missing");
      return;
    }
    try {
      if (!(window as any).showSaveFilePicker) {
        throw new Error("Your browser does not support the File System Access API. Downloading file directly instead.");
      }
      console.log("Calling showSaveFilePicker with name:", this.saveAsName);
      const pickerOptions: any = {
        suggestedName: this.saveAsName,
        id: 'gong-save-file-picker',
      };

      if (this.saveAsName.endsWith('.go')) {
        pickerOptions.types = [
          {
            description: 'Go Source File (*.go)',
            accept: { 'text/plain': ['.go'] },
          },
        ];
      } else if (this.saveAsName.endsWith('.html')) {
        pickerOptions.types = [
          {
            description: 'HTML Document (*.html)',
            accept: { 'text/html': ['.html'] },
          },
        ];
      }

      const handle = await (window as any).showSaveFilePicker(pickerOptions);

      const writable = await handle.createWritable();
      await writable.write(this.saveAsBlob);
      await writable.close();
      this.uploadStatus.set("File saved successfully.");
      this.saveAsReady.set(false);
    } catch (err: any) {
      if (err.name !== 'AbortError') {
        console.error(err);
        this.uploadStatus.set("Save fallback: downloading directly...");
        // Fallback to normal download
        const url = URL.createObjectURL(this.saveAsBlob);
        const link = document.createElement('a');
        link.href = url;
        link.download = this.saveAsName;
        link.click();
        URL.revokeObjectURL(url);
        this.saveAsReady.set(false);
        this.uploadStatus.set("File downloaded.");
      }
    }
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
          this.uploadStatus.set(`File "${fileName}" has been successfully uploaded.`);
          console.log("Upload successful");
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