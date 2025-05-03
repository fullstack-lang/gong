import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Doc2SpecificComponent } from './doc2-specific.component';

describe('Doc2SpecificComponent', () => {
  let component: Doc2SpecificComponent;
  let fixture: ComponentFixture<Doc2SpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Doc2SpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Doc2SpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
