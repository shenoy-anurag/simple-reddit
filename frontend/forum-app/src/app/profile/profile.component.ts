import { Component, OnInit } from '@angular/core';
import { ProfileService } from '../profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  profile;
  constructor(service: ProfileService) {
    this.profile = service.getProfile();
   }

  ngOnInit(): void {
  }

}
