import { TestBed } from '@angular/core/testing';

import { GongsimcontrolService } from './gongsimcontrol.service';

describe('GongsimcontrolService', () => {
  let service: GongsimcontrolService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsimcontrolService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
