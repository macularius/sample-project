import EmployeeWindowView from "./EmployeeWindowView.js"
import employeeModel from "../../../models/employeeModel.js"

// компонент окна для работы с сущностью сотрудника
export class CEmployeeWindow {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
        this.type       // тип текущего отображения окна
        this.onChange   // callback функция при CUD операциях над сотрудником
    }

    // метод инициализации компонента
    init() { }

    // метод получения webix конфигурации компонента
    config() {
        return EmployeeWindowView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() { 
        // инициализация используемых представлений
        this.view = {
            window: $$('employeeWindow'),
            windowLabel: $$('employeeWindowLabel'),
            windowCancelBtn: $$('employeeWindowCancelBtn'),
            windowConfirmBtn: $$('employeeWindowConfirmBtn'),
            form: $$('employeeWindowForm')
        }

        // обрабтка закрытия окна
        this.view.windowCancelBtn.attachEvent('onItemClick', () => {
            this.view.window.hide()
        })

        // обработка события "принять"
        this.view.windowConfirmBtn.attachEvent('onItemClick', () => {
            switch (this.type) {
                case EMPLOYEE_WINDOW_TYPE.create:
                    employeeModel.createEmployee(this.fetch()).then(() => {
                        this.onChange()
                        this.hide()
                    })
                    break;
                case EMPLOYEE_WINDOW_TYPE.update:
                    employeeModel.updateEmployee(this.fetch()).then(() => {
                        this.onChange()
                        this.hide()
                    })
                    break;
                case EMPLOYEE_WINDOW_TYPE.delete:
                    employeeModel.deleteEmployee(this.fetch()).then(() => {
                        this.onChange()
                        this.hide()
                    })
                    break;
            }
        })
    }

    // метод вызова модального окна
    switch(type) {
        switch (this.view.window.isVisible()) {
            case true:
                this.hide()
                break;
            case false:
                this.show(type)
                break;
        }
    }

    // метод отображения окна
    show(type) {
        switch (type) {
            case EMPLOYEE_WINDOW_TYPE.create:
                this.view.windowLabel.setHTML('Добавление сотрудника')
                break;
            case EMPLOYEE_WINDOW_TYPE.update:
                this.view.windowLabel.setHTML('Редактирование сотрудника')
                break;
            case EMPLOYEE_WINDOW_TYPE.delete:
                this.view.windowLabel.setHTML('Удаление сотрудника')
                break;
            default:
                console.error('Неизвестный тип отображения окна для работы с сущностью сотрудника');
                break;
        }

        this.type = type
        this.view.window.show()
    }

    // метод сокрытия окна
    hide() { 
        this.view.window.hide()
    }

    // метод получения сущности из формы окна
    fetch() {
        return this.view.form.getValues()
    }

    // метод размещения сущности в форме окна
    parse(values) {
        this.view.form.setValues(values)
    }
}

// типы отображения модального окна для работы с сущностью книги
export const EMPLOYEE_WINDOW_TYPE = {
    create: 'CREATE',
    update: 'UPDATE',
    delete: 'DELETe',
}

