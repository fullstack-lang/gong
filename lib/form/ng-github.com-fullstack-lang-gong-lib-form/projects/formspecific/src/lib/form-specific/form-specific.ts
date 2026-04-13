import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'lib-form2-specific',
  imports: [],
  templateUrl: './form-specific.html',
  styleUrl: './form-specific.css',
})
export class FormSpecific implements OnInit {
  @Input() Name: string = ""

  ngOnInit(): void {
    console.log("FormSpecific " + this.Name + " initialized")
  }
}
