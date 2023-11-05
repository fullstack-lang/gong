import { Component } from '@angular/core';
import { AstructDB, BstructDB, FrontRepo } from 'test';

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

  constructor() {

    let frontRepo = new (FrontRepo)

    frontRepo.Astructs_array.push(new (AstructDB))
    frontRepo.Astructs_array.push(new (AstructDB))

    let array = frontRepo.getArray(AstructDB.GONGSTRUCT_NAME)
    console.log("", array.length)
  }

}
