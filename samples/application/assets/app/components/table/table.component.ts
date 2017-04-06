import { Component } from '@angular/core';
import {getTemplatePath} from "../../services/utility";

@Component({
    templateUrl: getTemplatePath("table.component.html")
})
export class TableComponent {
    public pageTitleTable: string = 'Page Table';
}
