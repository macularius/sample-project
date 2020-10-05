import { EmployeeTabView, EmployeeTabContextMenu } from "./EmployeeTabView.js";
import { CEmployeeWindow, EMPLOYEE_WINDOW_TYPE } from "./employeeWindow/CEmployeeWindow.js";
import { Employee } from "../../models/entities/employee.js";
import employeeModel from "../../models/employeeModel.js";

export class CEmployeeTab {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
        this.window     // экземпляр окна для работы с книгами
        this.toEvents   // функция перехода к событиям
        this.toBooks    // функция перехода к книгам
    }

    // метод инициализации компонента
    init() {
        this.window = new CEmployeeWindow(); // инициализация компонента окна
        this.window.onChange = () => { this.refreshTable() } // установка метода, который будет вызван после нажатия на кнопку confirm
        this.window.init() // вызова инициализации компонента окна
    }

    // метод получения webix конфигурации компонента
    config() {
        // т.к. window и popup расположены не в дереве приложения, а поверх слоев, его нужно отрисовывать отдельно
        webix.ui(this.window.config())
        webix.ui(EmployeeTabContextMenu())

        // вызов функции представления
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

        // загрузка первичных данных в таблицу
        this.refreshTable()

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (id) => {
            this.handleContextMenu(id)
        });
    }

    // обработка выбора в контекстном меню
    handleContextMenu(id) {
        // получение значения пункта, на которое произошло нажатие
        let item = this.view.datatableContextMenu.getItem(id).value
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()

        switch (item) {
            case EMPLOYEE_CONTEXT_MENU.add: // добавление сотрудника
                this.window.parse(new Employee())
                this.window.switch(EMPLOYEE_WINDOW_TYPE.create)
                break;
            case EMPLOYEE_CONTEXT_MENU.edit: // редактирование сотрудника
                if (!selected) {
                    webix.message('Выделите строку')
                    return
                }
                if (!selected.ID) {
                    console.error('Incorrect ID of item:', selected.ID)
                    return
                }
                employeeModel.getEmployeeByID(selected.ID).then((employee) => {
                    this.window.parse(employee)
                    this.window.switch(EMPLOYEE_WINDOW_TYPE.update)
                })
                break;
            case EMPLOYEE_CONTEXT_MENU.remove: // удаление сотрудника
                if (!selected) {
                    webix.message('Выделите строку')
                    return
                }
                if (!selected.ID) {
                    console.error('Incorrect ID of item:', selected.ID)
                    return
                }
                employeeModel.getBookByID(selected.ID).then((employee) => {
                    this.window.parse(employee)
                    this.window.switch(EMPLOYEE_WINDOW_TYPE.delete)
                })
                break;
            default:
                console.error(`Неизвестное значение пункта меню: ${item}.`);
                break;
        }
    }

    // функция обновления таблицы сотрудников
    refreshTable(employees) {
        if (employees) {
            this.view.datatable.clearAll()
            this.view.datatable.parse(employees)
            return
        } else {
            employeeModel.getEmployees().then((employees) => {
                this.view.datatable.clearAll()
                this.view.datatable.parse(employees)
            })
        }
    }

    // метод отображения таба с фильтрацией по сотруднику
    showByEmployeeID(employeeID) {
        employeeModel.getEmployeeByID(employeeID).then((employee) => {
            // рпименение фильтров
            this.view.datatable.getFilter('lastname').value = employee.lastname;
            this.view.datatable.getFilter('firstname').value = employee.firstname;
            this.view.datatable.getFilter('middlename').value = employee.middlename;
            this.view.datatable.getFilter('position').value = employee.position;
            this.view.datatable.getFilter('phoneNumber').value = employee.phoneNumber;
            this.view.datatable.getFilter('email').value = employee.email;
            this.view.datatable.filterByAll();

            // выделение нужной строки
            for (let rowID = 0; rowID < this.view.datatable.serialize().length; rowID++) {
                let item = this.view.datatable.serialize()[rowID]

                if (item.ID === employeeID) {
                    this.view.datatable.select(item.id)
                    break
                }
            }
        })
    }
}

// допустимые значения пунктов контекстного меню таба Книги
export const EMPLOYEE_CONTEXT_MENU = {
    add: 'Добавить',
    edit: 'Изменить',
    remove: 'Удалить'
}