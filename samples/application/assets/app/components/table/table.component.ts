import { Component } from '@angular/core';
import {getTemplatePath} from "../../services/utility";
import {Http} from "@angular/http";

@Component({
    templateUrl: getTemplatePath("table.component.html")
})
export class TableComponent {
    public pageTitleTable: string = 'Page Table';
    public data;
    public filterQuery = "";
    public rowsOnPage = 10;
    public sortBy = "email";
    public sortOrder = "asc";

    constructor(private http: Http) {
    }

    ngOnInit(): void {
        this.http.get("/testGetAPI")
            .subscribe((data)=> {
                setTimeout(()=> {
                    this.data = data.json();
                }, 1000);
            });
    }

    public toInt(num: string) {
        return +num;
    }

    public sortByWordLength = (a: any) => {
        return a.city.length;
    }
}
