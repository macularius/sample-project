import { CEmployeeWindow } from "./employeeWindow/CEmployeeWindow";

export class CEmployeeTab {
    constructor() {
        this.#view
        this.#window
    }

    // метод инициализации компонента
    init() {
        this.#window = new CEmployeeWindow();
        this.#window.init()
    }

    // метод получения webix конфигурации компонента
    config() {
        EmployeeTabView(this.#window.config())
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {}
    
    // метод вызова модального окна
    #showByID(id) {}
}


