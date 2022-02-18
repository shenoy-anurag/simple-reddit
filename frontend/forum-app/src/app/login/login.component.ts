import { Component, OnInit } from '@angular/core';
import { SignupService } from '../signup.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private signupService: SignupService) { }

  usrLoggedIn: boolean = false;
  usr: string = "";

  ngOnInit(): void {
  }

  getLogIn(username: string, password: string) {
    console.log('Attempted login: USR:' + username + ' PWD:' + password);
    this.signupService.checkLogIn(username, password).subscribe((response: any) => {
      console.log(response);

      if (response.status == 200 && response.message == "success") {
        // LogIn Attempt Sucessful
        this.usrLoggedIn = true;
        this.usr = username;  
      }
      else {
        // Prompt user, incorrect login
        console.log("Failed login");
      }
    })
  }

  getLogOut() {
    this.usrLoggedIn = false;
    this.usr = "";
  }
}
