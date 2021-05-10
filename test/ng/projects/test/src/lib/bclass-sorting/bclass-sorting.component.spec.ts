import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BclassSortingComponent } from './bclass-sorting.component';

describe('BclassSortingComponent', () => {
  let component: BclassSortingComponent;
  let fixture: ComponentFixture<BclassSortingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BclassSortingComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BclassSortingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
