import { Component } from '@angular/core';
import { AstructDB, FrontRepo } from 'test';

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
export class TestspecificComponent {

  frontRepo = new FrontRepo

  titi() {

    let a = this.frontRepo.getArray<AstructDB>()

  }

}
