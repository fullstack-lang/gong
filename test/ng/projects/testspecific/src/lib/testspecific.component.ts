import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'lib-testspecific',
  template: `
    <p>
      testspecific works!
    </p>
  `,
  styles: [
  ]
})
export class TestspecificComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
