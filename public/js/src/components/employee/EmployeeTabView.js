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
                    { id: 'ID', header: ['', { content: "textFilter" }], hidden: true },
                    { id: 'lastname', header: ['Фамилия', { content: "textFilter" }], fillspace: true, },
                    { id: 'firstname', header: ['Имя', { content: "textFilter" }], fillspace: true, },
                    { id: 'middlename', header: ['Отчество', { content: "textFilter" }], fillspace: true, },
                    { id: 'position', header: ['Должность', { content: "textFilter" }], fillspace: true, },
                    { id: 'phone_number', header: ['Телефонный номер', { content: "textFilter" }], fillspace: true, },
                    { id: 'email', header: ['Email', { content: "textFilter" }], fillspace: true, },
                ],
                data: [],
                onContext: {
                    // обработка вызова контекстного меню при остутствии данных
                    webix_view: function (e, id) {
                        id = this.locate(e.target || e.srcElement);
                        if (!id) {
                            $$("employeeTabDatatableContextMenu").setContext({ obj: webix.$$(e) });
                            $$("employeeTabDatatableContextMenu").show(e);
                            webix.html.preventEvent(e);
                        }
                    }
                },
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
            EMPLOYEE_CONTEXT_MENU.add, 
            EMPLOYEE_CONTEXT_MENU.edit, 
            EMPLOYEE_CONTEXT_MENU.remove
        ],
    }
}