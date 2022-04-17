import { Injectable } from '@angular/core';
import { Storage } from './storage';
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

  deleteSubreddit(username: string, title: string) {
    return this.WebReqService.post("community/delete", {
      "username": username,
      "name": title
    });
  }

  subscribeToSubreddit(username: string, community_name: string) {
    return this.WebReqService.post("users/UpdateSubsciptions", {
      "username" : username,
      "communityname":community_name
    })
  }

  constructor(private WebReqService: WebRequestService) { }
}

// voteCommunity(community_id: string, username: string, vote: number) {
//   return this.WebReqService.patch("community/vote", {
//     "id": community_id,
//     "username": username,
//     "vote": vote
//   });
// }
