import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsvgComponent } from './gongsvg.component';

describe('MaterialSvgComponent', () => {
  let component: GongsvgComponent;
  let fixture: ComponentFixture<GongsvgComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GongsvgComponent]
    });
    fixture = TestBed.createComponent(GongsvgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
