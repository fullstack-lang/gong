import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProbeSpecificComponent } from './probe-specific.component';

describe('ProbeSpecificComponent', () => {
  let component: ProbeSpecificComponent;
  let fixture: ComponentFixture<ProbeSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProbeSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProbeSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
