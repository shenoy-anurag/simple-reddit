import { Injectable } from '@angular/core';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private WebReqService: WebRequestService) { }

  getPosts() {
    // get data from Backend
    return this.WebReqService.get('home');
  }
}
