import { TestBed } from '@angular/core/testing';

import { ButtonspecificService } from './buttonspecific.service';

describe('ButtonspecificService', () => {
  let service: ButtonspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ButtonspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
