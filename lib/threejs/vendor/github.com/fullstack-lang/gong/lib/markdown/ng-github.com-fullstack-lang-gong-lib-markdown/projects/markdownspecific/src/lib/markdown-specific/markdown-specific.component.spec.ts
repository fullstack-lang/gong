import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MarkdownSpecificComponent } from './markdown-specific.component';

describe('MarkdownSpecificComponent', () => {
  let component: MarkdownSpecificComponent;
  let fixture: ComponentFixture<MarkdownSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MarkdownSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MarkdownSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
