import employeeModel from '../../../models/employeeModel.js';
import positionModel from '../../../models/positionModel.js';
import EmployeeWindowView from './EmployeeWindowView.js';

// компонент окна для работы с сущностью сотрудника
export class CEmployeeWindow {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
        this.type       // тип текущего отображения окна
        this.onChange   // callback функция при CUD операциях над сотрудником
    }

    // метод инициализации компонента
    init(onChange) {
        this.onChange = onChange
    }

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
            form: $$('employeeWindowForm'),
            formfields: {
                lastname: $$('employeeWindowFormLastname'),
                firstname: $$('employeeWindowFormFirstname'),
                middlename: $$('employeeWindowFormMiddlename'),
                position: $$('employeeWindowFormPosition'),
                phoneNumber: $$('employeeWindowFormPhoneNumber'),
                email: $$('employeeWindowFormEmail'),
            }
        }

        // подгрузка должностей
        positionModel.getPositions().then((positions) => {
            positions.map((position) => {
                position.id = position.name
                position.value = position.name
            })

            this.view.formfields.position.define('options', positions)
            this.view.formfields.position.refresh()
        })

        // обрабтка закрытия окна
        this.view.windowCancelBtn.attachEvent('onItemClick', () => {
            this.view.window.hide()
        })

        // обработка события 'принять'
        this.view.windowConfirmBtn.attachEvent('onItemClick', () => {
            // при удалении не требуется валидировать данные формы
            // валидация введенных данных по обязательным полям
            if (this.type !== EMPLOYEE_WINDOW_TYPE.delete && !this.validate()) {
                webix.message('Заполните поля отмеченные *', 'error')
                return;
            }

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
                    // получение сотрудника
                    let employee = this.fetch()
                    employeeModel.getCardByEmployeeID(employee.ID).then((books) => {
                        // проверка наличия несданных книг
                        if (books) {
                            webix.message('Нельзя удалить сотрудника, который не сдал книги')
                            return
                        }

                        // удаление сотрудника
                        employeeModel.deleteEmployee(this.fetch()).then(() => {
                            this.onChange()
                            this.hide()
                        })
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
                this.view.formfields.lastname.disable()
                this.view.formfields.firstname.disable()
                this.view.formfields.middlename.disable()
                this.view.formfields.position.disable()
                this.view.formfields.phoneNumber.disable()
                this.view.formfields.email.disable()
                this.view.windowLabel.setHTML('Удаление сотрудника')
                this.view.window.resize()
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

    // функция валидации формы
    validate() {
        let isValid = false

        // удаление пробелов в полях формы
        this.view.formfields.lastname.setValue(this.view.formfields.lastname.getValue().trim())
        this.view.formfields.firstname.setValue(this.view.formfields.firstname.getValue().trim())
        this.view.formfields.middlename.setValue(this.view.formfields.middlename.getValue().trim())
        this.view.formfields.position.setValue(this.view.formfields.position.getValue().trim())
        this.view.formfields.phoneNumber.setValue(this.view.formfields.phoneNumber.getValue().trim())
        this.view.formfields.email.setValue(this.view.formfields.email.getValue().trim())

        // валидация webix
        isValid = this.view.form.validate()

        return isValid
    }
}

// типы отображения модального окна для работы с сущностью книги
export const EMPLOYEE_WINDOW_TYPE = {
    create: 'CREATE',
    update: 'UPDATE',
    delete: 'DELETe',
}

