import { Injectable } from '@angular/core';
import { Storage } from './storage';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private WebReqService: WebRequestService) { }

  getPosts() {
    // get data from Backend
    return this.WebReqService.post('home', {
      "pagenumber" : 1,
      "numberofposts" : 100,
      "mode" : "hot",
    });
  }

  deletePost(post_id: string) {
    return this.WebReqService.post('post/delete', {
      "id": post_id,
      "username": Storage.username
    });
  }

  votePost(post_id: string, username: string, vote: number) {
    return this.WebReqService.patch("post/vote", {
      "id": post_id,
      "username": username,
      "vote": vote
    });
  }
}
