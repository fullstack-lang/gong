import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class IsEditableService {

  isEditable = false

  constructor() { }

  setIsEditable(isEditable: boolean) {
    this.isEditable = isEditable
  }

  getIsEditable(): boolean {
    return this.isEditable
  }
}
