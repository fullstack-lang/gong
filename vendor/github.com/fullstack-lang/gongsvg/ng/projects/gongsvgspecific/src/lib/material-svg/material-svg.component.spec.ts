import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MaterialSvgComponent } from './material-svg.component';

describe('MaterialSvgComponent', () => {
  let component: MaterialSvgComponent;
  let fixture: ComponentFixture<MaterialSvgComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MaterialSvgComponent]
    });
    fixture = TestBed.createComponent(MaterialSvgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
