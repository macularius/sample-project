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
        position:'center',
        width: 400,
        body: {
            view: 'form',
            id: 'employeeWindowForm',
            elements: [
                {
                    view: 'text',
                    label: 'Фамилия',
                    name: 'lastname',
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Имя',
                    name: 'firstname',
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Отчество',
                    name: 'middlename',
                    labelWidth: 150,
                },
                {
                    view:'richselect',
                    id: 'employeeWindowFormPosition',
                    label: 'Должность',
                    name: 'position',
                    labelWidth: 150,
                    options:[]
                },
                {
                    view: 'text',
                    label: 'Телефонный номер',
                    name: 'phoneNumber',
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Электронная почта',
                    name: 'email',
                    labelWidth: 150,
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