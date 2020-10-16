import { EVENT_CONTEXT_MENU } from './CJournalTab.js';

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
                    { id: 'ID', header: ['', { content: 'textFilter' }], hidden: true },
                    { id: 'bookString', header: ['Книга', { content: 'textFilter' }], fillspace: true, },
                    { id: 'employeeString', header: ['Сотрудник', { content: 'textFilter' }], fillspace: true, },
                    { id: 'type', header: ['Тип события', { content:'selectFilter' }], width: 120, },
                    { id: 'dateString', header: ['Дата события', { content: 'textFilter' }], width: 160, },
                ],
                data: [],
                onContext: {
                    // обработка вызова контекстного меню при остутствии данных
                    webix_view: function (e, id) {
                        id = this.locate(e.target || e.srcElement);
                        if (!id) {
                            $$('eventTabDatatableContextMenu').setContext({ obj: webix.$$(e) });
                            $$('eventTabDatatableContextMenu').show(e);
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