import {NgModule} from '@angular/core';
import { FormsModule } from '@angular/forms';
import {BrowserModule} from '@angular/platform-browser';
import {HttpModule} from '@angular/http';

//Module
import {AppRoutingModule} from './router.module';
import { TableModule }   from './modules/table.module';
import { ChartsModule } from 'ng2-charts';

//Component
import {AppComponent} from './app.component';
import {HomeComponent} from './components/home/home.component';
import {InsertComponent} from './components/insert/insert.component';
import {ChartComponent} from './components/chart/chart.component';

//Service
import {AppService} from './services/app.service';

@NgModule({
  imports: [
    BrowserModule,
    HttpModule,
    AppRoutingModule,
    FormsModule,
    TableModule,
    ChartsModule
  ],
  declarations: [
    AppComponent,
    HomeComponent,
    InsertComponent,
    ChartComponent
  ],
  providers: [AppService],
  bootstrap: [AppComponent]
})

export class AppModule {}