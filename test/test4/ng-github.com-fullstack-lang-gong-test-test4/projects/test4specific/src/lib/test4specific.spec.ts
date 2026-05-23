import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test4specific } from './test4specific';

describe('Test4specific', () => {
  let component: Test4specific;
  let fixture: ComponentFixture<Test4specific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Test4specific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test4specific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
