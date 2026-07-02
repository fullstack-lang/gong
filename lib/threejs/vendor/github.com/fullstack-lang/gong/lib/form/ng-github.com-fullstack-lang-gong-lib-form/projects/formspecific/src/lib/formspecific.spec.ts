import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Formspecific } from './formspecific';

describe('Formspecific', () => {
  let component: Formspecific;
  let fixture: ComponentFixture<Formspecific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Formspecific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Formspecific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
