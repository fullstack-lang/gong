import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Test3SpecificComponent } from './test3-specific.component';

describe('Test3SpecificComponent', () => {
  let component: Test3SpecificComponent;
  let fixture: ComponentFixture<Test3SpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Test3SpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Test3SpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
