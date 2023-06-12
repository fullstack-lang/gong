import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsvgspecificComponent } from './gongsvgspecific.component';

describe('GongsvgspecificComponent', () => {
  let component: GongsvgspecificComponent;
  let fixture: ComponentFixture<GongsvgspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongsvgspecificComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GongsvgspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
