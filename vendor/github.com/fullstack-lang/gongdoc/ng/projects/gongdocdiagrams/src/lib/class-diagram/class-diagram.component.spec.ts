import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClassDiagramComponent } from './class-diagram.component';

describe('ClassDiagramComponent', () => {
  let component: ClassDiagramComponent;
  let fixture: ComponentFixture<ClassDiagramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClassDiagramComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ClassDiagramComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
