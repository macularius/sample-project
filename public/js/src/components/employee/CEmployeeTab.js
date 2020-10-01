import { EmployeeTabView, EmployeeTabContextMenu } from "./EmployeeTabView.js";
import { CEmployeeWindow, EMPLOYEE_WINDOW_TYPE } from "./employeeWindow/CEmployeeWindow.js";

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
        webix.ui(EmployeeTabContextMenu())

        return EmployeeTabView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            datatable: $$('employeeTabDatatable'),
            datatableContextMenu: $$('employeeTabDatatableContextMenu'),
        }

        // инициализация обработчиков событий модального окна
        this.window.attachEvents()

        // прикрепление контекстного меню к таблице
        this.view.datatableContextMenu.attachTo(this.view.datatable)

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (id) => {
            let item = this.view.datatableContextMenu.getItem(id).value
            switch (item) {
                case EMPLOYEE_CONTEXT_MENU.toEvent:
                    console.log(EMPLOYEE_CONTEXT_MENU.toEvent)
                    break;
                case EMPLOYEE_CONTEXT_MENU.toEmployee:
                    console.log(EMPLOYEE_CONTEXT_MENU.toEmployee)
                    break;
                case EMPLOYEE_CONTEXT_MENU.add:
                    this.window.switch(EMPLOYEE_WINDOW_TYPE.create)
                    break;
                case EMPLOYEE_CONTEXT_MENU.edit:
                    if (!this.view.datatable.getSelectedItem()) {
                        webix.message('Выделите строку')
                        return
                    }
                    this.window.switch(EMPLOYEE_WINDOW_TYPE.update)
                    break;
                case EMPLOYEE_CONTEXT_MENU.remove:
                    if (!this.view.datatable.getSelectedItem()) {
                        webix.message('Выделите строку')
                        return
                    }
                    this.window.switch(EMPLOYEE_WINDOW_TYPE.delete)
                    break;
                default:
                    console.error(`Неизвестное значение пункта меню: ${item}.`);
                    break;
            }
        });
    }
    
    // метод вызова модального окна
    showByID(id) {}
}

// допустимые значения пунктов контекстного меню таба Книги
export const EMPLOYEE_CONTEXT_MENU = {
    toEvents: 'К событиям',
    toBooks: 'К книгам',
    add: 'Добавить',
    edit: 'Изменить',
    remove: 'Удалить'
}