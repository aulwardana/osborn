import { Component } from '@angular/core';
import {getTemplatePath} from "../../services/utility";
import {Insert, InsertService} from "../../services/insert.services";

@Component({
    templateUrl: getTemplatePath("insert.component.html"),
    providers: [InsertService]
})
export class InsertComponent {
    public pageTitleInsert: string = 'Page Insert';

    kota = ['Malang', 'Bandung', 'Jakarta', 'Semarang'];

    inserts: Insert[];
    newInsert: Insert = new Insert();
    modifyingInsert: Insert;
    creating: boolean = false;

    constructor(
        private insertsService: InsertService) {}

    ngOnInit() {
        this.insertsService
            .list()
            .subscribe(
            inserts => this.inserts = inserts);
    }

    create() {
        if (this.creating && this.newInsert.name) {
        this.insertsService
            .create(this.newInsert)
            .subscribe(
                inserts => this.inserts = inserts,
                null,
                () => {
                this.creating = false;
                this.newInsert = new Insert();
                });

        return;
        }

        this.creating = !this.creating;
    }

    remove(id: string) {
        this.insertsService
            .remove(id)
            .subscribe(
            inserts => this.inserts = inserts);
    }

    update(insert: Insert) {
        if (insert.name) {
        this.insertsService
            .update(insert)
            .subscribe(
                inserts => this.inserts = inserts,
                null);
        }
    }
}
