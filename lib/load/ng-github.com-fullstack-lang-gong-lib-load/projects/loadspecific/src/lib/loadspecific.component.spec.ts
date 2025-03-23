import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadspecificComponent } from './loadspecific.component';

describe('LoadspecificComponent', () => {
  let component: LoadspecificComponent;
  let fixture: ComponentFixture<LoadspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoadspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoadspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
