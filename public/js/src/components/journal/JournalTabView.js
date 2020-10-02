import { EVENT_CONTEXT_MENU } from "./CJournalTab.js";

// возвращает webix конфигурации таба для работы с событиями
export function JournalTabView() {
    return {
        id: 'journalTab',
        rows: [
            {
                view: 'datatable',
                id: 'eventTabDatatable',
                select: true,
                columns: [
                    { id: 'bookString', header: 'Книга', fillspace: true, },
                    { id: 'employeeString', header: 'Сотрудник', fillspace: true, },
                    { id: 'type', header: 'Тип события', fillspace: true, },
                    { id: 'dateString', header: 'Дата события', fillspace: true, },
                ],
                data: [],
                onContext: {
                    // обработка вызова контекстного меню при остутствии данных
                    webix_view: function (e, id) {
                        id = this.locate(e.target || e.srcElement);
                        if (!id) {
                            $$("eventTabDatatableContextMenu").setContext({ obj: webix.$$(e) });
                            $$("eventTabDatatableContextMenu").show(e);
                            webix.html.preventEvent(e);
                        }
                    }
                },
            },
        ]
    }
}

// возвращает webix конфигурацию контекстного меню таба
export function EventTabContextMenu() {
    return {
        view: 'contextmenu',
        id: 'eventTabDatatableContextMenu',
        data: [EVENT_CONTEXT_MENU.toEmployee, EVENT_CONTEXT_MENU.toBook],
    }
}