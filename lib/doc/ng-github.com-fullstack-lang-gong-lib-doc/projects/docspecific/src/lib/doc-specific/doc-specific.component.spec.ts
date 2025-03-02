import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DocSpecificComponent } from './doc-specific.component';

describe('DocSpecificComponent', () => {
  let component: DocSpecificComponent;
  let fixture: ComponentFixture<DocSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DocSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DocSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
