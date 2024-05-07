import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MaterialFormComponent } from './material-form.component';

describe('MaterialFormComponent', () => {
  let component: MaterialFormComponent;
  let fixture: ComponentFixture<MaterialFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MaterialFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MaterialFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
