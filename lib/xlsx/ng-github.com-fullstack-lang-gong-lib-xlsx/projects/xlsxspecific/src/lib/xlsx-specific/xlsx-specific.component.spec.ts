import { ComponentFixture, TestBed } from '@angular/core/testing';

import { XlsxSpecificComponent } from './xlsx-specific.component';

describe('XlsxSpecificComponent', () => {
  let component: XlsxSpecificComponent;
  let fixture: ComponentFixture<XlsxSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [XlsxSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(XlsxSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
