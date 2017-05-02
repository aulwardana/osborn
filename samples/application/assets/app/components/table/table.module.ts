import { NgModule }      from '@angular/core';
import { CommonModule }      from '@angular/common';
import { FormsModule } from "@angular/forms";
import { DataTableModule } from "angular2-datatable";
import { HttpModule } from "@angular/http";

import { TableComponent }   from './table.component';
import { DataFilterPipe }   from "../../services/data-filter-pipe";

@NgModule({
  imports:      [ 
    CommonModule, 
    DataTableModule, 
    FormsModule,
    HttpModule
  ],
  declarations: [ TableComponent, DataFilterPipe ],
  exports: [TableComponent]
})

export class TableModule { }