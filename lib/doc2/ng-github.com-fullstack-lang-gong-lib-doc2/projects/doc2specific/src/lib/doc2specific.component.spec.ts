import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Doc2specificComponent } from './doc2specific.component';

describe('Doc2specificComponent', () => {
  let component: Doc2specificComponent;
  let fixture: ComponentFixture<Doc2specificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Doc2specificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(Doc2specificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
