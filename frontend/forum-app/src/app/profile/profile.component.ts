import { Component, OnInit, Inject, HostListener  } from '@angular/core';
import { ProfileService } from '../profile.service';
import { Storage } from '../storage';
import { Router } from '@angular/router';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  
  windowScrolled!: boolean;
  profile: any

  constructor(private service: ProfileService, @Inject(DOCUMENT) private document: Document) {}

  @HostListener("window:scroll", [])

onWindowScroll() {
    if (window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop > 100) {
        this.windowScrolled = true;
    } 
   else if (this.windowScrolled && window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop < 10) {
        this.windowScrolled = false;
    }
}


scrollToTop() {
  (function smoothscroll() {
      var currentScroll = document.documentElement.scrollTop || document.body.scrollTop;
      if (currentScroll > 0) {
          window.requestAnimationFrame(smoothscroll);
          window.scrollTo(0, currentScroll - (currentScroll / 8));
      }
  })();
}

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
