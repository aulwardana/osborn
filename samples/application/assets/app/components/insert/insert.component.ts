import { Component } from '@angular/core';
import {getTemplatePath} from "../../services/utility";

@Component({
    templateUrl: getTemplatePath("insert.component.html")
})
export class InsertComponent {
    public pageTitleInsert: string = 'Page Insert';

    kota = ['Malang', 'Bandung', 'Jakarta', 'Semarang'];
}
