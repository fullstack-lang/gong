import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongrouterOutletComponent } from './gongrouter-outlet.component';

describe('GongrouterOutletComponent', () => {
  let component: GongrouterOutletComponent;
  let fixture: ComponentFixture<GongrouterOutletComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongrouterOutletComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongrouterOutletComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
