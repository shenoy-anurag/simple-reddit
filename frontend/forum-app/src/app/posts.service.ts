import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private httpClient: HttpClient) { }

  getPosts() {
    // get data from Backend
    // this.httpClient.get<any>("httplocalhost:8080")
    return [
      {"title" : "Fake Post", "body" : "This is a fake post with fake body", "owner": "John", "created_on" : "02/11/2022"},
      {"title" : "Another Fake Post", "body" : "This is a second fake post with fake body", "owner": "John 2", "created_on" : "02/12/2022"}];
    }
}
