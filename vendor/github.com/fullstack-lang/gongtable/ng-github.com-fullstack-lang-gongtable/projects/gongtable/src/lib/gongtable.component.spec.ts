import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtableComponent } from './gongtable.component';

describe('GongtableComponent', () => {
  let component: GongtableComponent;
  let fixture: ComponentFixture<GongtableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongtableComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongtableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
