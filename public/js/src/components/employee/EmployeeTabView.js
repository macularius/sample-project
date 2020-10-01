import { EMPLOYEE_CONTEXT_MENU } from "./CEmployeeTab.js"

// возвращает webix конфигурации таба для работы с сотрудникамиы
export function EmployeeTabView() {
    return {
        id: 'employeeTab',
        rows: [
            {
                view: 'datatable',
                id: 'employeeTabDatatable',
                select: true,
                columns: [
                    { id: 'ID', hidden: true },
                    { id: 'lastname', header: 'Фамилия', fillspace: true, },
                    { id: 'firstname', header: 'Имя', fillspace: true, },
                    { id: 'middlename', header: 'Отчество', fillspace: true, },
                    { id: 'position', header: 'Должность', fillspace: true, },
                    { id: 'phoneNumber', header: 'Телефонный номер', fillspace: true, },
                    { id: 'email', header: 'Email', fillspace: true, },
                ],
                data: [
                    { ID: 1, lastname: 'lastname', firstname: 'firstname', middlename: 'middlename', position: 'position', phoneNumber: 'phone_number', email: 'email', },
                    { ID: 2, lastname: 'lastname', firstname: 'firstname', middlename: 'middlename', position: 'position', phoneNumber: 'phone_number', email: 'email', },
                    { ID: 3, lastname: 'lastname', firstname: 'firstname', middlename: 'middlename', position: 'position', phoneNumber: 'phone_number', email: 'email', },
                ],
                onContext: {},
            },
        ]
    }
}

// возвращает webix конфигурации контекстного меню таба
export function EmployeeTabContextMenu() {
    return {
        view: 'contextmenu',
        id: 'employeeTabDatatableContextMenu',
        data: [EMPLOYEE_CONTEXT_MENU.toBooks, EMPLOYEE_CONTEXT_MENU.toEvents, EMPLOYEE_CONTEXT_MENU.add, EMPLOYEE_CONTEXT_MENU.edit, EMPLOYEE_CONTEXT_MENU.remove],
    }
}