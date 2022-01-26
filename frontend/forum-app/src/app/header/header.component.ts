import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  title: string = "Simple Reddit";
  usrLoggedIn: boolean = false;
  usr: string = "";
  constructor() { }

  ngOnInit(): void {
  }

  getLogIn(username: string, password: string) {
    console.log('Attempted login: USR:' + username + ' PWD:' + password);
    this.usrLoggedIn = true;
    this.usr = username;
  }

  getLogOut() {
    this.usrLoggedIn = false;
    this.usr = "";
  }
}
