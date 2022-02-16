import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SubredditsService {

  getSubreddits() {
    // get data from Backend
    return [
      {"title": "Subreddit 1"},
      {"title": "Subreddit 2"},
      {"title": "Subreddit 3"},
      {"title": "Subreddit 4"}
    ]
  }

  constructor() { }
}
