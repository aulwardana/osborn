import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {HttpModule} from '@angular/http';

import {AppRoutingModule} from './router.module';

import {AppComponent} from './app.component';
import {HomeComponent} from './components/home/home.component';
import {InsertComponent} from './components/insert/insert.component';
import {ChartComponent} from './components/chart/chart.component';

import { TableModule }   from './components/table/table.module';
import { ChartsModule } from 'ng2-charts';

@NgModule({
  imports: [
    BrowserModule,
    HttpModule,
    AppRoutingModule,
    TableModule,
    ChartsModule
  ],
  declarations: [
    AppComponent,
    HomeComponent,
    InsertComponent,
    ChartComponent
  ],
  bootstrap: [AppComponent]
})

export class AppModule {}