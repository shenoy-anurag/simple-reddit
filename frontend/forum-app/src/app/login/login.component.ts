import { Component, OnInit } from '@angular/core';
import { SignupService } from '../signup.service';
import { FormControl, Validators, FormArray, FormBuilder, ValidatorFn, AbstractControl, Validator, ValidationErrors, FormGroup } from '@angular/forms';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { Storage } from '../storage';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  public showPassword: boolean = false;
  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private snackBar: MatSnackBar, private fb: FormBuilder, private router: Router) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
  }

  usrLoggedIn: boolean = Storage.isLoggedIn;
  usr: string = "";

  getLogInStatus() {
    return Storage.isLoggedIn;
  }

  ngOnInit(): void {
  }

  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }

  get f() {
    return this.form.controls;
  }

  getLogIn(username: string, password: string) {
    this.signupService.checkLogIn(username, password).subscribe((response: any) => {
      if (response.status == 200 && response.message == "success") {
        // LogIn Attempt Sucessful
        Storage.isLoggedIn = true;
        Storage.username = username;
        this.usr = username;
        this.snackBar.open("Logged in as " + username, "Dismiss", { duration: 1500 });

        // Route to Home
        this.router.navigate(['/home']);
      }
      else if (response.status == 200 && response.message == "failure") {
        // Prompt user, incorrect login
        this.snackBar.open("Failed login", "Dismiss", { duration: 1500 });
      }
      else {
        // Something else is wrong
        this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 1500 };
      }
    })
  }

  getLogOut() {
    Storage.isLoggedIn = false;
    Storage.username = "";
    this.usr = "";
    this.snackBar.open("Logged out ", "Dismiss",{ duration: 3000 });
  }
}
