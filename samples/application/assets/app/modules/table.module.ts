import { NgModule }      from '@angular/core';
import { CommonModule }      from '@angular/common';
import { FormsModule } from "@angular/forms";
import { DataTableModule } from "angular2-datatable";
import { HttpModule } from "@angular/http";

import { TableComponent }   from "../components/table/table.component";
import { DataFilterPipe }   from "../services/table-data-filter.service";

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