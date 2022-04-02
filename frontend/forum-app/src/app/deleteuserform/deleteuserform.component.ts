import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { SignupService } from '../signup.service';
import { Storage } from '../storage';

@Component({
  selector: 'app-deleteuserform',
  templateUrl: './deleteuserform.component.html',
  styleUrls: ['./deleteuserform.component.css']
})
export class DeleteuserformComponent implements OnInit {

  username:string = "";
  form: FormGroup = new FormGroup({});
  constructor(private router: Router, private signupService: SignupService, private snackBar: MatSnackBar, private fb: FormBuilder) {
    this.form = this.fb.group({
      username: ['', [Validators.required]],
      password: ['', [Validators.required]]
    })
   }

  ngOnInit(): void {
    this.username = this.getUsername();
  }

  getUsername() {
    if (Storage.isLoggedIn) {
      return Storage.username;
    }
    else {
      return "";
    }
  }

  get f() {
    return this.form.controls;
  }
  deleteUser(username: string, password: string) {
    this.signupService.deleteUser(username, password).subscribe((response: any) => {
      if (response.status == 200 && response.message == "success") {
        // Route to Home
        this.router.navigate(['/home']);
        this.snackBar.open(username + " deleted", "Dismiss");
      }
      else {
        this.snackBar.open("Failed to delete " + username, "Dismiss", { duration: 1500 });
      }
    });
  }

  // checkAuth(username: string, password: string) {
  //   this.signupService.checkLogIn(username, password).subscribe((response: any) => {

  //     if (response.status == 200 && response.message == "success") {
  //       this.deleteUser(username, password);
  //     }
  //     else if (response.status == 200 && response.message == "failure") {
  //       // Prompt user, incorrect login
  //       this.snackBar.open("Incorrect Credentials", "Dismiss", { duration: 2000 });
  //     }
  //     else {
  //       // Something else is wrong
  //       this.snackBar.open("Something is wrong", "Alert Adminstration"), { duration: 2000 };
  //     }
  //   })
  // }
}
