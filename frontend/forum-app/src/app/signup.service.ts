import { Injectable } from '@angular/core';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  constructor(private WebReqService: WebRequestService) { }
  addNewAccount(username: string) {
    return this.WebReqService.post('users/check-username', { username });
  }
}
