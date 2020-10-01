import EmployeeWindowView from "./EmployeeWindowView.js"

// компонент окна для работы с сущностью сотрудника
export class CEmployeeWindow {
    constructor() {
        this.view
        this.isShow = false
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
            form: $$('employeeWindowForm')
        }

        // обрабтка закрытия окна
        this.view.windowCancelBtn.attachEvent('onItemClick', () => {
            this.view.window.hide()
        })
    }

    // метод вызова модального окна
    switch(type) {
        switch (this.isShow) {
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

        this.view.window.show()
        this.isShow = true
    }

    // метод сокрытия окна
    hide() { 
        this.view.window.hide()
        this.isShow = false
    }
}

// типы отображения модального окна для работы с сущностью книги
export const EMPLOYEE_WINDOW_TYPE = {
    create: 'CREATE',
    update: 'UPDATE',
    delete: 'DELETe',
}

