import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {HttpModule} from '@angular/http';

import {AppRoutingModule} from './router.module';

import {AppComponent} from './app.component';
import {HomeComponent} from './components/home/home.component';
import {InsertComponent} from './components/insert/insert.component';

@NgModule({
  imports: [
    BrowserModule,
    HttpModule,
    AppRoutingModule
  ],
  declarations: [
    AppComponent,
    HomeComponent,
    InsertComponent 
  ],
  bootstrap: [AppComponent]
})

export class AppModule {}