import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DataModelPanelComponent } from './data-model-panel.component';

describe('DataModelPanelComponent', () => {
  let component: DataModelPanelComponent;
  let fixture: ComponentFixture<DataModelPanelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DataModelPanelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DataModelPanelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
