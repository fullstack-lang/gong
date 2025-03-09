import { TestBed } from '@angular/core/testing';

import { TreespecificService } from './treespecific.service';

describe('TreespecificService', () => {
  let service: TreespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TreespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
