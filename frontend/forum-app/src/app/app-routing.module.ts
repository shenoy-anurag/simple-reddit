import { SignupformComponent } from './signupform/signupform.component';
import { NgModule, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';
import { SubredditsComponent } from './subreddits/subreddits.component';
import { ProfileComponent } from './profile/profile.component';
import { DeleteuserformComponent } from './deleteuserform/deleteuserform.component';
import { NewpostformComponent } from './newpostform/newpostform.component';
import { NewsubredditsformComponent} from './newsubredditsform/newsubredditsform.component';
import { DeletesubredditsformComponent } from './deletesubredditsform/deletesubredditsform.component';
import { TermsandconditionsComponent } from './termsandconditions/termsandconditions.component';
import { PrivacypolicyComponent } from './privacypolicy/privacypolicy.component';
import { PostpageComponent } from './postpage/postpage.component';
import { CommunitypageComponent } from './communitypage/communitypage.component';
import { ContentpolicyComponent } from './contentpolicy/contentpolicy.component';
import { ModpolicyComponent } from './modpolicy/modpolicy.component';

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'home', component: HomeComponent},
  {path: 'subreddits', component: SubredditsComponent},
  {path: 'profile', component: ProfileComponent},
  {path: 'signupform', component: SignupformComponent},
  {path: 'delete-user', component: DeleteuserformComponent},
  {path: 'newpostform', component: NewpostformComponent},
  {path: 'newsubredditsform', component: NewsubredditsformComponent},
  {path: 'deletesubredditsform', component: DeletesubredditsformComponent},
  {path: 'termsandconditions', component: TermsandconditionsComponent},
  {path: 'privacypolicy' , component: PrivacypolicyComponent},
  {path:'contentpolicy' , component: ContentpolicyComponent},
  {path: 'modpolicy' , component: ModpolicyComponent},
  {path: 'post/:postID', component: PostpageComponent},
  {path: 'subreddits/:communityID', component: CommunitypageComponent},
  {path: '', redirectTo: '/home', pathMatch:'full'},
];



@NgModule({
  declarations: [],
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
