import { Component, OnInit } from '@angular/core';
import { Storage } from '../storage';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  getLoginStatus() {
    return Storage.isLoggedIn;
  }
  getLogOut() {
    Storage.isLoggedIn = false;
    Storage.username = "";
  }
}
