import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SubredditsService {

  getSubreddits() {
    // get data from Backend
    return ["Subreddit 1", "Subreddit 2", "Subreddit 3", "Subreddit 4"];
  }

  constructor() { }
}
