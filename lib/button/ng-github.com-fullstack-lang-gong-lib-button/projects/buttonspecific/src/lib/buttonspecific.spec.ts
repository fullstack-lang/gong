import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Buttonspecific } from './buttonspecific';

describe('Buttonspecific', () => {
  let component: Buttonspecific;
  let fixture: ComponentFixture<Buttonspecific>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Buttonspecific]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Buttonspecific);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
