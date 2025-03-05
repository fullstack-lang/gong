import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestSpecificComponent } from './test-specific.component';

describe('TestSpecificComponent', () => {
  let component: TestSpecificComponent;
  let fixture: ComponentFixture<TestSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TestSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
