import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongspecificComponent } from './gongspecific.component';

describe('GongspecificComponent', () => {
  let component: GongspecificComponent;
  let fixture: ComponentFixture<GongspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongspecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
