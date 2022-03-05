import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewpostformComponent } from './newpostform.component';

describe('NewpostformComponent', () => {
  let component: NewpostformComponent;
  let fixture: ComponentFixture<NewpostformComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewpostformComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewpostformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });
});
