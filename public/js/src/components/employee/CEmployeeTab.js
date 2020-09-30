import EmployeeTabView from "./EmployeeTabView.js";
import { CEmployeeWindow } from "./employeeWindow/CEmployeeWindow.js";

export class CEmployeeTab {
    constructor() {
        this.view
        this.window
    }

    // метод инициализации компонента
    init() {
        this.window = new CEmployeeWindow();
        this.window.init()
    }

    // метод получения webix конфигурации компонента
    config() {
        webix.ui(this.window.config())
        return EmployeeTabView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {}
    
    // метод вызова модального окна
    showByID(id) {}
}


