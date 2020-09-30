import EmployeeWindowView from "./EmployeeWindowView"

// компонент окна для работы с сущностью сотрудника
export class CEmployeeWindow {
    constructor() {
        this.#view
        this.#isShow
    }

    // метод инициализации компонента
    init() { }

    // метод получения webix конфигурации компонента
    config() {
        return EmployeeWindowView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() { }

    // метод вызова модального окна
    switch(type) {
        switch (this.#isShow) {
            case true:
                this.#hide()
                break;
            case false:
                this.#show(type)
                break;
        }
    }

    // метод отображения окна
    #show(type) {
        switch (type) {
            case EMPLOYEE_WINDOW_TYPE.create:

                break;
            case EMPLOYEE_WINDOW_TYPE.update:

                break;
            case EMPLOYEE_WINDOW_TYPE.delete:
                
                break;
            default:
                console.error('Неизвестный тип отображения окна для рабоыт с сущностью книги');
                break;
        }
    }

    // метод сокрытия окна
    #hide() { }
}

// типы отображения модального окна для работы с сущностью книги
export const EMPLOYEE_WINDOW_TYPE = {
    create: 'CREATE',
    update: 'UPDATE',
    delete: 'DELETe',
}

