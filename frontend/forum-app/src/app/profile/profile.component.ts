import { Component, OnInit } from '@angular/core';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  profile: any

  constructor(private service: ProfileService) {}

  ngOnInit(): void {
    console.log("onInit")
    this.service.getProfile().subscribe((response: any) => {
      console.log(response);
      console.log(response.data.user.username);
      if (response.status == 200) {
        this.profile = {
          "firstname" : "test",//response.data.post.firstname,
          "lastname": "test2", //response.data.post.lastname,
          "username": response.data.user.username,
          "email": response.data.post.email
        }
      }
    });
  }
}
