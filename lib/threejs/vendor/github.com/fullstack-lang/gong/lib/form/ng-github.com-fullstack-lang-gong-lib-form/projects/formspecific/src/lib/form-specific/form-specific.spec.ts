import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormSpecific } from './form-specific';

describe('FormSpecific', () => {
  let component: FormSpecific;
  let fixture: ComponentFixture<FormSpecific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormSpecific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FormSpecific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
