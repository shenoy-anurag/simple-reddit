import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewsubredditsformComponent } from './newsubredditsform.component';

describe('NewsubredditsformComponent', () => {
  let component: NewsubredditsformComponent;
  let fixture: ComponentFixture<NewsubredditsformComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewsubredditsformComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewsubredditsformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });
});
