import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsimcontrolComponent } from './gongsimcontrol.component';

describe('GongsimcontrolComponent', () => {
  let component: GongsimcontrolComponent;
  let fixture: ComponentFixture<GongsimcontrolComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [GongsimcontrolComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GongsimcontrolComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
