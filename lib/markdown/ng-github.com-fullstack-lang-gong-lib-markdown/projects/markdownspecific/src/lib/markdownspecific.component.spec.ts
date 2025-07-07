import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MarkdownspecificComponent } from './markdownspecific.component';

describe('MarkdownspecificComponent', () => {
  let component: MarkdownspecificComponent;
  let fixture: ComponentFixture<MarkdownspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MarkdownspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MarkdownspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
