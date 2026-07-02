import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TreespecificComponent } from './treespecific.component';

describe('TreespecificComponent', () => {
  let component: TreespecificComponent;
  let fixture: ComponentFixture<TreespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TreespecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TreespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
