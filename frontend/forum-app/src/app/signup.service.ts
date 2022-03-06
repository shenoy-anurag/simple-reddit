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

  createSubreddit(user_id: string, name: string, description: string) {
    return this.WebReqService.post('community', 
    {
      "user_id": "6217146a9b60fde368166137",
      "name": "science",
      "description": "This community is a place to share and discuss new scientific research. Read about the latest advances in astronomy, biology, medicine, physics, social science, and more. Find and submit new publications and popular science coverage of current research."
  })
  }

  createdeleteSubreddit(username: string, name: string) {
    return this.WebReqService.post('community/science', 
    {
      "username": "albert",
      "name": "science"
  })
  }

}
