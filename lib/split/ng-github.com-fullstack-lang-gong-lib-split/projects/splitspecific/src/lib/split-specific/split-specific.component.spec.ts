import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SplitSpecificComponent } from './split-specific.component';

describe('SplitSpecificComponent', () => {
  let component: SplitSpecificComponent;
  let fixture: ComponentFixture<SplitSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SplitSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SplitSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
