import { TestBed } from '@angular/core/testing';

import { CursorspecificService } from './cursorspecific.service';

describe('CursorspecificService', () => {
  let service: CursorspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CursorspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
