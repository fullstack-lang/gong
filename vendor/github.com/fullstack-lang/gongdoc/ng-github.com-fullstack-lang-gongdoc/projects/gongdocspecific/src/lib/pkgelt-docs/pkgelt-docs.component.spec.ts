import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PkgeltDocsComponent } from './pkgelt-docs.component';

describe('PkgeltDocsComponent', () => {
  let component: PkgeltDocsComponent;
  let fixture: ComponentFixture<PkgeltDocsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PkgeltDocsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PkgeltDocsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
