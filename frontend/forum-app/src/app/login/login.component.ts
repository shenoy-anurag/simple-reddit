import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor() { }

  usrLoggedIn: boolean = false;
  usr: string = "";

  ngOnInit(): void {
  }

  getLogIn(username: string, password: string) {
    console.log('Attempted login: USR:' + username + ' PWD:' + password);
    if (username.length > 0 && password.length > 0) {
      this.usrLoggedIn = true;
      this.usr = username;  
    }
  }

  getLogOut() {
    this.usrLoggedIn = false;
    this.usr = "";
  }
}
