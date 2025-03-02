import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormSpecificComponent } from './form-specific.component';

describe('FormSpecificComponent', () => {
  let component: FormSpecificComponent;
  let fixture: ComponentFixture<FormSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FormSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
