import { StringMap } from '@angular/compiler/src/compiler_facade_interface';
import { Injectable } from '@angular/core';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  constructor(private WebReqService: WebRequestService) { }
  addNewAccount(email: string, username: string, password: string, name: string) {
    return this.WebReqService.post('users/signup', 
    {
      "email": email,
      "username": username,
      "password": password,
      "name": name
    });
  }

  checkUsername(username: string) {
    return this.WebReqService.post('users/check-username', {"username": username});
  }

  checkLogIn(username: string, password: string) {
    return this.WebReqService.post('users/loginuser', 
    {
      "username": username,
      "password": password
    })
  }

  deleteUser(username: string) {
    // return this.WebReqService.delete('users/user',
    // {
    //   "username": username
    // });
  }

  createPost(username: string, community_id: string, title: string, body: string) {
    return this.WebReqService.post('post/create',
    {
      "username": username,
    	"community_id": community_id,
	    "title" : title,
      "body" : body
    })
  }
}
