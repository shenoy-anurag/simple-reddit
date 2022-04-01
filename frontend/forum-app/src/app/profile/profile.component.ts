import { Component, OnInit } from '@angular/core';
import { ProfileService } from '../profile.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  profile: any

  constructor(private service: ProfileService) {}

  ngOnInit(): void {
    this.getUserProfile();
  }

  getUserProfile() {
    this.service.getProfile(Storage.username).subscribe((response: any) => {
      console.log(response);
      console.log(response.status);
      if (response.status == 200) {
        console.log("TEST");
        console.log(response.data.user.username);
        this.profile = response.data;
        // this.profile = {
        //   "firstname" : "test",//response.data.post.firstname,
        //   "lastname": "test2", //response.data.post.lastname,
        //   "username": response.data.user.username,
        //   "email": response.data.post.email
        // }
      }
    });
  }
}
