import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ButtonSpecificComponent } from './button-specific.component';

describe('ButtonSpecificComponent', () => {
  let component: ButtonSpecificComponent;
  let fixture: ComponentFixture<ButtonSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ButtonSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ButtonSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
