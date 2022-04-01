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
        console.log(response.data.Profile.username);
        console.log(response.data.Profile.email);
        this.profile = {
          "firstname" : response.data.Profile.firstname,
          "lastname": response.data.Profile.lastname,
          "username": response.data.Profile.username,
          "email": response.data.Profile.email,
          "karma": response.data.Profile.karma,
        }
      }
    });
  }
}
