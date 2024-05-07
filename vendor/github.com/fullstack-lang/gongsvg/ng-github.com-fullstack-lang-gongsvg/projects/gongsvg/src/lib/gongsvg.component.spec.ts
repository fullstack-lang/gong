import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsvgComponent } from './gongsvg.component';

describe('GongsvgComponent', () => {
  let component: GongsvgComponent;
  let fixture: ComponentFixture<GongsvgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongsvgComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongsvgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
