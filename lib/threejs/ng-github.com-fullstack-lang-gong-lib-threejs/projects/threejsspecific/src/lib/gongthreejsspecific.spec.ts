import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Gongthreejsspecific } from './threejsspecific';

describe('Gongthreejsspecific', () => {
  let component: Gongthreejsspecific;
  let fixture: ComponentFixture<Gongthreejsspecific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Gongthreejsspecific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Gongthreejsspecific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
