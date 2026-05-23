import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test4Specific } from './test4-specific';

describe('Test4Specific', () => {
  let component: Test4Specific;
  let fixture: ComponentFixture<Test4Specific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Test4Specific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test4Specific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
