import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TextWidthCalculatorComponent } from './text-width-calculator.component';

describe('TextWidthCalculatorComponent', () => {
  let component: TextWidthCalculatorComponent;
  let fixture: ComponentFixture<TextWidthCalculatorComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TextWidthCalculatorComponent]
    });
    fixture = TestBed.createComponent(TextWidthCalculatorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
