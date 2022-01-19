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

  toggleLogIn(username: string, password: string) {
    console.log('Attempted login: USR:' + username + ' PWD:' + password);
    this.usrLoggedIn = true;
    this.usr = username;
  }

  toggleLogOut() {
    this.usrLoggedIn = false;
    this.usr = "";
  }
}
