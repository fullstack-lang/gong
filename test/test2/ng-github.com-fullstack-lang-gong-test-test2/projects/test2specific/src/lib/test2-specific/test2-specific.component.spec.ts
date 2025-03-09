import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test2SpecificComponent } from './test2-specific.component';

describe('Test2SpecificComponent', () => {
  let component: Test2SpecificComponent;
  let fixture: ComponentFixture<Test2SpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Test2SpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test2SpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
