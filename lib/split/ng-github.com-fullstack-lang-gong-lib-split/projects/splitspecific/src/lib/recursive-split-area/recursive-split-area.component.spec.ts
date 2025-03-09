import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecursiveSplitAreaComponent } from './recursive-split-area.component';

describe('RecursiveSplitAreaComponent', () => {
  let component: RecursiveSplitAreaComponent;
  let fixture: ComponentFixture<RecursiveSplitAreaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RecursiveSplitAreaComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RecursiveSplitAreaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
