import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeleteuserformComponent } from './deleteuserform.component';

describe('DeleteuserformComponent', () => {
  let component: DeleteuserformComponent;
  let fixture: ComponentFixture<DeleteuserformComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DeleteuserformComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DeleteuserformComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });
});
