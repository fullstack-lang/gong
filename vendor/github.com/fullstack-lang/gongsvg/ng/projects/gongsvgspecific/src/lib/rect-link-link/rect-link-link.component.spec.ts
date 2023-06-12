import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RectLinkLinkComponent } from './rect-link-link.component';

describe('RectLinkLinkComponent', () => {
  let component: RectLinkLinkComponent;
  let fixture: ComponentFixture<RectLinkLinkComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RectLinkLinkComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RectLinkLinkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
