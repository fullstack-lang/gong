import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongrouterspecificComponent } from './gongrouterspecific.component';

describe('GongrouterspecificComponent', () => {
  let component: GongrouterspecificComponent;
  let fixture: ComponentFixture<GongrouterspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongrouterspecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongrouterspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
