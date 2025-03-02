import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TableSpecificComponent } from './table-specific.component';

describe('TableSpecificComponent', () => {
  let component: TableSpecificComponent;
  let fixture: ComponentFixture<TableSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TableSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TableSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
