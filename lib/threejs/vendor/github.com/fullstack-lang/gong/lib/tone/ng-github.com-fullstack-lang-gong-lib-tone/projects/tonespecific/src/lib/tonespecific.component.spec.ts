import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TonespecificComponent } from './tonespecific.component';

describe('TonespecificComponent', () => {
  let component: TonespecificComponent;
  let fixture: ComponentFixture<TonespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TonespecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TonespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
