import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SliderSpecificComponent } from './slider-specific.component';

describe('SliderSpecificComponent', () => {
  let component: SliderSpecificComponent;
  let fixture: ComponentFixture<SliderSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SliderSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SliderSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
