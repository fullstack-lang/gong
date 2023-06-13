import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongdocspecificComponent } from './gongdocspecific.component';

describe('GongdocspecificComponent', () => {
  let component: GongdocspecificComponent;
  let fixture: ComponentFixture<GongdocspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongdocspecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongdocspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
