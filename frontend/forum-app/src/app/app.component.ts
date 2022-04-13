import { Component, HostListener } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Simple Reddit';
  public windowWidth: any;
  public windowHeight: any;

ngOnInit() {
  this.windowWidth = window.innerWidth;
  this.windowHeight = window.innerHeight;
}

@HostListener('window:resize', ['$event'])

resizeWindow() {
this.windowWidth = window.innerWidth;
this.windowHeight = window.innerHeight;
}

// interface MediaQueryList extends EventTarget {
//   matches: boolean; // => true if document matches the passed media query, false if not
//   media: string; // => the media query used for the matching
// }

}