import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PolygoneComponent } from './polygone.component';

describe('PolygoneComponent', () => {
  let component: PolygoneComponent;
  let fixture: ComponentFixture<PolygoneComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PolygoneComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PolygoneComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
