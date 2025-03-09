import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongSpecificComponent } from './gong-specific.component';

describe('GongSpecificComponent', () => {
  let component: GongSpecificComponent;
  let fixture: ComponentFixture<GongSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
