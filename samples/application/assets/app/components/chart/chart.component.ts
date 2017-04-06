import { Component } from '@angular/core';
import {getTemplatePath} from "../../services/utility";

@Component({
    templateUrl: getTemplatePath("chart.component.html")
})
export class ChartComponent {
    public pageTitleChart: string = 'Page Chart';
}
