import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SliderspecificComponent } from './sliderspecific.component';

describe('SliderspecificComponent', () => {
  let component: SliderspecificComponent;
  let fixture: ComponentFixture<SliderspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SliderspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SliderspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
