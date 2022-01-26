import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon'
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatFormFieldModule } from '@angular/material/form-field';


@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    MatButtonModule,
    MatIconModule,
    MatButtonToggleModule,
    MatSlideToggleModule,
    MatFormFieldModule
  ],
  exports: [
    MatButtonModule,
    MatIconModule,
    MatButtonToggleModule,
    MatSlideToggleModule,
    MatFormFieldModule
  ]
})
export class NgMaterialModule { }