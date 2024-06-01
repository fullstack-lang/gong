import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongComponent } from './gong.component';

describe('GongComponent', () => {
  let component: GongComponent;
  let fixture: ComponentFixture<GongComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
