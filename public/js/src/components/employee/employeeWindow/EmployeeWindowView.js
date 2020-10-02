import { EMPLOYEE_WINDOW_TYPE } from "./CEmployeeWindow.js";

// возвращает webix конфигурации окна для работы с сущностью сотрудника
export default function EmployeeWindowView(type) {
    let headText = 'Сотрудник' // текст заголовка модального окна

    switch (type) {
        case EMPLOYEE_WINDOW_TYPE.create:
            headText = 'Добавление сотрудника'
            break;
        case EMPLOYEE_WINDOW_TYPE.update:
            headText = 'Редактирование сотрудника'
            break;
        case EMPLOYEE_WINDOW_TYPE.delete:
            headText = 'Удаление сотрудника'
            break;
    }

    return { 
        view: 'window',
        id: 'employeeWindow',
        head: {
            view: 'template',
            id: 'employeeWindowLabel',
            template: headText,
            css: 'webix_template'
        },
        modal: true,
        position:"center",
        body: {
            view: 'form',
            id: 'employeeWindowForm',
            elements: [
                {
                    view: 'text',
                    label: 'Фамилия',
                    name: 'lastname',
                },
                {
                    view: 'text',
                    label: 'Имя',
                    name: 'firstname',
                },
                {
                    view: 'text',
                    label: 'Отчество',
                    name: 'middlename',
                },
                {
                    view: 'text',
                    label: 'Должность',
                    name: 'position',
                },
                {
                    view: 'text',
                    label: 'Телефонный номер',
                    name: 'phoneNumber',
                },
                {
                    view: 'text',
                    label: 'Электронная почта',
                    name: 'email',
                },
                {
                    cols: [
                        {
                            view: 'button',
                            id: 'employeeWindowConfirmBtn',
                            value: 'Применить',
                        },
                        {
                            view: 'button',
                            id: 'employeeWindowCancelBtn',
                            value: 'Отмена',
                        },
                    ]
                },
            ]
        }
    }
}