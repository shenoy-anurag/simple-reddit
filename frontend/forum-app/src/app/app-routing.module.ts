import { PopupMessageComponent } from './popup-message/popup-message.component';
import { SignupformComponent } from './signupform/signupform.component';
import { HeaderComponent } from './header/header.component';
import { NgModule, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { SubredditsComponent } from './subreddits/subreddits.component';
import { ProfileComponent } from './profile/profile.component';


const routes: Routes = [
  {path: 'header', component: HeaderComponent},
  {path: 'login', component: LoginComponent},
  {path: 'home', component: HomeComponent},
  {path: 'subreddits', component: SubredditsComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'signupform', component: SignupformComponent},
  {path: 'popup-message', component: PopupMessageComponent},
  {path: '', redirectTo: '/home', pathMatch:'full'},
];



@NgModule({
  declarations: [],
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
