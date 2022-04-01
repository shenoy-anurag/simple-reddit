import { Injectable } from '@angular/core';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private WebReqService: WebRequestService) { }

  getProfile(username: string) {
    // get data from Backend
    return this.WebReqService.post('profile',
    {
      "username": username
    });
  }
}
