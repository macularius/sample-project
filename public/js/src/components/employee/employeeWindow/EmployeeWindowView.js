import { EMPLOYEE_WINDOW_TYPE } from "./CEmployeeWindow.js";

// возвращает webix конфигурации окна для работы с сущностью сотрудника
export default function EmployeeWindowView(type) {
    let headText = 'Сотрудник'

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
                    label: 'ISBN',
                },
                {
                    view: 'text',
                    label: 'Название',
                },
                {
                    view: 'text',
                    label: 'Автор',
                },
                {
                    view: 'text',
                    label: 'Издательство',
                },
                {
                    view: 'text',
                    label: 'Статус',
                },
                {
                    view: 'text',
                    label: 'Год издания',
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