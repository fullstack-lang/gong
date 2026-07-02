import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CursorspecificComponent } from './cursorspecific.component';

describe('CursorspecificComponent', () => {
  let component: CursorspecificComponent;
  let fixture: ComponentFixture<CursorspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CursorspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CursorspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
