import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeletesubredditsformComponent } from './deletesubredditsform.component';

describe('DeletesubredditsformComponent', () => {
  let component: DeletesubredditsformComponent;
  let fixture: ComponentFixture<DeletesubredditsformComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DeletesubredditsformComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DeletesubredditsformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
