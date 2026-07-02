import { TestBed } from '@angular/core/testing';

import { IsEditableService } from './is-editable.service';

describe('IsEditableService', () => {
  let service: IsEditableService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IsEditableService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
