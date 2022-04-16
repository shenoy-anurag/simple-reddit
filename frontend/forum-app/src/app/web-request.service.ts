import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class WebRequestService {

  readonly ROOT_URL: string;
  constructor(private http: HttpClient) { 
    // this.ROOT_URL = 'http://3.145.146.32:27017';
    // this.ROOT_URL = 'http://3.145.146.32:8080'
    this.ROOT_URL = 'http://localhost:8080';
  }

  get(uri: string, quertyParams: HttpParams) {
    return this.http.get(`${this.ROOT_URL}/${uri}`, { params: quertyParams});
  }

  post(uri: string, payload: Object) {
    return this.http.post(`${this.ROOT_URL}/${uri}`, payload);
  }

  patch(uri: string, payload: Object) {
    return this.http.patch(`${this.ROOT_URL}/${uri}`, payload);
  }
  
  // May not work with payload
  delete(uri: string, payload: Object) {
    return this.http.delete(`${this.ROOT_URL}/${uri}`, payload);
  }
}
