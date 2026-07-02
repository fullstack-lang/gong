import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongthreejsSpecific } from './gong-lib-threejs-specific';

describe('GongthreejsSpecific', () => {
  let component: GongthreejsSpecific;
  let fixture: ComponentFixture<GongthreejsSpecific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongthreejsSpecific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongthreejsSpecific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
