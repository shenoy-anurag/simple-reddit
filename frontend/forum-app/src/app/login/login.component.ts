import { Component, OnInit } from '@angular/core';
import { SignupService } from '../signup.service';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private signupService: SignupService, private snackBar: MatSnackBar) { }

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
        this.snackBar.open("Logged in as " + username, "Dismiss", {duration: 2000});
      }
      else if (response.status == 200 && response.message == "failure" && response.data.data == 'Incorrect Credentials') {
        // Prompt user, incorrect login
        this.snackBar.open("Failed login", "Dismiss", {duration: 2000});
      }
      else {
        // Something else is wrong
        this.snackBar.open("Something is wrong", "Alert Adminstration"), {duration: 2000};
      }
    })
  }

  getLogOut() {
    this.usrLoggedIn = false;
    this.usr = "";
  }
}
