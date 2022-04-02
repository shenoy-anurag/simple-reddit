import { Injectable } from '@angular/core';
import { SignupService } from './signup.service';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class SubredditsService {

  getSubreddits() {
    // get data from Backend
    return this.WebReqService.post("community/all", {});
  }

  constructor(private WebReqService: WebRequestService) { }
}
