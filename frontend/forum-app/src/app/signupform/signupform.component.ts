import { Component, OnInit, Inject, HostListener } from '@angular/core';
import { Validators, FormBuilder, ValidatorFn, AbstractControl, ValidationErrors, FormGroup } from '@angular/forms';
import { SignupService } from '../signup.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Storage } from '../storage';
import { Router } from '@angular/router';
import { DOCUMENT } from '@angular/common';


export function checkIfMatchingPasswords(passwordKey: string, passwordConfirmationKey: string) {
  return (group: FormGroup) => {
    let passwordInput = group.controls[passwordKey],
        passwordConfirmationInput = group.controls[passwordConfirmationKey];
    if (passwordInput.value !== passwordConfirmationInput.value) {
      return passwordConfirmationInput.setErrors({notEquivalent: true})
    }
    else {
        return passwordConfirmationInput.setErrors(null);
    }
  }
}


export function passwordValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }

    const hasUpperCase = /[A-Z]+/.test(value);

    const hasLowerCase = /[a-z]+/.test(value);

    const hasNumeric = /[0-9]+/.test(value);

    const isPasswordValid = hasUpperCase && hasLowerCase && hasNumeric;

    return !isPasswordValid ? { passwordValid: true } : null;
  };
}

export function ConfirmedValidator(password: any): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }
    
    console.log(password);
    console.log(value);
    console.log(password == value);

    return password == value ? { passwordValid: true } : null;
  };
}

export function usernameValidator(signupService: any): ValidatorFn {
  let temp: any = null;
  return (control: AbstractControl): ValidationErrors | null => {
    const value = control.value;

    if (!value) {
      return null;
    }

    signupService.checkUsername(value).subscribe((response: any) => {
      if (response.data.usernameAlreadyExists == false) {
        temp = true;
      }
      else if (response.status == 200 && response.message == "success" && response.data.usernameAlreadyExists == true) {
        temp = false;
      }
    });
    return temp ? { usernameValid: true } : null;
  };
}


@Component({
  selector: 'app-signupform',
  templateUrl: './signupform.component.html',
  styleUrls: ['./signupform.component.css']
})

export class SignupformComponent implements OnInit {
  windowScrolled!: boolean;
  public showPassword: boolean = false;
  form: FormGroup = new FormGroup({});

  constructor(private router: Router, private snackBar: MatSnackBar, private signupService: SignupService, private fb: FormBuilder, @Inject(DOCUMENT) private document: Document ) {
    this.form = this.fb.group({
      firstname: ['', [Validators.required]],
      lastname: ['', [Validators.required]],
      username: ['', [Validators.required, usernameValidator(signupService)]],
      password: ['', [Validators.required, passwordValidator(), Validators.minLength(6)]],
      password2: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.email]]
    },
    {
      validator: checkIfMatchingPasswords('password', 'password2')
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


  ngOnInit(): void {
  }
  
  public togglePasswordVisibility(): void {
    this.showPassword = !this.showPassword;
  }

  get f() {
    return this.form.controls;
  }

  get password() {
    console.log("GETTING: "+ this.form.get('password'));
    return this.form.get('password');
  }

  getLogIn(username: string, password: string) {
    this.signupService.checkLogIn(username, password).subscribe((response: any) => {
      if (response.status == 200 && response.message == "success") {
        // LogIn Attempt Sucessful
        Storage.isLoggedIn = true;
        Storage.username = username;
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
  
  getSignUp(first: string, last: string, username: string, email: string, password: string): void {
    console.log(`sign up attempt with: ${first} ${last} ${username} ${email} ${password}`);
    this.signupService.addNewAccount(email.trim(), username.trim(), password.trim(), first.trim(), last.trim()).subscribe((response: any) => {
      console.log(response);
      if (response.status == 201 && response.message == "success") {
      // Log User In
      this.getLogIn(username, password);

        this.snackBar.open("Sign Up Successfull", "Dismiss", { duration: 1500});
      }
      
      else if (response.status == 200 && response.data.usernameAlreadyExists == true) {
          this.snackBar.open("Sign Up Failed. Username is already taken.", "Dismiss", { duration: 1500});
      }
    });
  }
}

