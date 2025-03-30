import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProbespecificComponent } from './probespecific.component';

describe('ProbespecificComponent', () => {
  let component: ProbespecificComponent;
  let fixture: ComponentFixture<ProbespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProbespecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProbespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
