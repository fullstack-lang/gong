import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ButtonspecificComponent } from './buttonspecific.component';

describe('ButtonspecificComponent', () => {
  let component: ButtonspecificComponent;
  let fixture: ComponentFixture<ButtonspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ButtonspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ButtonspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
