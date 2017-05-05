import {Injectable} from '@angular/core';
import {Http} from '@angular/http';

import {Observable} from 'rxjs/Observable';

import {AppService} from './app.service';


export class Insert {
  constructor(
    public id: string = '',
    public name: string = '',
    public email: string = '',
    public age: number = 0,
    public city: string = '') {}
}

@Injectable()
export class InsertService {
  private prefix = '/insert';

  constructor(private service: AppService) {
  }

  list(): Observable<Insert[]> {
    return this.service.get(this.prefix);
  }

  create(newInsert: Insert): Observable<Insert[]> {
    return this.service.post(this.prefix, newInsert);
  }

  remove(id: string): Observable<Insert[]> {
    return this.service.delete(`${this.prefix}/${id}`);
  }

  update(insert: Insert): Observable<Insert[]> {
    return this.service.put(`${this.prefix}/`, insert);
  }
}