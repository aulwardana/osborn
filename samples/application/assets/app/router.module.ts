import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

import { HomeComponent } from './components/home/home.component';
import { InsertComponent } from './components/insert/insert.component';
import { ChartComponent } from './components/chart/chart.component';
import { TableComponent } from './components/table/table.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'insert', component: InsertComponent },
  { path: 'chart', component: ChartComponent },
  { path: 'table', component: TableComponent },
  { path: '', redirectTo: '', pathMatch: 'full' },
  { path: '**', redirectTo: '', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule {}