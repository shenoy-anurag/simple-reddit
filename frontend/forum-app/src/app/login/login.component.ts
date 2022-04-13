import { Component, OnInit, Inject, HostListener } from '@angular/core';
import { SignupService } from '../signup.service';
import { FormControl, Validators, FormArray, FormBuilder, ValidatorFn, AbstractControl, Validator, ValidationErrors, FormGroup } from '@angular/forms';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { Storage } from '../storage';
import { Router } from '@angular/router';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  windowScrolled!: boolean;
  public showPassword: boolean = false;
  form: FormGroup = new FormGroup({});
  constructor(private signupService: SignupService, private snackBar: MatSnackBar, private fb: FormBuilder, private router: Router, @Inject(DOCUMENT) private document: Document) {
    
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
  }

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
