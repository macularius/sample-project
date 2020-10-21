import { EMPLOYEE_CONTEXT_MENU } from './CEmployeeTab.js'

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
                    { id: 'ID', header: ['', { content: 'textFilter' }], hidden: true },
                    { id: 'lastname', header: ['Фамилия', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'firstname', header: ['Имя', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'middlename', header: ['Отчество', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'position', header: ['Должность', { content:'selectFilter' }], sort: 'string', fillspace: true, },
                    { id: 'phone_number', header: ['Телефонный номер', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'email', header: ['Email', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                ],
                data: [],
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
        data: [
            EMPLOYEE_CONTEXT_MENU.libraryCard,
            EMPLOYEE_CONTEXT_MENU.edit, 
            EMPLOYEE_CONTEXT_MENU.remove
        ],
    }
}

// элементы управления для таба
export function TabControllsView() {
    return {
        id: 'employeetab-controlls',
        hidden: true,
        cols: [
            {
                id: 'employeetab-card-btn',
                view: 'icon',
                tooltip: 'Читательский билет',
                icon: 'book',
                width: 30,
            },
            {
                id: 'employeetab-add-btn',
                view: 'icon',
                tooltip: 'Добавить',
                icon: 'plus',
                width: 30,
            },
            {
                id: 'employeetab-edit-btn',
                view: 'icon',
                tooltip: 'Редактировать',
                icon: 'pencil',
                width: 30,
            },
            {
                id: 'employeetab-remove-btn',
                view: 'icon',
                tooltip: 'Удалить',
                icon: 'trash',
                width: 30,
            },
            { width: 30 },
        ]
    }
}