import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongdocdatamodelComponent } from './gongdocdatamodel.component';

describe('GongdocdatamodelComponent', () => {
  let component: GongdocdatamodelComponent;
  let fixture: ComponentFixture<GongdocdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongdocdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongdocdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
