import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SimSpecificComponent } from './sim-specific.component';

describe('SimSpecificComponent', () => {
  let component: SimSpecificComponent;
  let fixture: ComponentFixture<SimSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SimSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SimSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
