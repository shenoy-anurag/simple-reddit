import { Injectable } from '@angular/core';
import { Storage } from './storage';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private WebReqService: WebRequestService) { }

  getProfile() {
    // get data from Backend
    // return {"firstname" : "John", "lastname": "Doe", "username": "JohnDoe", "email": "JohnDoe@email.com"};
    console.log("getting profile of: " + Storage.username)
    return this.WebReqService.getPayload('profile', 
    {
      "username": Storage.username
    });
  }
}
