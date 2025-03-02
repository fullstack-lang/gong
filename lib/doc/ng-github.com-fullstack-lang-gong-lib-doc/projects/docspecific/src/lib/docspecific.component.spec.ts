import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DocspecificComponent } from './docspecific.component';

describe('DocspecificComponent', () => {
  let component: DocspecificComponent;
  let fixture: ComponentFixture<DocspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DocspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DocspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
