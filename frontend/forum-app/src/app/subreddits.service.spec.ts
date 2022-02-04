import { TestBed } from '@angular/core/testing';

import { SubredditsService } from './subreddits.service';

describe('SubredditsService', () => {
  let service: SubredditsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SubredditsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
