import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SvgspecificComponent } from './svgspecific.component';

describe('SvgspecificComponent', () => {
  let component: SvgspecificComponent;
  let fixture: ComponentFixture<SvgspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SvgspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SvgspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
