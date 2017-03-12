import { Component } from '@angular/core';
import {getTemplatePath} from "../../services/utility";

@Component({
    templateUrl: getTemplatePath("home.component.html")
})
export class HomeComponent {
    public pageTitle: string = 'Welcome';
}
