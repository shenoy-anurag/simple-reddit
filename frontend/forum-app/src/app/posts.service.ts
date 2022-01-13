import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  getPosts() {
    // get data from Backend
    return ["Post 1", "Post 2", "Post 3", "Post 4"];
  }

  constructor() { }
}
