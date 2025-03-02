import { TestBed } from '@angular/core/testing';

import { AngularDragEndEventService } from './angular-drag-end-event.service';

describe('AngularDragEndEventService', () => {
  let service: AngularDragEndEventService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AngularDragEndEventService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
