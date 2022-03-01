import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor() { }

  getProfile() {
    // get data from Backend
    return {"firstname" : "John", "lastname": "Doe", "username": "JohnDoe", "email": "JohnDoe@email.com"};
  }
}
