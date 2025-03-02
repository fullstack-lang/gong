import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SvgSpecificComponent } from './svg-specific.component';

describe('SvgSpecificComponent', () => {
  let component: SvgSpecificComponent;
  let fixture: ComponentFixture<SvgSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SvgSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SvgSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
