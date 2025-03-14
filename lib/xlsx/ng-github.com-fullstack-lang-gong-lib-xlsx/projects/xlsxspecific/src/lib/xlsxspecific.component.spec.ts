import { ComponentFixture, TestBed } from '@angular/core/testing';

import { XlsxspecificComponent } from './xlsxspecific.component';

describe('XlsxspecificComponent', () => {
  let component: XlsxspecificComponent;
  let fixture: ComponentFixture<XlsxspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [XlsxspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(XlsxspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
