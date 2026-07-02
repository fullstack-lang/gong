import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TreeSpecificComponent } from './tree-specific.component';

describe('TreeSpecificComponent', () => {
  let component: TreeSpecificComponent;
  let fixture: ComponentFixture<TreeSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TreeSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TreeSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
