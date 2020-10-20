import { BOOK_CONTEXT_MENU } from './CBookTab.js';

// возвращает webix конфигурацию таба для работы с книгами
export function BookTabView() {
    return {
        id: 'bookTab',
        rows: [
            {
                view: 'datatable',
                id: 'bookTabDatatable',
                select: true,
                columns: [
                    { id: 'ID', header: ['', { content: 'textFilter' }], hidden: true },
                    { id: 'ISBN', header: ['ISBN', { content: 'textFilter' }], sort: 'string', width: 140, },
                    { id: 'name', header: ['Название', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'author', header: ['Автор', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'publisher', header: ['Издательство', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'year', header: ['Год издания', { content: 'textFilter' }], sort: 'date', format: webix.Date.dateToStr("%Y"), width: 80, },
                    { id: 'status', header: ['Статус', { content: 'selectFilter' }], sort: 'string', width: 120, },
                ],
                data: [],
                onContext: {
                    // обработка вызова контекстного меню при остутствии данных
                    webix_view: function (e, id) {
                        id = this.locate(e.target || e.srcElement);
                        if (!id) {
                            $$('bookTabDatatableContextMenu').setContext({ obj: webix.$$(e) });
                            $$('bookTabDatatableContextMenu').show(e);
                            webix.html.preventEvent(e);
                        }
                    }
                },
            },
        ]
    }
}

// возвращает webix конфигурацию контекстного меню таба
export function BookTabContextMenu(employees) {
    return {
        view: 'contextmenu',
        id: 'bookTabDatatableContextMenu',
        data: [
            {
                value: BOOK_CONTEXT_MENU.give,
                id: BOOK_CONTEXT_MENU.give,
                submenu: employees,
            },
            BOOK_CONTEXT_MENU.add,
            BOOK_CONTEXT_MENU.edit,
            BOOK_CONTEXT_MENU.remove
        ],
    }
}