import { StringMap } from '@angular/compiler/src/compiler_facade_interface';
import { Injectable } from '@angular/core';
import { WebRequestService } from './web-request.service';

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  constructor(private WebReqService: WebRequestService) { }
  
  addNewAccount(email: string, username: string, password: string, firstname: string, lastname: string) {
    return this.WebReqService.post('users/signup', 
    {
      "email": email,
      "username": username,
      "password": password,
      "firstname": firstname,
      "lastname": lastname
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
    });
  }

  createcommunity(username: string, name: string, description: string) {
    console.log("into post block");
    return this.WebReqService.post('community',
    {
      "username": username,
      "name": name,
      "description": description
    });
  }

  deletecommunity(username: string, name: string) {
    console.log("delete post block")
    return this.WebReqService.delete('community', 
    {
      "username": 'albert',
      "name": 'science'
    });
  }

}
