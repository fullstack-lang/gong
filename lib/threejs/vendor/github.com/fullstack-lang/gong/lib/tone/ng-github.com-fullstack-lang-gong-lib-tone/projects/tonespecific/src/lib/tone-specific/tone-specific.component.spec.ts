import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ToneSpecificComponent } from './tone-specific.component';

describe('ToneSpecificComponent', () => {
  let component: ToneSpecificComponent;
  let fixture: ComponentFixture<ToneSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ToneSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ToneSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
