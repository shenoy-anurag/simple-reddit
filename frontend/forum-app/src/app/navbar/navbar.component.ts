import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Storage } from '../storage';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor(private snackBar: MatSnackBar) { }

  ngOnInit(): void {
  }

  getLoginStatus() {
    return Storage.isLoggedIn;
  }
  getLogOut() {
    Storage.isLoggedIn = false;
    Storage.username = "";
    // this.usr = "";
    this.snackBar.open("Logged out ", "Dismiss",{ duration: 3000 });
  }
}
