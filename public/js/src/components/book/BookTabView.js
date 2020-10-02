import { BOOK_CONTEXT_MENU } from "./CBookTab.js";

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
                    { id: 'ID', hidden: true },
                    { id: 'ISBN', header: 'ISBN', fillspace: true, },
                    { id: 'name', header: 'Название', fillspace: true, },
                    { id: 'author', header: 'Автор', fillspace: true, },
                    { id: 'publisher', header: 'Издательство', fillspace: true, },
                    { id: 'year', header: 'Год издания', fillspace: true, },
                    { id: 'status', header: 'Статус', fillspace: true, },
                ],
                data: [],
                onContext: {
                    // обработка вызова контекстного меню при остутствии данных
                    webix_view:function(e, id){
                        id = this.locate(e.target|| e.srcElement);
                        if(!id){
                            $$("bookTabDatatableContextMenu").setContext({ obj:webix.$$(e)});
                            $$("bookTabDatatableContextMenu").show(e);
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
            BOOK_CONTEXT_MENU.toEvent,
            BOOK_CONTEXT_MENU.toEmployee,
            BOOK_CONTEXT_MENU.add,
            BOOK_CONTEXT_MENU.edit,
            BOOK_CONTEXT_MENU.remove
        ],
    }
}