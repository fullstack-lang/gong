import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SplitterNewComponent } from './splitter-new.component';

describe('SplitterNewComponent', () => {
  let component: SplitterNewComponent;
  let fixture: ComponentFixture<SplitterNewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SplitterNewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SplitterNewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
